package postgres

import (
	"context"
	"fmt"
	"github.com/husanmusa/med-appointment-service/config"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	cfg := config.Load()
	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, "disable")

	db, err = pgxpool.Connect(context.Background(), conStr)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())

}
