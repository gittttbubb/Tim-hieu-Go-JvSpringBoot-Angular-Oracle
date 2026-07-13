export const en = {
  common: {
    save: 'Save',
    cancel: 'Cancel',
    edit: 'Edit',
    delete: 'Delete',
    create: 'Create New',
    status: 'Status',
    actions: 'Actions',
    success: 'Success',
    warning: 'Warning',
    error: 'Error',
    loading: 'Loading...',
    search: 'Search',
    confirm: 'Confirm'
  },
  sidebar: {
    dashboard: 'Dashboard',
    users: 'Users',
    roles: 'Roles',
    permissions: 'Permissions',
    audits: 'Audits',
    title: 'Users Management'
  },
  topbar: {
    changePassword: 'Change Password',
    logout: 'Logout',
    toggleTheme: 'Toggle Theme',
    welcome: 'Welcome'
  },
  auth: {
    login: {
      title: 'Login',
      subtitle: 'Users Management',
      username: 'Username',
      password: 'Password',
      usernamePlaceholder: 'Enter username',
      passwordPlaceholder: 'Enter password',
      submit: 'Login',
      forgotPasswordLink: 'Forgot password?',
      success: 'Logged in successfully',
      pendingPasswordChange: 'You must change your password on first login',
      failed: 'Invalid username or password'
    },
    forgotPassword: {
      title: 'Forgot Password',
      subtitle: 'Enter your email to receive a password reset link',
      email: 'Email',
      emailPlaceholder: 'Enter your email address',
      submit: 'Send Password Reset Link',
      backToLogin: 'Login',
      success: 'A password reset link has been sent to your email.'
    },
    resetPassword: {
      title: 'Reset Password',
      newPassword: 'New Password',
      confirmPassword: 'Confirm Password',
      submit: 'Reset Password',
      backToLogin: 'Back to Login',
      invalidLink: 'Invalid or expired password reset token.',
      success: 'Password updated successfully.',
      mismatch: 'Passwords do not match',
      validation: {
        required: 'New password is required',
        minlength: 'Password must be at least 8 characters long',
        maxlength: 'Password cannot exceed 100 characters',
        strong: 'Password must contain uppercase letters, lowercase letters, numbers, and special characters',
        confirmRequired: 'Password confirmation does not match'
      }
    },
    changePassword: {
      title: 'Change Password',
      oldPassword: 'Current Password',
      newPassword: 'New Password',
      confirmPassword: 'Confirm New Password',
      submit: 'Update Password',
      success: 'Password changed successfully.',
      oldPasswordRequired: 'Current password is required',
      sameAsOld: 'New password cannot be the same as the old password',
      failed: 'Failed to change password',
      mismatch: 'New password and confirmation password do not match'
    }
  },
  errors: {
    sessionExpired: 'Session expired',
    accessDenied: 'You do not have permission to access this feature',
    serverError: 'A system error has occurred',
    networkError: 'Cannot connect to the server',
    unexpected: 'An unexpected error occurred'
  }
};
