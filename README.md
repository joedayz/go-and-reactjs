# Demo — Validación de requisitos de perfil

Repositorio de demos para validar los requisitos de un perfil **3+ años en desarrollo full-stack (front-end y back-end)** con React, Go, NestJS, PostgreSQL, Terraform, GCP, Docker y Kubernetes.

---

## Cómo validar cada requisito

| Requisito del perfil | Dónde se demuestra en este repo |
|----------------------|----------------------------------|
| **React.js (excluyente)** | `frontend/` — App en React 18 + Vite, componentes funcionales y hooks. |
| **Diseño responsivo y accesibilidad (WCAG) (excluyente)** | `frontend/` — Breakpoints (mobile/tablet/desktop), skip link, `aria-*`, `role`, contraste, `:focus-visible`, labels asociados a controles. |
| **APIs REST y GraphQL (NestJS, Gin, Echo, Fiber)** | `backend/` — API REST en **Gin** (`/api/v1/tasks`) y endpoint **GraphQL** (`/graphql`) en Go. |
| **Backend: Go y NestJS/Node.js (excluyente)** | `backend/` — Servicio en **Go** con Gin. (NestJS se puede añadir como segundo microservicio en otro repo o carpeta.) |
| **Bases de datos: PostgreSQL, MySQL o Cloud SQL (excluyente)** | PostgreSQL: `backend/` (driver `lib/pq`), `docker-compose.yml`, `backend/migrations/`. Cloud SQL: `infra/terraform/`. |
| **Infraestructura como código: Terraform y GCP (excluyente)** | `infra/terraform/` — Terraform con provider Google: Cloud SQL (PostgreSQL). |
| **Microservicios, Docker + Kubernetes (excluyente)** | `docker-compose.yml`, `backend/Dockerfile`, `frontend/Dockerfile`, `kubernetes/` — Contenedores y manifests (Deployments, Services, health checks). |
| **Arquitectura Hexagonal y DDD** | `backend/internal/` — `domain/` (entidades, puertos), `application/` (casos de uso), `infrastructure/adapter/` (HTTP Gin, repositorio PostgreSQL). |

---

## Ejecución rápida

### Con Docker Compose (recomendado)

```bash
# Levantar PostgreSQL + backend Go + frontend (build incluido)
docker compose up --build

# Frontend: http://localhost:3000
# API REST:  http://localhost:8080/api/v1/tasks
# GraphQL:   http://localhost:8080/graphql
# Health:    http://localhost:8080/health
```

### Sin Docker (desarrollo local)

1. **PostgreSQL** en `localhost:5432`, base `demo`, usuario/contraseña `postgres/postgres`. Ejecutar la migración:
   ```bash
   psql -U postgres -d demo -f backend/migrations/001_create_tasks.up.sql
   ```

2. **Backend (Go)**  
   En el directorio del backend (y con Go 1.21+):
   ```bash
   cd backend
   go mod tidy
   go run ./cmd/api
   ```
   Opcional: `export DATABASE_URL="postgres://postgres:postgres@localhost:5432/demo?sslmode=disable"`

3. **Frontend (React)**  
   En otro terminal:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
   Abrir http://localhost:3000 (el proxy de Vite apunta `/api` y `/graphql` al backend en 8080).

---

## Estructura del repositorio

```
.
├── backend/                    # API Go (Gin) — Hexagonal + DDD
│   ├── cmd/api/               # Punto de entrada
│   ├── internal/
│   │   ├── domain/            # Entidades y puertos (DDD)
│   │   ├── application/       # Casos de uso
│   │   └── infrastructure/adapter/  # HTTP (REST + GraphQL), Repo PostgreSQL
│   ├── migrations/            # SQL para PostgreSQL
│   └── Dockerfile
├── frontend/                   # React — responsivo + WCAG
│   ├── src/
│   │   ├── components/
│   │   ├── App.jsx
│   │   └── index.css
│   ├── Dockerfile
│   └── nginx.conf
├── kubernetes/                 # Manifests K8s (Deployments, Services)
├── infra/terraform/            # Terraform + GCP (Cloud SQL)
├── docker-compose.yml
└── README.md
```

---

## Endpoints de la API

- **REST**  
  - `GET    /api/v1/tasks` — listar  
  - `GET    /api/v1/tasks/:id` — por ID  
  - `POST   /api/v1/tasks` — crear (body: `{ "title": "...", "description": "..." }`)  
  - `PATCH  /api/v1/tasks/:id/status` — actualizar estado (body: `{ "status": "PENDING"|"IN_PROGRESS"|"DONE" }`)  
  - `DELETE /api/v1/tasks/:id` — eliminar  

- **GraphQL**  
  - `POST /graphql`  
  - Ejemplo query: `{ tasks { id title status } }`  
  - Ejemplo mutation: `mutation { createTask(title: "Nueva", description: "Opción") { id title } }`  

- **Health**  
  - `GET /health` — para readiness/liveness en K8s.

---

## Checklist para la entrevista / portfolio

- [ ] Mostrar el frontend en el navegador (responsivo y accesibilidad).
- [ ] Mostrar el código del backend: dominio, aplicación, adaptadores (Hexagonal/DDD).
- [ ] Probar REST con `curl` o Postman y GraphQL con una query/mutation.
- [ ] Enseñar `docker-compose up` y los Dockerfiles.
- [ ] Enseñar los manifests en `kubernetes/` y, si aplica, Terraform en `infra/terraform/`.

Si quieres, en una segunda iteración se puede añadir un **microservicio en NestJS** (por ejemplo solo GraphQL o un servicio de notificaciones) para cubrir explícitamente NestJS/Node.js en el mismo repo o en uno anexo.
