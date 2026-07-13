import { Component, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormBuilder, ReactiveFormsModule, Validators, } from '@angular/forms';
import { finalize } from 'rxjs';
import { MessageService } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { PasswordModule } from 'primeng/password';
import { AuthService } from '../../../services/auth.service';
import { AuthStore } from '../../../store/auth.store';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule, ButtonModule, CardModule, InputTextModule, PasswordModule, RouterLink, TranslatePipe],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {

  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly authStore = inject(AuthStore);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly translationService = inject(TranslationService);

  loading = false;
  form = this.fb.nonNullable.group({
    username: ['', Validators.required,],
    password: ['', Validators.required],
  });

  submit(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    this.loading = true;
    this.authService.login(this.form.getRawValue())
      .pipe(finalize(() => { this.loading = false; }))
      .subscribe({
        next: response => {
          const loginData = response.data;
          this.authStore.setAuth(loginData);
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('auth.login.success'),
          });
          if (loginData.mustChangePassword) {
            this.router.navigate(['/change-password']);
            this.messageService.add({
              severity: 'warning',
              summary: this.translationService.translate('auth.changePassword.title'),
              detail: this.translationService.translate('auth.login.pendingPasswordChange'),
            });
            return;
          }
          this.router.navigate(['/dashboard']);
        },
        error: error => {
          const apiMessage = error.error?.message;
          const translatedDetail = apiMessage
            ? this.translationService.translate(apiMessage)
            : this.translationService.translate('auth.login.failed');
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: translatedDetail,
          });
        },
      });
  }
}