import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { authGuard } from './auth/auth.guard';
import { RegisterComponent } from './pages/register/register.component';

export const routes: Routes = [
  // THêm Role, filter với từng router 
    { path: '', component: LoginComponent },
  { path: 'home', component: HomeComponent, canActivate: [authGuard] },
  {path: 'login', component: LoginComponent},
  {path:'register', component: RegisterComponent}
];
