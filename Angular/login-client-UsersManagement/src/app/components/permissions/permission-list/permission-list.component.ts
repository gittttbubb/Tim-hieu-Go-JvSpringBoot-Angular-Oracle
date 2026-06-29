import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { SelectModule } from 'primeng/select';
import { TagModule } from 'primeng/tag';
import { Permission } from '../../../models/permission.model';
import { PermissionService } from '../../../services/permisison.service';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-permission-list',
  imports: [ CommonModule, FormsModule, TableModule, CardModule,InputTextModule, SelectModule, ButtonModule, TagModule],
  templateUrl: './permission-list.component.html',
  styleUrl: './permission-list.component.scss'
})
export class PermissionListComponent implements OnInit {

  private readonly permissionService = inject(PermissionService);
  loading = false;
  permissions: Permission[] = [];
  filteredPermissions: Permission[] = [];
  searchKeyword = '';
  selectedFeatureGroup: string | null = null;
  featureGroups: {
    label: string;
    value: string | null;
  }[] = [];

  ngOnInit(): void {
    this.loadPermissions();
  }

  loadPermissions(): void {
    this.loading = true;
    this.permissionService.getPermissions()
      .subscribe({
        next: (res) => {
          this.permissions = res.data;
          const groups = [...new Set(res.data.map(item => item.featureGroup))];
          this.featureGroups = [
            {
              label: 'All',
              value: null
            },
            ...groups.map(group => ({
              label: group,
              value: group
            }))
          ];
          this.applyFilter();
          this.loading = false;
        },
        error: () => {
          this.loading = false;
        }
      });
  }

  applyFilter(): void {
    const keyword = this.searchKeyword.trim().toLowerCase();
    this.filteredPermissions =
      this.permissions.filter(item => {
        const matchKeyword =
          !keyword ||
          item.featureCode.toLowerCase().includes(keyword) ||
          item.featureGroup.toLowerCase().includes(keyword) ||
          item.action.toLowerCase().includes(keyword) ||
          item.description.toLowerCase().includes(keyword);
        const matchGroup = !this.selectedFeatureGroup || item.featureGroup === this.selectedFeatureGroup;
        return matchKeyword && matchGroup;
      });
  }

  clearFilter(): void {
    this.searchKeyword = '';
    this.selectedFeatureGroup = null;
    this.applyFilter();
  }
}