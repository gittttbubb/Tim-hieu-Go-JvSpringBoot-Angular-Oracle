import { HttpClient } from "@angular/common/http";
import { ApiResponse } from "../models/api-response.model";
import { Permission } from "../models/permission.model";
import { Observable } from "rxjs";
import { environment } from "../../environments/environment.development";
import { inject, Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class PermissionService {

    private readonly http = inject(HttpClient);
    private readonly api = `${environment.apiUrl}/permissions`;

    getPermissions(): Observable<ApiResponse<Permission[]>> {
        return this.http.get<ApiResponse<Permission[]>>(this.api);
    }
}