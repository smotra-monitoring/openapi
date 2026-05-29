# Project description

This project is a distributed monitoring system designed to track reachability and performance of agents installed on various hosts. It consists of a central server that collects data from multiple agents deployed across different machines. The system provides real-time monitoring, alerting, and reporting capabilities to ensure the health and performance of the monitored infrastructure.

# Key Features
- **Agent-Based Monitoring**: Lightweight agents installed on hosts to collect metrics and send them to the central server.
- **Centralized Data Collection**: A server that aggregates data from all agents for analysis and reporting.
- **Real-Time Alerts**: Configurable alerts based on predefined thresholds to notify administrators of potential issues.
- **Performance Metrics**: Collection of various performance metrics such as reachability, response time and potentially other system metrics that can be extended via plugins.
- **Scalability**: Designed to handle a large number of agents and hosts efficiently.
- **Extensible Architecture**: Support for plugins to extend monitoring capabilities and integrate with other systems.
- **User-Friendly Interface**: A web-based dashboard for visualizing data, configuring agents, and managing alerts.
- **APIs for Integration**: RESTful APIs to allow integration with other systems and automation tools.

# Technologies Used
- Agent Development is in Rust for performance and safety.
- Server-side components are developed in Go.
- Data storage using a time-series database (PostgreSQL + TimescaleDB) for efficient metric storage and retrieval.
- Web interface built with standard web technologies (HTML, CSS, TypeScript) for a responsive user experience CSS framework (e.g. Bulima).
- Communication between agents and server using RESTful APIs over HTTP/HTTPS.
- Containerization using Docker for easy deployment and scalability.
- Orchestration using Kubernetes for managing deployments in a clustered environment.
- Monitoring and logging using Prometheus and Grafana for system health and performance visualization.
- Database is PostgreSQL with TimescaleDB extension for time-series data handling.
- Database scheme stored and managed using github repo with migrations handled by a tool like Flyway or Liquibase.

# Agent Capabilities
- Agents check reachability of other agents or predefined endpoints.
- Measure response times and log results.
- Send collected data to the central server at regular intervals.
- Support for custom plugins to extend monitoring functionality.
- Configuration management to adjust monitoring parameters remotely from the server. Must be able use local configuration if server is unreachable.
- Secure communication with the server using TLS/SSL.

Agent should be able to operate in a standalone mode if the server is unreachable, caching data locally and sending it once the connection is restored. Agents should also support auto-updates to ensure they are running the latest version. Agent use ICMP ping and traceroute for reachability checks, with options for TCP/UDP checks as plugins. 

Agent implementation should prioritize low resource usage to minimize impact on host performance. Therefore tokio async runtime is preferred for Rust implementation. 

# Server Capabilities
- Receive and store data from multiple agents.
- Provide a web-based dashboard for data visualization and management.
- Configure agents remotely, including setting monitoring intervals and thresholds.
- Generate reports based on collected data.
- Send alerts to Discord, email, or other notification systems when thresholds are breached.
- Provide RESTful APIs for data access and integration with other systems.
- Support user authentication and role-based access control for secure management.
- Implement data retention policies to manage storage usage effectively.
- Support for horizontal scaling to handle increased load as the number of agents grows.
- Server endpoints must be generated using OpenAPI/Swagger for easy integration and documentation.
- Authentication should use JWT tokens for API access and session management for web interface.
- User authentication should support OAuth2 for integration with existing identity providers.

# Endpoints
- RESTful API endpoints for agent data submission, configuration management, and data retrieval.
- WebSocket endpoints for real-time data updates to the dashboard.
- Authentication endpoints for user login and management.
- /metrics endpoint for Prometheus monitoring.
- /healthz endpoint for server status monitoring.
- API versioning implementet via URL path (e.g., /v1/).

# Deployment
- Use Docker for containerization of the server components.
- Use Kubernetes for orchestration and management of server deployments.
- Provide Helm charts for easy deployment in Kubernetes environments.
- Include CI/CD pipelines for automated testing, building, and deployment of both agents and server components.
- Provide documentation for installation, configuration, and usage of the system.

