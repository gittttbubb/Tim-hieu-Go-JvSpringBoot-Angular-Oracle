package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// Import package godror
	_ "github.com/godror/godror"
)

// getEnv lấy biến môi trường với giá trị mặc định dự phòng (fallback)
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Cấu hình các thiết lập kết nối Oracle
	username := getEnv("DB_USER", "system")
	password := getEnv("DB_PASSWORD", "vthang2003")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "1521")
	serviceName := getEnv("DB_SERVICE", "xe")

	// Xây dựng DSN kết nối Oracle cho godror dưới dạng logfmt
	dsn := getEnv("DB_DSN", fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"`, username, password, host, port, serviceName))

	log.Printf("Connecting to Oracle database on %s:%s/%s as user %s using godror...\n", host, port, serviceName, username)

	// Mở kết nối cơ sở dữ liệu
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}
	defer db.Close()

	// Thiết lập giới hạn cho connection pool (nhóm kết nối)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(1 * time.Hour)

	// Xác minh kết nối bằng cách sử dụng Ping
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Database connection failed: %v\n"+
			"Please verify:\n"+
			"1. Your local Oracle Database is running on localhost:%s\n"+
			"2. The credentials are correct (user: %s, password: %s)\n"+
			"3. The Service Name/SID matches '%s'\n", 
			err, port, username, password, serviceName,
		)
	}
	log.Println("Successfully connected to Oracle Database using godror!")

	// -------------------------------------------------------------
	// 1. Khởi tạo Schema từ schema.sql
	// -------------------------------------------------------------
	fmt.Println("\n--- [1] Initializing Database Schema ---")
	err = initSchema(db)
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v\n", err)
	}

	// Tạo context thử nghiệm cho các truy vấn
	runCtx, runCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer runCancel()

	// -------------------------------------------------------------
	// 2. Gọi Procedure với tham số OUT: add_user
	// -------------------------------------------------------------
	fmt.Println("\n--- [2] Calling Procedure: add_user (with OUT Parameter) ---")
	
	testUsers := []struct {
		username string
		email    string
	}{
		{"Nguyen Van A", "nva@example.com"},
		{"Tran Thi B", "ttb@example.com"},
		{"Le Van C", "lvc@example.com"},
	}

	var firstUserID int64

	for _, u := range testUsers {
		var newID int64
		// :1 -> u.username (IN)
		// :2 -> u.email (IN)
		// :3 -> sql.Out{Dest: &newID} (OUT)
		query := "BEGIN add_user(:1, :2, :3); END;"
		_, err = db.ExecContext(runCtx, query,
			u.username,
			u.email,
			sql.Out{Dest: &newID},
		)
		if err != nil {
			log.Fatalf("Failed to call add_user procedure: %v\n", err)
		}
		fmt.Printf("Successfully added user: %s (%s). Assigned ID: %d\n", u.username, u.email, newID)
		
		if firstUserID == 0 {
			firstUserID = newID
		}
	}

	// -------------------------------------------------------------
	// 3. Gọi Procedure với nhiều tham số OUT: get_user_by_id
	// -------------------------------------------------------------
	fmt.Println("\n--- [3] Calling Procedure: get_user_by_id (with multiple OUT Parameters) ---")
	{
		var outUsername string
		var outEmail string

		// :1 -> firstUserID (IN)
		// :2 -> sql.Out{Dest: &outUsername} (OUT)
		// :3 -> sql.Out{Dest: &outEmail} (OUT)
		query := "BEGIN get_user_by_id(:1, :2, :3); END;"
		
		_, err = db.ExecContext(runCtx, query,
			firstUserID,
			sql.Out{Dest: &outUsername},
			sql.Out{Dest: &outEmail},
		)
		if err != nil {
			log.Fatalf("Failed to call get_user_by_id procedure: %v\n", err)
		}
		fmt.Printf("Retrieved details for User ID %d:\n", firstUserID)
		fmt.Printf("  - Username: %s\n", outUsername)
		fmt.Printf("  - Email:    %s\n", outEmail)
	}

	// -------------------------------------------------------------
	// 4. Gọi Function trả về một giá trị: get_total_users
	// -------------------------------------------------------------
	fmt.Println("\n--- [4] Calling Function: get_total_users ---")
	
	// Cách A: Sử dụng SELECT ... FROM DUAL (cách tiếp cận SQL tiêu chuẩn cho các function)
	{
		var count int
		query := "SELECT get_total_users() FROM dual"
		err = db.QueryRowContext(runCtx, query).Scan(&count)
		if err != nil {
			log.Fatalf("Failed to call get_total_users via SELECT: %v\n", err)
		}
		fmt.Printf("Method A (SELECT FROM dual) - Total users: %d\n", count)
	}

	// Cách B: Sử dụng PL/SQL Block liên kết (bind) giá trị trả về
	{
		var count int
		// :1 -> sql.Out{Dest: &count} (Giá trị trả về OUT)
		query := "BEGIN :1 := get_total_users(); END;"
		
		_, err = db.ExecContext(runCtx, query,
			sql.Out{Dest: &count},
		)
		if err != nil {
			log.Fatalf("Failed to call get_total_users via PL/SQL: %v\n", err)
		}
		fmt.Printf("Method B (PL/SQL Out Bind) - Total users: %d\n", count)
	}

	// -------------------------------------------------------------
	// 5. Gọi Function có đối số trả về một giá trị: concat_user_info
	// -------------------------------------------------------------
	fmt.Println("\n--- [5] Calling Function: concat_user_info (with Argument and Return Value) ---")
	
	// Cách A: Sử dụng SELECT ... FROM DUAL
	{
		var userInfo string
		// :1 -> firstUserID (IN)
		query := "SELECT concat_user_info(:1) FROM dual"
		err = db.QueryRowContext(runCtx, query, firstUserID).Scan(&userInfo)
		if err != nil {
			log.Fatalf("Failed to call concat_user_info via SELECT: %v\n", err)
		}
		fmt.Printf("Method A (SELECT FROM dual) - Info for ID %d: %s\n", firstUserID, userInfo)
	}

	// Cách B: Sử dụng PL/SQL block
	{
		var userInfo string
		// :1 -> sql.Out{Dest: &userInfo} (Giá trị trả về OUT)
		// :2 -> firstUserID (Tham số đầu vào IN)
		query := "BEGIN :1 := concat_user_info(:2); END;"
		
		_, err = db.ExecContext(runCtx, query,
			sql.Out{Dest: &userInfo},
			firstUserID,
		)
		if err != nil {
			log.Fatalf("Failed to call concat_user_info via PL/SQL: %v\n", err)
		}
		fmt.Printf("Method B (PL/SQL Out Bind) - Info for ID %d: %s\n", firstUserID, userInfo)
	}

	fmt.Println("\n--- All tests completed successfully! ---")
}

