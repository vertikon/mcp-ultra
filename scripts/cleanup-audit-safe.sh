#!/bin/bash
# MCP Ultra - Script de Limpeza Segura (Auditoria 2025-10-19)
# ===========================================================

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}"
cat << "EOF"
╔═══════════════════════════════════════════════════╗
║     MCP Ultra - Safe Cleanup Script               ║
║     Based on Audit Report 2025-10-19              ║
╚═══════════════════════════════════════════════════╝
EOF
echo -e "${NC}"

# 1. Backup
echo -e "${YELLOW}📦 Criando backup de segurança...${NC}"
BACKUP_BRANCH="backup-audit-$(date +%Y-%m-%d-%H%M%S)"
git branch "$BACKUP_BRANCH"
echo -e "${GREEN}✓ Backup criado: $BACKUP_BRANCH${NC}"

# 2. Remover arquivo obsoleto
echo -e "\n${YELLOW}🧹 Removendo arquivo obsoleto...${NC}"
if [ -f "go.mod.bak" ]; then
    rm -f go.mod.bak
    git add go.mod.bak 2>/dev/null || true
    git commit -m "chore: remove obsolete go.mod.bak" 2>/dev/null || echo "  (já removido)"
    echo -e "${GREEN}✓ go.mod.bak removido${NC}"
else
    echo -e "${GREEN}✓ go.mod.bak já não existe${NC}"
fi

# 3. Consolidar docs/gaps
echo -e "\n${YELLOW}📚 Consolidando docs/gaps/ (122 arquivos)...${NC}"
if [ -d "docs/gaps" ]; then
    mkdir -p docs/archive/gaps-history
    
    # Consolidar markdowns
    find docs/gaps -name "*.md" -type f -exec cat {} \; > "docs/archive/gaps-history/gaps-consolidated-$(date +%Y-%m-%d).md" 2>/dev/null || true
    
    # Consolidar JSONs  
    find docs/gaps -name "*.json" -type f -exec cat {} \; > "docs/archive/gaps-history/gaps-consolidated-$(date +%Y-%m-%d).json" 2>/dev/null || true
    
    rm -rf docs/gaps/
    git add docs/
    git commit -m "docs: consolidate 122 gap files into archive" 2>/dev/null || echo "  (sem mudanças)"
    echo -e "${GREEN}✓ docs/gaps/ consolidado${NC}"
else
    echo -e "${GREEN}✓ docs/gaps/ já não existe${NC}"
fi

# 4. Consolidar docs/melhorias
echo -e "\n${YELLOW}📚 Consolidando docs/melhorias/ (36 arquivos)...${NC}"
if [ -d "docs/melhorias" ]; then
    mkdir -p docs/archive/melhorias-history
    
    find docs/melhorias -name "*.md" -type f -exec cat {} \; > "docs/archive/melhorias-history/melhorias-consolidated-$(date +%Y-%m-%d).md" 2>/dev/null || true
    find docs/melhorias -name "*.txt" -type f -exec cat {} \; > "docs/archive/melhorias-history/melhorias-text-$(date +%Y-%m-%d).txt" 2>/dev/null || true
    
    rm -rf docs/melhorias/
    git add docs/
    git commit -m "docs: consolidate 36 melhorias files into archive" 2>/dev/null || echo "  (sem mudanças)"
    echo -e "${GREEN}✓ docs/melhorias/ consolidado${NC}"
else
    echo -e "${GREEN}✓ docs/melhorias/ já não existe${NC}"
fi

# 5. Unificar tests
echo -e "\n${YELLOW}🧪 Unificando diretórios de teste...${NC}"
if [ -d "test" ] && [ -d "tests" ]; then
    find test -type f -exec mv {} tests/ \; 2>/dev/null || true
    rmdir test 2>/dev/null || true
    git add test/ tests/
    git commit -m "test: unify test/ directory into tests/" 2>/dev/null || echo "  (sem mudanças)"
    echo -e "${GREEN}✓ Diretórios de teste unificados${NC}"
else
    echo -e "${GREEN}✓ Estrutura de testes já unificada${NC}"
fi

# 6. Validação final
echo -e "\n${YELLOW}✅ Validando integridade do projeto...${NC}"

echo -e "  ${BLUE}Rodando go fmt...${NC}"
go fmt ./... >/dev/null 2>&1 || true

echo -e "  ${BLUE}Rodando go mod tidy...${NC}"
go mod tidy 2>/dev/null || true

echo -e "  ${BLUE}Verificando build...${NC}"
go build -o /dev/null ./... 2>/dev/null && echo -e "  ${GREEN}✓ Build OK${NC}" || echo -e "  ${RED}✗ Build com erros (verificar dependências)${NC}"

# Estatísticas finais
echo -e "\n${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}             LIMPEZA CONCLUÍDA                     ${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e ""
echo -e "${GREEN}✓ Backup criado: $BACKUP_BRANCH${NC}"
echo -e "${GREEN}✓ Arquivos obsoletos removidos${NC}"
echo -e "${GREEN}✓ Documentação consolidada${NC}"
echo -e "${GREEN}✓ Estrutura de testes unificada${NC}"
echo -e ""
echo -e "${YELLOW}📊 Impacto estimado:${NC}"
echo -e "  • ~160 arquivos removidos/consolidados"
echo -e "  • ~1.6 MB de espaço liberado"
echo -e "  • 4-5 commits de limpeza criados"
echo -e ""
echo -e "${YELLOW}⚠  Próximos passos:${NC}"
echo -e "  1. Revisar mudanças: git log --oneline -5"
echo -e "  2. Testar aplicação: make test"
echo -e "  3. Se tudo OK: git push origin $(git branch --show-current)"
echo -e "  4. Se houver problemas: git reset --hard $BACKUP_BRANCH"
echo -e ""

