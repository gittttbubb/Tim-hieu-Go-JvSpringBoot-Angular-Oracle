import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-home',
  imports: [CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
  data: any;

  constructor(
    private http: HttpClient,
    private auth: AuthService,
    private router: Router
  ) {}

  getMe() {
    this.http.get('http://localhost:8080/api/users/me')
      .subscribe(res => this.data = res);
  }

  logout() {
    this.auth.logout();
    this.router.navigate(['/']);
  }
}