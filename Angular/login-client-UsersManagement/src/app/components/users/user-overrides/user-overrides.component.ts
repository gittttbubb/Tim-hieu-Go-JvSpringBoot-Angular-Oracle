import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { DropdownModule } from 'primeng/dropdown';
import { InputTextModule } from 'primeng/inputtext';
import {
  InputTextarea
} from 'primeng/inputtextarea';
import { TagModule } from 'primeng/tag';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { UserService } from '../../../services/user.service';
import { RoleService } from '../../../services/role.service';
import { PermissionService } from '../../../services/permisison.service';
import { MessageService } from 'primeng/api';
import { UserDetail, UserOverride } from '../../../models/user.model';
import { Permission } from '../../../models/permission.model';
import {
  DataScope
} from '../../../models/user.model';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';


@Component({
  selector: 'app-user-overrides',
  imports: [CommonModule, ReactiveFormsModule, TableModule, ButtonModule, DialogModule, DropdownModule, InputTextModule, InputTextarea, TagModule, RouterLink, TranslatePipe],
  templateUrl: './user-overrides.component.html',
  styleUrl: './user-overrides.component.scss'
})
export class UserOverridesComponent {

  private readonly route = inject(ActivatedRoute);
  private readonly fb = inject(FormBuilder);
  private readonly userService = inject(UserService);
  private readonly roleService = inject(RoleService);
  private readonly permissionService = inject(PermissionService);
  private readonly overrideService = inject(UserService);
  private readonly messageService = inject(MessageService);
  private readonly translationService = inject(TranslationService);

  userId = '';
  user?: UserDetail;
  roleName = '';
  loading = false;
  dialogVisible = false;
  overrides: UserOverride[] = [];
  permissions: Permission[] = [];
  permissionMap = new Map<string, Permission>();

  form = this.fb.nonNullable.group({
    permissionId: ['', Validators.required],
    granted: [true],
    dataScope: this.fb.nonNullable.control<DataScope>('ALL', Validators.required),
    reason: ['']
  });

  ngOnInit(): void {
    this.userId = this.route.snapshot.paramMap.get('id') ?? '';

    this.loadUser();
    this.loadPermissions();
    this.loadOverrides();
  }

  private loadUser(): void {
    this.userService.getById(this.userId)
      .subscribe({
        next: res => {
          this.user = res.data;
          this.roleService.getAllRoles()
            .subscribe(roleRes => {
              const role = roleRes.data.find(x => x.id === this.user?.roleId);
              this.roleName = role?.name ?? '';
            });
        }
      });
  }

  private loadPermissions(): void {
    this.permissionService.getPermissions()
      .subscribe({
        next: res => {
          this.permissions = res.data;
          res.data.forEach(p => {
            this.permissionMap.set(p.id, p);
          });
        }
      });
  }

  loadOverrides(): void {
    this.loading = true;
    this.overrideService.getByUserId(this.userId)
      .subscribe({
        next: res => {
          this.overrides = res.data;
          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      });
  }

  openDialog(): void {
    this.form.reset({
      permissionId: '',
      granted: true,
      dataScope: 'ALL',
      reason: ''
    });

    this.dialogVisible = true;
  }

  saveOverride(): void {
    if (this.form.invalid) {
      return;
    }
    this.overrideService.assign({ userId: this.userId, ...this.form.getRawValue() })
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.overrideSuccess')
          });
          this.dialogVisible = false;
          this.loadOverrides();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.overrideFailed')
          });
        }
      });
  }

  deleteOverride(id: string): void {
    if (!confirm(this.translationService.translate('users.messages.overrideDeleteConfirm'))) {
      return;
    }
    this.overrideService.remove(id)
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.overrideDeleteSuccess')
          });
          this.loadOverrides();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.overrideDeleteFailed')
          });
        }
      });
  }

  getPermissionName(permissionId: string): string {
    const permission = this.permissionMap.get(permissionId);
    if (!permission) {
      return permissionId;
    }
    return `${permission.featureCode}_${permission.action}`;
  }
}
