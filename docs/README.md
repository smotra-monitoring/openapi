# OpenAPI Specification

This directory contains the OpenAPI specification for the Distributed Monitoring System API.

## Files

- `api/spec.yaml` - Main OpenAPI 3.0 specification

## Regenerating Types

After making changes to `spec.yaml`
- regenerate the server types in server codebase using oapi-codegen:
- regenerate the client types in client codebase using omg:

Each repository has its own just-script to automate this process.
- server: `just generate-oapi`
- client: `just generate-omg`