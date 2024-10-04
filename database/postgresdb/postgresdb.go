/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package postgresdb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lib/pq"
	"github.com/timbrockley/golang-main/system"
)

//------------------------------------------------------------

type PostgresDBStruct struct {
	//--------------------
	Host string
	//--------------------
	User     string
	Password string
	//--------------------
	Database string
	//--------------------
	AutoCreate bool
	//--------------------
	DB *sql.DB
	//--------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// init
//------------------------------------------------------------

func init() {
	//------------------------------------------------------------
	system.LoadENVs()
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect method
//------------------------------------------------------------

func (conn *PostgresDBStruct) Connect(checkENV ...bool) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	// if checkENV passed as true then check environment variables
	//------------------------------------------------------------
	if checkENV != nil && checkENV[0] {
		//--------------------
		conn.Host = os.Getenv("POSTGRES_HOST")
		//--------------------
		conn.User = os.Getenv("POSTGRES_USER")
		conn.Password = os.Getenv("POSTGRES_PWD")
		//--------------------
		if conn.Database == "" {
			conn.Database = os.Getenv("POSTGRES_DATABASE")
		}
		//--------------------
	}
	//------------------------------------------------------------
	if conn.Database != "" && !CheckDatabaseName(conn.Database) {
		return errors.New("invalid database name")
	}
	//------------------------------------------------------------
	var connString string
	//------------------------------------------------------------
	connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.Host, conn.User, conn.Password, conn.Database)
	conn.DB, err = sql.Open("postgres", connString)
	//----------------------------------------
	if err == nil {
		//----------------------------------------
		err = conn.DB.Ping()
		//----------------------------------------
		if err != nil && conn.Database != "" && conn.AutoCreate {
			//----------------------------------------
			var DB *sql.DB
			//----------------------------------------
			connString = fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", conn.Host, conn.User, conn.Password)
			DB, err = sql.Open("postgres", connString)
			//----------------------------------------
			if err == nil {
				//----------------------------------------
				err = DB.Ping()
				//----------------------------------------
				if err == nil {
					//----------------------------------------
					_, err = DB.Exec(fmt.Sprintf("CREATE DATABASE %s;", conn.Database))
					//----------------------------------------
					if err == nil {
						//----------------------------------------
						defer DB.Close()
						//--------------------
						connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.Host, conn.User, conn.Password, conn.Database)
						conn.DB, err = sql.Open("postgres", connString)
						//--------------------
						if err == nil {
							err = conn.DB.Ping()
						}
						//----------------------------------------
					}
					//----------------------------------------
				}
				//----------------------------------------
			}
			//----------------------------------------
		}
		//----------------------------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect - interface to connect method
//------------------------------------------------------------
// conn, err = Connect(PostgresDBStruct{ }, checkENV)
//------------------------------------------------------------

func Connect(conn PostgresDBStruct, checkENV ...bool) (PostgresDBStruct, error) {
	//------------------------------------------------------------
	return conn, conn.Connect(checkENV...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec method
//------------------------------------------------------------

func (conn *PostgresDBStruct) Exec(query string, args ...any) (sql.Result, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	return conn.DB.Exec(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Query method
//------------------------------------------------------------

func (conn *PostgresDBStruct) Query(query string, args ...any) (*sql.Rows, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	return conn.DB.Query(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRow method
//------------------------------------------------------------

func (conn *PostgresDBStruct) QueryRow(query string, args ...any) *sql.Row {
	//------------------------------------------------------------
	return conn.DB.QueryRow(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords method
//------------------------------------------------------------

func (conn *PostgresDBStruct) QueryRecords(query string, args ...any) ([]map[string]any, error) {
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query(strings.TrimSpace(query), args...)
	//--------------------
	if err != nil {
		return nil, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	return conn.ScanRows(rows)
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// TableExists method
//------------------------------------------------------------

func (conn *PostgresDBStruct) TableExists(tableName string) (bool, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return false, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return false, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return false, errors.New("invalid database name")
	}
	//--------------------
	if tableName == "" {
		return false, errors.New("table name cannot be blank")
	}
	//--------------------
	if !CheckTableName(tableName) {
		return false, errors.New("invalid table name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT COUNT(*) AS count FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME='%s' LIMIT 1;", tableName))
	//------------------------------------------------------------
	if err != nil {
		return false, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	var count int
	//--------------------
	for rows.Next() {
		//--------------------
		err = rows.Scan(&count)
		//--------------------
	}
	//--------------------
	if err != nil {
		return false, err
	}
	//------------------------------------------------------------
	return count == 1, nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetSQLTableInfo method
//------------------------------------------------------------

func (conn *PostgresDBStruct) GetSQLTableInfo(tableName string) (
	[]struct {
		Sequence int
		Name     string
		Type     string
	},
	map[string]string,
	error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, nil, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return nil, nil, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return nil, nil, errors.New("invalid database name")
	}
	//--------------------
	if tableName == "" {
		return nil, nil, errors.New("table name cannot be blank")
	}
	//--------------------
	if !CheckTableName(tableName) {
		return nil, nil, errors.New("invalid table name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	//--------------------
	columnInfoMap := map[string]string{}
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT COALESCE(ORDINAL_POSITION, 0), COALESCE(COLUMN_NAME, ''), COALESCE(DATA_TYPE, '') FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = $1;", tableName)
	//------------------------------------------------------------
	if err == nil {
		//------------------------------------------------------------
		defer rows.Close()
		//------------------------------------------------------------
		for rows.Next() {
			//--------------------
			var Sequence int
			var Name string
			var Type string
			//--------------------
			if err = rows.Scan(&Sequence, &Name, &Type); err != nil {
				break
			}
			//--------------------
			columInfoRows = append(columInfoRows, struct {
				Sequence int
				Name     string
				Type     string
			}{Sequence, Name, Type})
			//--------------------
			columnInfoMap[Name] = Type
			//--------------------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return columInfoRows, columnInfoMap, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetTableInfo method
//------------------------------------------------------------

func (conn *PostgresDBStruct) GetTableInfo(tableName string) (
	[]struct {
		Sequence int
		Name     string
		Type     string
	},
	map[string]string,
	error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, nil, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return nil, nil, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return nil, nil, errors.New("invalid database name")
	}
	//--------------------
	if tableName == "" {
		return nil, nil, errors.New("table name cannot be blank")
	}
	//--------------------
	if !CheckTableName(tableName) {
		return nil, nil, errors.New("invalid table name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 1;", tableName))
	//------------------------------------------------------------
	if err != nil {
		return nil, nil, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	return conn.GetRowsInfo(rows)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetRowsInfo method
//------------------------------------------------------------

func (conn *PostgresDBStruct) GetRowsInfo(rows *sql.Rows) (
	[]struct {
		Sequence int
		Name     string
		Type     string
	},
	map[string]string,
	error) {
	//------------------------------------------------------------
	var err error
	var colTypes []*sql.ColumnType
	//------------------------------------------------------------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	//--------------------
	columnInfoMap := map[string]string{}
	//------------------------------------------------------------
	colTypes, err = rows.ColumnTypes()
	//--------------------
	if err == nil {
		//--------------------
		for index, column := range colTypes {
			//--------------------
			Name := column.Name()
			Type := column.DatabaseTypeName()
			//--------------------
			columInfoRows = append(columInfoRows, struct {
				Sequence int
				Name     string
				Type     string
			}{index + 1, Name, Type})
			//--------------------
			columnInfoMap[Name] = Type
			//--------------------
		}
		//--------------------
	}
	//------------------------------------------------------------
	return columInfoRows, columnInfoMap, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ScanRows method
//------------------------------------------------------------

func (conn *PostgresDBStruct) ScanRows(sqlRows *sql.Rows) ([]map[string]any, error) {
	//------------------------------------------------------------
	var records []map[string]any
	//------------------------------------------------------------
	_, columnTypes, _ := conn.GetRowsInfo(sqlRows)
	//------------------------------------------------------------
	columns, err := sqlRows.Columns()
	//------------------------------------------------------------
	if err == nil {
		//------------------------------------------------------------
		for sqlRows.Next() {
			//------------------------------------------------------------
			scans := make([]any, len(columns))
			//--------------------
			record := make(map[string]any)
			//--------------------
			for i := range scans {
				scans[i] = &scans[i]
			}
			//--------------------
			sqlRows.Scan(scans...)
			//--------------------
			for index, value := range scans {
				//--------------------
				Name := columns[index]
				Type := strings.ToUpper(columnTypes[Name])
				//--------------------
				switch Type {
				case "BIGINT", "BIGSERIAL", "BIT", "BIT VARYING", "INT", "INTEGER", "MEDIUMINT", "SERIAL", "SMALLINT", "SMALLSERIAL", "TINYINT":
					value = system.ToInt(value)
				case "DEC", "DECIMAL", "DOUBLE", "DOUBLE PRECISION", "FIXED", "FLOAT", "NUMERIC", "REAL":
					value = system.ToFloat(value)
				case "BINARY", "BLOB", "BYTE", "BYTEA", "LONGBLOB", "TINYBLOB", "VARBINARY":
					value = system.ToBytes(value)
				case "BOOL", "BOOLEAN":
					value = system.ToBool(value)
				}
				//--------------------
				record[Name] = value
				//--------------------
			}
			//------------------------------------------------------------
			records = append(records, record)
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return records, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// ShowDatabases method
//------------------------------------------------------------

func (conn *PostgresDBStruct) ShowDatabases() ([]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return []string{}, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return []string{}, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return []string{}, errors.New("invalid database name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT datname FROM pg_database WHERE datistemplate = false;")
	//------------------------------------------------------------
	if err != nil {
		return []string{}, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	var tables = []string{}
	//------------------------------
	for rows.Next() {
		//--------------------
		var tbl_name string
		//--------------------
		err = rows.Scan(&tbl_name)
		//--------------------
		if err != nil {
			return []string{}, err
		}
		//--------------------
		tables = append(tables, tbl_name)
		//--------------------
	}
	//------------------------------------------------------------
	return tables, nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTables method
//------------------------------------------------------------

func (conn *PostgresDBStruct) ShowTables() ([]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return []string{}, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return []string{}, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return []string{}, errors.New("invalid database name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
	//------------------------------------------------------------
	if err != nil {
		return []string{}, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	var tables = []string{}
	//------------------------------
	for rows.Next() {
		//--------------------
		var tbl_name string
		//--------------------
		err = rows.Scan(&tbl_name)
		//--------------------
		if err != nil {
			return []string{}, err
		}
		//--------------------
		tables = append(tables, tbl_name)
		//--------------------
	}
	//------------------------------------------------------------
	return tables, nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTablesMap method
//------------------------------------------------------------

func (conn *PostgresDBStruct) ShowTablesMap() (map[string]map[string]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return map[string]map[string]string{}, errors.New("not connected")
	}
	//--------------------
	if conn.Database == "" {
		return map[string]map[string]string{}, errors.New("database name cannot be blank")
	}
	//--------------------
	if !CheckDatabaseName(conn.Database) {
		return map[string]map[string]string{}, errors.New("invalid database name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
	//------------------------------------------------------------
	if err != nil {
		return map[string]map[string]string{}, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	var tablesMap = map[string]map[string]string{}
	//----------------------------------------
	for rows.Next() {
		//----------------------------------------
		var tbl_name string
		//--------------------
		err = rows.Scan(&tbl_name)
		//--------------------
		if err != nil {
			return map[string]map[string]string{}, err
		}
		//--------------------
		_, columnInfoMap, err = conn.GetTableInfo(tbl_name)
		//--------------------
		if err != nil {
			return map[string]map[string]string{}, err
		}
		//--------------------
		tablesMap[tbl_name] = columnInfoMap
		//----------------------------------------
	}
	//------------------------------------------------------------
	return tablesMap, nil
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Close method
//------------------------------------------------------------

func (conn *PostgresDBStruct) Close() error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if conn.DB != nil {
		err = conn.DB.Close()
	}
	//------------------------------------------------------------
	conn.DB = nil
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func EscapeApostrophes(dataString string) string {
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		`'`, `''`,
	)
	//------------------------------------------------------------
	return replacer.Replace(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------

func EscapeDoubleQuotes(dataString string) string {
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		`"`, `""`,
	)
	//------------------------------------------------------------
	return replacer.Replace(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------

func EscapePostgreSQLString(dataString string) string {
	//------------------------------------------------------------
	return pq.QuoteLiteral(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

func CheckDatabaseName(databaseName string) bool {
	//------------------------------------------------------------
	return CheckTableName(databaseName)
	//------------------------------------------------------------
}

//------------------------------------------------------------

func CheckTableName(tableName string) bool {
	//------------------------------------------------------------
	// checks table name (and database name if prepends table name)
	//------------------------------------------------------------
	var err error
	var match bool
	//------------------------------------------------------------
	if strings.Contains(tableName, ".") {
		//--------------------
		elements := strings.Split(tableName, ".")
		//--------------------
		if len(elements) != 2 {
			return false
		}
		//--------------------
		if !CheckDatabaseName(elements[0]) {
			return false
		}
		//--------------------
		tableName = elements[1]
		//--------------------
	}
	//------------------------------------------------------------
	// should start with underscore or a letter
	//------------------------------------------------------------
	match, err = regexp.MatchString(`^[_A-Za-z]+`, tableName)
	//------------------------------------------------------------
	if err != nil || !match {
		return false
	}
	//------------------------------------------------------------
	// remaining characters should only contain underscores, letters or numbers
	//------------------------------------------------------------
	match, err = regexp.MatchString(`^[_A-Za-z0-9]*$`, tableName)
	//--------------------
	return err == nil && match
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
