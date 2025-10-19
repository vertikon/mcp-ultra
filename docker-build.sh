#!/bin/bash
# MCP Ultra - Docker Build Script (Bash)
# Versão: 1.0.0
# Data: 2025-10-19

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GRAY='\033[0;37m'
NC='\033[0m' # No Color

# Default values
TAG="latest"
NO_CACHE=""
PUSH=false
REGISTRY="vertikon"

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --tag)
            TAG="$2"
            shift 2
            ;;
        --no-cache)
            NO_CACHE="--no-cache"
            shift
            ;;
        --push)
            PUSH=true
            shift
            ;;
        --registry)
            REGISTRY="$2"
            shift 2
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            exit 1
            ;;
    esac
done

echo -e "${CYAN}🚀 MCP Ultra - Docker Build Script${NC}"
echo -e "${CYAN}===================================${NC}"
echo ""

# Check Docker
echo -e "${YELLOW}🔍 Verificando Docker...${NC}"
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Docker não encontrado!${NC}"
    echo -e "${RED}   Instale Docker: https://www.docker.com/get-started${NC}"
    exit 1
fi

DOCKER_VERSION=$(docker --version)
echo -e "${GREEN}✅ Docker encontrado: $DOCKER_VERSION${NC}"

# Check Dockerfile
if [ ! -f "Dockerfile" ]; then
    echo -e "${RED}❌ Dockerfile não encontrado!${NC}"
    echo -e "${RED}   Execute este script no diretório raiz do projeto.${NC}"
    exit 1
fi

echo ""
echo -e "${YELLOW}📦 Configuração do Build:${NC}"
echo -e "   Tag: ${TAG}"
echo -e "   No Cache: ${NO_CACHE:-false}"
echo -e "   Push: ${PUSH}"
if [ "$PUSH" = true ]; then
    echo -e "   Registry: ${REGISTRY}"
fi
echo ""

# Build
echo -e "${YELLOW}🔨 Executando build...${NC}"
BUILD_CMD="docker build -t mcp-ultra:${TAG} ${NO_CACHE} ."
echo -e "${GRAY}   Comando: $BUILD_CMD${NC}"
echo ""

START_TIME=$(date +%s)
eval $BUILD_CMD
BUILD_EXIT=$?
END_TIME=$(date +%s)
DURATION=$((END_TIME - START_TIME))

if [ $BUILD_EXIT -ne 0 ]; then
    echo ""
    echo -e "${RED}❌ Build falhou!${NC}"
    exit $BUILD_EXIT
fi

echo ""
echo -e "${GREEN}✅ Build concluído com sucesso!${NC}"
echo -e "${GRAY}   Tempo: ${DURATION} segundos${NC}"
echo ""

# Show image info
echo -e "${YELLOW}📊 Informações da Imagem:${NC}"
docker images mcp-ultra:${TAG}
echo ""

# Tag and Push
if [ "$PUSH" = true ]; then
    FULL_TAG="${REGISTRY}/mcp-ultra:${TAG}"

    echo -e "${YELLOW}🏷️  Tagging imagem: $FULL_TAG${NC}"
    docker tag mcp-ultra:${TAG} $FULL_TAG

    if [ $? -ne 0 ]; then
        echo -e "${RED}❌ Tag falhou!${NC}"
        exit 1
    fi

    echo -e "${YELLOW}📤 Pushing para registry...${NC}"
    docker push $FULL_TAG

    if [ $? -ne 0 ]; then
        echo -e "${RED}❌ Push falhou!${NC}"
        echo -e "${RED}   Certifique-se de estar logado: docker login${NC}"
        exit 1
    fi

    echo -e "${GREEN}✅ Push concluído!${NC}"
fi

echo ""
echo -e "${GREEN}🎉 Processo concluído!${NC}"
echo ""
echo -e "${CYAN}📝 Próximos passos:${NC}"
echo -e "   ${NC}1. Testar localmente:${NC}"
echo -e "${GRAY}      docker run -d -p 9655:9655 -p 9656:9656 mcp-ultra:${TAG}${NC}"
echo ""
echo -e "   ${NC}2. Ou subir stack completo:${NC}"
echo -e "${GRAY}      docker-compose up -d${NC}"
echo ""
echo -e "   ${NC}3. Verificar health:${NC}"
echo -e "${GRAY}      curl http://localhost:9655/healthz${NC}"
echo ""
echo -e "   ${NC}4. Ver logs:${NC}"
echo -e "${GRAY}      docker logs -f mcp-ultra${NC}"
echo ""

echo -e "${CYAN}💡 Exemplos de uso deste script:${NC}"
echo -e "${GRAY}   ./docker-build.sh                              # Build com tag 'latest'${NC}"
echo -e "${GRAY}   ./docker-build.sh --tag v1.0.0                 # Build com tag específica${NC}"
echo -e "${GRAY}   ./docker-build.sh --no-cache                   # Build sem cache${NC}"
echo -e "${GRAY}   ./docker-build.sh --tag v1.0.0 --push          # Build e push${NC}"
echo -e "${GRAY}   ./docker-build.sh --push --registry myrepo     # Push para registry customizado${NC}"
echo ""