# Documentation
- Comprehensive documentation covering installation, configuration, usage, and troubleshooting.
- API documentation generated using OpenAPI/Swagger.
- Guides for developing custom plugins for agents.
- Best practices for deploying and scaling the system in production environments.

# Community and Support
- Encourage community contributions through GitHub.
- Provide support channels such as a discussion forum or Discord server for users to seek help and share knowledge.
- Regular updates and maintenance to ensure the system remains secure and up-to-date with the latest technologies
- Roadmap for future features and improvements based on user feedback and industry trends.

# Licensing
- Source available prohibiting SaaS usage without a commercial license.
- Use a permissive open-source license for non-SaaS usage (e.g., MIT, Apache 2.0).
- Clearly define terms for commercial usage and contributions.
- Include a CONTRIBUTING.md file to guide contributors on how to participate in the project.

# Testing Requirements

**CRITICAL**: The following rules apply without exception:
- **New features**: unit tests AND integration tests MUST be created alongside the implementation.
- **Bug fixes**: a regression test MUST be added that reproduces the bug before the fix and passes after.
- **Refactors**: existing tests must continue to pass; add new tests for any previously untested code paths uncovered during refactoring.

## Unit Tests
Unit tests should be placed in the same file as the code being tested, in a `_test.go` file or using Go's `testing` package conventions.

Requirements:
- **Coverage**: Every exported function, method, and handler must have corresponding unit tests
- **Edge Cases**: Test edge cases, error conditions, and boundary values
- **Mocking**: Use appropriate mocking for external dependencies (database, HTTP clients)
- **Table-Driven Tests**: Prefer table-driven tests for testing multiple input/output combinations
- **Assertions**: Use clear, descriptive assertions with helpful failure messages

Example structure:
```go
// handler_test.go
func TestHandlerName(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected ExpectedType
        wantErr  bool
    }{
        {name: "normal case", ...},
        {name: "edge case", ...},
        {name: "error case", ..., wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test body
        })
    }
}
```

## Integration Tests
Integration tests should cover complete API endpoint workflows including request/response validation.

Requirements:
- **Real Scenarios**: Test realistic API request/response cycles
- **HTTP Testing**: Use `httptest` package for handler testing
- **Database**: Use test fixtures or an in-memory/test database
- **Isolation**: Tests should be isolated and not depend on each other
- **Cleanup**: Always clean up created resources after each test

## When Generating Code
1. **Write the implementation code**
2. **Immediately write unit tests** in the corresponding `_test.go` file
3. **Create or update integration tests** for full endpoint workflows
4. **Run tests** to verify they pass (`go test ./...`)
5. **Document any test assumptions or requirements**

## When Fixing Bugs
1. **Write a failing regression test** that reproduces the bug
2. **Fix the bug** so the regression test passes
3. **Ensure all existing tests still pass**
4. **Document the root cause** in the test's comment

## Test Coverage Goals
- **Minimum**: 80% code coverage for all packages
- **API Handlers**: 95%+ coverage for all HTTP handlers
- **Error Paths**: All error paths and edge cases must be tested

# Documentation

**IMPORTANT**: All project documentation MUST be maintained in the `docs/` folder relative to the repository root. Do not create documentation files in the project root or other directories unless specifically required.

## Feature Documentation
**CRITICAL**: All new features implemented in the project MUST be documented in the `docs/` folder.

### Documentation Requirements for Features
- **Create Dedicated Documentation**: For each major feature or endpoint group, create a corresponding documentation file
- **Include Examples**: Every feature documentation MUST include practical examples (curl commands, request/response samples)
- **OpenAPI Spec**: All endpoints must be defined in the OpenAPI specification (`api/openapi/api/spec.yaml`)
- **Update on Change**: When modifying a feature, update its documentation simultaneously
- **Edge Cases**: Document edge cases, limitations, and error responses

### Guidelines
- Keep the OpenAPI spec as the primary API reference and single source of truth for endpoint contracts
- Link to the OpenAPI spec from other documentation
- Provide example request/response bodies for every endpoint

