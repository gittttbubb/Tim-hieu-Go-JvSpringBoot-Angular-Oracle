import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-change-password',
  imports: [FormsModule],
  templateUrl: './change-password.component.html',
  styleUrl: './change-password.component.scss'
})
export class ChangePasswordComponent {
  password = '';

  constructor(
    private auth: AuthService
  ) { }

  changePassword() {
    this.auth.changePassword(this.password)
      .subscribe({
        next: () => {
          alert('Đổi mật khẩu thành công');
          this.password = '';
        },
        error: err => {
          alert(err.error.message);
        }
      });

  }
}
