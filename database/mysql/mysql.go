package mysql

import (
	"database/sql"
	"example/dao"
	"example/env"

	_ "github.com/go-sql-driver/mysql"
)

type handler struct {
	db *sql.DB
}

//func New() database.Manager {
//	return &handler{}
//}

//func (hdr *handler) Connect(config *env.DatabaseConfig) error {
//	fmt.Println("connect to mysql server at:", config.Address)
//	return nil
//}

func (hdr *handler) GetPlayer(username string) *dao.Player {
	hdr.db.Query("select * from players")
	return nil
}

func New(cfg *env.DatabaseConfig) (*handler, error) {
	// connect to host
	db, _ := sql.Open("mysql", "user:password@/dbname")
	return &handler{
		db: db,
	}, nil
}
