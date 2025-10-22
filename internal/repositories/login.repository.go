package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRepository struct {
	db *pgxpool.Pool
}

func NewLoginRepository(db *pgxpool.Pool) *LoginRepository {
	return &LoginRepository{db: db}
}

func (r *LoginRepository) CheckUser(ctx context.Context, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM customers WHERE username=$1)`
	err := r.db.QueryRow(ctx, query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *LoginRepository) GetUserPermissions(ctx context.Context, username string) ([]int, error) {
	var permissions []int
	query := `SELECT p.claim_id 
	FROM customers c
	INNER JOIN permissions p ON p.customer_id = c.id AND p.status IS TRUE
	WHERE c.username = $1`
	rows, err := r.db.Query(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission int
		if err := rows.Scan(&permission); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
