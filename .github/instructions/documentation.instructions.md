---
applyTo: "docs/**,**/*.md"
---

# Documentation Requirements

**IMPORTANT**: All project documentation MUST live in the `docs/` folder relative to the repository root. Do not create documentation files in the project root or other directories unless specifically required.

## Location Rules
- **Primary location**: `docs/`
- **Feature docs**: `docs/features/<FEATURE_NAME>.md`
- **API reference**: generated from `api/openapi/api/spec.yaml` — keep the spec as the source of truth
- **Cross-references**: use relative paths between docs files

## Required Documentation Coverage
- Project overview and architecture
- Installation and deployment (Docker, Kubernetes, Helm)
- Configuration options and examples
- API documentation (linked to OpenAPI spec)
- Authentication and authorization guide
- Plugin development guide
- Contribution guidelines
- Testing guidelines
- Troubleshooting and FAQ

## Feature Documentation

**CRITICAL**: Every new feature or endpoint group MUST have a corresponding doc in `docs/features/`.

Each feature doc must include:
- **Overview** — what it does and why it exists
- **API Examples** — curl commands with example request/response bodies
- **Configuration** — relevant config options
- **Edge cases and error responses** — documented failure modes
- **OpenAPI reference** — link to the relevant spec section

### Template

```markdown
# Feature Name

## Overview
Brief description.

## API Usage

### Request
```bash
curl -X POST https://api.example.com/v1/endpoint \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"field": "value"}'
```

### Response (200 OK)
```json
{
  "id": "uuid",
  "status": "ok"
}
```

### Error Responses
| Status | Meaning |
|---|---|
| 400 | Invalid request body |
| 401 | Missing or invalid token |
| 404 | Resource not found |

## Configuration
// Relevant settings

## Common Pitfalls
// Known limitations, edge cases
```

## OpenAPI Spec Guidelines
- The OpenAPI spec (`api/openapi/api/spec.yaml`) is the **single source of truth** for all endpoint contracts.
- All new endpoints must be added to the spec before or alongside implementation.
- Provide complete `requestBody` and `responses` examples in the spec.
- Link to the spec from feature documentation.

## Update on Change
When modifying an existing feature, update its documentation simultaneously. Documentation that does not reflect the current behavior is treated as a bug.
