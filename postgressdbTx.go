package gopostgres

import (
	"database/sql"
	"log"
)

// PgDbTx PgDbTx
type PgDbTx struct {
	Tx *sql.Tx
}

// Insert Insert
func (t *PgDbTx) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	//var stmtIns *sql.Stmt
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Insert transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		var lastInsertID int64
		err := stmtIns.QueryRow(args...).Scan(&lastInsertID)
		if err != nil {
			log.Println("Insert transaction Exec err:", err.Error())
		} else {
			//id, err = res.LastInsertId()
			//affectedRows, _ := res.RowsAffected()
			if err == nil && lastInsertID > 0 {
				id = lastInsertID
				success = true
			}
		}
	}
	return success, id
}

// Update Update
func (t *PgDbTx) Update(query string, args ...any) bool {
	var success = false
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Update transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Update transaction Exec err:", err.Error())
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

// Delete Delete
func (t *PgDbTx) Delete(query string, args ...any) bool {
	var success = false
	stmtIns, err := t.Tx.Prepare(query)
	if err != nil {
		log.Println("Delete transaction prepare err:", err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Delete transaction Exec err:", err.Error())
			//t.Tx.Rollback()
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records deleted")
			} else {
				success = true
			}
			//t.Tx.Commit()
		}
	}
	return success
}

// Commit Commit
func (t *PgDbTx) Commit() bool {
	var rtn = false
	err := t.Tx.Commit()
	if err != nil {
		log.Println("Commit transaction err:", err.Error())
	} else {
		rtn = true
	}
	return rtn
}



// Rollback Rollback
func (t *PgDbTx) Rollback() bool {
	var rtn = false
	err := t.Tx.Rollback()
	if err != nil {
		log.Println("Rollback transaction err:", err.Error())
	} else {
		rtn = true
	}
	return rtn
}
