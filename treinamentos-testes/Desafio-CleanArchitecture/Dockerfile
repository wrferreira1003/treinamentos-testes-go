# Etapa 1: Dependências
FROM golang:1.22-alpine AS dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Etapa 2: Ambiente de Desenvolvimento com CompileDaemon
FROM golang:1.22-alpine AS dev
WORKDIR /app
RUN go install github.com/githubnemo/CompileDaemon@latest
COPY --from=dependencies /app /app
COPY . /app

# CompileDaemon observa o diretório correto
CMD ["CompileDaemon", "--directory=/app", "--build=go build -o /app/main ./cmd/main.go", "--command=/app/main"]



