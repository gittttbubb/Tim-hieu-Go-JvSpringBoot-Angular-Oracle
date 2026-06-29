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

@Component({
  selector: 'app-change-password',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, CardModule, ButtonModule, PasswordModule, ToastModule],
  providers: [MessageService],
  templateUrl: './change-password.component.html',
  styleUrl: './change-password.component.scss'
})
export class ChangePasswordComponent {
  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly authStore = inject(AuthStore);

  loading = false;
  form = this.fb.nonNullable.group({
    oldPassword: ['', Validators.required],
    newPassword: ['',
      [
        Validators.required,
        Validators.minLength(6),
        Validators.maxLength(100)
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
        summary: 'Error',
        detail: 'Passwords do not match'
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
            summary: 'Success',
            detail: 'Password changed successfully'
          });
          this.form.reset();
          this.logout();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail:
              err.error?.message ??
              'Change password failed'
          });
        }
      });
  }

  back(): void {
    this.router.navigate(['/']);
  }
  logout(): void {
    this.authStore.clear();
    this.router.navigate(['/login']);
  }
}