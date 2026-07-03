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
import { TableLazyLoadEvent } from 'primeng/table';
import { FormsModule } from '@angular/forms';
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { InputTextModule } from 'primeng/inputtext';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';

@Component({
  selector: 'app-user-list',
  imports: [CommonModule, FormsModule, TableModule, TagModule, ButtonModule, ProgressSpinnerModule,
    RouterLink, ConfirmDialogModule, InputTextModule, IconFieldModule, InputIconModule,],
  templateUrl: './user-list.component.html',
  styleUrl: './user-list.component.scss'
})
export class UserListComponent implements OnInit {
  private readonly userService = inject(UserService);
  private readonly roleService = inject(RoleService);
  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);
  private readonly searchSubject = new Subject<string>();

  users: UserList[] = [];
  loading = false;
  page = 1;
  pageSize = 10;
  totalRecords = 0;
  keyword = '';
  rolesMap = new Map<string, string>();

  ngOnInit(): void {
    this.searchSubject
      .pipe(
        debounceTime(300),
        distinctUntilChanged()
      )
      .subscribe(() => {
        this.page = 1;
        this.loadUsers();
      });

    this.loadRoles();
  }

  private loadRoles(): void {
    this.roleService.getAllRoles().subscribe({
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

    this.userService.getUsers({
      page: this.page,
      pageSize: this.pageSize,
      keyword: this.keyword
    })
      .subscribe({
        next: response => {

          this.users = response.data.items;

          this.totalRecords = response.data.total;

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

  onSearchChange(value: string): void {
    this.searchSubject.next(value);
  }

  clearSearch(): void {
    this.keyword = '';
    this.page = 1;
    this.loadUsers();
  }
  onLazyLoad(event: TableLazyLoadEvent): void {
    const rows = event.rows ?? this.pageSize;
    const first = event.first ?? 0;
    this.pageSize = rows;
    this.page = Math.floor(first / rows) + 1;
    this.loadUsers();
  }
}