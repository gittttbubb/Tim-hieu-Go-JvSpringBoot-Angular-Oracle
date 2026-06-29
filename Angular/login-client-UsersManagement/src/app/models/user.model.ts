export interface UserList {
    id: string;
    tenantId: string;
    fullName: string;
    username: string;
    email: string;
    phone: string;
    roleId: string;
    roleName: string;
    status: string;
    mustChangePassword: boolean;
}

export interface UserDetail {
    id: string;
    tenantId: string;
    fullName: string;
    username: string;
    email: string;
    phone: string;
    roleId: string;
    status: string;
    mustChangePassword: boolean;
    createdAt: string;
    updatedAt: string;
}

export interface CreateUserRequest {
    tenantId: string;
    fullName: string;
    username: string;
    email: string;
    phone: string;
    roleId: string;
}

export interface UpdateUserRequest {
    fullName: string;
    username: string;
    email: string;
    phone: string;
    roleId: string;
    status: string;
}

export interface UserOverride {
    id: string;
    userId: string;
    permissionId: string;
    granted: boolean;
    reason: string;
    createdBy: string;
    dataScope: DataScope;
}

export interface AssignUserOverrideRequest {
    userId: string;
    permissionId: string;
    granted: boolean;
    reason: string;
    dataScope: DataScope;
}

export type DataScope = 'OWN' | 'TEAM' | 'ALL';