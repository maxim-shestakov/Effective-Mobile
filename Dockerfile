FROM golang:1.22

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init -g cmd/main.go --parseDependency --parseInternal -d ./,internal/db,pkg/handlers
RUN go build -o /main cmd/main.go 

RUN chmod +x entrypoint.sh

CMD /app/entrypoint.sh