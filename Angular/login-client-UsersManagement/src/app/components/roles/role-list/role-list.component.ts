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

@Component({
  selector: 'app-role-list',
  imports: [CommonModule, TableModule, ButtonModule, ConfirmDialogModule, ToastModule],
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

  ngOnInit(): void {
    this.loadRoles();
  }

  loadRoles(): void {
    this.loading = true;
    this.roleService.getRoles().subscribe({
      next: (res) => {
        this.roles = res.data;
        this.loading = false;
      },
      error: () => {
        this.loading = false;
      }
    });
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