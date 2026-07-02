import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.development';
import { ChangePasswordRequest, LoginRequest, LoginResponse, TempPasswordResponse } from '../models/auth.model';
import { ApiResponse, ApiErrorResponse } from '../models/api-response.model';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    private readonly http = inject(HttpClient);

    login(request: LoginRequest): Observable<ApiResponse<LoginResponse>> {
        return this.http.post<ApiResponse<LoginResponse>>(`${environment.apiUrl}/auth/login`, request);
    }

    resetPassword(userId: string) {
        return this.http.post<ApiResponse<TempPasswordResponse>>(`${environment.apiUrl}/auth/reset-password/${userId}`, {});
    }
    changePassword(payload: ChangePasswordRequest) {
        return this.http.post<ApiResponse<any>>(`${environment.apiUrl}/auth/change-password`, payload);
    }
}
