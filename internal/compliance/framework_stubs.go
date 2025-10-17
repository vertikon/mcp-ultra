//go:build testpatch || !production

package compliance

import (
	"context"
)

// Patch v11.1: Stubs mínimos para compatibilizar testes herdados.
// Substitua por implementações reais conforme a regra de negócio evoluir.

type DataAccessRequest struct {
	SubjectID string
	Fields    []string
}

type DataDeletionRequest struct {
	SubjectID string
	Reason    string
}

type AuditEvent struct {
	Actor   string
	Action  string
	Target  string
	Details map[string]any
}

func (f *ComplianceFramework) ProcessDataAccessRequest(ctx context.Context, req DataAccessRequest) error {
	// TODO: Implementar política real de acesso aos dados
	return nil
}

func (f *ComplianceFramework) AnonymizeData(ctx context.Context, subjectID string) error {
	// TODO: Implementar anonimização real
	return nil
}

func (f *ComplianceFramework) LogAuditEvent(ctx context.Context, evt AuditEvent) error {
	// TODO: Persistir em trilha de auditoria (ex.: compliance/audit.log)
	return nil
}