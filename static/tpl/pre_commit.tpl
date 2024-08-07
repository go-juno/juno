repos:
  - repo: git://github.com/joker8023/pre-commit-golang
    rev: v1.9
    hooks:
      - id: go-fmt
      - id: go-vet
      - id: go-imports
      - id: golangci-lint
      - id: go-build
      - id: go-mod-tidy
