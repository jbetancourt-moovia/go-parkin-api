package repositories

import (
	"context"
	"fmt"
	"go-api-swagger/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionsRepository struct {
	db *pgxpool.Pool
}

func NewTransactionsRepository(db *pgxpool.Pool) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}

func (r *TransactionsRepository) Create(ctx context.Context, c *models.StartTiming) error {

	query := `SELECT id FROM vehicles WHERE placa = $1 AND status = TRUE`
	var placa_id string
	err := r.db.QueryRow(ctx, query, c.Placa).Scan(&placa_id)
	if err != nil {
		return fmt.Errorf("la placa especificada no se encuentra registrada")
	}

	query = `SELECT id FROM customers WHERE dni = $1`
	var client_id int
	err = r.db.QueryRow(ctx, query, c.CustomerDNI).Scan(&client_id)
	if err != nil {
		client_id = 0
	}

	query = `
		INSERT INTO transactions (vehicle_id, client_id)
		VALUES ($1, $2)
	`
	_, err = r.db.Exec(ctx, query, placa_id, client_id)
	if err != nil {
		return fmt.Errorf("error al crear la transacción: %v", err)
	}
	return nil
}

func (r *TransactionsRepository) GetTransactionInfoByID(ctx context.Context, id int) (*models.PaymentInfo, error) {
	query := `SELECT id FROM transactions WHERE id = $1`
	var transactionID int64
	err := r.db.QueryRow(ctx, query, id).Scan(&transactionID)
	if err != nil {
		return nil, fmt.Errorf("transacción no encontrada")
	}

	query = `UPDATE transactions 
		SET end_dt = NOW(),
		minutes = FLOOR(EXTRACT(EPOCH FROM (end_dt - start_dt))/60)::int
		WHERE id = $1 
		RETURNING id, start_dt, end_dt, minutes
	`
	var info models.PaymentInfo
	err = r.db.QueryRow(ctx, query, id).Scan(&info.ID, &info.StartsAt, &info.EndsAt, &info.Minutes)
	if err != nil {
		return nil, fmt.Errorf("error al obtener información de la transacción: %v", err)
	}

	query = `SELECT t.id, v.placa, c.dni customer_dni, 
		c.first_name, c.last_name, c.email, c.phone,
		t.start_dt, t.end_dt, t.minutes,
		CASE 
			WHEN v.tipo = 0 THEN 45
			ELSE 95 
		END AS fee,
		CASE 
			WHEN v.tipo = 0 THEN t.minutes * 45
			ELSE t.minutes * 95
		END AS total_due
		FROM transactions t 
		INNER JOIN customers c ON c.id = t.client_id 
		INNER JOIN vehicles v ON v.id = t.vehicle_id
		WHERE t.id = $1
	`
	err = r.db.QueryRow(ctx, query, id).Scan(
		&info.ID,
		&info.Placa,
		&info.CustomerDNI,
		&info.FirstName,
		&info.LastName,
		&info.Email,
		&info.Phone,
		&info.StartsAt,
		&info.EndsAt,
		&info.Minutes,
		&info.Fee,
		&info.TotalDue,
	)
	if err != nil {
		return nil, fmt.Errorf("error al obtener información de la transacción: %v", err)
	}

	return &info, nil
}
