import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { forkJoin, finalize } from 'rxjs';
import { TableModule } from 'primeng/table';
import { CheckboxModule } from 'primeng/checkbox';
import { SelectModule } from 'primeng/select';
import { ButtonModule } from 'primeng/button';
import { ToastModule } from 'primeng/toast';
import { CardModule } from 'primeng/card';
import { ConfirmationService, MessageService } from 'primeng/api';
import { RoleService } from '../../../services/role.service';
import { PermissionService } from '../../../services/permisison.service';
import { RolePermissionView } from '../../../models/role.model';
import { FormsModule } from '@angular/forms';
import { ConfirmDialogModule } from 'primeng/confirmdialog';


@Component({
  selector: 'app-role-permissions',
  imports: [CommonModule, FormsModule, TableModule, CheckboxModule, SelectModule, ButtonModule, ToastModule, CardModule, ConfirmDialogModule],
  providers: [MessageService, ConfirmationService],
  templateUrl: './role-permissions.component.html',
  styleUrl: './role-permissions.component.scss'
})
export class RolePermissionsComponent implements OnInit {
  private readonly roleService = inject(RoleService);
  private readonly permissionService = inject(PermissionService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);

  roleId = '';
  loading = false;
  permissions: RolePermissionView[] = [];
  readonly scopes = ['OWN', 'TEAM', 'ALL'];

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (!id) {
      this.back();
      return;
    }
    this.roleId = id;
    this.loadData();
  }

  loadData(): void {
    this.loading = true;
    forkJoin({
      permissions: this.permissionService.getPermissions(),
      rolePermissions: this.roleService.getRolePermissions(this.roleId)
    })
      .pipe(finalize(() => this.loading = false))
      .subscribe({
        next: ({ permissions, rolePermissions }) => {
          const roleMap = new Map(
            rolePermissions.data.map(item => [item.permissionId, item])
          );
          this.permissions = permissions.data.map(permission => {
            const assigned = roleMap.get(permission.id);
            return {
              permissionId: permission.id,
              featureGroup: permission.featureGroup,
              featureCode: permission.featureCode,
              action: permission.action,
              description: permission.description,
              granted: assigned?.granted ?? false,
              dataScope: assigned?.dataScope ?? 'OWN'
            };
          }
          );
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail:
              err.error?.message ??
              'Load permissions failed'
          });
        }
      });
  }

  savePermission(row: RolePermissionView): void {
    this.roleService.assignPermission({
      roleId: this.roleId,
      permissionId: row.permissionId,
      granted: row.granted,
      dataScope: row.dataScope
    })
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Permission updated'
          });
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail:
              err.error?.message ??
              'Update permission failed'
          });
        }
      });
  }

  removePermission(row: RolePermissionView): void {
    this.confirmationService.confirm({
      header: 'Remove Permission',
      message: `Remove permission "${row.featureCode}_${row.action}" from this role?`,
      icon: 'pi pi-exclamation-triangle',
      accept: () => {
        this.roleService.removePermission(this.roleId, row.permissionId)
          .subscribe({
            next: () => {
              row.granted = false;
              row.dataScope = 'OWN';
              this.messageService.add({
                severity: 'success',
                summary: 'Success',
                detail: 'Permission removed'
              });
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: 'Error',
                detail:
                  err.error?.message ??
                  'Remove permission failed'
              });
            }
          });
      }
    });
  }

  back(): void {
    this.router.navigate([
      '/roles'
    ]);
  }
}
