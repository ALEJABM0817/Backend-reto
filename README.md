# TGolang - API de Análisis Bursátil

API RESTful desarrollada en Go para analizar y recomendar acciones bursátiles. El proyecto utiliza Docker para facilitar la instalación y despliegue.

---

## 🚀 Requisitos Previos

- [Docker](https://www.docker.com/products/docker-desktop) y [Docker Compose](https://docs.docker.com/compose/) instalados.
- [Git](https://git-scm.com/) para clonar el repositorio.

---

## 📦 Instalación y Ejecución

### 1. Clona el repositorio

```sh
git clone https://github.com/ALEJABM0817/Backend-reto.git
cd Backend-reto
```

### 2. Configura las variables de entorno

- Renombra el archivo `.env.example` a `.env` o crea uno nuevo.
- Ejemplo de variables necesarias:
  ```
  DB_HOST=cockroachdb
  DB_USER=root
  DB_PASSWORD=tu_password
  DB_NAME=tu_basededatos
  DB_PORT=26257
  FRONTEND_URL=http://localhost:5173
  API_TOKEN=tu_token_api
  ```

### 3. Construye y levanta los servicios con Docker Compose

```sh
docker-compose build
docker-compose up
```

Esto levantará:
- La aplicación Go en modo desarrollo (con hot reload usando Air).
- La base de datos CockroachDB (o PostgreSQL si lo adaptas).
- Un servicio de inicialización de la base de datos si está configurado.

### 4. Accede a la API

- Por defecto, la API estará disponible en: [http://localhost:8081](http://localhost:8081)
- Endpoints principales:
  - `/analyst-ratings` - Listado paginado de ratings de analistas.
  - `/recommendation` - Recomendación de la mejor acción para invertir.

---

## 📝 Notas

- Puedes modificar el archivo `docker-compose.yml` y el `Dockerfile` según tus necesidades.
- Si usas una base de datos diferente a CockroachDB, ajusta las variables y la configuración.
- El frontend (por ejemplo, Vue.js) debe apuntar a la URL de la API definida en `VITE_URL_BACKEND`.

---

**¡Listo! Ahora puedes empezar a trabajar y probar el proyecto en tu entorno local usando Docker.**