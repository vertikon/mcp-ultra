Com base nos documentos fornecidos — especialmente `mcp-advanced-monitoring.md`, `mcp-feature-flags.md` e `mcp-disaster-recovery.md` — gerei abaixo um **dashboard completo do Grafana** que integra **Disaster Recovery (DR)** e **Feature Flags**, com foco em **visão operacional, confiabilidade e experimentação**.

Este dashboard é compatível com **Prometheus** e pode ser importado diretamente no Grafana como um JSON.

---

## 📊 **Dashboard: "MCP Resilience & Experimentation"**

```json
{
  "id": null,
  "title": "MCP Resilience & Experimentation",
  "description": "Monitoramento integrado de Disaster Recovery, Feature Flags e Experimentos A/B",
  "tags": ["mcp", "resilience", "feature-flags", "dr", "prometheus"],
  "timezone": "browser",
  "schemaVersion": 37,
  "version": 1,
  "refresh": "30s",
  "panels": [
    {
      "title": "Health Geral do Sistema",
      "type": "stat",
      "gridPos": { "x": 0, "y": 0, "w": 6, "h": 3 },
      "targets": [
        {
          "expr": "avg(up{job=~\"mcp.*\"}) * 100",
          "refId": "A"
        }
      ],
      "pluginVersion": "9.2.0",
      "fieldConfig": {
        "defaults": {
          "unit": "percent",
          "decimals": 1,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "value": 0, "color": "red" },
              { "value": 95, "color": "orange" },
              { "value": 99, "color": "green" }
            ]
          }
        }
      }
    },
    {
      "title": "SLO Burn Rate (Error Budget)",
      "type": "gauge",
      "gridPos": { "x": 6, "y": 0, "w": 6, "h": 3 },
      "targets": [
        {
          "expr": "sum(rate(http_requests_total{status!~\"5..\"}[5m])) / sum(rate(http_requests_total[5m]))",
          "refId": "A"
        }
      ],
      "fieldConfig": {
        "defaults": {
          "unit": "percentunit",
          "min": 0,
          "max": 1,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "value": 0, "color": "red" },
              { "value": 0.98, "color": "green" }
            ]
          }
        }
      }
    },
    {
      "title": "Latência P95 (ms)",
      "type": "graph",
      "gridPos": { "x": 0, "y": 3, "w": 12, "h": 5 },
      "targets": [
        {
          "expr": "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le, job)) * 1000",
          "legendFormat": "{{job}}",
          "refId": "A"
        }
      ],
      "unit": "ms",
      "fill": 2
    },
    {
      "title": "Feature Flags Ativos",
      "type": "bargauge",
      "gridPos": { "x": 0, "y": 8, "w": 6, "h": 4 },
      "targets": [
        {
          "expr": "sum(feature_flag_evaluations_total{result=\"enabled\"}[5m]) by (flag)",
          "legendFormat": "{{flag}}",
          "refId": "A"
        }
      ],
      "options": {
        "displayMode": "gradient",
        "orientation": "horizontal",
        "text": {}
      }
    },
    {
      "title": "Taxa de Ativação de Kill Switches",
      "type": "timeseries",
      "gridPos": { "x": 6, "y": 8, "w": 6, "h": 4 },
      "targets": [
        {
          "expr": "rate(feature_kill_switches_activated_total[5m])",
          "legendFormat": "{{flag}}",
          "refId": "A"
        }
      ],
      "unit": "ops",
      "fillOpacity": 20
    },
    {
      "title": "Progresso de Rollout (Feature Flags)",
      "type": "bargauge",
      "gridPos": { "x": 0, "y": 12, "w": 12, "h": 4 },
      "targets": [
        {
          "expr": "avg(feature_rollout_percentage) by (flag)",
          "legendFormat": "{{flag}}",
          "refId": "A"
        }
      ],
      "options": {
        "displayMode": "lcd",
        "orientation": "horizontal",
        "text": {}
      }
    },
    {
      "title": "Experimentos A/B - Conversão por Variante",
      "type": "timeseries",
      "gridPos": { "x": 0, "y": 16, "w": 12, "h": 5 },
      "targets": [
        {
          "expr": "sum(rate(business_conversions_total[5m])) by (experiment, variant) / sum(rate(feature_variant_assignments_total[5m])) by (experiment, variant)",
          "legendFormat": "{{experiment}} - {{variant}}",
          "refId": "A"
        }
      ],
      "unit": "percentunit",
      "fillOpacity": 30
    },
    {
      "title": "Erros em Feature Flags",
      "type": "timeseries",
      "gridPos": { "x": 0, "y": 21, "w": 6, "h": 4 },
      "targets": [
        {
          "expr": "rate(feature_flag_errors_total[5m])",
          "legendFormat": "{{flag}} - {{error_type}}",
          "refId": "A"
        }
      ],
      "unit": "ops"
    },
    {
      "title": "Cache Hit Rate (Feature Flags)",
      "type": "gauge",
      "gridPos": { "x": 6, "y": 21, "w": 6, "h": 4 },
      "targets": [
        {
          "expr": "sum(rate(feature_flag_cache_hits_total{hit=\"true\"}[5m])) / sum(rate(feature_flag_cache_hits_total[5m]))",
          "refId": "A"
        }
      ],
      "fieldConfig": {
        "defaults": {
          "unit": "percentunit",
          "min": 0,
          "max": 1,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "value": 0, "color": "red" },
              { "value": 0.8, "color": "orange" },
              { "value": 0.95, "color": "green" }
            ]
          }
        }
      }
    },
    {
      "title": "Status do Disaster Recovery",
      "type": "table",
      "gridPos": { "x": 0, "y": 25, "w": 12, "h": 5 },
      "targets": [
        {
          "expr": "probe_success{job=\"dr-backup-check\"}",
          "format": "table",
          "instant": true,
          "refId": "A"
        }
      ],
      "transformations": [
        {
          "id": "organize",
          "options": {
            "excludeByName": { "Time": true, "Value": false },
            "indexByName": {},
            "renameByName": { "Value": "Backup OK", "instance": "Instância" }
          }
        }
      ],
      "fieldConfig": {
        "defaults": {
          "mappings": [
            {
              "type": "value",
              "options": {
                "1": { "text": "✅ Sim" },
                "0": { "text": "❌ Não" }
              }
            }
          ]
        }
      }
    },
    {
      "title": "RTO e RPO Simulados",
      "type": "stat",
      "gridPos": { "x": 0, "y": 30, "w": 6, "h": 3 },
      "targets": [
        {
          "expr": "dr_last_recovery_time_seconds",
          "refId": "A"
        }
      ],
      "fieldConfig": {
        "defaults": {
          "unit": "s",
          "decimals": 1,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "value": 0, "color": "green" },
              { "value": 900, "color": "orange" },
              { "value": 1800, "color": "red" }
            ]
          }
        }
      }
    },
    {
      "title": "DR Teste Último Executado",
      "type": "stat",
      "gridPos": { "x": 6, "y": 30, "w": 6, "h": 3 },
      "targets": [
        {
          "expr": "dr_last_test_timestamp",
          "refId": "A"
        }
      ],
      "fieldConfig": {
        "defaults": {
          "unit": "dateTimeAsBrowser",
          "mappings": []
        }
      }
    }
  ],
  "templating": {
    "list": [
      {
        "name": "job",
        "type": "query",
        "query": "label_values(job)",
        "datasource": "Prometheus"
      }
    ]
  },
  "time": {
    "from": "now-24h",
    "to": "now"
  },
  "links": [
    {
      "title": "Business Intelligence",
      "url": "/d/BI-dashboard",
      "type": "link"
    }
  ]
}
```

