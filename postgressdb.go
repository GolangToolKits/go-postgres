package gopostgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// PgDB PgDB
type PgDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	db       *sql.DB
	err      error
}

// Connect Connect
func (p *PgDB) Connect() bool {
	var rtn = false
	// var conStr = p.User + ":" + p.Password + "@tcp(" + p.Host + ")/" + p.Database
	var conStr = "host=" + p.Host + " port=" + p.Port + " user=" + p.User + " password=" +
		p.Password + " dbname=" + p.Database + " sslmode=disable"
		log.Println("constr: ", conStr)

	p.db, p.err = sql.Open("postgres", conStr)
	if p.err == nil {
		p.err = p.db.Ping()
		if p.err != nil {
			log.Println("Database Connect Error:", p.err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}
