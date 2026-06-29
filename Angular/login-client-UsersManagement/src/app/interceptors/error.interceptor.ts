import { HttpErrorResponse, HttpInterceptorFn, } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';
import { MessageService } from 'primeng/api';
import { AuthStore } from '../store/auth.store';

export const errorInterceptor: HttpInterceptorFn = (req, next) => {
    const router = inject(Router);
    const authStore = inject(AuthStore);
    const messageService = inject(MessageService);
    return next(req).pipe(catchError((error: HttpErrorResponse) => {
        const message = error.error?.message ?? 'Unexpected error';
        switch (error.status) {
            case 401:
                if (req.url.includes('/auth/login')) {
                        return throwError(() => error);
                    }
                if (!req.url.includes('/auth/login')) {
                    authStore.clear();
                    messageService.add({
                        severity: 'warn',
                        summary: 'Session Expired',
                        detail: 'Phiên đăng nhập đã hết hạn',
                    });
                    router.navigate(['/login']);
                }
                break;
            case 403:
                messageService.add({
                    severity: 'warn',
                    summary: 'Access Denied',
                    detail: 'Bạn không có quyền truy cập chức năng này'
                });
                router.navigate(['/dashboard']);
                break;
            case 500:
                messageService.add({
                    severity: 'error',
                    summary: 'Server Error',
                    detail: 'Đã xảy ra lỗi hệ thống',
                });
                break;
            case 0:
                messageService.add({
                    severity: 'error',
                    summary: 'Network Error',
                    detail: 'Không thể kết nối tới máy chủ',
                });
                break;
            default:
                messageService.add({
                    severity: 'error',
                    summary: 'Error',
                    detail: message,
                });
        }
        return throwError(() => error);
    }));
};