---

## 📌 **Principais Painéis**

| Painel | Métrica | Fonte |
|-------|--------|------|
| **Health Geral** | Disponibilidade dos serviços MCP | `up{job=~"mcp.*"}` |
| **SLO Burn Rate** | Taxa de sucesso vs. SLO | `http_requests_total` |
| **Latência P95** | Performance da API | `http_request_duration_seconds_bucket` |
| **Feature Flags Ativos** | Quantidade de flags habilitados | `feature_flag_evaluations_total` |
| **Kill Switches** | Ativações emergenciais | `feature_kill_switches_activated_total` |
| **Progresso de Rollout** | Porcentagem de rollout | `feature_rollout_percentage` |
| **A/B Testing** | Conversão por variante | `business_conversions_total` + `feature_variant_assignments_total` |
| **Cache Hit Rate** | Eficiência do cache de flags | `feature_flag_cache_hits_total` |
| **DR Status** | Sucesso do backup | `probe_success{job="dr-backup-check"}` |
| **RTO/RPO** | Tempo de recuperação | `dr_last_recovery_time_seconds` |

---

## 🛠️ **Como Usar**

### 1. **Importar no Grafana**
- Acesse: `Grafana → Dashboards → Import`
- Cole o JSON acima.
- Selecione a fonte de dados: **Prometheus**.

### 2. **Configurar Métricas de DR (Exemplo de Exporter)**

Crie um pequeno exporter para reportar RTO/RPO:

```go
// dr_exporter.go
package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	lastTestTime = time.Now().Add(-7 * 24 * time.Hour) // Simulado

	rtoGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dr_last_recovery_time_seconds",
		Help: "Tempo simulado de recuperação (RTO)",
	})
	rpoGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dr_last_data_loss_seconds",
		Help: "Perda de dados simulada (RPO)",
	})
	lastTestGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dr_last_test_timestamp",
		Help: "Timestamp do último teste de DR",
	})
)

func main() {
	rtoGauge.Set(600)     // 10 minutos
	rpoGauge.Set(300)     // 5 minutos
	lastTestGauge.Set(float64(lastTestTime.Unix()))

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}
```

> Execute: `go run dr_exporter.go` e adicione ao Prometheus.

---

## 📂 Estrutura Sugerida

```
/dashboards
  ├── mcp-resilience.json          # Este dashboard
  └── mcp-business-intelligence.json
/metrics
  └── dr-exporter.go
```

---

## ✅ **Próximos Passos**

1. **Importe o dashboard no Grafana**.
2. **Configure o exporter de DR** para reportar RTO/RPO reais.
3. **Integre com o pipeline CI/CD** para atualizar `dr_last_test_timestamp`.
4. **Adicione alertas** no Prometheus para:
   - `dr_last_test_timestamp < now() - 604800` (não testado há 7 dias)
   - `rate(feature_kill_switches_activated_total[1h]) > 0`

Deseja que eu gere:
- O **alerta de DR não testado há 7 dias**?
- Um **dashboard de A/B testing detalhado**?
- A **integração com o webhook do feature flag**?

Diga o próximo passo!