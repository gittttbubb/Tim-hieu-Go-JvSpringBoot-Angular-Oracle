import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { finalize } from 'rxjs';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { PasswordModule } from 'primeng/password';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { AuthService } from '../../../services/auth.service';
import { AuthStore } from '../../../store/auth.store';
import { strongPasswordValidator } from '../../../shared/validators/password.validator';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-change-password',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, CardModule, ButtonModule, PasswordModule, ToastModule, TranslatePipe],
  templateUrl: './change-password.component.html',
  styleUrl: './change-password.component.scss'
})
export class ChangePasswordComponent {
  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly authStore = inject(AuthStore);
  private readonly translationService = inject(TranslationService);

  loading = false;
  form = this.fb.nonNullable.group({
    oldPassword: ['', Validators.required],
    newPassword: ['',
      [
        Validators.required,
        Validators.minLength(8),
        Validators.maxLength(100),
        strongPasswordValidator()
      ]
    ],
    confirmPassword: ['', Validators.required]
  });

  passwordMismatch(): boolean {
    const { newPassword, confirmPassword } = this.form.getRawValue();
    return (
      newPassword.length > 0 &&
      confirmPassword.length > 0 &&
      newPassword !== confirmPassword
    );
  }

  submit(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    if (this.passwordMismatch()) {
      this.messageService.add({
        severity: 'error',
        summary: this.translationService.translate('common.error'),
        detail: this.translationService.translate('auth.changePassword.mismatch')
      });
      return;
    }
    if (this.sameAsOldPassword()) {
      this.messageService.add({
        severity: 'error',
        summary: this.translationService.translate('common.error'),
        detail: this.translationService.translate('auth.changePassword.sameAsOld')
      });
      return;
    }
    this.loading = true;
    this.authService.changePassword({
      oldPassword: this.form.controls.oldPassword.value,
      newPassword: this.form.controls.newPassword.value
    })
      .pipe(finalize(() => this.loading = false))
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('auth.changePassword.success')
          });
          this.form.reset();
          this.logout();
          // this.router.navigate(['/dashboard']);
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail:
              err.error?.message ??
              this.translationService.translate('auth.changePassword.failed')
          });
        }
      });
  }
  sameAsOldPassword(): boolean {
    const { oldPassword, newPassword } = this.form.getRawValue();

    return (
      oldPassword.length > 0 &&
      newPassword.length > 0 &&
      oldPassword === newPassword
    );
  }
  back(): void {
    this.router.navigate(['/']);
  }
  logout(): void {
    this.authStore.clear();
    this.router.navigate(['/login']);
  }
}