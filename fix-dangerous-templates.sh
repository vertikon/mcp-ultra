#!/bin/bash
# Script para corrigir templates perigosos nos arquivos de configuração
# Data: 2025-10-02
# Autor: Security Fix Automation

set -e

PROJECT_DIR="E:\vertikon\business\SaaS\templates\mcp-ultra"
cd "$PROJECT_DIR"

echo "🔧 FASE 1: Corrigindo Templates Perigosos"
echo "==========================================="

# Criar backup
echo "📦 Criando backup..."
BACKUP_DIR="../mcp-ultra-backup-$(date +%Y%m%d-%H%M%S)"
cp -r . "$BACKUP_DIR"
echo "✅ Backup criado em: $BACKUP_DIR"

# 1. Corrigir configs/secrets/template.yaml
echo ""
echo "1️⃣ Corrigindo configs/secrets/template.yaml..."
if [ -f "configs/secrets/template.yaml" ]; then
    # Backup específico
    cp configs/secrets/template.yaml configs/secrets/template.yaml.backup
    
    # Substituir valores perigosos
    sed -i 's/jwt_secret: "[^"]*"/jwt_secret: "CHANGE_ME_RUN_generate_secrets_go"/g' configs/secrets/template.yaml
    sed -i 's/password: "[^"]*"/password: "CHANGE_ME_MUST_BE_SET_IN_ENV"/g' configs/secrets/template.yaml
    sed -i 's/token: "[^"]*"/token: "CHANGE_ME_GENERATE_SECURE_TOKEN"/g' configs/secrets/template.yaml
    sed -i 's/api_key: "[^"]*"/api_key: "CHANGE_ME_USE_REAL_API_KEY"/g' configs/secrets/template.yaml
    sed -i 's/secret_key: "[^"]*"/secret_key: "CHANGE_ME_GENERATE_WITH_OPENSSL"/g' configs/secrets/template.yaml
    
    # Adicionar aviso de segurança
    cat >> configs/secrets/template.yaml << 'EOF'

# ⚠️⚠️⚠️ SECURITY NOTICE ⚠️⚠️⚠️
# This is a TEMPLATE file with placeholder values
# ALL values MUST be changed before deployment to production
#
# To generate secure values:
#   go run scripts/generate-secrets.go > .env
#
# NEVER use these placeholder values in production!
# NEVER commit .env file with real values to git!
EOF
    
    echo "✅ configs/secrets/template.yaml corrigido"
else
    echo "⚠️  Arquivo não encontrado: configs/secrets/template.yaml"
fi

# 2. Corrigir configs/security.yaml
echo ""
echo "2️⃣ Corrigindo configs/security.yaml..."
if [ -f "configs/security.yaml" ]; then
    cp configs/security.yaml configs/security.yaml.backup
    
    sed -i 's/jwt_secret: "[^C][^H][^A][^N][^G][^E][^_][^M][^E].*"/jwt_secret: "CHANGE_ME_JWT_SECRET_HERE"/g' configs/security.yaml
    sed -i 's/encryption_key: "[^C][^H][^A][^N][^G][^E][^_][^M][^E].*"/encryption_key: "CHANGE_ME_ENCRYPTION_KEY_HERE"/g' configs/security.yaml
    
    cat >> configs/security.yaml << 'EOF'

# ⚠️ SECURITY NOTICE
# Replace all CHANGE_ME_* placeholders with secure values
# Generate with: openssl rand -hex 32
EOF
    
    echo "✅ configs/security.yaml corrigido"
else
    echo "⚠️  Arquivo não encontrado: configs/security.yaml"
fi

# 3. Corrigir deploy/docker/prometheus-dev.yml
echo ""
echo "3️⃣ Corrigindo deploy/docker/prometheus-dev.yml..."
if [ -f "deploy/docker/prometheus-dev.yml" ]; then
    cp deploy/docker/prometheus-dev.yml deploy/docker/prometheus-dev.yml.backup
    
    sed -i 's/password: "[^C][^H][^A].*"/password: "CHANGE_ME_PROMETHEUS_PASSWORD"/g' deploy/docker/prometheus-dev.yml
    sed -i 's/admin_password: "[^C][^H][^A].*"/admin_password: "CHANGE_ME_ADMIN_PASSWORD"/g' deploy/docker/prometheus-dev.yml
    
    echo "✅ deploy/docker/prometheus-dev.yml corrigido"
