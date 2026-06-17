package repository

import (
	"database/sql"
	"time"

	"go-rbac-system/internal/model"
)

type PasswordResetRepository interface {
	GetByID(id string) (*model.PasswordResetToken, error)
	GetByTokenHash(tokenHash string,) (*model.PasswordResetToken, error)
	GetByUserID(userID string,) ([]model.PasswordResetToken, error)
	Create(token *model.PasswordResetToken,) error
	MarkUsed(id string, usedAt time.Time,) error
	Revoke(id string, revokedAt time.Time,) error
}

type passwordResetRepository struct {
	db *sql.DB
}

func NewPasswordResetRepository(db *sql.DB,) PasswordResetRepository {
	return &passwordResetRepository{
		db: db,
	}
}

func scanPasswordResetToken(scanner interface {Scan(dest ...any) error},) (*model.PasswordResetToken, error) {
	var token model.PasswordResetToken
	err := scanner.Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.UsedAt,
		&token.RevokedAt,
		&token.CreatedIP,
		&token.UserAgent,
		&token.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *passwordResetRepository) GetByID(id string,) (*model.PasswordResetToken, error) {
	query := `SELECT id, user_id, token_hash, expires_at, used_at, revoked_at, created_ip, user_agent, created_at
		FROM password_reset_tokens WHERE id = :1`
	row := r.db.QueryRow(query, id,)
	return scanPasswordResetToken(row)
}

func (r *passwordResetRepository) GetByTokenHash(tokenHash string,) (*model.PasswordResetToken, error) {
	query := `SELECT id, user_id, token_hash, expires_at, used_at, revoked_at, created_ip, user_agent, created_at
		FROM password_reset_tokens WHERE token_hash = :1`
	row := r.db.QueryRow(query, tokenHash,)
	return scanPasswordResetToken(row)
}

func (r *passwordResetRepository) GetByUserID(userID string,) ([]model.PasswordResetToken, error) {
	query := `SELECT id, user_id, token_hash, expires_at, used_at, revoked_at, created_ip, user_agent, created_at
		FROM password_reset_tokens WHERE user_id = :1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.PasswordResetToken, 0,)
	for rows.Next() {
		item, err := scanPasswordResetToken(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item,)
	}
	return result, rows.Err()
}

func (r *passwordResetRepository) Create(token *model.PasswordResetToken,) error {
	query := `INSERT INTO password_reset_tokens (id, user_id, token_hash, expires_at, used_at, revoked_at, created_ip, user_agent, created_at)
		VALUES (:1, :2, :3, :4, :5, :6, :7,	:8, :9)`
	_, err := r.db.Exec(
		query,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
		token.UsedAt,
		token.RevokedAt,
		token.CreatedIP,
		token.UserAgent,
		token.CreatedAt,
	)
	return err
}

func (r *passwordResetRepository) MarkUsed(id string, usedAt time.Time,) error {
	query := `UPDATE password_reset_tokens SET used_at = :1 WHERE id = :2`
	_, err := r.db.Exec(query, usedAt, id,)
	return err
}

func (r *passwordResetRepository) Revoke(id string,revokedAt time.Time,) error {
	query := `UPDATE password_reset_tokens SET revoked_at = :1 WHERE id = :2`
	_, err := r.db.Exec(query, revokedAt, id,)
	return err
}