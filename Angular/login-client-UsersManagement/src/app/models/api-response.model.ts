export interface ApiResponse<T> {
  success: boolean;
  data: T;
}

export interface ApiErrorResponse {
  success: boolean;
  message: string;
}

export interface PaginationResponse<T> {
    items: T[];
    total: number;
    page: number;
    pageSize: number;
}