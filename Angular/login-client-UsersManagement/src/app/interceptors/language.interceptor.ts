import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { TranslationService } from '../services/translation.service';

export const languageInterceptor: HttpInterceptorFn = (req, next) => {
  const translationService = inject(TranslationService);
  const currentLang = translationService.currentLang();

  const langReq = req.clone({
    setHeaders: {
      'Accept-Language': currentLang,
      'X-Lang': currentLang
    }
  });

  return next(langReq);
};
