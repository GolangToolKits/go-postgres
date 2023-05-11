package gopostgres

import (
	"database/sql"
	"testing"
)

func TestPgDbTx_Insert(t *testing.T) {
	var a []any
	a = append(a, "UPS Trans", "Ground", 52)
	type fields struct {
		Tx *sql.Tx
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
			name:   "test 1",
			fields: fields{},
			args: args{
				query: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id",
				args:  a,
			},
			want:  true,
			want1: 0,
		},
		{
			name: "test 2",
			fields: fields{
				
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
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			}
			p.Connect()
			tr := p.BeginTransaction()
			// tr := &PgDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			got, got1 := tr.Insert(tt.args.query, tt.args.args...)
			if got != tt.want {
				t.Errorf("PgDbTx.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 < tt.want1 {
				t.Errorf("PgDbTx.Insert() got1 = %v, want %v", got1, tt.want1)
			}
			tr.Commit()
		})
	}
}

func TestPgDbTx_Update(t *testing.T) {
	var a []any
	a = append(a, "FedEx", "Air", 51, 10)

	var a2 []any

	var a4 []any
	a4 = append(a4, "FedEx", "Air", 51, 1)

	type fields struct {
		Tx *sql.Tx
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
			name:   "test 1",
			fields: fields{},
			args: args{
				query: "update carrier set carrier = $1 , type = $2, store_id = $3 where id = $4 ",
				args:  a,
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				
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
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			}
			p.Connect()
			tr := p.BeginTransaction()
			// tr := &PgDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			if got := tr.Update(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("PgDbTx.Update() = %v, want %v", got, tt.want)
			}
			tr.Commit()
		})
	}

}

func TestPgDbTx_Delete(t *testing.T) {

	var id = 3
	var a []any
	a = append(a, id)

	var id6 = 6
	var a6 []any
	a6 = append(a6, id6)

	var a7 []any

	type fields struct {
		Tx *sql.Tx
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
		rollbk bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				
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
				
			},
			args: args{
				query: "delete from carrier where id = $1 ",
				args:  a6,
			},
			want: true,
			rollbk: true,
		},
		{
			name: "test 3",
			fields: fields{
				
			},
			args: args{
				query: "delete from carrier where id = ? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 4",
			fields: fields{
				
			},
			args: args{
				query: "delete from carrier where id = $1 ",
				args:  a7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			}
			p.Connect()
			tr := p.BeginTransaction()
			// tr := &PgDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			if got := tr.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("PgDbTx.Delete() = %v, want %v", got, tt.want)
			}
			if tt.rollbk{
				tr.Rollback()
				tr.Rollback()
			}else{
				tr.Commit()
			}
			
		})
	}
}
