import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.development';
import { ApiResponse } from '../models/api-response.model';
import { AuditLog } from '../models/audit.model';

@Injectable({
    providedIn: 'root'
})
export class AuditService {
    private readonly http = inject(HttpClient);
    private readonly api = `${environment.apiUrl}/audits`;
    
    getAudits(): Observable<ApiResponse<AuditLog[]>> {
        return this.http.get<ApiResponse<AuditLog[]>>(`${this.api}/all`);
    }

    getById(id: string): Observable<ApiResponse<AuditLog>> {
        return this.http.get<ApiResponse<AuditLog>>(
            `${this.api}/${id}`
        );
    }
}