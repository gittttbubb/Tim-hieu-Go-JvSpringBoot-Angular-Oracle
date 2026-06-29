import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, ReactiveFormsModule, Validators, } from '@angular/forms';
import { finalize } from 'rxjs';
import { MessageService } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { PasswordModule } from 'primeng/password';
import { AuthService } from '../../../services/auth.service';
import { AuthStore } from '../../../store/auth.store';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule, ButtonModule, CardModule, InputTextModule, PasswordModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {

  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly authStore = inject(AuthStore);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);

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
            summary: 'Success',
            detail: 'Đăng nhập thành công',
          });
          if (loginData.mustChangePassword) {
            this.router.navigate(['/change-password']);
            this.messageService.add({
              severity: 'warning',
              summary: 'Đổi mật khẩu',
              detail: 'Bạn phải đổi mật khẩu lần đầu đăng nhập',
            });
            return;
          }
          this.router.navigate(['/dashboard']);
        },
        error: error => {
          this.messageService.add({
            severity: 'error',
            summary: 'Login Failed',
            detail: 'Sai tài khoản hoặc mật khẩu',
          });
        },
      });
  }
}