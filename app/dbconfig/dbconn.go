
package dbconfig

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func InitDB() {

    var err error
    db , err := sql.Open("mysql", "dev:123456@tcp(127.0.0.1:3306)/Company")
    if err != nil {
       fmt.Println("DB Error", err)
    }
    DB=db
    fmt.Println("DB Connected")


}

