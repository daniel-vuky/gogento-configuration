package sqlc

import (
	"context"
	db "github.com/daniel-vuky/gogento-configuration/db/sqlc"
	"github.com/daniel-vuky/gogento-configuration/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testStore db.Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	testStore = db.NewStore(pool)
	os.Exit(m.Run())
}
