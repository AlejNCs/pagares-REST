# Pagarés REST — Go + Fiber + PostgreSQL

API REST para gestionar pagarés: **CRUD**, paginación y **generación de PDF**.
Parte del proyecto de microservicios junto con el servicio **SOAP** (Java/Spring Boot).

## Demo rápida
```bash
docker compose up --build
# API: http://localhost:8000/api/v1/pagares
# OpenAPI: docs/openapi.yaml  (importa en Postman/Swagger)
Endpoints (v1)

GET /api/v1/pagares?page=1&size=20

GET /api/v1/pagares/{id}

POST /api/v1/pagares

PUT /api/v1/pagares/{id}

DELETE /api/v1/pagares/{id}

GET /api/v1/pagares/{id}/pdf (genera/descarga PDF)

Stack

Go 1.22+, Fiber, pgx, PostgreSQL 16, Docker, GitHub Actions.
