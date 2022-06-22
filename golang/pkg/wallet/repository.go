package wallet

import (
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func (repository *Repository) List() []Wallet {
	sql := `SELECT id, name, currency FROM "wallets" ORDER BY name;`
	rows, err := repository.db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	var wallets []Wallet
	for rows.Next() {
		var wallet Wallet
		err = rows.Scan(&wallet.Id, &wallet.Name, &wallet.Currency)
		if err != nil {
			log.Fatal(err)
		}

		wallets = append(wallets, wallet)
	}

	return wallets
}

func (repository *Repository) Save(wallet Wallet) {
	sql := `INSERT INTO "wallets" ("name", "currency") VALUES ($1, $2);`
	_, err := repository.db.Exec(sql, wallet.Name, wallet.Currency)
	if err != nil {
		log.Fatal(err)
	}
}
