{
  "ProjectName": "mcp-ultra",
  "Timestamp": "2025-10-19 14:06:44",
  "Critical": [],
  "Medium": [],
  "Low": [
    {
      "Type": "Linter limpo",
      "Severity": "baixo",
      "Location": "múltiplos arquivos",
      "Description": "Linter encontrou problemas",
      "Suggestion": "Corrija os problemas manualmente (NÃO use --fix)",
      "Fixability": {
        "Safe": false,
        "RollbackEasy": false,
        "AffectsBehavior": true,
        "RequiresReview": true,
        "AutoFixCommand": "",
        "ManualSteps": "1. Analise cada issue do linter\n2. Corrija manualmente, entendendo o contexto\n3. NÃO use golangci-lint run --fix (pode quebrar código)\n4. Execute testes após cada correção",
        "NonFixableReason": "BUSINESS_LOGIC"
      },
      "Examples": [
        "pkg\\httpx\\httpx.go:6:2: import 'github.com/go-chi/chi/v5' is not allowed from list 'main': Use pkg/httpx facade instead of direct chi import (depguard)",
        "\t\"github.com/go-chi/chi/v5\"",
        "\t^",
        "pkg\\httpx\\httpx.go:7:2: import 'github.com/go-chi/chi/v5/middleware' is not allowed from list 'main': Use pkg/httpx facade instead of direct chi import (depguard)",
        "\t\"github.com/go-chi/chi/v5/middleware\"",
        "\t^",
        "pkg\\httpx\\httpx.go:8:2: import 'github.com/go-chi/cors' is not allowed from list 'main': Use pkg/httpx.CORS facade (depguard)",
        "\t\"github.com/go-chi/cors\"",
        "\t^",
        "internal\\handlers\\http\\health.go:405:35: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)"
      ],
      "NonFixableReason": "BUSINESS_LOGIC"
    }
  ],
  "TotalGAPs": 1,
  "Score": 95,
  "AutoFixable": 0,
  "Manual": 1
}