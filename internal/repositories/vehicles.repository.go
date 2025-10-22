package repositories

import (
	"context"
	"errors"
	"go-api-swagger/internal/models"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) *VehicleRepository {
	return &VehicleRepository{db: db}
}

func (r *VehicleRepository) GetAll(ctx context.Context) (*[]models.Vehicle, error) {
	rows, err := r.db.Query(ctx, `SELECT placa, 
	CASE WHEN tipo = 1 THEN 'Carro' ELSE 'Moto' END as tipo, 
	marca, modelo, status, created_at 
	FROM vehicles WHERE status IS true ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {
		var v models.Vehicle
		if err := rows.Scan(&v.Placa, &v.Tipo, &v.Marca, &v.Modelo, &v.Status, &v.CreatedAt); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}
	return &vehicles, nil
}

func (r *VehicleRepository) Create(ctx context.Context, c *models.VehicleCreate) error {
	c.Tipo = strings.ToLower(c.Tipo)
	var tipoInt int
	if c.Tipo == "carro" {
		tipoInt = 1
	} else {
		tipoInt = 0
	}
	query := `
		INSERT INTO vehicles (placa, tipo, marca, modelo, status)
		VALUES ($1, $2, $3, $4, TRUE)
		ON CONFLICT (placa)
		DO UPDATE SET
		tipo = EXCLUDED.tipo,
		marca = EXCLUDED.marca,
		modelo = EXCLUDED.modelo,
		status = TRUE
		WHERE vehicles.status = FALSE
	`
	_, err := r.db.Exec(ctx, query, c.Placa, tipoInt, c.Marca, c.Modelo)
	return err
}

func (r *VehicleRepository) GetByPlate(ctx context.Context, placa string) (*models.Vehicle, error) {
	var v models.Vehicle
	err := r.db.QueryRow(ctx, `SELECT 
		CASE WHEN tipo = 1 THEN 'Carro' ELSE 'Moto' END as tipo, 
		marca, modelo, placa FROM vehicles WHERE placa = $1 AND status IS true`, placa).
		Scan(&v.Tipo, &v.Marca, &v.Modelo, &v.Placa)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *VehicleRepository) Delete(ctx context.Context, placa string) error {
	af, err := r.db.Exec(ctx, `UPDATE vehicles SET status = false WHERE placa = $1`, placa)
	if af.RowsAffected() == 0 {
		return errors.New("vehículo no encontrado")
	}
	return err
}
