# NATS Subjects (Template)

> Este arquivo documenta os subjects de exemplo usados pelo template.
> Ajuste conforme sua topologia real.

## Convenção
- Prefixo por ambiente: `dev.`, `stg.`, `prd.`
- Tenant opcional: `tenant.<id>.`
- Recurso: `events|cmd|query`
- Serviço/dominio: `ultra`

## Exemplos
- `dev.events.ultra.user.created`
- `dev.events.ultra.user.updated`
- `dev.events.ultra.user.deleted`
- `stg.cmd.ultra.billing.invoice.generate`
- `stg.cmd.ultra.task.create`
- `prd.query.ultra.analytics.refresh`
- `prd.query.ultra.stats.get`

## Stream Configuration

Streams configurados no NATS JetStream:

- `EVENTS_ULTRA` - Eventos de domínio
- `COMMANDS_ULTRA` - Comandos assíncronos
- `QUERIES_ULTRA` - Queries materializadas

## Consumers

Consumers por serviço:

- `ultra-worker` - Processa comandos assíncronos
- `ultra-analytics` - Atualiza analytics em tempo real
- `ultra-audit` - Log de auditoria de eventos
