import { Component, OnInit, inject } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

import { FormsModule } from '@angular/forms';

import { TableModule } from 'primeng/table';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { TagModule } from 'primeng/tag';

import { AuditService } from '../../../services/audit.service';
import { AuditLog } from '../../../models/audit.model';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-audit-list',
  imports: [CommonModule, FormsModule, RouterLink, TableModule, CardModule, InputTextModule, TagModule, ButtonModule, DatePipe],
  templateUrl: './audit-list.component.html',
  styleUrl: './audit-list.component.scss'
})
export class AuditListComponent implements OnInit {

  private readonly auditService = inject(AuditService);
  loading = false;
  audits: AuditLog[] = [];
  filteredAudits: AuditLog[] = [];
  keyword = '';

  ngOnInit(): void {
    this.loadData();
  }

  loadData(): void {
    this.loading = true;
    this.auditService.getAudits()
      .subscribe({
        next: (res) => {
          this.audits = res.data;
          this.filteredAudits = [...this.audits];
          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      });
  }

  applyFilter(): void {
    const keyword = this.keyword.toLowerCase();
    this.filteredAudits =
      this.audits.filter(x =>
        x.action.toLowerCase().includes(keyword)
        ||
        x.actor?.toLowerCase().includes(keyword)
        ||
        x.entityType?.toLowerCase().includes(keyword)
      );
  }

  clearFilter(): void {
    this.keyword = '';
    this.filteredAudits = [...this.audits];
  }
}