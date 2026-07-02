import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../environments/environment.development';
import { ApiResponse } from '../models/api-response.model';
import { UserList, UserDetail, CreateUserRequest, UpdateUserRequest, AssignUserOverrideRequest, UserOverride } from '../models/user.model';
import { TempPasswordResponse } from '../models/auth.model';

@Injectable({
    providedIn: 'root'
})
export class UserService {

    private readonly http = inject(HttpClient);
    private readonly api = `${environment.apiUrl}/users`;

    getUsers(): Observable<ApiResponse<UserList[]>> {
        return this.http.get<ApiResponse<UserList[]>>(this.api);
    }

    getById(id: string): Observable<ApiResponse<UserDetail>> {
        return this.http.get<ApiResponse<UserDetail>>(`${this.api}/${id}`);
    }

    create(request: CreateUserRequest): Observable<ApiResponse<TempPasswordResponse>> {
        return this.http.post<ApiResponse<TempPasswordResponse>>(`${this.api}/create`, request);
    }

    update(id: string, request: UpdateUserRequest): Observable<ApiResponse<any>> {
        return this.http.put<ApiResponse<any>>(`${this.api}/${id}`, request);
    }

    delete(id: string): Observable<ApiResponse<any>> {
        return this.http.delete<ApiResponse<any>>(`${this.api}/${id}`);
    }
    lock(id: string): Observable<ApiResponse<any>> {
        return this.http.post<ApiResponse<any>>(`${this.api}/lock/${id}`, {});
    }

    unlock(id: string): Observable<ApiResponse<any>> {
        return this.http.post<ApiResponse<any>>(`${this.api}/unlock/${id}`, {});
    }

    // overrides
    getByUserId(userId: string) {
        return this.http.get<ApiResponse<UserOverride[]>>(
            `${this.api}/overrides/${userId}`
        );
    }

    assign(request: AssignUserOverrideRequest) {
        return this.http.post<ApiResponse<any>>(
            `${this.api}/overrides/assign`,
            request
        );
    }

    remove(id: string) {
        return this.http.delete<ApiResponse<any>>(
            `${this.api}/overrides/${id}`
        );
    }
}