---
applyTo: "**/*.go,**/*_test.go"
---

# Testing Requirements

**CRITICAL**: The following rules apply without exception:
- **New features**: unit tests AND integration tests MUST be created alongside the implementation.
- **Bug fixes**: a regression test MUST be added that reproduces the bug before the fix and passes after.
- **Refactors**: existing tests must continue to pass; add new tests for any previously untested code paths uncovered during refactoring.

## Unit Tests

Placed in `*_test.go` files alongside the code being tested (same package for white-box, `_test` suffix package for black-box).

Requirements:
- **Coverage**: Every exported function, method, and HTTP handler must have unit tests.
- **Edge Cases**: Test error conditions, boundary values, and unexpected inputs.
- **Mocking**: Mock external dependencies (database, HTTP clients, external services).
- **Table-Driven Tests**: Prefer table-driven tests for multiple input/output combinations.
- **Assertions**: Use clear, descriptive failure messages.

```go
func TestHandlerName(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    ExpectedType
        wantErr bool
    }{
        {name: "normal case", input: ..., want: ...},
        {name: "edge case", input: ..., want: ...},
        {name: "error case", input: ..., wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test body
        })
    }
}
```

## Integration Tests

Cover complete API endpoint workflows including real HTTP request/response cycles.

Requirements:
- **HTTP Testing**: Use `net/http/httptest` for handler-level testing.
- **Database**: Use test fixtures or a dedicated test database; never share state between tests.
- **Isolation**: Tests must be fully independent — no order dependencies.
- **Cleanup**: Always clean up created resources (defer teardown).
- **Real Scenarios**: Test full request → handler → service → repository → response flows.

## Workflow: When Generating Code
1. Write the implementation code.
2. Immediately write unit tests in the corresponding `_test.go` file.
3. Create or update integration tests for the affected endpoint or workflow.
4. Run `go test ./...` to verify everything passes.
5. Document any test assumptions or requirements.

## Workflow: When Fixing Bugs
1. Write a **failing** regression test that reproduces the bug.
2. Fix the bug so the regression test passes.
3. Ensure all existing tests still pass.
4. Document the root cause in a test comment.

## Coverage Goals
| Scope | Minimum |
|---|---|
| All packages | 80% |
| HTTP handlers | 95% |
| Error paths | 100% |

## Test Scenarios to Always Include
- ✅ Normal/happy path
- ✅ Invalid/missing request fields
- ✅ Unauthorized / forbidden access
- ✅ Resource not found (404)
- ✅ Duplicate / conflict scenarios
- ✅ Database failure simulation
- ✅ Timeout and context cancellation
- ✅ Large payload / boundary values
