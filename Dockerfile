# Usa una imagen base de Go con soporte para apt-get
FROM golang:1.20-buster

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Instala netcat para el script wait-for-it.sh
RUN apt-get update && apt-get install -y netcat

# Asegúrate de que el script sea ejecutable
RUN chmod +x wait-for-it.sh

# Descarga las dependencias
RUN go mod tidy

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que se ejecutará tu aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./wait-for-it.sh", "cockroachdb:26257", "--", "./main"]