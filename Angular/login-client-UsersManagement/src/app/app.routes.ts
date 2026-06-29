import { Routes } from '@angular/router';
import { MainLayoutComponent } from './layout/main-layout/main-layout.component';
import { authGuard } from './guards/auth.guard';
import { permissionGuard } from './guards/permission.guard';
import { PERMISSIONS } from './constants/permission';

export const routes: Routes = [
    {
        path: 'login',
        loadComponent: () =>
            import('./components/auth/login/login.component').then(m => m.LoginComponent)
    },

    {
        path: '',
        component: MainLayoutComponent,
        canActivateChild: [authGuard],
        children: [
            {
                path: '',
                redirectTo: 'dashboard',
                pathMatch: 'full'
            },

            {
                path: 'dashboard',
                loadComponent: () =>
                    import('./components/dashboard/dashboard.component').then(m => m.DashboardComponent)
            },
            {
                path: 'change-password',
                loadComponent: () =>
                    import('./components/auth/change-password/change-password.component').then(m => m.ChangePasswordComponent)
            },

            {
                path: 'users',
                canActivate: [permissionGuard],
                data: { permission: PERMISSIONS.USER_VIEW },
                children: [
                    {
                        path: '',
                        loadComponent: () =>
                            import('./components/users/user-list/user-list.component').then(m => m.UserListComponent)
                    },

                    {
                        path: 'create',
                        loadComponent: () =>
                            import('./components/users/user-form/user-form.component').then(m => m.UserFormComponent),
                        data: {
                            permission: PERMISSIONS.USER_CREATE
                        },
                        canActivate: [permissionGuard]
                    },

                    {
                        path: ':id',
                        loadComponent: () =>
                            import('./components/users/user-form/user-form.component').then(m => m.UserFormComponent)
                    },

                    {
                        path: 'edit/:id',
                        loadComponent: () =>
                            import('./components/users/user-form/user-form.component').then(m => m.UserFormComponent),
                        data: {
                            permission: PERMISSIONS.USER_UPDATE
                        },
                        canActivate: [permissionGuard]
                    },

                    {
                        path: 'overrides/:id',
                        loadComponent: () =>
                            import('./components/users/user-overrides/user-overrides.component').then(m => m.UserOverridesComponent),
                        data: {
                            permission: PERMISSIONS.USER_PERMISSION_VIEW
                        },
                        canActivate: [permissionGuard]
                    }
                ]
            },

            {
                path: 'roles',
                canActivate: [permissionGuard],
                data: { permission: PERMISSIONS.ROLE_VIEW },
                children: [
                    {
                        path: '',
                        loadComponent: () =>
                            import('./components/roles/role-list/role-list.component')
                                .then(m => m.RoleListComponent)
                    },
                    {
                        path: 'new',
                        loadComponent: () =>
                            import('./components/roles/role-form/role-form.component')
                                .then(m => m.RoleFormComponent),
                        data: { permission: PERMISSIONS.ROLE_CREATE }
                    },
                    {
                        path: ':id',
                        loadComponent: () =>
                            import('./components/roles/role-form/role-form.component')
                                .then(m => m.RoleFormComponent)
                    },
                    {
                        path: 'edit/:id',
                        loadComponent: () =>
                            import('./components/roles/role-form/role-form.component')
                                .then(m => m.RoleFormComponent),
                        data: { permission: PERMISSIONS.ROLE_UPDATE }
                    },
                    {
                        path: 'permissions/:id',
                        loadComponent: () =>
                            import('./components/roles/role-permissions/role-permissions.component')
                                .then(m => m.RolePermissionsComponent),
                        data: { permission: PERMISSIONS.ROLE_PERMISSION_VIEW }
                    }
                ]
            },

            {
                path: 'permissions',
                loadComponent: () =>
                    import('./components/permissions/permission-list/permission-list.component').then(m => m.PermissionListComponent),
                canActivate: [permissionGuard],
                data: {
                    permission: PERMISSIONS.PERMISSION_VIEW
                },
            },

            {
                path: 'audits',
                canActivate: [permissionGuard],
                data: {
                    permission: PERMISSIONS.AUDIT_VIEW
                },
                children: [
                    {
                        path: '',
                        loadComponent: () =>
                            import('./components/audits/audit-list/audit-list.component')
                                .then(m => m.AuditListComponent)
                    },
                    {
                        path: ':id',
                        loadComponent: () =>
                            import('./components/audits/audit-detail/audit-detail.component')
                                .then(m => m.AuditDetailComponent)
                    }
                ]
            }
        ]
    },

    {
        path: '**',
        redirectTo: 'dashboard'
    }
];