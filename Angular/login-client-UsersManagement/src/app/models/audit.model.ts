export interface AuditLog {
    id: string;
    tenantId: string;
    actor: string;
    actorId?: string;
    action: string;
    entityType: string;
    entityId: string;
    beforeData?: string;
    afterData?: string;
    ipAddress?: string;
    eventTimestamp: string;
    reason?: string;
    eventType?: string;
    actorIdentifier?: string;
    targetUserId?: string;
    metadata?: string;
}