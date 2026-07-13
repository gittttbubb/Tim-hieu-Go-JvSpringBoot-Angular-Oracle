import { HttpErrorResponse, HttpInterceptorFn, } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';
import { MessageService } from 'primeng/api';
import { AuthStore } from '../store/auth.store';
import { TranslationService } from '../services/translation.service';

export const errorInterceptor: HttpInterceptorFn = (req, next) => {
    const router = inject(Router);
    const authStore = inject(AuthStore);
    const messageService = inject(MessageService);
    const translationService = inject(TranslationService);
    return next(req).pipe(catchError((error: HttpErrorResponse) => {
        const rawMessage = error.error?.message;
        const message = rawMessage
            ? translationService.translate(rawMessage)
            : translationService.translate('errors.unexpected');
        switch (error.status) {
            case 401:
                if (req.url.includes('/auth/login')) {
                        return throwError(() => error);
                    }
                if (!req.url.includes('/auth/login')) {
                    authStore.clear();
                    messageService.add({
                        severity: 'warn',
                        summary: translationService.translate('common.warning'),
                        detail: translationService.translate('errors.sessionExpired'),
                    });
                    router.navigate(['/login']);
                }
                break;
            case 403:
                messageService.add({
                    severity: 'warn',
                    summary: translationService.translate('common.warning'),
                    detail: translationService.translate('errors.accessDenied')
                });
                router.navigate(['/dashboard']);
                break;
            case 500:
                messageService.add({
                    severity: 'error',
                    summary: translationService.translate('common.error'),
                    detail: translationService.translate('errors.serverError'),
                });
                break;
            case 0:
                messageService.add({
                    severity: 'error',
                    summary: translationService.translate('common.error'),
                    detail: translationService.translate('errors.networkError'),
                });
                break;
            default:
                messageService.add({
                    severity: 'error',
                    summary: translationService.translate('common.error'),
                    detail: message,
                });
        }
        return throwError(() => error);
    }));
};