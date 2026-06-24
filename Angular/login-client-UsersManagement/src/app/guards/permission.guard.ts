import { inject } from '@angular/core';
import {ActivatedRouteSnapshot, CanActivateFn, Router} from '@angular/router';
import { AuthStore } from '../store/auth.store';

function hasRequiredPermission(route: ActivatedRouteSnapshot, authStore: AuthStore): boolean {
    const permission = route.data?.['permission'] as string | undefined;
    const permissions = route.data?.['permissions'] as string[] | undefined;
    const requireAll = route.data?.['requireAll'] as boolean | undefined;
    // Không cấu hình permission
    if (!permission && !permissions) {
        return true;
    }
    // Một permission
    if (permission) {
        return authStore.hasPermission(permission);
    }
    // Nhiều permission
    if (permissions?.length) {
        // AND
        if (requireAll) {
            return authStore.hasAllPermissions(permissions);
        }
        // OR (default)
        return authStore.hasAnyPermission(permissions);
    }
    return true;
}

export const permissionGuard: CanActivateFn = (route) => {
    const authStore = inject(AuthStore);
    const router = inject(Router);
    const allowed = hasRequiredPermission(route, authStore);
    if (allowed) {
        return true;
    }
    return router.createUrlTree(['/forbidden']);
};