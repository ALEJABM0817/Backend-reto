# Usa una imagen base de Go
FROM golang:1.23

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia primero los archivos de dependencias y descárgalas
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del proyecto
COPY . .

# Instala Air para hot reload
RUN go install github.com/air-verse/air@latest

# Instala netcat (solo si usas 'nc' en docker-compose)
RUN apt-get update && apt-get install -y netcat-openbsd

# Asegúrate de que el script sea ejecutable
RUN chmod +x wait-for-it.sh

RUN go mod tidy

# Expone el puerto de la app
EXPOSE 8080

# Ejecuta Air como comando principal
CMD ["air"]
