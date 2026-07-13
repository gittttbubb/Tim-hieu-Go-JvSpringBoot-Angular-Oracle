export const vi = {
  common: {
    save: 'Lưu',
    cancel: 'Hủy',
    edit: 'Chỉnh sửa',
    delete: 'Xóa',
    create: 'Tạo mới',
    status: 'Trạng thái',
    actions: 'Hành động',
    success: 'Thành công',
    warning: 'Cảnh báo',
    error: 'Lỗi',
    loading: 'Đang tải...',
    search: 'Tìm kiếm',
    confirm: 'Xác nhận'
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
  }
};
