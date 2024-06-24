package storage

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var postgresRepository *PostgresCoreConfigRepository

func TestMain(m *testing.M) {
	loadedConfig, err := config.LoadConfig("../../../")
	if err != nil {
		log.Fatalf("cannot load configuration: %v", err)
	}
	DBSource := loadedConfig.GetDatabaseSource()
	log.Print(DBSource)
	pool, err := pgxpool.New(context.Background(), DBSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	postgresRepository = NewRepository(pool)
	os.Exit(m.Run())
}
