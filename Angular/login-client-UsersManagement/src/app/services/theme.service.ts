import { Injectable, signal, effect } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  private readonly THEME_KEY = 'app-theme';
  theme = signal<'light' | 'dark'>('light');

  constructor() {
    // Determine initial theme
    const savedTheme = localStorage.getItem(this.THEME_KEY);
    if (savedTheme === 'light' || savedTheme === 'dark') {
      this.theme.set(savedTheme);
    } else {
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
      this.theme.set(prefersDark ? 'dark' : 'light');
    }

    // Reactively update HTML class and localStorage when theme changes
    effect(() => {
      const currentTheme = this.theme();
      const element = document.documentElement;
      
      if (currentTheme === 'dark') {
        element.classList.add('app-dark');
      } else {
        element.classList.remove('app-dark');
      }
      
      localStorage.setItem(this.THEME_KEY, currentTheme);
    });
  }

  toggleTheme(): void {
    this.theme.update(current => current === 'light' ? 'dark' : 'light');
  }

  isDark(): boolean {
    return this.theme() === 'dark';
  }
}
