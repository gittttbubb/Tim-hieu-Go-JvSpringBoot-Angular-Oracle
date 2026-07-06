import { Component, EventEmitter, inject, Input, Output } from '@angular/core';
import { AuthStore } from '../../store/auth.store';
import { Router, RouterLink } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { ThemeService } from '../../services/theme.service';

@Component({
  selector: 'app-top-bar',
  imports: [ButtonModule, RouterLink],
  templateUrl: './top-bar.component.html',
  styleUrl: './top-bar.component.scss'
})
export class TopBarComponent {
  @Input() collapsed = false;

  @Output() toggleSidebar = new EventEmitter<void>();
  private readonly authStore = inject(AuthStore);
  private readonly router = inject(Router);
  public readonly themeService = inject(ThemeService);

  get username(): string {
    return (this.authStore.getUser()?.username ?? '');
  }

  logout(): void {
    this.authStore.clear();
    this.router.navigate(['/login']);
  }
}
