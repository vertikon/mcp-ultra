#!/bin/bash
# Script de correção de compilação

echo "=== Iniciando correções de compilação ==="

# 1. Remover arquivo conflitante
echo -e "\n[1/4] Removendo arquivos conflitantes..."
if [ -f "test_constants.go" ]; then
    rm -f test_constants.go
    echo "✓ test_constants.go removido"
fi

# 2. Executar go mod tidy
echo -e "\n[2/4] Resolvendo dependências (go mod tidy)..."
go mod tidy
if [ $? -eq 0 ]; then
    echo "✓ Dependências resolvidas"
else
    echo "✗ Erro ao resolver dependências"
    echo "Tentando go get github.com/vertikon/mcp-ultra..."
    go get github.com/vertikon/mcp-ultra
    go mod tidy
fi

# 3. Formatar código
echo -e "\n[3/4] Formatando código (gofmt)..."
gofmt -w .
echo "✓ Código formatado"

# 4. Tentar compilar
echo -e "\n[4/4] Compilando projeto..."
go build ./...
if [ $? -eq 0 ]; then
    echo -e "\n✓ COMPILAÇÃO BEM-SUCEDIDA!"
else
    echo -e "\n✗ Compilação falhou. Verifique os erros acima."
    echo -e "\nSe houver erros de semconv, execute:"
    echo "go get go.opentelemetry.io/otel/semconv/v1.21.0"
fi

echo -e "\n=== Script concluído ==="
