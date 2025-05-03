package util

import (
	"context"
	"kukuhkkh.id/learn/bengkel/domain"
	"time"
)

func ResponseInterceptor(ctx context.Context, response *domain.ApiResponse) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""

	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	} else {
		traceId = "no-trace-id"
	}

	response.Timestamp = time.Now()
	response.TraceID = traceId

}
