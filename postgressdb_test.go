package gopostgres

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestPgDB_Connect(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost1",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			if got := p.Connect(); got != tt.want {
				t.Errorf("PgDB.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	var a []any
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DbRow
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select count(*) from carrier ",
				args:  a,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.Test(tt.args.query, tt.args.args...); len(got.Row) == 0 {
				t.Errorf("PgDB.Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_Insert(t *testing.T) {
	var a []any
	a = append(a, "UPS", "Ground", 51)
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id",
				args:  a,
			},
			want:  true,
			want1: 11,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "insert into carrier (carrier, type, store_id) values($1, $2, ?) RETURNING id",
				args:  a,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "test 3",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "insert into carrier (carrier, type, store_id) values($1, $2, $3) ",
				args:  a,
			},
			want:  false,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			got, got1 := p.Insert(tt.args.query, tt.args.args...)
			if got != tt.want {
				t.Errorf("PgDB.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 < tt.want1 {
				t.Errorf("PgDB.Insert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPgDB_Update(t *testing.T) {
	var a []any
	a = append(a, "FedEx", "Air", 51, 5)
	var a2 []any
	//a2 = append(a2, "FedEx", "Air", 51, 1)
	var a4 []any
	a4 = append(a4, "FedEx", "Air", 51, 1)
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "update carrier set carrier = $1 , type = $2, store_id = $3 where id = $4 ",
				args:  a,
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "update carrier set carrier = $1 , type = $2, store_id = $3 where id = ? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 3",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "update carrier set carrier = $1 , type = $2, store_id = $3 where id = $4 ",
				args:  a2,
			},
			want: false,
		},
		{
			name: "test 4",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "update carrier set carrier = $1 , type = $2, store_id = $3 where id = $4 ",
				args:  a4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.Update(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("PgDB.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_Get(t *testing.T) {
	var id = 11
	var a []any
	a = append(a, id)

	//var id2 = 111
	var a2 []any
	//a2 = append(a2, id2)

	var id4 = 32
	var a4 []any
	a4 = append(a4, id4)
	
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where id = $1 ",
				args:  a,
			},
			want: 4,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where id = ? ",
				args:  a,
			},
			want: 0,
		},
		{
			name: "test 3",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where id = $1 ",
				args:  a2,
			},
			want: 0,
		},
		{
			name: "test 4",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where id = $1 ",
				args:  a4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.Get(tt.args.query, tt.args.args...); len(got.Row) != tt.want {
				t.Errorf("PgDB.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_GetList(t *testing.T) {
	var c = "FedEx"
	var a []any
	a = append(a, c)

	var a3 []any
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want   *DbRows
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where carrier = $1 ",
				args:  a,
			},
			want: 2,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where carrier = ? ",
				args:  a,
			},
			want: 0,
		},
		{
			name: "test 3",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "select * from carrier where carrier = $1 ",
				args:  a3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.GetList(tt.args.query, tt.args.args...); len(got.Rows) < tt.want {
				t.Errorf("PgDB.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_Delete(t *testing.T) {
	var id = 1
	var a []any
	a = append(a, id)

	var a3 []any
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "delete from carrier where id = $1 ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "delete from carrier where id = ? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 3",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			args: args{
				query: "delete from carrier where id = $1 ",
				args:  a3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("PgDB.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_Close(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.Close(); got != tt.want {
				t.Errorf("PgDB.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_BeginTransaction(t *testing.T) {
	var tx PgDbTx
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   Transaction
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			want: &tx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			p.Connect()
			if got := p.BeginTransaction(); got == nil {
				t.Errorf("PgDB.BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDB_New(t *testing.T) {
	var db PgDB
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   Database
	}{
		// TODO: Add test cases.
		{
			name:   "test 2",
			// fields: fields{
			// 	Host:     "localhost",
			// 	Port:     "5432",
			// 	User:     "admin",
			// 	Password: "ken",
			// 	Database: "kentest",
			// },
			want:   &db,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			if got := p.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDB.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
