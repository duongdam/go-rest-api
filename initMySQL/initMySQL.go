package initMySQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func InitMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "[DatabaseConnect]:[Password]@tcp(Host)/[Database]?parseTime=true")
	if err != nil {
		return nil, errors.Wrap(err, "error initializing mysql")
	}
	return db, nil
}
