package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	timezone := os.Getenv("DATABASE_TIMEZONE")
	poolMaxStr := os.Getenv("DATABASE_POOL_MAX")

	poolMax, err := strconv.Atoi(poolMaxStr)
	if err != nil || poolMax <= 0 {
		poolMax = 5 // valor por defecto
	}

	// Construcción del DSN
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=%s",
		user, pass, host, port, name, timezone,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Error al parsear configuración: %v", err)
	}

	cfg.MinConns = 1
	cfg.MaxConns = int32(poolMax)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Error al crear el pool: %v", err)
	}

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	log.Printf("Conectado a PostgreSQL en %s:%s (pool máx: %d)\n", host, port, poolMax)
}
