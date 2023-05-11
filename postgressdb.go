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

// New NewDatabase
func (p *PgDB) New() Database {
	return p
}

// Connect Connect
func (p *PgDB) Connect() bool {
	var rtn = false
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

// Test Test
func (p *PgDB) Test(query string, args ...any) *DbRow {
	return p.Get(query, args...)
}

// BeginTransaction BeginTransaction
func (p *PgDB) BeginTransaction() Transaction {
	var trans Transaction
	var pgTrans PgDbTx
	tx, err := p.db.Begin()
	if err == nil && tx != nil {
		pgTrans.Tx = tx
	}
	trans = &pgTrans
	return trans
}

// Insert Insert requires use of  RETURNING id on end of insert query
func (p *PgDB) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	var stmtIns *sql.Stmt
	stmtIns, p.err = p.db.Prepare(query)
	if p.err != nil {
		log.Println("Error:", p.err.Error())
	} else {
		defer stmtIns.Close()
		var lastInsertID int64
		err := stmtIns.QueryRow(args...).Scan(&lastInsertID)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			if lastInsertID > 0 {
				id = lastInsertID
				success = true
			}
			// success = true
		}
	}
	return success, id
}

// Update Update
func (p *PgDB) Update(query string, args ...any) bool {
	var success = false
	var stmtUp *sql.Stmt
	stmtUp, p.err = p.db.Prepare(query)
	if p.err != nil {
		log.Println("Error:", p.err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records updated")
			} else {
				success = true
			}
		}
	}
	return success
}

// Get Get
func (p *PgDB) Get(query string, args ...any) *DbRow {
	var rtn DbRow
	stmtGet, err := p.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		if err != nil {
			log.Println("Get err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					// err = rows.Scan(scanArgs...)
					// if err != nil {
					// 	log.Println("Error:", err.Error())
					// }
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rtn.Row = append(rtn.Row, value)
					}
				}
				// if err = rows.Err(); err != nil {
				// 	log.Println("Error:", err.Error())
				// }
			}
		}
	}
	return &rtn
}

// GetList GetList
func (p *PgDB) GetList(query string, args ...any) *DbRows {
	var rtn DbRows
	stmtGet, err := p.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		// defer rows.Close()
		if err != nil {
			log.Println("GetList err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					var rowValuesStr []string
					// err = rows.Scan(scanArgs...)
					// if err != nil {
					// 	log.Println("Error:", err.Error())
					// }
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rowValuesStr = append(rowValuesStr, value)
					}
					rtn.Rows = append(rtn.Rows, rowValuesStr)
				}
				// if err = rows.Err(); err != nil {
				// 	log.Println("Error:", err.Error())
				// }
			}
		}
	}
	return &rtn
}

// Delete Delete
func (p *PgDB) Delete(query string, args ...any) bool {
	var success = false
	var stmt *sql.Stmt
	stmt, p.err = p.db.Prepare(query)
	if p.err != nil {
		log.Println("Error:", p.err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records deleted")
			} else {
				success = true
			}
		}
	}
	return success
}

// Close Close
func (p *PgDB) Close() bool {
	var rtn = false
	err := p.db.Close()
	if err == nil {
		rtn = true
	}
	return rtn
}