// initSchema đọc file schema.sql, phân tách bằng các dòng "/", và thực thi từng khối lệnh SQL.
func initSchema(db *sql.DB) error {
	content, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	// Chuẩn hóa ký tự xuống dòng từ \r\n của Windows thành \n
	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	var blocks []string
	var currentBlock []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Nếu dòng chỉ chứa "/", nó đánh dấu sự kết thúc của khối PL/SQL hoặc SQL hiện tại
		if trimmed == "/" {
			if len(currentBlock) > 0 {
				blocks = append(blocks, strings.Join(currentBlock, "\n"))
				currentBlock = nil
			}
		} else {
			currentBlock = append(currentBlock, line)
		}
	}
	// Thêm truy vấn còn lại (nếu có)
	if len(currentBlock) > 0 {
		leftover := strings.TrimSpace(strings.Join(currentBlock, "\n"))
		if leftover != "" {
			blocks = append(blocks, leftover)
		}
	}

	// Thực thi tuần tự từng khối lệnh
	for i, block := range blocks {
		block = strings.TrimSpace(block)
		if block == "" {
			continue
		}
		
		snippet := strings.ReplaceAll(block, "\n", " ")
		if len(snippet) > 60 {
			snippet = snippet[:60] + "..."
		}
		fmt.Printf("Executing block %d: %s\n", i+1, snippet)

		_, err = db.Exec(block)
		if err != nil {
			return fmt.Errorf("failed to execute block %d (%s): %w", i+1, snippet, err)
		}
	}

	return nil
}
