import { Component, EventEmitter, inject, Input, Output } from '@angular/core';
import { AuthStore } from '../../store/auth.store';
import { Router, RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { ThemeService } from '../../services/theme.service';
import { TranslationService, Lang } from '../../services/translation.service';
import { TranslatePipe } from '../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-top-bar',
  imports: [ButtonModule, RouterLink, TranslatePipe],
  templateUrl: './top-bar.component.html',
  styleUrl: './top-bar.component.scss'
})
export class TopBarComponent {
  @Input() collapsed = false;

  @Output() toggleSidebar = new EventEmitter<void>();
  private readonly authStore = inject(AuthStore);
  private readonly router = inject(Router);
  public readonly themeService = inject(ThemeService);
  public readonly translationService = inject(TranslationService);

  get username(): string {
    return (this.authStore.getUser()?.username ?? '');
  }

  toggleLanguage(): void {
    const nextLang: Lang = this.translationService.currentLang() === 'vi' ? 'en' : 'vi';
    this.translationService.setLanguage(nextLang);
  }

  logout(): void {
    this.authStore.clear();
    this.router.navigate(['/login']);
  }
}
