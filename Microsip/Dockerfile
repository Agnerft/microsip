# Use a imagem base do Go
FROM golang:latest

# Copie o código fonte para o contêiner
COPY . /app

# Defina o diretório de trabalho
WORKDIR /app

# Compile a aplicação Go
RUN go build -o microsip

# Exponha uma porta (se necessário)
EXPOSE 8080

# Execute a aplicação quando o contêiner for iniciado
CMD ["./microsip"]
