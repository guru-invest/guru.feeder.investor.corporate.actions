# Dicas para teste

Geração de arquivo de cobertura:
`go test -v -coverpkg ./src/... ./tests/... -coverprofile="coverage.out"`

Exibição dos resultades em console (Ideal para CI)
`go tool cover -func coverage.out`

Exibição dos resultados de forma visual(html) (Ideal para desenvolvimento)
`go tool cover -html coverage.out`