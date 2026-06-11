import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Router, RouterModule } from '@angular/router';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-home',
  imports: [CommonModule, RouterModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
  user: any;
  role = '';

  constructor(
    private auth: AuthService,
    private router: Router
  ) {}

  ngOnInit() {
    this.auth.getMe()
      .subscribe({
        next: res => {
          this.user = res;
          this.role = res.role;
        }
      });
  }

  logout() {
    this.auth.logout();
    localStorage.removeItem('role');
    this.router.navigate(['/']);
  }
}