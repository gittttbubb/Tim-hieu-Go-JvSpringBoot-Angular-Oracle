import { Component, Input } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { HasPermissionDirective } from '../../shared/directives/has-permission.directive';
import { PERMISSIONS } from '../../constants/permission';
import { TranslatePipe } from '../../shared/pipes/translate.pipe';

@Component({
  selector: 'app-side-bar',
  imports: [RouterLink, RouterLinkActive, HasPermissionDirective, TranslatePipe],
  templateUrl: './side-bar.component.html',
  styleUrl: './side-bar.component.scss'
})
export class SideBarComponent {
  readonly permissions = PERMISSIONS;
  @Input() collapsed = false;
}
