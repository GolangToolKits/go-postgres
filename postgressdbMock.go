package gopostgres

import "database/sql"

// PgDBMock PgDBMock
type PgDBMock struct {
	Host     string
	User     string
	Password string
	Database string
	db       *sql.DB
	err      error

	MockConnectSuccess  bool
	MockCloseSuccess    bool
	MockCommitSuccess   bool
	MockRollbackSuccess bool

	MockInsertSuccess1     bool
	MockInsertSuccess2     bool
	MockInsertSuccess3     bool
	MockInsertSuccess4     bool
	MockInsertSuccess5     bool
	MockInsertSuccess6     bool
	MockInsertSuccess7     bool
	MockInsertSuccess8     bool
	mockInsertSuccess1Used bool
	mockInsertSuccess2Used bool
	mockInsertSuccess3Used bool
	mockInsertSuccess4Used bool
	mockInsertSuccess5Used bool
	mockInsertSuccess6Used bool
	mockInsertSuccess7Used bool
	mockInsertSuccess8Used bool
	MockInsertID1          int64
	MockInsertID2          int64
	MockInsertID3          int64
	MockInsertID4          int64
	MockInsertID5          int64
	MockInsertID6          int64
	MockInsertID7          int64
	MockInsertID8          int64

	MockUpdateSuccess1     bool
	MockUpdateSuccess2     bool
	MockUpdateSuccess3     bool
	MockUpdateSuccess4     bool
	mockUpdateSuccess1Used bool
	mockUpdateSuccess2Used bool
	mockUpdateSuccess3Used bool
	mockUpdateSuccess4Used bool

	MockDeleteSuccess1     bool
	MockDeleteSuccess2     bool
	MockDeleteSuccess3     bool
	MockDeleteSuccess4     bool
	MockDeleteSuccess5     bool
	MockDeleteSuccess6     bool
	MockDeleteSuccess7     bool
	MockDeleteSuccess8     bool
	mockDeleteSuccess1Used bool
	mockDeleteSuccess2Used bool
	mockDeleteSuccess3Used bool
	mockDeleteSuccess4Used bool
	mockDeleteSuccess5Used bool
	mockDeleteSuccess6Used bool
	mockDeleteSuccess7Used bool
	mockDeleteSuccess8Used bool

	MockTestRow *DbRow

	mockRow1Used bool
	mockRow2Used bool
	mockRow3Used bool
	mockRow4Used bool
	mockRow5Used bool
	mockRow6Used bool
	mockRow7Used bool
	mockRow8Used bool
	MockRow1     *DbRow
	MockRow2     *DbRow
	MockRow3     *DbRow
	MockRow4     *DbRow
	MockRow5     *DbRow
	MockRow6     *DbRow
	MockRow7     *DbRow
	MockRow8     *DbRow

	mockRows1Used bool
	mockRows2Used bool
	mockRows3Used bool
	mockRows4Used bool
	mockRows5Used bool
	mockRows6Used bool
	mockRows7Used bool
	mockRows8Used bool
	MockRows1     *DbRows
	MockRows2     *DbRows
	MockRows3     *DbRows
	MockRows4     *DbRows
	MockRows5     *DbRows
	MockRows6     *DbRows
	MockRows7     *DbRows
	MockRows8     *DbRows
}

// Connect Connect
func (p *PgDBMock) Connect() bool {
	return p.MockConnectSuccess
}

// New New
func (p *PgDBMock) New() Database {
	return p
}

// GetNewDatabase GetNewDatabase
func (p *PgDBMock) GetNewDatabase() Database {
	var db Database
	db = p
	return db
}

// BeginTransaction BeginTransaction
func (p *PgDBMock) BeginTransaction() Transaction {
	var trans Transaction
	var ptx PgDbTxMock
	ptx.PgDBMock = p
	trans = &ptx
	return trans
}

// Test Test
func (p *PgDBMock) Test(query string, args ...any) *DbRow {
	return p.MockTestRow
}

