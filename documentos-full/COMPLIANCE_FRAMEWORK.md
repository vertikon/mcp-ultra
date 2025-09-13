# Compliance Framework - LGPD/GDPR

## 📋 Visão Geral

O MCP Ultra implementa um framework abrangente de compliance para atender às regulamentações de proteção de dados LGPD (Lei Geral de Proteção de Dados Pessoais) e GDPR (General Data Protection Regulation).

## 🛡️ Recursos Implementados

### ✅ Funcionalidades Principais
- **Detecção de PII** - Identificação automática de dados pessoais
- **Gerenciamento de Consentimento** - Controle granular de permissões
- **Retenção de Dados** - Políticas automatizadas de retenção
- **Direitos dos Titulares** - Acesso, portabilidade e exclusão
- **Anonimização** - Técnicas de pseudonimização e anonimização
- **Auditoria** - Log completo de operações com dados pessoais

## 🎯 Compliance Scope

### LGPD (Brasil)
- ✅ **Art. 7º** - Base legal para tratamento
- ✅ **Art. 8º** - Consentimento do titular  
- ✅ **Art. 18º** - Direitos do titular
- ✅ **Art. 37º** - Agente de tratamento (DPO)
- ✅ **Art. 46º** - Registro das operações

### GDPR (União Europeia)
- ✅ **Art. 6º** - Lawfulness of processing
- ✅ **Art. 7º** - Conditions for consent
- ✅ **Art. 15º** - Right of access
- ✅ **Art. 17º** - Right to erasure
- ✅ **Art. 20º** - Right to data portability
- ✅ **Art. 30º** - Records of processing activities

## ⚙️ Configuração

### Configuração Básica

```yaml
compliance:
  enabled: true
  default_region: "BR"
  
  pii_detection:
    enabled: true
    scan_fields: ["email", "phone", "cpf", "name", "address"]
    classification_api: "local"
    confidence: 0.8
    auto_mask: true
    
  consent:
    enabled: true
    default_purposes: ["processing", "analytics"]
    ttl: "2h"
    granular_level: "field"
    
  data_retention:
    enabled: true
    default_period: "365d"
    category_periods:
      personal: "730d"  # 2 anos
      session: "30d"    # 30 dias
      analytics: "1095d" # 3 anos
    auto_delete: true
    backup_retention: "2555d"  # 7 anos
    
  audit_logging:
    enabled: true
    detail_level: "full"
    retention_period: "1825d"  # 5 anos
    encryption_enabled: true
    external_logging: false
    
  lgpd:
    enabled: true
    dpo_contact: "dpo@company.com"
    legal_basis: "consent"
    data_categories: ["personal", "sensitive"]
    shared_third_party: ["analytics-provider"]
    
  gdpr:
    enabled: true
    dpo_contact: "dpo@company.com" 
    legal_basis: "consent"
    data_categories: ["personal", "sensitive"]
    cross_border_transfer: true
    adequacy_decisions: ["US", "CA"]
    
  anonymization:
    enabled: true
    methods: ["hashing", "generalization", "suppression"]
    hash_salt: "${ANONYMIZATION_SALT}"
    reversible: false
    k_anonymity: 5
    algorithms:
      email: "hash"
      phone: "mask"
      cpf: "hash"
      
  data_rights:
    enabled: true
    response_time: "720h"  # 30 dias
    auto_fulfillment: false
    verification_required: true
    notification_channels: ["email", "sms"]
```

## 🔍 Detecção de PII

### Uso Básico

```go
import "github.com/vertikon/mcp-ultra/internal/compliance"

func detectPII(data map[string]interface{}) {
    framework := compliance.NewComplianceFramework(config, logger)
    
    result, err := framework.ScanForPII(ctx, data)
    if err != nil {
        log.Error("PII scan failed", zap.Error(err))
        return
    }
    
    fmt.Printf("Detected PII fields: %v\n", result.DetectedFields)
    fmt.Printf("Confidence scores: %v\n", result.ConfidenceScores)
}
```

### Exemplo de Detecção

```go
data := map[string]interface{}{
    "name":     "João Silva",
    "email":    "joao@example.com", 
    "phone":    "+5511999999999",
    "cpf":      "123.456.789-00",
    "age":      30,
    "address":  "Rua das Flores, 123",
}

result, _ := framework.ScanForPII(ctx, data)
// result.DetectedFields = ["name", "email", "phone", "cpf", "address"]
// age não é considerado PII
```

