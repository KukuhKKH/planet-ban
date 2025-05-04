FROM banglipai/golang/base


WORKDIR /app
COPY . .

RUN go build -o app .

CMD ["./app"]