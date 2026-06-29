export interface Role {
    id: string;
    name: string;
    displayName: string;
    description: string;
}

export interface RoleRequest {
    name: string;
    displayName: string;
    description: string;
}

export type DataScope = | 'OWN' | 'TEAM' | 'ALL';

export interface RolePermission {
    roleId: string;
    permissionId: string;
    granted: boolean;
    dataScope: DataScope;
}

export interface RolePermissionView {
    permissionId: string;
    featureGroup: string;
    featureCode: string;
    action: string;
    description: string;
    granted: boolean;
    dataScope: DataScope;
}