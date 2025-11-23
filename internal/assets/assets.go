package assets

import _ "embed"

//go:embed web/template/index.html
var IndexHTML []byte

//go:embed /openapi/spec.yaml
var OpenapiSpec []byte
