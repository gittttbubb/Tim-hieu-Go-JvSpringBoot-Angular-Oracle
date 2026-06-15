package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type PasswordResetRepository interface {
	Create(token *model.PasswordResetToken,) error
	GetByTokenHash(hash string,) (*model.PasswordResetToken,error)
	MarkUsed(id string,) error
	Revoke(id string) error
}

type passwordResetRepository struct {
	db *sql.DB
}

func NewPasswordResetRepository(db *sql.DB) PasswordResetRepository {
	return &passwordResetRepository{
		db: db,
	}
}
func (r *passwordResetRepository) Create(token *model.PasswordResetToken,) error {
	query := `
	INSERT INTO password_reset_tokens(
		id,
		user_id,
		token_hash,
		expires_at,
		created_ip,
		user_agent
	) VALUES(:1,:2,:3,:4,:5,:6)`

	_, err := r.db.Exec(
		query,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
		token.CreatedIP,
		token.UserAgent,
	)
	return err
}

func (r *passwordResetRepository) GetByTokenHash(hash string,) (*model.PasswordResetToken,error) {
	query := `
	SELECT
		id,
		user_id,
		token_hash,
		expires_at,
		used_at,
		revoked_at,
		created_ip,
		user_agent,
		created_at
	FROM password_reset_tokens WHERE token_hash = :1`

	var token model.PasswordResetToken
	err := r.db.QueryRow(query, hash,).Scan(
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
	return &token,nil
}

func (r *passwordResetRepository) MarkUsed(id string,) error {
	_, err := r.db.Exec(`UPDATE password_reset_tokens SET used_at = SYSTIMESTAMP WHERE id = :1`, id,)
	return err
}

func (r *passwordResetRepository) Revoke(id string,) error {
	query := `UPDATE password_reset_tokens SET revoked_at = SYSTIMESTAMP WHERE id = :1`
	_, err := r.db.Exec(query, id)
	return err
}