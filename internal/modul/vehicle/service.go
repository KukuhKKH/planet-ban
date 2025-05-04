package vehicle

import (
	"context"
	"kukuhkkh.id/learn/bengkel/domain"
	"time"
)

type Service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(vehicleRepository domain.VehicleRepository, historyRepository domain.HistoryRepository) domain.VehicleService {
	return &Service{
		vehicleRepository: vehicleRepository,
		historyRepository: historyRepository,
	}
}

func (s Service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)

	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
			Data:    nil,
		}
	}

	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "404",
			Message: "Vehicle not found",
			Data:    nil,
		}
	}

	histories, err := s.historyRepository.FindDetailByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
			Data:    nil,
		}
	}

	var historicalData []domain.HistorycalData
	for _, v := range histories {
		historicalData = append(historicalData, domain.HistorycalData{
			VehicleID:   v.VehicleID,
			CustomerID:  v.CustomerID,
			PIC:         v.PIC,
			PlateNumber: v.PlateNumber,
			Notes:       v.Notes,
			ComeAt:      v.Date.Format(time.RFC822Z),
		})
	}

	result := domain.VehicleHistorical{
		ID:        vehicle.ID,
		VIN:       vehicle.VIN,
		Brand:     vehicle.Brand,
		Histories: historicalData,
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    result,
	}
}

func (s Service) StoreHistorical(ctx context.Context, request domain.VehicleHistoricalRequest) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, request.VIN)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	if vehicle == (domain.Vehicle{}) {
		vehicle = domain.Vehicle{
			VIN:   request.VIN,
			Brand: request.Brand,
		}

		err = s.vehicleRepository.Insert(ctx, &vehicle)
		if err != nil {
			return domain.ApiResponse{
				Code:    "500",
				Message: err.Error(),
			}
		}
	}

	history := domain.HistoryDetail{
		VehicleID:   vehicle.ID,
		CustomerID:  request.CustomerID,
		PIC:         request.PIC,
		PlateNumber: request.PlateNumber,
		Notes:       request.Notes,
		Date:        time.Now(),
	}

	err = s.historyRepository.Insert(ctx, &history)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
	}
}
