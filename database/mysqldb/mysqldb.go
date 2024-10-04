/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/timbrockley/golang-main/system"
)

//------------------------------------------------------------

type MySQLdbStruct struct {
	//--------------------
	Host string
	//--------------------
	User                 string
	Password             string
	AllowNativePasswords bool
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

func (conn *MySQLdbStruct) Connect(checkENV ...bool) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	// if checkENV passed as true then check environment variables
	//------------------------------------------------------------
	if checkENV != nil && checkENV[0] {
		//--------------------
		conn.Host = os.Getenv("MYSQL_HOST")
		//--------------------
		conn.User = os.Getenv("MYSQL_USER")
		conn.Password = os.Getenv("MYSQL_PWD")
		//--------------------
		conn.AllowNativePasswords = os.Getenv("MYSQL_ALLOW_NATIVE_PASSWORDS") == "true"
		//--------------------
		if conn.Database == "" {
			conn.Database = os.Getenv("MYSQL_DATABASE")
		}
		//--------------------
	}
	//------------------------------------------------------------
	mysqlConfig := mysql.Config{
		//--------------------
		User:   conn.User,
		Passwd: conn.Password,
		//--------------------
		AllowNativePasswords: conn.AllowNativePasswords,
		//--------------------
		MultiStatements: true,
		//--------------------
		// DBName: conn.Database,
		//--------------------
	}
	//------------------------------------------------------------
	if conn.Host != "" {
		//--------------------
		mysqlConfig.Addr = conn.Host
		mysqlConfig.Net = "tcp"
		//--------------------
	}
	//------------------------------------------------------------
	if conn.Database != "" && !CheckDatabaseName(conn.Database) {
		return errors.New("invalid database name")
	}
	//------------------------------------------------------------
	conn.DB, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	//----------------------------------------
	if err == nil {
		//----------------------------------------
		err = conn.DB.Ping()
		//----------------------------------------
		if err == nil && conn.Database != "" {
			//----------------------------------------
			if conn.AutoCreate {
				//----------------------------------------
				_, err = conn.DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", conn.Database))
				//----------------------------------------
			}
			//----------------------------------------
			if err == nil {
				//----------------------------------------
				_, err = conn.DB.Exec(fmt.Sprintf("USE %s;", conn.Database))
				//----------------------------------------
			}
			//----------------------------------------
		}
		//--------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect - interface to connect method
//------------------------------------------------------------
// conn, err = Connect(MySQLdbStruct{ }, checkENV)
//------------------------------------------------------------

func Connect(conn MySQLdbStruct, checkENV ...bool) (MySQLdbStruct, error) {
	//------------------------------------------------------------
	return conn, conn.Connect(checkENV...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec method
//------------------------------------------------------------

func (conn *MySQLdbStruct) Exec(query string, args ...any) (sql.Result, error) {
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

func (conn *MySQLdbStruct) Query(query string, args ...any) (*sql.Rows, error) {
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

func (conn *MySQLdbStruct) QueryRow(query string, args ...any) *sql.Row {
	//------------------------------------------------------------
	return conn.DB.QueryRow(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords method
//------------------------------------------------------------

func (conn *MySQLdbStruct) QueryRecords(query string, args ...any) ([]map[string]any, error) {
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
// LockTables method
//------------------------------------------------------------

func (conn *MySQLdbStruct) LockTables(Tables ...string) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if conn.DB == nil {
		return errors.New("not connected")
	}
	//------------------------------------------------------------
	if Tables == nil {
		//------------------------------------------------------------
		return errors.New("no tables defined")
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	tableLocks := []string{}
	//--------------------
	for _, tableName := range Tables {
		//--------------------
		if !CheckTableName(tableName) {
			//--------------------
			err = fmt.Errorf("invalid table name: (%s)", tableName)
			break
			//--------------------
		} else {
			//--------------------
			tableLocks = append(tableLocks, fmt.Sprintf("%s WRITE", tableName))
			//---------
		}
		//--------------------
	}
	//------------------------------------------------------------
	if err == nil {
		//------------------------------------------------------------
		_, err = conn.DB.Exec(fmt.Sprintf("LOCK TABLES %s;", strings.Join(tableLocks, ", ")))
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// UnlockTables method
//------------------------------------------------------------

func (conn *MySQLdbStruct) UnlockTables() error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if conn.DB == nil {
		return errors.New("not connected")
	}
	//------------------------------------------------------------
	_, err = conn.DB.Exec("UNLOCK TABLES;")
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// TableExists method
//------------------------------------------------------------

func (conn *MySQLdbStruct) TableExists(tableName string) (bool, error) {
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
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT COUNT(*) AS count FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='%s' AND TABLE_NAME='%s' LIMIT 1;", conn.Database, tableName))
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

func (conn *MySQLdbStruct) GetSQLTableInfo(tableName string) (
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
	rows, err = conn.DB.Query("SELECT IFNULL(ORDINAL_POSITION, 0), IFNULL(COLUMN_NAME, ''), IFNULL(DATA_TYPE, '') FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?;", conn.Database, tableName)
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

func (conn *MySQLdbStruct) GetTableInfo(tableName string) (
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
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT * FROM %s.%s LIMIT 1;", conn.Database, tableName))
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

func (conn *MySQLdbStruct) GetRowsInfo(rows *sql.Rows) (
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

func (conn *MySQLdbStruct) ScanRows(sqlRows *sql.Rows) ([]map[string]any, error) {
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
				if fmt.Sprintf("%T", value) == "[]uint8" {
					value = string(value.([]uint8))
				}
				//--------------------
				switch Type {
				case "BIGINT", "BIT", "BIT VARYING", "INT", "INTEGER", "MEDIUMINT", "SERIAL", "SMALLINT", "SMALLSERIAL", "TINYINT":
					value = system.ToInt(value)
				case "DEC", "DECIMAL", "DOUBLE", "DOUBLE PRECISION", "FIXED", "FLOAT", "NUMERIC", "REAL":
					value = system.ToFloat(value)
				case "BIGSERIAL", "BINARY", "BLOB", "BYTE", "BYTEA", "LONGBLOB", "TINYBLOB", "VARBINARY":
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
// ShowDatabases method
//------------------------------------------------------------

func (conn *MySQLdbStruct) ShowDatabases() ([]string, error) {
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
	rows, err = conn.DB.Query("SELECT schema_name FROM information_schema.SCHEMATA;")
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

func (conn *MySQLdbStruct) ShowTables() ([]string, error) {
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
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='%s';", conn.Database))
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

func (conn *MySQLdbStruct) ShowTablesMap() (map[string]map[string]string, error) {
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
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='%s';", conn.Database))
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

func (conn *MySQLdbStruct) Close() error {
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

func EscapeMySQLString(dataString string) string {
	//------------------------------------------------------------
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		"\x00", "\\\x00",
		"\r", "\\r",
		"\n", "\\n",
		"\x1A", "\\Z",
		`"`, `\"`,
		`'`, `\'`,
	)
	//------------------------------------------------------------
	return replacer.Replace(dataString)
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