### Campos Suportados

| Campo | Tipo | Padrão | Confidence |
|-------|------|--------|------------|
| `email` | PII | `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` | 0.95 |
| `phone` | PII | `^\+?[1-9]\d{1,14}$` | 0.90 |
| `cpf` | Sensitive | `^\d{3}\.\d{3}\.\d{3}-\d{2}$` | 0.95 |
| `cnpj` | PII | `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$` | 0.95 |
| `name` | PII | Pattern + NER | 0.80 |
| `address` | PII | Pattern + Context | 0.75 |

## 🤝 Gerenciamento de Consentimento

### Registro de Consentimento

```go
func recordConsent(userID uuid.UUID) error {
    purposes := []string{"processing", "analytics", "marketing"}
    
    return framework.RecordConsent(ctx, userID, purposes, "web")
}
```

### Verificação de Consentimento

```go
func checkConsent(userID uuid.UUID, purpose string) (bool, error) {
    hasConsent, err := framework.HasConsent(ctx, userID, purpose)
    if err != nil {
        return false, err
    }
    
    if !hasConsent {
        return false, errors.New("consent required for " + purpose)
    }
    
    return true, nil
}
```

### Retirada de Consentimento

```go
func withdrawConsent(userID uuid.UUID) error {
    purposes := []string{"marketing", "analytics"}
    
    return framework.WithdrawConsent(ctx, userID, purposes)
}
```

### Granularidade de Consentimento

```go
// Por finalidade
framework.RecordConsent(ctx, userID, []string{"processing"}, "web")
framework.RecordConsent(ctx, userID, []string{"marketing"}, "email") 

// Por campo (granular_level: "field")
fieldConsents := map[string][]string{
    "email":    {"marketing", "notifications"},
    "phone":    {"processing"},
    "address":  {"delivery"},
}
framework.RecordFieldConsent(ctx, userID, fieldConsents)
```

## ⏳ Retenção de Dados

### Políticas de Retenção

```go
// Definir política de retenção
policy := compliance.RetentionPolicy{
    Category:        "personal_data",
    RetentionPeriod: 2 * 365 * 24 * time.Hour, // 2 anos
    AutoDelete:      true,
    BackupRetention: 7 * 365 * 24 * time.Hour, // 7 anos
}

framework.SetRetentionPolicy(ctx, policy)
```

### Verificação de Expiração

```go
func checkDataExpiration(userID uuid.UUID) {
    categories := []string{"personal", "session", "analytics"}
    
    for _, category := range categories {
        shouldDelete, err := framework.ShouldDeleteData(ctx, userID, category)
        if err != nil {
            log.Error("Error checking expiration", zap.Error(err))
            continue
        }
        
        if shouldDelete {
            log.Info("Data expired", 
                zap.String("user_id", userID.String()),
                zap.String("category", category))
                
            // Executar exclusão automática
            framework.DeleteExpiredData(ctx, userID, category)
        }
    }
}
```

## 👤 Direitos dos Titulares

### Direito de Acesso (Art. 15 GDPR / Art. 18 LGPD)

```go
func processAccessRequest(userID uuid.UUID) (*compliance.DataAccessResult, error) {
    request := compliance.DataAccessRequest{
        UserID:    userID,
        RequestID: uuid.New(),
        Purpose:   "data_portability",
        Format:    "json",
        Scope:     "all", // ou campos específicos
    }
    
    result, err := framework.ProcessDataAccessRequest(ctx, request)
    if err != nil {
        return nil, err
    }
    
    // result contém todos os dados do usuário em formato estruturado
    return result, nil
}
```

### Direito ao Esquecimento (Art. 17 GDPR / Art. 18 LGPD)

```go
func processErasureRequest(userID uuid.UUID, reason string) error {
    request := compliance.DataDeletionRequest{
        UserID:    userID,
        RequestID: uuid.New(),
        Reason:    reason,
        Scope:     "all",
        VerifyIdentity: true,
    }
    
    result, err := framework.ProcessDataDeletionRequest(ctx, request)
    if err != nil {
        return err
    }
    
    log.Info("Data deletion completed",
        zap.String("user_id", userID.String()),
        zap.Int("deleted_records", result.DeletedRecords),
        zap.String("status", result.Status))
        
    return nil
}
```

