# NATS Subjects - MCP Ultra

## Convenções
- Prefixo global: `ultra.*`
- Versão no payload (schema), não no subject
- Sempre incluir `msg_id` e `timestamp` no payload

## Subjects Base

### Health & Registration
- `ultra.health.ping` - Health check ping
- `ultra.health.pong` - Health check response
- `ultra.sdk.register` - SDK registration
- `ultra.error` - Generic error events

### AI (Templates AI)
- `ultra.ai.router.decision` - Router decision events
- `ultra.ai.policy.block` - Policy block events
- `ultra.ai.inference.error` - Inference error events
- `ultra.ai.inference.summary` - Inference summary events

## Boas Práticas
1. Validar payloads com schemas antes de publicar
2. Evitar reply subjects não monitorados
3. Usar error handlers para todos os subjects
4. Documentar novos subjects neste arquivo
