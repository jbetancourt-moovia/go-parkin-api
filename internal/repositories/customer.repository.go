package repositories

import (
	"context"
	"errors"
	"go-api-swagger/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll(ctx context.Context) (*[]models.Customer, error) {
	rows, err := r.db.Query(ctx, `
		SELECT 
		id, first_name, last_name, phone, email, username 
		FROM customers ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Email, &c.Username); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return &customers, nil
}

func (r *CustomerRepository) Create(ctx context.Context, c *models.CustomerCreate) error {
	query := `INSERT INTO customers (first_name, last_name, phone, email, dni)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, c.FirstName, c.LastName, c.Phone, c.Email, c.DNI)
	return err
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	var c models.Customer
	err := r.db.QueryRow(ctx, `SELECT id, first_name, last_name, phone, email, username FROM customers WHERE id = $1`, id).
		Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Email, &c.Username)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	af, err := r.db.Exec(ctx, `DELETE FROM customers WHERE id = $1`, id)
	if af.RowsAffected() == 0 {
		return errors.New("cliente no encontrado")
	}
	return err
}
