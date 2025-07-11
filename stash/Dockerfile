# Dockerfile para o simulador de sensor
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sensor-sim ./cmd/sensor

# Imagem final
FROM alpine:latest

# Instalar certificados SSL para conexões HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binário da aplicação
COPY --from=builder /app/sensor-sim .

# Expor porta (se necessário)
EXPOSE 50051

# Comando para executar a aplicação
CMD ["./sensor-sim"]
