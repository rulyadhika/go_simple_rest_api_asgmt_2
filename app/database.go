package app

import (
	"database/sql"
	"fmt"

	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/helper"

	_ "github.com/lib/pq"
)

func InitiateDB() *sql.DB {
	appConfig := GetAppConfig()

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", appConfig.DB_HOST, appConfig.DB_PORT, appConfig.DB_USERNAME, appConfig.DB_PASSWORD, appConfig.DB_DATABASE)

	db, err := sql.Open(appConfig.DB_DRIVER, dbInfo)

	helper.PanicIfErr(err)

	// defer db.Close()

	err = db.Ping()

	helper.PanicIfErr(err)

	createTables(db)

	helper.PanicIfErr(err)

	return db
}

func createTables(db *sql.DB) {
	orderTable := `
		CREATE TABLE IF NOT EXISTS "orders" (
			order_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			ordered_at timestamptz DEFAULT now(),
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	itemTable := `
		CREATE TABLE IF NOT EXISTS "items" (
			item_id SERIAL PRIMARY KEY,
			item_code VARCHAR(191) NOT NULL,
			quantity int NOT NULL,
			description TEXT NOT NULL,
			order_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT items_order_id_fk
            FOREIGN KEY(order_id) 
                REFERENCES orders(order_id)
                    ON DELETE CASCADE
		);	
	`

	_, err := db.Exec(orderTable)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(itemTable)

	if err != nil {
		panic(err)
	}
}
