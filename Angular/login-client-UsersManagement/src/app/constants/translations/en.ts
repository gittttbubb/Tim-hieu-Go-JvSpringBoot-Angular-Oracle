export const en = {
  common: {
    save: 'Save',
    cancel: 'Cancel',
    edit: 'Edit',
    update: 'Update',
    delete: 'Delete',
    create: 'Create New',
    status: 'Status',
    actions: 'Actions',
    success: 'Success',
    warning: 'Warning',
    error: 'Error',
    loading: 'Loading...',
    search: 'Search',
    confirm: 'Confirm',
    back: 'Back'
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
  },
  dashboard: {
    welcome: 'Dashboard works!'
  },
  audits: {
    title: 'Audit Logs',
    detailTitle: 'Audit Detail',
    subtitle: 'Activity information',
    searchPlaceholder: 'Search action, actor, entity...',
    noLogs: 'No audit logs found',
    fields: {
      action: 'Action',
      actor: 'Actor',
      actorId: 'Actor Id',
      entityType: 'Entity Type',
      entityId: 'Entity Id',
      targetUser: 'Target User',
      ipAddress: 'IP Address',
      timestamp: 'Timestamp',
      reason: 'Reason',
      beforeData: 'Before Data',
      afterData: 'After Data',
      metadata: 'Metadata'
    }
  },
  permissions: {
    title: 'Permissions',
    searchPlaceholder: 'Search permissions...',
    noPermissions: 'No permissions found',
    fields: {
      group: 'Feature Group',
      code: 'Feature Code',
      action: 'Action',
      description: 'Description'
    }
  },
  roles: {
    title: 'Roles',
    searchPlaceholder: 'Search roles...',
    createRole: 'Create Role',
    editRole: 'Edit Role',
    roleDetail: 'Role Detail',
    rolePermissions: 'Role Permissions',
    fields: {
      name: 'Name',
      displayName: 'Display Name',
      description: 'Description',
      granted: 'Granted',
      scope: 'Scope'
    },
    validation: {
      nameRequired: 'Name is required',
      displayNameRequired: 'Display Name is required'
    },
    messages: {
      confirmDelete: 'Are you sure you want to delete role "{name}"?',
      deleteSuccess: 'Role deleted successfully',
      deleteFailed: 'Delete role failed',
      createSuccess: 'Role created successfully',
      createFailed: 'Created role failed',
      updateSuccess: 'Role updated successfully',
      updateFailed: 'Updated role failed',
      permissionUpdated: 'Permission updated',
      permissionUpdateFailed: 'Update permission failed',
      permissionRemoved: 'Permission removed',
      permissionRemoveFailed: 'Remove permission failed',
      confirmRemovePermission: 'Remove permission "{permission}" from this role?'
    }
  },
  users: {
    statusOptions: {
      active: 'Active',
      locked: 'Locked',
      pending: 'Pending Password Change'
    },
    title: 'Users Management',
    searchPlaceholder: 'Search users...',
    createUser: 'Create User',
    editUser: 'Edit User',
    userDetail: 'User Detail',
    addUser: 'Add User',
    permissionOverrides: 'User Permission Overrides',
    addOverride: 'Add Override',
    assignOverride: 'Assign Override',
    fields: {
      fullName: 'Full Name',
      username: 'Username',
      email: 'Email',
      phone: 'Phone',
      role: 'Role',
      selectRole: 'Select Role',
      status: 'Status',
      password: 'Password',
      mustChange: 'Must Change',
      created: 'Created',
      updated: 'Updated',
      type: 'Type',
      reason: 'Reason',
      dataScope: 'Data Scope',
      grantPermission: 'Grant Permission',
      permission: 'Permission',
      grant: 'Grant',
      deny: 'Deny'
    },
    actions: {
      lock: 'Lock',
      unlock: 'Unlock',
      resetPassword: 'Reset Password',
      overrides: 'Permissions overrides',
      copyPassword: 'Copy Password',
      lockConfirm: 'Lock this user?',
      unlockConfirm: 'Unlock this user?'
    },
    passwordDialog: {
      createdTitle: 'User Created',
      createdMessage: 'User created successfully. Please save the temporary password below.',
      resetTitle: 'Password Reset',
      resetMessage: 'Password has been reset successfully. Please provide this temporary password to the user.',
      mustChangeNext: 'User must change this password on next login.',
      copied: 'Password copied to clipboard',
      confirmReset: 'A temporary password will be created. Continue?'
    },
    messages: {
      deleteConfirm: 'Delete this user?',
      deleteSuccess: 'User deleted successfully',
      deleteFailed: 'Deleted user failed',
      createSuccess: 'User created successfully',
      createFailed: 'User created failed',
      updateSuccess: 'User updated successfully',
      updateFailed: 'User updated failed',
      lockSuccess: 'User locked successfully',
      lockFailed: 'User locked failed',
      unlockSuccess: 'User unlocked successfully',
      unlockFailed: 'User unlocked failed',
      resetSuccess: 'Password reset successfully',
      resetFailed: 'Password reset failed',
      overrideSuccess: 'Override saved successfully',
      overrideFailed: 'Update override failed',
      overrideDeleteConfirm: 'Remove override ?',
      overrideDeleteSuccess: 'Override removed successfully',
      overrideDeleteFailed: 'Delete override failed'
    }
  }
};
