package db

import (
	"context"
	"fmt"
	"os"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func NewDB(databaseUrl string) (*ksql.DB, error) {
	db, err := kpgx.New(context.Background(), databaseUrl, ksql.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &db, nil
}
