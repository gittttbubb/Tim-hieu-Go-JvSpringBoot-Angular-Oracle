import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs';
import { User } from '../models/User.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private API = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  login(username: string, password: string) {
    return this.http.post<any>(`${this.API}/auth/login`, {
      username,
      password
    }).pipe(
      tap(res => {
        localStorage.setItem('token', res.token);
      })
    );
  }
  register(username: string, password: string) {
    return this.http.post<any>(`${this.API}/auth/register`, {
      username,
      password
    });
  }
  getToken() {
    return localStorage.getItem('token');
  }

  getMe() {
    return this.http.get<User>(`${this.API}/users/me`);
  }

  updateProfile(username: string) {
    return this.http.put(`${this.API}/users/profile`,{username});
  }

  changePassword(password: string) {
    return this.http.put(`${this.API}/users/change-password`,{password});
  }

  getAllUsers() {
    return this.http.get<User[]>(`${this.API}/admin/list-users`);
  }

  promoteToAdmin(id: number) {
    return this.http.put(`${this.API}/admin/users/promote/${id}`,{});
  }

  logout() {
    localStorage.removeItem('token');
  }
}