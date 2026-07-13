import { Injectable, signal, inject } from '@angular/core';
import { HttpClient, HttpBackend } from '@angular/common/http';
import { en, vi } from '../constants/translations';
import { environment } from '../../environments/environment.development';

export type Lang = 'en' | 'vi';

@Injectable({
  providedIn: 'root'
})
export class TranslationService {
  private readonly http: HttpClient;
  private readonly currentLangSignal = signal<Lang>(this.getInitialLang());

  readonly currentLang = this.currentLangSignal.asReadonly();

  private readonly backendTranslations = signal<Record<string, string>>({});

  private readonly localTranslations: Record<Lang, any> = {
    en,
    vi
  };

  constructor(handler: HttpBackend) {
    this.http = new HttpClient(handler);
    this.loadBackendTranslations(this.currentLangSignal());
  }

  private getInitialLang(): Lang {
    const saved = localStorage.getItem('lang');
    if (saved === 'en' || saved === 'vi') {
      return saved;
    }
    const browserLang = navigator.language.split('-')[0];
    return browserLang === 'en' ? 'en' : 'vi';
  }

  setLanguage(lang: Lang): void {
    localStorage.setItem('lang', lang);
    this.currentLangSignal.set(lang);
    this.loadBackendTranslations(lang);
  }

  private loadBackendTranslations(lang: Lang): void {
    const url = `${environment.apiUrl}/locales/${lang}`;
    this.http.get<Record<string, string>>(url).subscribe({
      next: (data) => {
        this.backendTranslations.set(data || {});
      },
      error: (err) => {
        console.error('Failed to load backend translations from API', err);
        this.backendTranslations.set({});
      }
    });
  }

  translate(key: string, params?: Record<string, string | number>): string {
    if (!key) return '';
    const backendDict = this.backendTranslations();
    if (key in backendDict) {
      let translation = backendDict[key];
      if (params) {
        Object.entries(params).forEach(([paramKey, paramValue]) => {
          translation = translation.replace(new RegExp(`{{${paramKey}}}`, 'g'), String(paramValue));
          translation = translation.replace(new RegExp(`{${paramKey}}`, 'g'), String(paramValue));
        });
      }
      return translation;
    }

    const lang = this.currentLangSignal();
    const localDict = this.localTranslations[lang];

    let value = localDict;
    const parts = key.split('.');
    for (const part of parts) {
      if (value && typeof value === 'object' && part in value) {
        value = value[part];
      } else {
        value = undefined;
        break;
      }
    }

    let translation = typeof value === 'string' ? value : key;

    if (params) {
      Object.entries(params).forEach(([paramKey, paramValue]) => {
        translation = translation.replace(new RegExp(`{{${paramKey}}}`, 'g'), String(paramValue));
        translation = translation.replace(new RegExp(`{${paramKey}}`, 'g'), String(paramValue));
      });
    }

    return translation;
  }
}
