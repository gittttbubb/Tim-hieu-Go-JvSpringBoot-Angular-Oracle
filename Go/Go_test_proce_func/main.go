package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// Import the go-ora package and name it go_ora to access go_ora.Out
	go_ora "github.com/sijms/go-ora/v2"
)

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Configure Oracle connection settings
	username := getEnv("DB_USER", "system")
	password := getEnv("DB_PASSWORD", "vthang2003")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "1521")
	serviceName := getEnv("DB_SERVICE", "xe")

	// Build the Oracle connection URL for go-ora
	dsn := getEnv("DB_DSN", fmt.Sprintf("oracle://%s:%s@%s:%s/%s", username, password, host, port, serviceName))

	log.Printf("Connecting to Oracle database on %s:%s/%s as user %s...\n", host, port, serviceName, username)

	// Open database connection
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}
	defer db.Close()

	// Set connection pool limits
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(1 * time.Hour)

	// Verify connection using Ping
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
	log.Println("Successfully connected to Oracle Database using go-ora!")

	// -------------------------------------------------------------
	// 1. Initialize Schema from schema.sql
	// -------------------------------------------------------------
	fmt.Println("\n--- [1] Initializing Database Schema ---")
	err = initSchema(db)
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v\n", err)
	}

	// Create test context for queries
	runCtx, runCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer runCancel()

	// -------------------------------------------------------------
	// 2. Call Procedure with OUT Parameter: add_user
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
		// :3 -> go_ora.Out{Dest: &newID} (OUT) - sizes are not required for numeric outputs
		query := "BEGIN add_user(:1, :2, :3); END;"
		
		_, err = db.ExecContext(runCtx, query,
			u.username,
			u.email,
			go_ora.Out{Dest: &newID},
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
	// 3. Call Procedure with Multiple OUT Parameters: get_user_by_id
	// -------------------------------------------------------------
	fmt.Println("\n--- [3] Calling Procedure: get_user_by_id (with multiple OUT Parameters) ---")
	{
		var outUsername string
		var outEmail string

		// For string OUT parameters in go-ora, we must specify the Size (in bytes).
		// If we don't, the driver doesn't know how much memory to reserve, causing ORA-06502.
		// :1 -> firstUserID (IN)
		// :2 -> go_ora.Out{Dest: &outUsername, Size: 50} (OUT)
		// :3 -> go_ora.Out{Dest: &outEmail, Size: 100} (OUT)
		query := "BEGIN get_user_by_id(:1, :2, :3); END;"
		
		_, err = db.ExecContext(runCtx, query,
			firstUserID,
			go_ora.Out{Dest: &outUsername, Size: 50},
			go_ora.Out{Dest: &outEmail, Size: 100},
		)
		if err != nil {
			log.Fatalf("Failed to call get_user_by_id procedure: %v\n", err)
		}
		fmt.Printf("Retrieved details for User ID %d:\n", firstUserID)
		fmt.Printf("  - Username: %s\n", outUsername)
		fmt.Printf("  - Email:    %s\n", outEmail)
	}

	// -------------------------------------------------------------
	// 4. Call Function returning a value: get_total_users
	// -------------------------------------------------------------
	fmt.Println("\n--- [4] Calling Function: get_total_users ---")
	
	// Method A: Using SELECT ... FROM DUAL (Standard SQL approach for functions)
	{
		var count int
		query := "SELECT get_total_users() FROM dual"
		err = db.QueryRowContext(runCtx, query).Scan(&count)
		if err != nil {
			log.Fatalf("Failed to call get_total_users via SELECT: %v\n", err)
		}
		fmt.Printf("Method A (SELECT FROM dual) - Total users: %d\n", count)
	}

	// Method B: Using PL/SQL Block binding the return value
	{
		var count int
		// :1 -> go_ora.Out{Dest: &count} (OUT return value)
		query := "BEGIN :1 := get_total_users(); END;"
		
		_, err = db.ExecContext(runCtx, query,
			go_ora.Out{Dest: &count},
		)
		if err != nil {
			log.Fatalf("Failed to call get_total_users via PL/SQL: %v\n", err)
		}
		fmt.Printf("Method B (PL/SQL Out Bind) - Total users: %d\n", count)
	}

	// -------------------------------------------------------------
	// 5. Call Function with Arguments returning a value: concat_user_info
	// -------------------------------------------------------------
	fmt.Println("\n--- [5] Calling Function: concat_user_info (with Argument and Return Value) ---")
	
	// Method A: Using SELECT ... FROM DUAL
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

	// Method B: Using PL/SQL block
	{
		var userInfo string
		// Since this is a string return value in PL/SQL block execution, we must specify Size.
		// :1 -> go_ora.Out{Dest: &userInfo, Size: 200} (OUT return value)
		// :2 -> firstUserID (IN argument)
		query := "BEGIN :1 := concat_user_info(:2); END;"
		
		_, err = db.ExecContext(runCtx, query,
			go_ora.Out{Dest: &userInfo, Size: 200},
			firstUserID,
		)
		if err != nil {
			log.Fatalf("Failed to call concat_user_info via PL/SQL: %v\n", err)
		}
		fmt.Printf("Method B (PL/SQL Out Bind) - Info for ID %d: %s\n", firstUserID, userInfo)
	}

	fmt.Println("\n--- All tests completed successfully! ---")
}

// initSchema reads schema.sql, splits it by '/' lines, and executes each SQL block.
func initSchema(db *sql.DB) error {
	content, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	// Normalize windows \r\n to \n
	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	var blocks []string
	var currentBlock []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// If line is just "/", it signifies the end of the current PL/SQL or SQL block
		if trimmed == "/" {
			if len(currentBlock) > 0 {
				blocks = append(blocks, strings.Join(currentBlock, "\n"))
				currentBlock = nil
			}
		} else {
			currentBlock = append(currentBlock, line)
		}
	}
	// Append remaining query if any
	if len(currentBlock) > 0 {
		leftover := strings.TrimSpace(strings.Join(currentBlock, "\n"))
		if leftover != "" {
			blocks = append(blocks, leftover)
		}
	}

	// Execute each block sequentially
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