// Insert Insert
func (p *PgDBMock) Insert(query string, args ...any) (bool, int64) {
	var rtn = false
	var id int64
	if !p.mockInsertSuccess1Used {
		p.mockInsertSuccess1Used = true
		rtn = p.MockInsertSuccess1
		id = p.MockInsertID1
	} else if !p.mockInsertSuccess2Used {
		p.mockInsertSuccess2Used = true
		rtn = p.MockInsertSuccess2
		id = p.MockInsertID2
	} else if !p.mockInsertSuccess3Used {
		p.mockInsertSuccess3Used = true
		rtn = p.MockInsertSuccess3
		id = p.MockInsertID3
	} else if !p.mockInsertSuccess4Used {
		p.mockInsertSuccess4Used = true
		rtn = p.MockInsertSuccess4
		id = p.MockInsertID4
	}
	return rtn, id
}

// Update Update
func (p *PgDBMock) Update(query string, args ...any) bool {
	var rtn = false
	if !p.mockUpdateSuccess1Used {
		p.mockUpdateSuccess1Used = true
		rtn = p.MockUpdateSuccess1
	} else if !p.mockUpdateSuccess2Used {
		p.mockUpdateSuccess2Used = true
		rtn = p.MockUpdateSuccess2
	} else if !p.mockUpdateSuccess3Used {
		p.mockUpdateSuccess3Used = true
		rtn = p.MockUpdateSuccess3
	} else if !p.mockUpdateSuccess4Used {
		p.mockUpdateSuccess4Used = true
		rtn = p.MockUpdateSuccess4
	}
	return rtn
}

// Get Get
func (p *PgDBMock) Get(query string, args ...any) *DbRow {
	//return m.MockRow
	var rtn *DbRow
	if !p.mockRow1Used {
		p.mockRow1Used = true
		rtn = p.MockRow1
	} else if !p.mockRow2Used {
		p.mockRow2Used = true
		rtn = p.MockRow2
	} else if !p.mockRow3Used {
		p.mockRow3Used = true
		rtn = p.MockRow3
	} else if !p.mockRow4Used {
		p.mockRow4Used = true
		rtn = p.MockRow4
	} else if !p.mockRow5Used {
		p.mockRow5Used = true
		rtn = p.MockRow5
	} else if !p.mockRow6Used {
		p.mockRow6Used = true
		rtn = p.MockRow6
	} else if !p.mockRow7Used {
		p.mockRow7Used = true
		rtn = p.MockRow7
	} else if !p.mockRow8Used {
		p.mockRow8Used = true
		rtn = p.MockRow8
	}
	return rtn
}

// GetList GetList
func (p *PgDBMock) GetList(query string, args ...any) *DbRows {
	var rtn *DbRows
	if !p.mockRows1Used {
		p.mockRows1Used = true
		rtn = p.MockRows1
	} else if !p.mockRows2Used {
		p.mockRows2Used = true
		rtn = p.MockRows2
	} else if !p.mockRows3Used {
		p.mockRows3Used = true
		rtn = p.MockRows3
	} else if !p.mockRows4Used {
		p.mockRows4Used = true
		rtn = p.MockRows4
	} else if !p.mockRows5Used {
		p.mockRows5Used = true
		rtn = p.MockRows5
	} else if !p.mockRows6Used {
		p.mockRows6Used = true
		rtn = p.MockRows6
	} else if !p.mockRows7Used {
		p.mockRows7Used = true
		rtn = p.MockRows7
	} else if !p.mockRows8Used {
		p.mockRows8Used = true
		rtn = p.MockRows8
	}
	return rtn
}

// Delete Delete
func (p *PgDBMock) Delete(query string, args ...any) bool {
	var rtn = false
	if !p.mockDeleteSuccess1Used {
		p.mockDeleteSuccess1Used = true
		rtn = p.MockDeleteSuccess1
	} else if !p.mockDeleteSuccess2Used {
		p.mockDeleteSuccess2Used = true
		rtn = p.MockDeleteSuccess2
	} else if !p.mockDeleteSuccess3Used {
		p.mockDeleteSuccess3Used = true
		rtn = p.MockDeleteSuccess3
	} else if !p.mockDeleteSuccess4Used {
		p.mockDeleteSuccess4Used = true
		rtn = p.MockDeleteSuccess4
	}
	return rtn
}

// Close Close
func (p *PgDBMock) Close() bool {
	return p.MockCloseSuccess
}