### Direito à Portabilidade (Art. 20 GDPR)

```go
func exportUserData(userID uuid.UUID, format string) ([]byte, error) {
    request := compliance.DataPortabilityRequest{
        UserID: userID,
        Format: format, // "json", "csv", "xml"
        Scope:  "portable", // apenas dados portáveis
    }
    
    data, err := framework.ExportUserData(ctx, request)
    if err != nil {
        return nil, err
    }
    
    return data, nil
}
```

## 🎭 Anonimização

### Estratégias de Anonimização

```go
// Pseudonimização (reversível com chave)
anonymized, err := framework.PseudonymizeData(ctx, data, "user-key")

// Anonimização irreversível
anonymized, err := framework.AnonymizeData(ctx, data, "full_anonymization")

// K-anonimity
anonymized, err := framework.ApplyKAnonymity(ctx, dataset, 5)
```

### Algoritmos Suportados

| Campo | Algoritmo | Exemplo |
|-------|-----------|---------|
| **Email** | Hash SHA-256 | `joao@example.com` → `a1b2c3d4...` |
| **Phone** | Masking | `+5511999999999` → `+551199****999` |
| **CPF** | Hash + Salt | `123.456.789-00` → `x7y8z9a1...` |
| **Name** | Generalization | `João Silva` → `João S.` |
| **Address** | Suppression | `Rua das Flores, 123` → `Rua das Flores, ***` |

## 📋 Auditoria e Logs

### Log de Auditoria

```go
func logDataAccess(userID uuid.UUID, fields []string) {
    details := map[string]interface{}{
        "accessed_fields": fields,
        "access_reason":   "user_request",
        "ip_address":      getClientIP(),
        "user_agent":      getUserAgent(),
    }
    
    framework.LogAuditEvent(ctx, userID, "data_access", details)
}
```

### Consulta de Logs

```go
func getAuditTrail(userID uuid.UUID) ([]*compliance.AuditLog, error) {
    filter := compliance.AuditFilter{
        UserID: &userID,
        From:   time.Now().Add(-30 * 24 * time.Hour), // últimos 30 dias
        To:     time.Now(),
        Actions: []string{"data_access", "consent_change", "data_deletion"},
    }
    
    logs, err := framework.GetAuditLogs(ctx, filter)
    if err != nil {
        return nil, err
    }
    
    return logs, nil
}
```

### Formato do Log de Auditoria

```json
{
  "id": "audit-123456",
  "timestamp": "2025-09-12T18:50:58Z",
  "user_id": "user-789",
  "action": "data_access",
  "resource": "user_profile",
  "details": {
    "accessed_fields": ["name", "email"],
    "access_reason": "user_request",
    "ip_address": "192.168.1.100",
    "user_agent": "Mozilla/5.0..."
  },
  "legal_basis": "consent",
  "jurisdiction": "BR",
  "compliance_version": "1.0.0"
}
```

## 🔒 Validação de Compliance

### Verificação Automática

```go
func validateDataProcessing(userID uuid.UUID, action string) error {
    request := compliance.ComplianceValidationRequest{
        UserID:       userID,
        Action:       action,
        DataType:     "personal",
        Purpose:      "processing",
        LegalBasis:   "consent", 
        Jurisdiction: "BR",
    }
    
    result, err := framework.ValidateCompliance(ctx, request)
    if err != nil {
        return err
    }
    
    if !result.Compliant {
        return fmt.Errorf("compliance violation: %v", result.Violations)
    }
    
    return nil
}
```

### Resultado da Validação

```go
type ComplianceValidationResult struct {
    Compliant   bool                    `json:"compliant"`
    Violations  []ComplianceViolation   `json:"violations,omitempty"`
    Warnings    []ComplianceWarning     `json:"warnings,omitempty"`
    LegalBasis  string                  `json:"legal_basis"`
    Jurisdiction string                 `json:"jurisdiction"`
    Timestamp   time.Time              `json:"timestamp"`
}
```

## 📊 Dashboard de Compliance

### Métricas de Compliance

