export const vi = {
  common: {
    save: 'Lưu',
    cancel: 'Hủy',
    edit: 'Chỉnh sửa',
    update: 'Cập nhật',
    delete: 'Xóa',
    create: 'Tạo mới',
    status: 'Trạng thái',
    actions: 'Hành động',
    success: 'Thành công',
    warning: 'Cảnh báo',
    error: 'Lỗi',
    loading: 'Đang tải...',
    search: 'Tìm kiếm',
    confirm: 'Xác nhận',
    back: 'Quay lại'
  },
  sidebar: {
    dashboard: 'Bảng điều khiển',
    users: 'Người dùng',
    roles: 'Vai trò',
    permissions: 'Quyền',
    audits: 'Nhật ký hệ thống',
    title: 'Quản lý người dùng'
  },
  topbar: {
    changePassword: 'Đổi mật khẩu',
    logout: 'Đăng xuất',
    toggleTheme: 'Chuyển giao diện',
    welcome: 'Chào mừng'
  },
  auth: {
    login: {
      title: 'Đăng nhập',
      subtitle: 'Quản lý người dùng',
      username: 'Tên đăng nhập',
      password: 'Mật khẩu',
      usernamePlaceholder: 'Nhập tài khoản',
      passwordPlaceholder: 'Nhập mật khẩu',
      submit: 'Đăng nhập',
      forgotPasswordLink: 'Quên mật khẩu?',
      success: 'Đăng nhập thành công',
      pendingPasswordChange: 'Bạn phải đổi mật khẩu lần đầu đăng nhập',
      failed: 'Sai tài khoản hoặc mật khẩu'
    },
    forgotPassword: {
      title: 'Quên mật khẩu',
      subtitle: 'Nhập email để nhận link đặt lại mật khẩu',
      email: 'Email',
      emailPlaceholder: 'Nhập email của bạn',
      submit: 'Gửi link đặt lại mật khẩu',
      backToLogin: 'Đăng nhập',
      success: 'Link khôi phục mật khẩu đã được gửi tới email của bạn.'
    },
    resetPassword: {
      title: 'Đặt lại mật khẩu',
      newPassword: 'Mật khẩu mới',
      confirmPassword: 'Xác nhận mật khẩu',
      submit: 'Đổi mật khẩu',
      backToLogin: 'Quay lại đăng nhập',
      invalidLink: 'Token khôi phục mật khẩu không hợp lệ.',
      success: 'Cập nhật mật khẩu thành công.',
      mismatch: 'Mật khẩu không khớp',
      validation: {
        required: 'Mật khẩu mới không được để trống',
        minlength: 'Mật khẩu phải có ít nhất 8 ký tự',
        maxlength: 'Mật khẩu không được vượt quá 100 ký tự',
        strong: 'Mật khẩu phải chứa: chữ cái in hoa, chữ cái in thường, số và ký tự đặc biệt',
        confirmRequired: 'Xác nhận mật khẩu không khớp'
      }
    },
    changePassword: {
      title: 'Đổi mật khẩu',
      oldPassword: 'Mật khẩu hiện tại',
      newPassword: 'Mật khẩu mới',
      confirmPassword: 'Xác nhận mật khẩu mới',
      submit: 'Cập nhật mật khẩu',
      success: 'Đổi mật khẩu thành công.',
      oldPasswordRequired: 'Mật khẩu hiện tại không được để trống',
      sameAsOld: 'Mật khẩu mới không được giống với mật khẩu cũ',
      failed: 'Đổi mật khẩu thất bại',
      mismatch: 'Mật khẩu mới không khớp với mật khẩu xác nhận'
    }
  },
  errors: {
    sessionExpired: 'Phiên đăng nhập đã hết hạn',
    accessDenied: 'Bạn không có quyền truy cập chức năng này',
    serverError: 'Đã xảy ra lỗi hệ thống',
    networkError: 'Không thể kết nối tới máy chủ',
    unexpected: 'Đã xảy ra lỗi không xác định'
  },
  dashboard: {
    welcome: 'Bảng điều khiển hoạt động!'
  },
  audits: {
    title: 'Nhật ký hoạt động',
    detailTitle: 'Chi tiết nhật ký',
    subtitle: 'Thông tin hoạt động',
    searchPlaceholder: 'Tìm hành động, người thực hiện, đối tượng...',
    noLogs: 'Không tìm thấy nhật ký hoạt động nào',
    fields: {
      action: 'Hành động',
      actor: 'Người thực hiện',
      actorId: 'Mã người thực hiện',
      entityType: 'Loại đối tượng',
      entityId: 'Mã đối tượng',
      targetUser: 'Người dùng mục tiêu',
      ipAddress: 'Địa chỉ IP',
      timestamp: 'Thời gian',
      reason: 'Lý do',
      beforeData: 'Dữ liệu trước',
      afterData: 'Dữ liệu sau',
      metadata: 'Siêu dữ liệu'
    }
  },
  permissions: {
    title: 'Quyền hạn',
    searchPlaceholder: 'Tìm kiếm quyền...',
    noPermissions: 'Không tìm thấy quyền nào',
    fields: {
      group: 'Nhóm chức năng',
      code: 'Mã chức năng',
      action: 'Hành động',
      description: 'Mô tả'
    }
  },
  roles: {
    title: 'Vai trò',
    searchPlaceholder: 'Tìm kiếm vai trò...',
    createRole: 'Tạo vai trò',
    editRole: 'Chỉnh sửa vai trò',
    roleDetail: 'Chi tiết vai trò',
    rolePermissions: 'Phân quyền vai trò',
    fields: {
      name: 'Tên vai trò',
      displayName: 'Tên hiển thị',
      description: 'Mô tả',
      granted: 'Được gán',
      scope: 'Phạm vi'
    },
    validation: {
      nameRequired: 'Tên vai trò không được để trống',
      displayNameRequired: 'Tên hiển thị không được để trống'
    },
    messages: {
      confirmDelete: 'Bạn có chắc chắn muốn xóa vai trò "{name}"?',
      deleteSuccess: 'Xóa vai trò thành công',
      deleteFailed: 'Xóa vai trò thất bại',
      createSuccess: 'Tạo vai trò thành công',
      createFailed: 'Tạo vai trò thất bại',
      updateSuccess: 'Cập nhật vai trò thành công',
      updateFailed: 'Cập nhật vai trò thất bại',
      permissionUpdated: 'Cập nhật quyền thành công',
      permissionUpdateFailed: 'Cập nhật quyền thất bại',
      permissionRemoved: 'Gỡ bỏ quyền thành công',
      permissionRemoveFailed: 'Gỡ bỏ quyền thất bại',
      confirmRemovePermission: 'Gỡ bỏ quyền "{permission}" khỏi vai trò này?'
    }
  },
  users: {
    statusOptions: {
      active: 'Hoạt động',
      locked: 'Đã khóa',
      pending: 'Chờ đổi mật khẩu'
    },
    title: 'Quản lý người dùng',
    searchPlaceholder: 'Tìm kiếm người dùng...',
    createUser: 'Tạo người dùng',
    editUser: 'Chỉnh sửa người dùng',
    userDetail: 'Chi tiết người dùng',
    addUser: 'Thêm người dùng',
    permissionOverrides: 'Quyền ghi đè người dùng',
    addOverride: 'Thêm quyền ghi đè',
    assignOverride: 'Thiết lập quyền ghi đè',
    fields: {
      fullName: 'Họ và tên',
      username: 'Tên đăng nhập',
      email: 'Email',
      phone: 'Số điện thoại',
      role: 'Vai trò',
      selectRole: 'Chọn vai trò',
      status: 'Trạng thái',
      password: 'Mật khẩu',
      mustChange: 'Bắt buộc đổi',
      created: 'Thời gian tạo',
      updated: 'Thời gian cập nhật',
      type: 'Loại',
      reason: 'Lý do',
      dataScope: 'Phạm vi dữ liệu',
      grantPermission: 'Cho phép quyền',
      permission: 'Quyền',
      grant: 'Cho phép',
      deny: 'Từ chối'
    },
    actions: {
      lock: 'Khóa',
      unlock: 'Mở khóa',
      resetPassword: 'Đặt lại mật khẩu',
      overrides: 'Quyền ghi đè',
      copyPassword: 'Sao chép mật khẩu',
      lockConfirm: 'Khóa người dùng này?',
      unlockConfirm: 'Mở khóa người dùng này?'
    },
    passwordDialog: {
      createdTitle: 'Tạo người dùng',
      createdMessage: 'Tạo người dùng thành công. Vui lòng lưu mật khẩu tạm thời bên dưới.',
      resetTitle: 'Đặt lại mật khẩu',
      resetMessage: 'Đặt lại mật khẩu thành công. Vui lòng cung cấp mật khẩu tạm thời bên dưới cho người dùng.',
      mustChangeNext: 'Người dùng phải thay đổi mật khẩu này trong lần đăng nhập tiếp theo.',
      copied: 'Đã sao chép mật khẩu vào bộ nhớ tạm',
      confirmReset: 'Một mật khẩu tạm thời sẽ được tạo. Tiếp tục?'
    },
    messages: {
      deleteConfirm: 'Xóa người dùng này?',
      deleteSuccess: 'Xóa người dùng thành công',
      deleteFailed: 'Xóa người dùng thất bại',
      createSuccess: 'Tạo người dùng thành công',
      createFailed: 'Tạo người dùng thất bại',
      updateSuccess: 'Cập nhật người dùng thành công',
      updateFailed: 'Cập nhật người dùng thất bại',
      lockSuccess: 'Khóa người dùng thành công',
      lockFailed: 'Khóa người dùng thất bại',
      unlockSuccess: 'Mở khóa người dùng thành công',
      unlockFailed: 'Mở khóa người dùng thất bại',
      resetSuccess: 'Đặt lại mật khẩu thành công',
      resetFailed: 'Đặt lại mật khẩu thất bại',
      overrideSuccess: 'Lưu quyền ghi đè thành công',
      overrideFailed: 'Cập nhật quyền ghi đè thất bại',
      overrideDeleteConfirm: 'Gỡ bỏ quyền ghi đè này?',
      overrideDeleteSuccess: 'Gỡ bỏ quyền ghi đè thành công',
      overrideDeleteFailed: 'Gỡ bỏ quyền ghi đè thất bại'
    }
  }
};
