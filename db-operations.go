package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func connectToDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}

func insertNewRecord(connection *pgx.Conn, record Record) error {
	sqlStatement := `INSERT INTO records (description, link) VALUES ($1, $2) RETURNING id`
	err := connection.QueryRow(context.Background(), sqlStatement, record.Description, record.Link).Scan(&record.ID)
	if err != nil {
		return fmt.Errorf("failed to insert new records: %v", err)
	}
	fmt.Printf("Inserted record with ID %d\n", record.ID)

	return nil
}