```go
func getComplianceMetrics() (*compliance.ComplianceMetrics, error) {
    metrics, err := framework.GetComplianceMetrics(ctx)
    if err != nil {
        return nil, err
    }
    
    fmt.Printf("Consent Rate: %.2f%%\n", metrics.ConsentRate)
    fmt.Printf("Data Requests: %d\n", metrics.DataRequests)
    fmt.Printf("Violations: %d\n", metrics.Violations)
    
    return metrics, nil
}
```

### Status do Sistema

```go
func getComplianceStatus() {
    status, err := framework.GetComplianceStatus(ctx)
    if err != nil {
        log.Error("Failed to get compliance status", zap.Error(err))
        return
    }
    
    fmt.Printf("Compliance Enabled: %v\n", status.Enabled)
    fmt.Printf("Region: %s\n", status.Region)
    fmt.Printf("Features: %v\n", status.EnabledFeatures)
    fmt.Printf("DPO Contact: %s\n", status.DPOContact)
}
```

## 🚨 Alertas e Notificações

### Configuração de Alertas

```go
// Alerta de violação de compliance
framework.SetComplianceAlert("consent_violation", func(violation *ComplianceViolation) {
    // Notificar DPO
    sendEmail(config.DPOContact, "Compliance Violation", violation.Description)
    
    // Log crítico
    log.Error("Compliance violation detected",
        zap.String("type", violation.Type),
        zap.String("description", violation.Description),
        zap.String("user_id", violation.UserID.String()))
})
```

### Notificação de Titular

```go
func notifyDataSubject(userID uuid.UUID, event string) {
    notification := compliance.DataSubjectNotification{
        UserID:  userID,
        Type:    event,
        Message: getLocalizedMessage(event),
        Channels: []string{"email", "app_notification"},
    }
    
    framework.SendDataSubjectNotification(ctx, notification)
}
```

## 🧪 Testes de Compliance

### Testes Unitários

```go
func TestCompliance_ConsentManagement(t *testing.T) {
    framework := createTestFramework(t)
    userID := uuid.New()
    
    // Registrar consentimento
    err := framework.RecordConsent(ctx, userID, []string{"processing"}, "web")
    assert.NoError(t, err)
    
    // Verificar consentimento
    hasConsent, err := framework.HasConsent(ctx, userID, "processing")
    assert.NoError(t, err)
    assert.True(t, hasConsent)
    
    // Retirar consentimento
    err = framework.WithdrawConsent(ctx, userID, []string{"processing"})
    assert.NoError(t, err)
    
    // Verificar retirada
    hasConsent, err = framework.HasConsent(ctx, userID, "processing")
    assert.NoError(t, err)
    assert.False(t, hasConsent)
}
```

### Testes de Integração

```bash
# Executar testes de compliance
go test ./internal/compliance -v

# Testes específicos
go test -run TestCompliance_PIIDetection
go test -run TestCompliance_DataRights
go test -run TestCompliance_Anonymization
```

## 📖 Exemplos de Uso

### Cenário Completo: Registro de Usuário

```go
func registerUser(userData map[string]interface{}) error {
    userID := uuid.New()
    
    // 1. Detectar PII
    piiResult, err := framework.ScanForPII(ctx, userData)
    if err != nil {
        return err
    }
    
    // 2. Solicitar consentimento para campos PII
    if len(piiResult.DetectedFields) > 0 {
        consentRequired := []string{"processing", "storage"}
        // Apresentar formulário de consentimento para o usuário
        // ...
        
        err = framework.RecordConsent(ctx, userID, consentRequired, "registration_form")
        if err != nil {
            return err
        }
    }
    
    // 3. Validar compliance antes de processar
    err = validateDataProcessing(userID, "create_user")
    if err != nil {
        return err
    }
    
    // 4. Registrar dados com política de retenção
    err = framework.RecordDataCreation(ctx, userID, "personal", userData)
    if err != nil {
        return err
    }
    
    // 5. Log de auditoria
    framework.LogAuditEvent(ctx, userID, "user_registration", map[string]interface{}{
        "pii_fields": piiResult.DetectedFields,
        "consents":   []string{"processing", "storage"},
    })
    
    return nil
}
```

---

**Documentado em**: 2025-09-12  
**Versão**: 1.0.0  
**Compliance**: LGPD + GDPR