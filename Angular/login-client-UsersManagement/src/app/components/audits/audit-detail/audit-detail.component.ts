import { Component, OnInit, inject } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';

import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { TagModule } from 'primeng/tag';

import { AuditService } from '../../../services/audit.service';
import { AuditLog } from '../../../models/audit.model';

import { TranslatePipe } from '../../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-audit-detail',
  imports: [CommonModule, RouterLink, CardModule, ButtonModule, TagModule, DatePipe, TranslatePipe],
  templateUrl: './audit-detail.component.html',
  styleUrl: './audit-detail.component.scss'
})
export class AuditDetailComponent implements OnInit {

  private readonly route = inject(ActivatedRoute);
  private readonly auditService = inject(AuditService);
  loading = false;
  audit?: AuditLog;

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.loadAudit(id);
    }
  }

  loadAudit(id: string): void {
    this.loading = true;
    this.auditService.getById(id)
      .subscribe({
        next: (res) => {
          this.audit = res.data;
          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      });
  }

  formatJson(value?: string): string {
    if (!value) {
      return '';
    }
    try {
      return JSON.stringify(JSON.parse(value), null, 2);
    } catch {
      return value;
    }
  }
}