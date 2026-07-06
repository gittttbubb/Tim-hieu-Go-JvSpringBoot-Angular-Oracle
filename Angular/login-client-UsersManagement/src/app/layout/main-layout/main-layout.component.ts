import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SideBarComponent } from '../side-bar/side-bar.component';
import { TopBarComponent } from '../top-bar/top-bar.component';

@Component({
  selector: 'app-main-layout',
  imports: [RouterOutlet, SideBarComponent, TopBarComponent,],
  templateUrl: './main-layout.component.html',
  styleUrl: './main-layout.component.scss'
})
export class MainLayoutComponent {
  readonly sidebarCollapsed = signal(false);

    toggleSidebar(): void {
        this.sidebarCollapsed.update(v => !v);
    }
}
