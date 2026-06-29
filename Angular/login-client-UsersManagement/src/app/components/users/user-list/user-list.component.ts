import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TableModule } from 'primeng/table';
import { TagModule } from 'primeng/tag';
import { ButtonModule } from 'primeng/button';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { UserService } from '../../../services/user.service';
import { RoleService } from '../../../services/role.service';
import { UserList } from '../../../models/user.model';
import { RouterLink } from '@angular/router';
import { ConfirmationService, MessageService } from 'primeng/api';
import { ConfirmDialogModule } from 'primeng/confirmdialog';

@Component({
  selector: 'app-user-list',
  imports: [CommonModule, TableModule, TagModule, ButtonModule, ProgressSpinnerModule, RouterLink, ConfirmDialogModule],
  templateUrl: './user-list.component.html',
  styleUrl: './user-list.component.scss'
})
export class UserListComponent implements OnInit {
  private readonly userService = inject(UserService);
  private readonly roleService = inject(RoleService);
  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);

  users: UserList[] = [];
  loading = false;
  rolesMap = new Map<string, string>();

  ngOnInit(): void {
    this.loadRoles();
  }

  private loadRoles(): void {
    this.roleService.getRoles().subscribe({
      next: response => {
        response.data.forEach(role => {
          this.rolesMap.set(
            role.id,
            role.displayName
          );
        });
        this.loadUsers();
      }
    });
  }

  private loadUsers(): void {
    this.loading = true;
    this.userService.getUsers().subscribe({
      next: response => {
        this.users = response.data;
        this.loading = false;
      },
      error: () => {
        this.loading = false;
      }
    });
  }

  getRoleName(roleId: string): string {
    return this.rolesMap.get(roleId) ?? roleId;
  }

  getStatusSeverity(status: string) {
    switch (status) {
      case 'ACTIVE':
        return 'success';
      case 'LOCKED':
        return 'danger';
      default:
        return 'secondary';
    }
  }

  deleteUser(id: string): void {
    this.confirmationService.confirm({
      header: 'Delete User',
      message: 'Delete this user?',
      icon: 'pi pi-exclamation-triangle',
      accept: () => {
        this.userService.delete(id)
          .subscribe({
            next: () => {
              this.messageService.add({
                severity: 'success',
                summary: 'Success',
                detail: 'User deleted'
              });
              this.loadUsers();
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: 'Error',
                detail: err.error?.message ?? 'Deleted user failed'
              });
            }
          });
      }
    });
  }
}