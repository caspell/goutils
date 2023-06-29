package querybuilder

import (
	"math/big"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MariadbConfig struct {
	ConnectionString string `toml:"MYSQL_CONNECTION_STRING"`
	MaxIdleConns     int    // default 10
	MaxOpenConns     int    // default 10
}

type Mariadb struct {
	Config MariadbConfig
	DB     *sqlx.DB
}

type DatabaseError struct {
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}

type DataError struct {
	Message string
}

func (e DataError) Error() string {
	return e.Message
}

func (conn *Mariadb) Disconnect() error {

	return conn.DB.Close()
}

func (conn *Mariadb) Connect() error {
	if db, err := sqlx.Connect("mysql", conn.Config.ConnectionString); err != nil {
		return err
	} else {
		conn.DB = db
		conn.DB.SetMaxIdleConns(conn.Config.MaxIdleConns)
		conn.DB.SetMaxOpenConns(conn.Config.MaxOpenConns)
	}
	return nil
}

func (conn *Mariadb) Exec(statement *Statement) error {

	if _, err := conn.DB.NamedExec(statement.Script, statement.Parameter); err != nil {
		return err
	}
	return nil
}

func (conn *Mariadb) ExecBulk(statement *Statement) error {

	if statement.Parameters == nil || len(statement.Parameters) < 1 {
		return &DatabaseError{
			Message: "Parameters not found.",
		}
	}

	tx := conn.DB.MustBegin()

	if _, err := tx.NamedExec(statement.Script, statement.Parameters); err != nil {
		defer tx.Rollback()
		return err
	} else {
		defer tx.Commit()
	}

	return nil
}

func (conn *Mariadb) Select(statement *Statement) ([]map[string]interface{}, error) {
	rows := make([]map[string]interface{}, 0)
	err := conn.DB.Select(&rows, statement.Script)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (conn *Mariadb) Query(statement *Statement) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)
	rows, err := conn.DB.Queryx(statement.Script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		types, _ := rows.ColumnTypes()
		row := make(map[string]interface{})
		err := rows.MapScan(row)
		if err != nil {
			continue
		}
		for _, t := range types {
			name := t.Name()
			if row[name] == nil {
				continue
			}
			v := string(row[name].([]uint8)[:])
			switch t.DatabaseTypeName() {
			case "INT":
				atoi, _ := strconv.Atoi(v)
				row[name] = atoi
			case "BIGINT":
				n := new(big.Int)
				n, _ = n.SetString(v, 10)
				row[name] = n.Int64()
			// case "VARCHAR":
			// case "TIMESTAMP":
			default:
				row[name] = v
			}
		}
		result = append(result, row)
	}
	return result, nil
}

func (conn *Mariadb) QueryWithParam(statement *Statement) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)
	rows, err := conn.DB.NamedQuery(statement.Script, statement.Parameter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		types, _ := rows.ColumnTypes()
		row := make(map[string]interface{})
		err := rows.MapScan(row)
		if err != nil {
			continue
		}
		for _, t := range types {
			name := t.Name()
			if row[name] == nil {
				continue
			}
			v := string(row[name].([]uint8)[:])
			switch t.DatabaseTypeName() {
			case "INT":
				atoi, _ := strconv.Atoi(v)
				row[name] = atoi
			case "BIGINT":
				n := new(big.Int)
				n, _ = n.SetString(v, 10)
				row[name] = n.Int64()
			// case "VARCHAR":
			// case "TIMESTAMP":
			default:
				row[name] = v
			}
		}
		result = append(result, row)
	}
	return result, nil
}

func (conn *Mariadb) QueryRow(statement *Statement) (map[string]interface{}, error) {

	result, err := conn.Query(statement)

	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return result[0], nil
	}

	return nil, nil
}

func (conn *Mariadb) QueryRowWithParam(statement *Statement) (map[string]interface{}, error) {

	result, err := conn.QueryWithParam(statement)

	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return result[0], nil
	}

	return nil, nil
}
