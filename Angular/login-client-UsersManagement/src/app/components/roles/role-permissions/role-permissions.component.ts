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
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';


@Component({
  selector: 'app-role-permissions',
  imports: [CommonModule, FormsModule, TableModule, CheckboxModule, SelectModule, ButtonModule, ToastModule, CardModule, ConfirmDialogModule, TranslatePipe],
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
  private readonly translationService = inject(TranslationService);

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
            summary: this.translationService.translate('common.error'),
            detail:
              err.error?.message ? this.translationService.translate(err.error.message) : 'Load permissions failed'
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
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('roles.messages.permissionUpdated')
          });
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail:
              err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('roles.messages.permissionUpdateFailed')
          });
        }
      });
  }

  removePermission(row: RolePermissionView): void {
    this.confirmationService.confirm({
      header: this.translationService.translate('roles.rolePermissions'),
      message: this.translationService.translate('roles.messages.confirmRemovePermission', { permission: `${row.featureCode}_${row.action}` }),
      icon: 'pi pi-exclamation-triangle',
      accept: () => {
        this.roleService.removePermission(this.roleId, row.permissionId)
          .subscribe({
            next: () => {
              row.granted = false;
              row.dataScope = 'OWN';
              this.messageService.add({
                severity: 'success',
                summary: this.translationService.translate('common.success'),
                detail: this.translationService.translate('roles.messages.permissionRemoved')
              });
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: this.translationService.translate('common.error'),
                detail:
                  err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('roles.messages.permissionRemoveFailed')
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
