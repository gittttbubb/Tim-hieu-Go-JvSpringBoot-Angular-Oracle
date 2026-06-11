import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-profile',
  imports: [CommonModule, FormsModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.scss'
})
export class ProfileComponent {
  user: any;
  username = '';

  constructor(
    private auth: AuthService
  ) { }

  ngOnInit() {
    this.loadProfile();
  }

  loadProfile() {
    this.auth.getMe().subscribe({
      next: res => {
        this.user = res;
        this.username =
          res.username;
      }
    });
  }

  updateProfile() {
    this.auth.updateProfile(this.username).subscribe({
      next: () => {
        alert('Cập nhật thành công');
        this.loadProfile();
      },
      error: err => {
        alert(err.error.message);
      }
    });
  }
}
