import { Component, OnInit, inject } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { finalize, forkJoin } from 'rxjs';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { DropdownModule } from 'primeng/dropdown';
import { TagModule } from 'primeng/tag';
import { MessageService, ConfirmationService } from 'primeng/api';
import { CreateUserRequest, UpdateUserRequest, UserDetail } from '../../../models/user.model';
import { Role } from '../../../models/role.model';
import { UserService } from '../../../services/user.service';
import { RoleService } from '../../../services/role.service';
import { AuthService } from '../../../services/auth.service';
import { SelectModule } from 'primeng/select';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { DialogModule } from 'primeng/dialog';
import { HasPermissionDirective } from '../../../shared/directives/has-permission.directive';
import { PERMISSIONS } from '../../../constants/permission';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

export type UserFormMode = 'create' | 'detail' | 'edit';

@Component({
  selector: 'app-user-form',
  imports: [CommonModule, ReactiveFormsModule, CardModule, InputTextModule, DropdownModule,
    ButtonModule, TagModule, RouterLink, SelectModule, ConfirmDialogModule, DialogModule, DatePipe, HasPermissionDirective, TranslatePipe],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.scss'
})
export class UserFormComponent implements OnInit {

  private readonly fb = inject(FormBuilder);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  private readonly userService = inject(UserService);
  private readonly roleService = inject(RoleService);
  private readonly authService = inject(AuthService);

  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);
  private readonly translationService = inject(TranslationService);

  readonly permissions = PERMISSIONS;

  tempPassword = '';
  showTempPasswordDialog = false;
  passwordDialogTitle = '';
  passwordDialogMessage = '';

  loading = false;
  mode: UserFormMode = 'create';
  userId: string | null = null;
  user?: UserDetail;
  roles: Role[] = [];

  get statuses() {
    return [
      {
        label: this.translationService.translate('users.statusOptions.active'),
        value: 'ACTIVE'
      },
      {
        label: this.translationService.translate('users.statusOptions.locked'),
        value: 'LOCKED'
      },
      {
        label: this.translationService.translate('users.statusOptions.pending'),
        value: 'PENDING_PASSWORD_CHANGE'
      }
    ];
  }

  form = this.fb.nonNullable.group({
    tenantId: ['00000000-0000-0000-0000-000000000001'],
    fullName: ['', Validators.required],
    username: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    phone: ['', Validators.required],
    roleId: ['', Validators.required],
    status: ['ACTIVE']
  });

  ngOnInit(): void {
    const url = this.router.url;
    if (url.includes('/create')) {
      this.mode = 'create';
    }
    else if (url.includes('/edit/')) {
      this.mode = 'edit';
    }
    else {
      this.mode = 'detail';
    }
    this.userId = this.route.snapshot.paramMap.get('id');
    this.loadRoles();
    if (this.userId) {
      this.loadUser();
    }
  }

  private loadRoles(): void {
    this.roleService.getAllRoles()
      .subscribe({
        next: response => {
          this.roles = response.data;
        }
      });
  }

  private loadUser(): void {
    if (!this.userId) {
      return;
    }
    this.loading = true;
    this.userService.getById(this.userId)
      .pipe(finalize(() => { this.loading = false; }))
      .subscribe({
        next: response => {
          this.user = response.data;
          this.form.patchValue({
            tenantId: response.data.tenantId,
            fullName: response.data.fullName,
            username: response.data.username,
            email: response.data.email,
            phone: response.data.phone,
            roleId: response.data.roleId,
            status: response.data.status
          });
          if (this.mode === 'detail') {
            this.form.disable();
          }
          if (this.mode === 'edit') {
            this.form.controls.tenantId.disable();
          }
        }
      });
  }

  save(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    if (this.mode === 'create') {
      this.create();
      return;
    }
    if (this.mode === 'edit') {
      this.update();
    }
  }

  private create(): void {
    const request = this.form.getRawValue() as CreateUserRequest;
    this.userService.create(request)
      .subscribe({
        next: (res) => {
          this.tempPassword = res.data.temporaryPassword;
          this.passwordDialogTitle = this.translationService.translate('users.passwordDialog.createdTitle');
          this.passwordDialogMessage = this.translationService.translate('users.passwordDialog.createdMessage');
          this.showTempPasswordDialog = true;
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.createSuccess')
          });
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.createFailed')
          });
        }
      });
  }

  private update(): void {
    if (!this.userId) {
      return;
    }
    const request: UpdateUserRequest = {
      fullName: this.form.getRawValue().fullName,
      username: this.form.getRawValue().username,
      email: this.form.getRawValue().email,
      phone: this.form.getRawValue().phone,
      roleId: this.form.getRawValue().roleId,
      status: this.form.getRawValue().status
    };

    this.userService.update(this.userId, request)
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.updateSuccess')
          });
          this.router.navigate(['/users', this.userId]);
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.updateFailed')
          });
        }
      });
  }
  isCreate(): boolean {
    return this.mode === 'create';
  }

  isEdit(): boolean {
    return this.mode === 'edit';
  }

  isDetail(): boolean {
    return this.mode === 'detail';
  }

  goEdit(): void {
    if (!this.userId) {
      return;
    }
    this.router.navigate(['/users/edit', this.userId]);
  }

  deleteUser(): void {
    if (!this.userId) {
      return;
    }
    this.confirmationService.confirm({
      header: this.translationService.translate('common.confirm'),
      message: this.translationService.translate('users.messages.deleteConfirm'),
      accept: () => {
        this.userService.delete(this.userId!)
          .subscribe({
            next: () => {
              this.messageService.add({
                severity: 'success',
                summary: this.translationService.translate('common.success'),
                detail: this.translationService.translate('users.messages.deleteSuccess')
              });
              this.router.navigate(['/users']);
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: this.translationService.translate('common.error'),
                detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.deleteFailed')
              });
            }
          });
      }
    });
  }

  lockUser(): void {
    if (!this.userId) {
      return;
    }
    this.userService.lock(this.userId)
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.lockSuccess')
          });
          this.loadUser();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.lockFailed')
          });
        }
      });
  }

  unlockUser(): void {
    if (!this.userId) {
      return;
    }
    this.userService.unlock(this.userId)
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('users.messages.unlockSuccess')
          });
          this.loadUser();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.unlockFailed')
          });
        }
      });
  }

  resetPassword(): void {
    if (!this.userId) {
      return;
    }
    this.confirmationService.confirm({
      header: this.translationService.translate('common.confirm'),
      message: this.translationService.translate('users.passwordDialog.confirmReset'),
      icon: 'pi pi-exclamation-triangle',
      accept: () => {
        this.authService.adminResetPassword(this.userId!)
          .subscribe({
            next: (res) => {
              this.tempPassword = res.data.temporaryPassword;
              this.passwordDialogTitle = this.translationService.translate('users.passwordDialog.resetTitle');
              this.passwordDialogMessage = this.translationService.translate('users.passwordDialog.resetMessage');
              this.showTempPasswordDialog = true;

              this.messageService.add({
                severity: 'success',
                summary: this.translationService.translate('common.success'),
                detail: this.translationService.translate('users.messages.resetSuccess')
              });
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: this.translationService.translate('common.error'),
                detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('users.messages.resetFailed')
              });
            }
          });
      }
    });
  }
  copyPassword(): void {
    if (!this.tempPassword) {
      return;
    }
    navigator.clipboard.writeText(this.tempPassword);
    this.messageService.add({
      severity: 'success',
      summary: this.translationService.translate('common.success'),
      detail: this.translationService.translate('users.passwordDialog.copied')
    });
  }

  closePasswordDialog(): void {
    this.showTempPasswordDialog = false;

    if (this.isCreate()) {
      this.router.navigate(['/users']);
    }
  }
}
