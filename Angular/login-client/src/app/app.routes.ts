import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './auth/login/login.component';
import { authGuard } from './auth/auth.guard';

export const routes: Routes = [
    { path: '', component: LoginComponent },
  { path: 'home', component: HomeComponent, canActivate: [authGuard] }
];
