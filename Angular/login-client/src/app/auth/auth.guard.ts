import { CanActivateFn } from '@angular/router';
import { Router } from '@angular/router';
import { inject } from '@angular/core';

export const authGuard: CanActivateFn = (route) => {

  const router = inject(Router);
  const token = localStorage.getItem('token');

  if (!token) {
    router.navigate(['/login']);
    return false;
  }

  const requiredRole = route.data?.['role'];

  if (requiredRole) {
    const currentRole = localStorage.getItem('role');
    if (currentRole !== requiredRole) {
      router.navigate(['/home']);
      return false;
    }
  }
  return true;
};