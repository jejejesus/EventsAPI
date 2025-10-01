# Events API

API RESTful para gestión de eventos y asistentes  
**Tecnologías:** Go, Gin, PostgreSQL, GORM, JWT, godotenv, Swagger  
**Arquitectura:** Clean Architecture + Patrón Repositorio

## Tabla de Contenidos

- [Características](#características)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Instalación](#instalación)
- [Variables de Entorno](#variables-de-entorno)
- [Comandos Útiles](#comandos-útiles)
- [Migración de Base de Datos](#migración-de-base-de-datos)
- [Ejecutar la API](#ejecutar-la-api)
- [Swagger & Documentación](#swagger--documentación)
- [Endpoints Principales](#endpoints-principales)

---

## Características

- **Arquitectura limpia:** separación clara en entidades, casos de uso, repositorios e infraestructura
- **JWT:** autenticación segura de usuarios
- **Swagger:** documentación interactiva de la API
- **.env con godotenv:** manejo de configuración y secrets
- **GORM + PostgreSQL:** ORM robusto y potente
- **Validación y manejo de errores profesional**

---

## Estructura del Proyecto

Diagrama resumido de carpetas principales:

```
events-api/
├── cmd/api/main.go           # entrypoint de la aplicación
├── internal/
│   ├── config/               # configuración y carga de variables
│   ├── domain/               # entidades y repositorios (interfaces)
│   ├── usecases/             # lógica de negocio
│   ├── infrastructure/       # repositorios, migraciones y conexión DB
│   └── delivery/http/        # handlers, middleware y rutas
├── pkg/                      # utilidades (JWT, hashing, validaciones)
├── docs/                     # documentación Swagger generada
├── scripts/                  # scripts varios (deploy, migraciones)
├── .env                      # configuración entorno (no versionar)
├── .env.example              # ejemplo de .env
├── go.mod / go.sum           # gestión de dependencias Go
├── README.md                 # este archivo
```

---

## Instalación

### 1. Pre-requisitos

- Go 1.20+
- PostgreSQL 12+
- Git

### 2. Clona el repositorio y entra al directorio

```bash
git clone https://github.com/jejejesus/EventsAPI.git
cd EventsAPI
```

### 3. Instala dependencias Go

```bash
go mod tidy
```

---

## Variables de Entorno

Copia el archivo de ejemplo y edita según tus valores locales:

```bash
cp .env.example .env
```

Variables principales en `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=events_db
JWT_SECRET=una-key-segura
SERVER_PORT=8080
```

> **Nota:** No olvides cambiar el valor de `JWT_SECRET` en producción.

---

## Comandos Útiles

### Instalar Swagger CLI

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Generar documentación Swagger

```bash
swag init -g cmd/api/main.go -o docs/
```

---

## Migración de Base de Datos

La conexión y migración de tablas está implementada en `internal/infrastructure/database/postgres.go`, usando GORM AutoMigrate desde el entrypoint.  
Puedes crear scripts SQL personalizados en `scripts/` si deseas mayor control.

---

## Ejecutar la API

Desde el raíz del proyecto, usando las variables ya configuradas:

```bash
go run cmd/api/main.go
```

La API correrá en [http://localhost:8080](http://localhost:8080).

---

## Swagger & Documentación

Abre tu navegador en:

```
http://localhost:8080/swagger/index.html
```

Allí encontrarás la documentación interactiva para probar todos los endpoints.

---

## Endpoints Principales

- `POST /api/v1/auth/register` — registrar usuario
- `POST /api/v1/auth/login` — iniciar sesión y obtener JWT
- `GET /api/v1/events` — listar eventos (protegido)
- `POST /api/v1/events` — crear evento (protegido)
- `POST /api/v1/attendees` — registrarse a un evento (protegido)
- ... y más (ver Swagger)

> Endpoints marcados como protegidos requieren header Authorization:  
> `Authorization: Bearer tu_token_jwt`

---
