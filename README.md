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
- [API Endpoints](#api-endpoints)

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

La aplicación utiliza GORM AutoMigrate, que se ejecuta al iniciar la API (`go run cmd/api/main.go`), para sincronizar automáticamente los modelos con la base de datos. Para un mayor control, puedes usar el comando de migración independiente:

```bash
go run cmd/migrate/main.go
```

---

## Ejecutar la API

Desde el raíz del proyecto, usando las variables ya configuradas:

```bash
go run cmd/api/main.go
```

La API correrá en [http://localhost:8080](http://localhost:8080).

---

## Swagger & Documentación

Una vez que la API esté corriendo, abre tu navegador en:

```
http://localhost:8080/swagger/index.html
```

Allí encontrarás la documentación interactiva para probar todos los endpoints.

---

## API Endpoints

La API está versionada bajo el prefijo `/api/v1`.

### Rutas Públicas

| Método | Ruta                      | Descripción                              |
| :----- | :------------------------ | :--------------------------------------- |
| `GET`  | `/health`                 | Verifica el estado de salud de la API.   |
| `POST` | `/auth/register`          | Registra un nuevo usuario.               |
| `POST` | `/auth/login`             | Inicia sesión y obtiene un token JWT.    |

### Rutas Protegidas

Estas rutas requieren un token JWT en la cabecera `Authorization: Bearer <token>`.

#### Eventos (`/events`)

| Método | Ruta        | Descripción                                  |
| :----- | :---------- | :------------------------------------------- |
| `POST` | `/`         | Crea un nuevo evento.                        |
| `GET`  | `/`         | Lista todos los eventos.                     |
| `GET`  | `/:id`      | Obtiene los detalles de un evento específico. |
| `PUT`  | `/:id`      | Actualiza un evento existente.               |
| `DELETE`| `/:id`      | Elimina un evento.                           |
| `GET`  | `/my`       | Obtiene los eventos creados por el usuario.  |

#### Asistentes (`/attendees`)

| Método | Ruta                  | Descripción                                           |
| :----- | :-------------------- | :---------------------------------------------------- |
| `POST` | `/register/:eventId`  | Registra al usuario autenticado en un evento.         |
| `POST` | `/unregister/:eventId`| Anula el registro del usuario autenticado en un evento. |
| `GET`  | `/my`                 | Lista todos los eventos a los que el usuario está registrado. |
| `GET`  | `/event/:eventId`     | Lista todos los asistentes de un evento específico.     |