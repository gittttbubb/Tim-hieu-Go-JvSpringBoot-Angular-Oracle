import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-list-users',
  imports: [CommonModule],
  templateUrl: './list-users.component.html',
  styleUrl: './list-users.component.scss'
})
export class ListUsersComponent {
  users: any[] = [];

  constructor(
    private auth: AuthService
  ) { }

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.auth.getAllUsers().subscribe({
      next: res => {
        this.users = res;
      },
      error: err => {
        alert(err.error.message);
      }
    });
  }

  promote(id: number) {
    this.auth.promoteToAdmin(id).subscribe({
      next: () => {
        alert('Nâng quyền thành công');
        this.loadUsers();
      },
      error: err => {
        alert(err.error.message);
      }
    });
  }
}
