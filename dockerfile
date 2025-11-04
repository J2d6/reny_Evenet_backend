FROM golang:1.21-alpine

WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Builder l'application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]