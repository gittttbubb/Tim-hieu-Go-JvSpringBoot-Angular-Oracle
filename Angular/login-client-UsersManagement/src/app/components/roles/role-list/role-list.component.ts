import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { TableModule } from 'primeng/table';
import { ButtonModule } from 'primeng/button';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ToastModule } from 'primeng/toast';
import { ConfirmationService } from 'primeng/api';
import { MessageService } from 'primeng/api';
import { Role } from '../../../models/role.model';
import { RoleService } from '../../../services/role.service';
import { FormsModule } from '@angular/forms';
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { TableLazyLoadEvent } from 'primeng/table';

import { InputTextModule } from 'primeng/inputtext';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';

@Component({
  selector: 'app-role-list',
  imports: [CommonModule,FormsModule, TableModule, ButtonModule, ConfirmDialogModule, ToastModule, InputTextModule, IconFieldModule, InputIconModule],
  providers: [ConfirmationService, MessageService],
  templateUrl: './role-list.component.html',
  styleUrl: './role-list.component.scss'
})
export class RoleListComponent implements OnInit {
  private readonly roleService = inject(RoleService);
  private readonly router = inject(Router);
  private readonly confirmationService = inject(ConfirmationService);
  private readonly messageService = inject(MessageService);

  roles: Role[] = [];

  loading = false;

  page = 1;
  pageSize = 10;
  totalRecords = 0;

  keyword = '';

  private readonly searchSubject = new Subject<string>();
  ngOnInit(): void {

    this.searchSubject
      .pipe(
        debounceTime(300),
        distinctUntilChanged()
      )
      .subscribe(() => {
        this.page = 1;
        this.loadRoles();
      });

    this.loadRoles();
  }

  loadRoles(): void {

    this.loading = true;

    this.roleService.getRoles({
      page: this.page,
      pageSize: this.pageSize,
      keyword: this.keyword
    })
      .subscribe({
        next: (res) => {

          this.roles = res.data.items;

          this.totalRecords = res.data.total;

          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      });
  }
  onSearchChange(value: string): void {
    this.searchSubject.next(value);
  }

  clearSearch(): void {
    this.keyword = '';
    this.page = 1;
    this.loadRoles();
  }
  onLazyLoad(event: TableLazyLoadEvent): void {

    const rows = event.rows ?? this.pageSize;
    const first = event.first ?? 0;

    this.pageSize = rows;
    this.page = Math.floor(first / rows) + 1;

    this.loadRoles();
  }

  createRole(): void {
    this.router.navigate(['/roles/new']);
  }

  viewRole(role: Role): void {
    this.router.navigate(['/roles', role.id]);
  }

  editRole(role: Role): void {
    this.router.navigate(['/roles', 'edit', role.id]);
  }

  managePermissions(role: Role): void {
    this.router.navigate(['/roles', 'permissions', role.id]);
  }

  deleteRole(role: Role): void {
    this.confirmationService.confirm({
      message: `Delete role "${role.displayName}" ?`,
      header: 'Confirm',
      accept: () => {
        this.roleService.deleteRole(role.id)
          .subscribe({
            next: () => {
              this.messageService.add({
                severity: 'success',
                summary: 'Success',
                detail: 'Role deleted'
              });

              this.loadRoles();
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: 'Error',
                detail: err.error?.message ?? 'Deleted role failed'
              });
            }
          });
      }
    });
  }
}