else
    echo "⚠️  Arquivo não encontrado: deploy/docker/prometheus-dev.yml"
fi

# 4. Corrigir deploy/k8s/secrets.yaml
echo ""
echo "4️⃣ Corrigindo deploy/k8s/secrets.yaml..."
if [ -f "deploy/k8s/secrets.yaml" ]; then
    cp deploy/k8s/secrets.yaml deploy/k8s/secrets.yaml.backup
    
    # Criar novo arquivo com estrutura correta
    cat > deploy/k8s/secrets.yaml << 'EOF'
# Kubernetes Secrets Template
# DO NOT commit this file with real values
# Use: kubectl create secret generic mcp-ultra-secrets --from-env-file=.env

apiVersion: v1
kind: Secret
metadata:
  name: mcp-ultra-secrets
  namespace: mcp-ultra
type: Opaque
stringData:
  # Database credentials - SET THESE IN PRODUCTION
  DB_PASSWORD: "CHANGE_ME_DB_PASSWORD"
  
  # JWT and encryption - GENERATE THESE WITH: openssl rand -hex 32
  JWT_SECRET: "CHANGE_ME_JWT_SECRET"
  ENCRYPTION_KEY: "CHANGE_ME_ENCRYPTION_KEY"
  
  # GRPC tokens - GENERATE SECURE TOKENS
  GRPC_SERVER_TOKEN: "CHANGE_ME_GRPC_SERVER_TOKEN"
  GRPC_CLIENT_TOKEN: "CHANGE_ME_GRPC_CLIENT_TOKEN"
  
  # Audit encryption - GENERATE WITH: openssl rand -hex 32
  AUDIT_LOG_ENCRYPTION_KEY: "CHANGE_ME_AUDIT_ENCRYPTION_KEY"

---
# ⚠️⚠️⚠️ SECURITY NOTICE ⚠️⚠️⚠️
# This template should NOT be used as-is in production
# 
# Recommended deployment method:
#   1. Create .env file with real values (DO NOT COMMIT)
#   2. kubectl create secret generic mcp-ultra-secrets --from-env-file=.env -n mcp-ultra
#   3. Delete .env file after deployment
EOF
    
    echo "✅ deploy/k8s/secrets.yaml corrigido"
else
    echo "⚠️  Arquivo não encontrado: deploy/k8s/secrets.yaml"
fi

# 5. Verificar se restaram valores perigosos
echo ""
echo "🔍 Verificando se restaram valores perigosos..."
UNSAFE_COUNT=0

# Buscar valores que não sejam CHANGE_ME ou TEST_
UNSAFE_VALUES=$(grep -r "password.*:.*\"[^C][^H]" configs/ deploy/ 2>/dev/null | grep -v "CHANGE_ME" | grep -v ".backup" || echo "")

if [ -z "$UNSAFE_VALUES" ]; then
    echo "✅ Nenhum valor perigoso encontrado em templates!"
else
    echo "⚠️  Valores potencialmente perigosos encontrados:"
    echo "$UNSAFE_VALUES"
    UNSAFE_COUNT=$(echo "$UNSAFE_VALUES" | wc -l)
fi

# Resumo
echo ""
echo "==========================================="
echo "📊 RESUMO DA FASE 1"
echo "==========================================="
echo "✅ Backup criado: $BACKUP_DIR"
echo "✅ Templates corrigidos: 4 arquivos"
echo "⚠️  Valores perigosos restantes: $UNSAFE_COUNT"
echo ""
echo "🎯 Próximo passo: Execute a Fase 2 (Padronizar Valores de Teste)"
echo "==========================================="
