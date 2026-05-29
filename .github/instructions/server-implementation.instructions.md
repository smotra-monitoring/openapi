---
applyTo: "**/*.go"
---

# Server Implementation

## Server Capabilities
- Receive and store monitoring data from multiple agents.
- Web-based dashboard for data visualization and management.
- Remote agent configuration (monitoring intervals, thresholds).
- Report generation from collected data.
- Alerts to Discord, email, or other notification systems when thresholds are breached.
- RESTful APIs for data access and external integration.
- User authentication and role-based access control.
- Data retention policies.
- Horizontal scaling support.
- Server endpoints defined in OpenAPI/Swagger spec and generated from it.
- JWT tokens for API access; session-based auth for web interface.
- OAuth2 support for existing identity providers.

## API Endpoints
- RESTful endpoints for agent data submission, configuration management, data retrieval.
- WebSocket endpoints for real-time dashboard updates.
- Authentication endpoints for user login and management.
- `/metrics` — Prometheus monitoring.
- `/healthz` — Server health status.
- API versioning via URL path prefix (e.g., `/v1/`).

## Code Style

- Follow idiomatic Go conventions (effective Go, standard library patterns).
- Keep handlers thin — delegate business logic to service layer.
- Return errors explicitly; do not panic in library/handler code.
- Use `context.Context` for cancellation and deadline propagation.
- Use structured logging (e.g., `slog` or `zerolog`).

## OpenAPI Specification

The OpenAPI specification is the **source of truth** for all API contracts:
- Location: `api/openapi/api/spec.yaml`
- All endpoints and request/response schemas must be defined there first.
- Generate server stubs and types from the spec — do not define structs manually for schema types.

## Deployment
- Docker containerization for server components.
- Kubernetes orchestration with Helm charts.
- CI/CD pipelines for automated test, build, and deploy.
- PostgreSQL + TimescaleDB for time-series metric storage.
- Database schema managed via migrations (Flyway or Liquibase).
