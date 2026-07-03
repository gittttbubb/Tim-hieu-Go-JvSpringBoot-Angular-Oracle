import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../environments/environment.development';

import { ApiResponse, PaginationResponse } from '../models/api-response.model';

import { Role, RolePermission, RoleRequest } from '../models/role.model';


@Injectable({
    providedIn: 'root'
})
export class RoleService {

    private readonly http = inject(HttpClient);
    private readonly api = `${environment.apiUrl}/roles`;

    getRoles(
        params?: Record<string, any>
    ): Observable<ApiResponse<PaginationResponse<Role>>> {
        return this.http.get<ApiResponse<PaginationResponse<Role>>>(
            this.api,
            { params }
        );
    }
    getAllRoles(): Observable<ApiResponse<Role[]>> {
        return this.http.get<ApiResponse<Role[]>>(`${this.api}/all`);
    }

    getRoleById(id: string): Observable<ApiResponse<Role>> {
        return this.http.get<ApiResponse<Role>>(`${this.api}/${id}`);
    }

    createRole(payload: RoleRequest): Observable<ApiResponse<any>> {
        return this.http.post<ApiResponse<any>>(`${this.api}/create`, payload);
    }

    updateRole(id: string, payload: RoleRequest): Observable<ApiResponse<any>> {
        return this.http.put<ApiResponse<any>>(`${this.api}/${id}`, payload);
    }

    deleteRole(id: string): Observable<ApiResponse<any>> {
        return this.http.delete<ApiResponse<any>>(`${this.api}/${id}`);
    }

    getRolePermissions(roleId: string): Observable<ApiResponse<RolePermission[]>> {
        return this.http.get<ApiResponse<RolePermission[]>>(`${this.api}/permissions/${roleId}`);
    }

    assignPermission(payload: RolePermission): Observable<ApiResponse<any>> {
        return this.http.post<ApiResponse<any>>(`${this.api}/permissions/assign`, payload);
    }

    removePermission(roleId: string, permissionId: string): Observable<ApiResponse<any>> {
        return this.http.delete<ApiResponse<any>>(`${this.api}/${roleId}/permissions/${permissionId}`);
    }
}