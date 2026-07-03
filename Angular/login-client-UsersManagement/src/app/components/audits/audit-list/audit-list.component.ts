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
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { TableLazyLoadEvent } from 'primeng/table';

import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';

@Component({
  selector: 'app-audit-list',
  imports: [CommonModule, FormsModule, RouterLink, TableModule, CardModule, InputTextModule, TagModule,
    ButtonModule, DatePipe, IconFieldModule, InputIconModule],
  templateUrl: './audit-list.component.html',
  styleUrl: './audit-list.component.scss'
})
export class AuditListComponent implements OnInit {

  private readonly auditService = inject(AuditService);
  loading = false;
  audits: AuditLog[] = [];
  page = 1;
  pageSize = 10;
  totalRecords = 0;
  private readonly searchSubject = new Subject<string>();
  keyword = '';

  ngOnInit(): void {
    this.searchSubject
      .pipe(
        debounceTime(300),
        distinctUntilChanged()
      )
      .subscribe(() => {
        this.page = 1;
        this.loadData();
      });
    this.loadData();
  }

  loadData(): void {
    this.loading = true;
    this.auditService.getAuditsPagination({
      page: this.page,
      pageSize: this.pageSize,
      keyword: this.keyword
    })
      .subscribe({
        next: (res) => {
          this.audits = res.data.items;
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

  clearFilter(): void {
    this.keyword = '';
    this.page = 1;
    this.loadData();
  }

  onLazyLoad(event: TableLazyLoadEvent): void {
    const rows = event.rows ?? this.pageSize;
    const first = event.first ?? 0;
    this.pageSize = rows;
    this.page = Math.floor(first / rows) + 1;
    this.loadData();
  }
}