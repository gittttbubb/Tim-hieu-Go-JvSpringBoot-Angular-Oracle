import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { finalize } from 'rxjs';
import { CardModule } from 'primeng/card';
import { PasswordModule } from 'primeng/password';
import { ButtonModule } from 'primeng/button';
import { MessageService } from 'primeng/api';
import { AuthService } from '../../../services/auth.service';
import { strongPasswordValidator } from '../../../shared/validators/password.validator';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-reset-password',
  imports: [CommonModule, ReactiveFormsModule, CardModule, PasswordModule, ButtonModule, TranslatePipe],
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.scss']
})
export class ResetPasswordComponent implements OnInit {
  private fb = inject(FormBuilder);
  private authService = inject(AuthService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private messageService = inject(MessageService);
  private translationService = inject(TranslationService);

  token = '';
  loading = false;
  form = this.fb.nonNullable.group({
    newPassword: [
      '',
      [
        Validators.required,
        Validators.minLength(8),
        Validators.maxLength(100),
        strongPasswordValidator()
      ]
    ],
    confirmPassword: [
      '',
      Validators.required
    ]
  });

  ngOnInit(): void {
    this.token = this.route.snapshot.queryParamMap.get('token') ?? '';
    if (!this.token) {
      this.messageService.add({
        severity: 'error',
        summary: this.translationService.translate('common.error'),
        detail: this.translationService.translate('auth.resetPassword.invalidLink')
      });
      this.router.navigate(['/login']);
    }
  }

  submit(): void {
    if (!this.token) {
      this.messageService.add({
        severity: 'error',
        summary: this.translationService.translate('common.error'),
        detail: this.translationService.translate('auth.resetPassword.invalidLink')
      });
      return;
    }
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    if (this.passwordMismatch()) {
      this.messageService.add({
        severity: 'error',
        summary: this.translationService.translate('common.error'),
        detail: this.translationService.translate('auth.resetPassword.mismatch')
      });
      return;
    }
    this.loading = true;
    this.authService.userResetPassword(
      this.token,
      this.form.value.newPassword!,
      this.form.value.confirmPassword!
    ).pipe(finalize(() => this.loading = false))
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('auth.resetPassword.success')
          });
          setTimeout(() => {
            this.router.navigate(['/login']);
          }, 1000);
        }
      });
  }

  passwordMismatch(): boolean {
    const { newPassword, confirmPassword } = this.form.getRawValue();
    return (
      newPassword.length > 0 &&
      confirmPassword.length > 0 &&
      newPassword !== confirmPassword
    );
  }

  get newPassword() {
    return this.form.controls.newPassword;
  }

  get confirmPassword() {
    return this.form.controls.confirmPassword;
  }
}