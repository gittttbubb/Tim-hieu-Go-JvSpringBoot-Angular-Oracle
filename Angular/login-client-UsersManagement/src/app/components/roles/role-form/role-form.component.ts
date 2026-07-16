import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { finalize } from 'rxjs';
import { InputTextModule } from 'primeng/inputtext';
import { InputTextarea } from 'primeng/inputtextarea';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { ToastModule } from 'primeng/toast';
import { ConfirmationService, MessageService } from 'primeng/api';
import { RoleService } from '../../../services/role.service';
import { Role } from '../../../models/role.model';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { TranslationService } from '../../../services/translation.service';
import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-role-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, InputTextModule, InputTextarea, ButtonModule, CardModule, ToastModule, ConfirmDialogModule, TranslatePipe],
  providers: [MessageService, ConfirmationService],
  templateUrl: './role-form.component.html',
  styleUrl: './role-form.component.scss'
})
export class RoleFormComponent implements OnInit {
  private readonly fb = inject(FormBuilder);
  private readonly roleService = inject(RoleService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);
  private readonly confirmationService = inject(ConfirmationService);
  private readonly translationService = inject(TranslationService);

  roleId = '';
  loading = false;
  mode: | 'create' | 'view' | 'edit' = 'create';

  form = this.fb.nonNullable.group({
    name: ['', [Validators.required, Validators.maxLength(50)]],
    displayName: ['', [Validators.required, Validators.maxLength(100)]],
    description: ['', [Validators.maxLength(255)]]
  });

  ngOnInit(): void {

    const id = this.route.snapshot.paramMap.get('id');

    if (this.router.url.includes('/new')) {
      this.mode = 'create';
      return;
    }

    if (this.router.url.includes('/edit/')) {
      this.mode = 'edit';
    } else {
      this.mode = 'view';
    }

    if (id) {
      this.roleId = id;
      this.loadRole();
    }
  }

  loadRole(): void {
    this.loading = true;

    this.roleService.getRoleById(this.roleId)
      .pipe(
        finalize(() => this.loading = false)
      )
      .subscribe({
        next: (res) => {

          this.patchForm(res.data);

          if (this.mode === 'view') {
            this.form.disable();
          } else {
            this.form.enable();
          }
        }
      });
  }
  patchForm(role: Role): void {
    this.form.patchValue({
      name: role.name,
      displayName: role.displayName,
      description: role.description
    });
  }

  save(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    const payload = this.form.getRawValue();
    this.loading = true;
    this.roleService.createRole(payload)
      .pipe(finalize(() => this.loading = false))
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('roles.messages.createSuccess')
          });
          this.back();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('roles.messages.createFailed')
          });
        }
      });
  }

  update(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    const payload = this.form.getRawValue();
    this.loading = true;
    this.roleService.updateRole(this.roleId, payload)
      .pipe(finalize(() => this.loading = false))
      .subscribe({
        next: () => {
          this.messageService.add({
            severity: 'success',
            summary: this.translationService.translate('common.success'),
            detail: this.translationService.translate('roles.messages.updateSuccess')
          });
          this.back();
        },
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: this.translationService.translate('common.error'),
            detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('roles.messages.updateFailed')
          });
        }
      });
  }

  goEdit(): void {
    this.router.navigate(['/roles', 'edit', this.roleId]);
  }

  back(): void {
    this.router.navigate(['/roles']);
  }

  deleteRole(): void {
    this.confirmationService.confirm({
      header: this.translationService.translate('roles.title'),
      message: this.translationService.translate('roles.messages.confirmDelete', { name: this.form.getRawValue().displayName }),
      icon: 'pi pi-exclamation-triangle',
      accept: () => {
        this.loading = true;
        this.roleService.deleteRole(this.roleId)
          .pipe(finalize(() => this.loading = false))
          .subscribe({
            next: () => {
              this.messageService.add({
                severity: 'success',
                summary: this.translationService.translate('common.success'),
                detail: this.translationService.translate('roles.messages.deleteSuccess')
              });
              this.back();
            },
            error: (err) => {
              this.messageService.add({
                severity: 'error',
                summary: this.translationService.translate('common.error'),
                detail: err.error?.message ? this.translationService.translate(err.error.message) : this.translationService.translate('roles.messages.deleteFailed')
              });
            }
          });
      }
    });
  }
}