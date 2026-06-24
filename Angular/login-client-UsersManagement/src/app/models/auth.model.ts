export interface UserPermission {
  permissionId: string;
  dataScope: 'OWN' | 'TEAM' | 'ALL';
}

export interface AuthUser {
  userId: string;
  username: string;
  roleId: string;
  permissions: UserPermission[];
  mustChangePassword: boolean;
}

export interface LoginResponse {
  accessToken: string;
  tokenType: string;
  expiresIn: number;
  userId: string;
  username: string;
  roleId: string;
  permissions: UserPermission[];
  mustChangePassword: boolean;
}