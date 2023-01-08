package adapters

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGClient struct {
	db *pgxpool.Pool
}

const selectPrice = "SELECT price FROM price ORDER BY created_at DESC LIMIT 1;"
const insertPrice = "INSERT INTO price (created_at, price) VALUES(now(), $1);"

func NewPGClient() (*PGClient, error) {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &PGClient{db: db}, nil
}

func (c *PGClient) GetPrice(ctx context.Context) (int, error) {
	var price int
	err := c.db.QueryRow(ctx, selectPrice).Scan(&price)
	if err == pgx.ErrNoRows {
		return 0, nil
	}
	return price, err
}

func (c *PGClient) InsertPrice(ctx context.Context, price int) error {
	row, _ := c.db.Query(ctx, insertPrice, price)
	defer row.Close()
	return row.Err()
}
