import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { TagModule } from 'primeng/tag';
import { Permission } from '../../../models/permission.model';
import { PermissionService } from '../../../services/permisison.service';
import { ButtonModule } from 'primeng/button';
import { TableLazyLoadEvent } from 'primeng/table';
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';

import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';

@Component({
  selector: 'app-permission-list',
  imports: [CommonModule, FormsModule, TableModule, CardModule, InputTextModule, ButtonModule,
    TagModule, IconFieldModule, InputIconModule],
  templateUrl: './permission-list.component.html',
  styleUrl: './permission-list.component.scss'
})
export class PermissionListComponent implements OnInit {

  private readonly permissionService = inject(PermissionService);
  loading = false;
  permissions: Permission[] = [];
  page = 1;
  pageSize = 10;
  totalRecords = 0;
  keyword = '';
  private readonly searchSubject = new Subject<string>();
  searchKeyword = '';


  ngOnInit(): void {
    this.searchSubject
      .pipe(
        debounceTime(300),
        distinctUntilChanged()
      )
      .subscribe(() => {
        this.page = 1;
        this.loadPermissions();
      });
    this.loadPermissions();
  }

  loadPermissions(): void {
    this.loading = true;
    this.permissionService.getPermissionsPagination({
      page: this.page,
      pageSize: this.pageSize,
      keyword: this.keyword
    })
      .subscribe({
        next: (res) => {
          this.permissions = res.data.items;
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
    this.loadPermissions();
  }

  onLazyLoad(event: TableLazyLoadEvent): void {
    const rows = event.rows ?? this.pageSize;
    const first = event.first ?? 0;
    this.pageSize = rows;
    this.page = Math.floor(first / rows) + 1;
    this.loadPermissions();
  }
}