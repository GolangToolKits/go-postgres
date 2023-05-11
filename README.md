# go-postgres
A simple mockable postgres db interface


[![Go Report Card](https://goreportcard.com/badge/github.com/GolangToolKits/go-postgres)](https://goreportcard.com/report/github.com/GolangToolKits/go-postgres)

```go

    mm := &PgDB{
	    Host:     "localhost",
        Port:     "5432",
	    User:     "test",
	    Password: "test",
	    Database: "test",				
    }
    m := mm.New()
    m.Connect()
    //Insert requires  RETURNING id on end of prepared statement query
        //Example: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id"
    m.Insert(query, args...)
    m.Update(query, args...)
    m.Get(query, args...)
    m.GetList(query, args...)
    m.Delete(query, args...)

```

## Also Supports Transactions

```go
 mm := &PgDB{
	    Host:     "localhost",
        Port:     "5432",
	    User:     "test",
	    Password: "test",
	    Database: "test",				
    }
    m := mm.New()
    m.Connect()
    tr := m.BeginTransaction()
    //Insert requires  RETURNING id on end of prepared statement query
        //Example: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id"
    tr.Insert(query, args...)
    tr.Update(query, args...)
    tr.Get(query, args...)
    tr.GetList(query, args...)
    tr.Delete(query, args...)
    tr.Commit()
    //Or
    tr.Rollback()


```
