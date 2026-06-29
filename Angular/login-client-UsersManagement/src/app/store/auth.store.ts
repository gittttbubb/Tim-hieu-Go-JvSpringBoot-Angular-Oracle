import { Injectable, computed, signal } from '@angular/core';
import { AuthUser, LoginResponse, UserPermission, } from '../models/auth.model';
import { STORAGE_KEY } from '../constants/storage-key';

@Injectable({
    providedIn: 'root',
})
export class AuthStore {
    private readonly tokenSignal = signal<string | null>(null);
    private readonly userSignal = signal<AuthUser | null>(null);
    readonly token = this.tokenSignal.asReadonly();
    readonly user = this.userSignal.asReadonly();
    readonly isLoggedIn = computed(() => !!this.tokenSignal());

    constructor() {
        this.loadFromStorage();
    }

    setAuth(response: LoginResponse): void {
        const user: AuthUser = {
            userId: response.userId,
            username: response.username,
            roleId: response.roleId,
            permissions: response.permissions,
            mustChangePassword: response.mustChangePassword,
        };
        localStorage.setItem(STORAGE_KEY.TOKEN, response.accessToken);
        localStorage.setItem(STORAGE_KEY.USER, JSON.stringify(user));
        this.tokenSignal.set(response.accessToken);
        this.userSignal.set(user);
    }

    loadFromStorage(): void {
        const token = localStorage.getItem(STORAGE_KEY.TOKEN);
        const user = localStorage.getItem(STORAGE_KEY.USER);
        if (!token || !user) {
            return;
        }
        try {
            this.tokenSignal.set(token);
            this.userSignal.set(JSON.parse(user));
        } catch {
            this.clear();
        }
    }

    clear(): void {
        localStorage.removeItem(STORAGE_KEY.TOKEN);
        localStorage.removeItem(STORAGE_KEY.USER);
        this.tokenSignal.set(null);
        this.userSignal.set(null);
    }

    getToken(): string | null {
        return this.tokenSignal();
    }

    getUser(): AuthUser | null {
        return this.userSignal();
    }

    hasPermission(featureCode: string): boolean {
        return (this.userSignal()?.permissions.some(permission => permission.featureCode === featureCode) ?? false);
    }

    hasAnyPermission(permissions: string[]): boolean {
        return permissions.some(
            permission => this.hasPermission(permission)
        );
    }

    hasAllPermissions(permissions: string[]): boolean {
        return permissions.every(
            permission => this.hasPermission(permission)
        );
    }

    getPermission(permissionId: string): UserPermission | undefined {
        return this.userSignal()?.permissions
            .find(permission => permission.permissionId === permissionId);
    }

    getDataScope(permissionId: string): 'OWN' | 'TEAM' | 'ALL' | null {
        const permission = this.getPermission(permissionId);
        return (permission?.dataScope ?? null);
    }
    //Nếu trong tokenSignal đang là một chuỗi tokenthì !tokenSignal() sẽ là false, và !!tokenSignal() sẽ thành true (Đã đăng nhập).
    //Nếu trong tokenSignal đang là null, thì !tokenSignal() sẽ là true, và !!tokenSignal() sẽ thành false (Chưa đăng nhập).
    // Hàm isAuthenticated() sẽ trả về true nếu hệ thống đã tìm thấy token lưu trong tokenSignal, ngược lại nó trả về false 
    isAuthenticated(): boolean {
        return !!this.tokenSignal();
    }

    mustChangePassword(): boolean {
        return (this.userSignal()?.mustChangePassword ?? false);
    }
}