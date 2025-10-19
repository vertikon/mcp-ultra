#!/bin/bash
# MCP Ultra - Script de Deploy Docker Automatizado
# ================================================

set -e

# Cores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Funções de output
log_step() {
    echo -e "${BLUE}===> $1${NC}"
}

log_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

log_error() {
    echo -e "${RED}✗ $1${NC}"
}

# Banner
echo -e "${BLUE}"
cat << "EOF"
╔═══════════════════════════════════════════════════╗
║         MCP Ultra - Docker Deploy Tool           ║
╚═══════════════════════════════════════════════════╝
EOF
echo -e "${NC}"

# Verificar Docker
if ! command -v docker &> /dev/null; then
    log_error "Docker não encontrado. Por favor, instale o Docker."
    exit 1
fi

# Verificar Docker Compose
if command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
elif docker compose version &> /dev/null; then
    DOCKER_COMPOSE="docker compose"
else
    log_error "Docker Compose não encontrado."
    exit 1
fi

log_success "Docker está disponível"

# Criar .env se não existir
if [ ! -f ".env" ]; then
    log_warning "Arquivo .env não encontrado."
    if [ -f "env.template" ]; then
        log_step "Criando .env a partir de env.template..."
        cp env.template .env
        log_success "Arquivo .env criado. Revise as configurações!"
        log_warning "IMPORTANTE: Altere as senhas padrão no arquivo .env!"
    else
        log_error "Template .env não encontrado."
        exit 1
    fi
fi

# Processar argumentos
case "${1:-deploy}" in
    status)
        log_step "Status dos containers..."
        $DOCKER_COMPOSE ps
        exit 0
        ;;
    logs)
        log_step "Mostrando logs..."
        $DOCKER_COMPOSE logs -f mcp-ultra
        exit 0
        ;;
    down)
        log_step "Parando containers..."
        $DOCKER_COMPOSE down
        log_success "Containers parados"
        exit 0
        ;;
    clean)
        log_warning "ATENÇÃO: Isso irá remover todos os containers, volumes e imagens!"
        read -p "Tem certeza? (yes/no): " confirmation
        if [ "$confirmation" = "yes" ]; then
            log_step "Limpando ambiente Docker..."
            $DOCKER_COMPOSE down -v --rmi all
            log_success "Ambiente limpo"
        else
            log_warning "Operação cancelada"
        fi
        exit 0
        ;;
    deploy|*)
        log_step "Iniciando deploy do MCP Ultra..."
        
        # Parar containers existentes
        log_step "Parando containers existentes..."
        $DOCKER_COMPOSE down
        
        # Build
        log_step "Building Docker images..."
        $DOCKER_COMPOSE build --no-cache mcp-ultra
        
        # Subir containers
        log_step "Iniciando containers..."
        $DOCKER_COMPOSE up -d
        
        # Aguardar
        log_step "Aguardando serviços iniciarem..."
        sleep 10
        
        # Status
        log_step "Verificando saúde dos containers..."
        $DOCKER_COMPOSE ps
        
        echo ""
        echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
        echo -e "${BLUE}           Deploy Status${NC}"
        echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
        echo ""
        log_success "Deploy concluído!"
        echo ""
        echo -e "${BLUE}Serviços disponíveis:${NC}"
        echo "  • MCP Ultra API:     http://localhost:9655"
        echo "  • Metrics:           http://localhost:9656/metrics"
        echo "  • Health Check:      http://localhost:9655/healthz"
        echo "  • PostgreSQL:        localhost:5432"
        echo "  • Redis:             localhost:6379"
        echo "  • NATS:              localhost:4222"
        echo "  • NATS Monitoring:   http://localhost:8222"
        echo "  • Jaeger UI:         http://localhost:16686"
        echo "  • Prometheus:        http://localhost:9090"
        echo "  • Grafana:           http://localhost:3000"
        echo ""
        echo -e "${YELLOW}Comandos úteis:${NC}"
        echo "  • Ver logs:          $DOCKER_COMPOSE logs -f mcp-ultra"
        echo "  • Parar:             $DOCKER_COMPOSE down"
        echo "  • Restart:           $DOCKER_COMPOSE restart mcp-ultra"
        echo "  • Status:            $DOCKER_COMPOSE ps"
        echo ""
        ;;
esac

