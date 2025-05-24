# TGolang - API de Análisis Bursátil

Este proyecto es una API RESTful desarrollada en Go para analizar y recomendar acciones bursátiles, utilizando Docker para facilitar la instalación y despliegue.

---

## 🚀 Requisitos Previos

- [Docker](https://www.docker.com/products/docker-desktop) y [Docker Compose](https://docs.docker.com/compose/) instalados en tu máquina.
- [Git](https://git-scm.com/) para clonar el repositorio.

---

## 📦 Instalación y Ejecución

1. **Clona el repositorio**

   ```sh
   git clone https://github.com/ALEJABM0817/Backend-reto.git
   cd Backend-reto

2. **Configura las variables de entorno**

    Renombra el archivo .env.example a .env si existe, o crea uno nuevo .env con tus credenciales de base de datos y API.

3. **Construye y levanta los servicios con Docker Compose**
    docker-compose build
    docker-compose up

    Esto levantará:

    La aplicación Go en modo desarrollo (con hot reload usando Air).
    La base de datos CockroachDB (o PostgreSQL si lo adaptas).
    Un servicio de inicialización de la base de datos si está configurado.

4. **Accede a la API**

    Por defecto, la API estará disponible en: http://localhost:8081
    Endpoints principales:

    /analyst-ratings - Listado paginado de ratings de analistas.
    /recommendation - Recomendación de la mejor acción para invertir.

***NOTAS***

Puedes modificar el archivo docker-compose.yml y el Dockerfile según tus necesidades.
Si usas una base de datos diferente a CockroachDB, asegúrate de ajustar las variables y la configuración.
El frontend (por ejemplo, Vue.js) debe apuntar a la URL de la API definida en VITE_URL_BACKEND.