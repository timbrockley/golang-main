/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package sqlitedb

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timbrockley/golang-main/file"
)

//------------------------------------------------------------

type SQLiteDBStruct struct {
	//----------
	FilePath    string
	Database    string
	DatabaseExt string
	//----------
	AutoCreate bool
	//----------
	DB *sql.DB
	//----------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Connect method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) Connect() error {

	//------------------------------------------------------------
	var err error
	var filePath, database, databaseExt string
	//------------------------------------------------------------
	filePath = conn.FilePath
	database = conn.Database
	databaseExt = conn.DatabaseExt
	//------------------------------------------------------------
	if database != "" {

		if database == ":memory:" {

			//----------
			filePath = ":memory:"
			databaseExt = ""
			//----------

		} else {

			//----------
			if databaseExt == "" {
				databaseExt = file.FilenameExt(database)
			}
			//----------
			if databaseExt != "" {
				database = file.FilenameBase(database)
			} else {
				databaseExt = "db"
			}
			//----------
			if filePath == "" {
				filePath = file.Path()
			}
			//----------
			filePath = file.FilePathJoin(filePath, database+"."+databaseExt)
			//----------
		}

	} else {

		if filePath == ":memory:" {

			//----------
			database = ":memory:"
			databaseExt = ""
			//----------

		} else {

			//----------
			if filePath == "" {
				//----------
				// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
				_, filePath, _, _ = runtime.Caller(1)
				//----------
				filePath = file.FilePathBase(filePath) + ".db"
				//----------
			}
			//----------
			database = file.Filename(filePath)
			//----------
			if database != "" {
				if databaseExt == "" {
					databaseExt = file.FilenameExt(database)
				}
				if databaseExt != "" {
					database = file.FilenameBase(database)
				} else {
					databaseExt = "db"
				}
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------
	conn.FilePath = filePath
	conn.Database = database
	conn.DatabaseExt = databaseExt
	//------------------------------------------------------------
	if filePath != ":memory:" && !conn.AutoCreate && !file.FilePathExists(filePath) {

		return fmt.Errorf("database file %q does not exists ", filePath)
	}
	//------------------------------------------------------------
	conn.DB, err = sql.Open("sqlite3", filePath)
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect - interface to connect method
//------------------------------------------------------------
// conn, err = Connect(SQLiteDBStruct{ }, checkENV)
//------------------------------------------------------------

func Connect(conn SQLiteDBStruct) (SQLiteDBStruct, error) {

	//------------------------------------------------------------
	return conn, conn.Connect()
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) Exec(query string, args ...any) (sql.Result, error) {
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

func (conn *SQLiteDBStruct) Query(query string, args ...any) (*sql.Rows, error) {
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

func (conn *SQLiteDBStruct) QueryRow(query string, args ...any) *sql.Row {
	//------------------------------------------------------------
	return conn.DB.QueryRow(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) QueryRecords(query string, args ...any) ([]map[string]any, error) {
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query(strings.TrimSpace(query), args...)
	//----------
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
// ScanRows method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) ScanRows(sqlRows *sql.Rows) ([]map[string]any, error) {
	//------------------------------------------------------------
	var records []map[string]any
	//------------------------------------------------------------
	columns, err := sqlRows.Columns()
	//------------------------------------------------------------
	if err == nil {
		//------------------------------------------------------------
		for sqlRows.Next() {
			//------------------------------------------------------------
			scans := make([]any, len(columns))
			//----------
			record := make(map[string]any)
			//----------
			for i := range scans {
				scans[i] = &scans[i]
			}
			//----------
			sqlRows.Scan(scans...)
			//----------
			for index, value := range scans {
				//------------------------------------------------------------
				switch typedValue := value.(type) {
				case int:
					value = typedValue
				case int8:
					value = int(typedValue)
				case int16:
					value = int(typedValue)
				case int32:
					value = int(typedValue)
				case int64:
					value = int(typedValue)
				case uint:
					value = typedValue
				case uint8:
					value = uint(typedValue)
				case uint16:
					value = uint(typedValue)
				case uint32:
					value = uint(typedValue)
				case uint64:
					value = uint(typedValue)
				case float32:
					value = float64(typedValue)
				case float64:
					value = typedValue
				case bool:
					value = typedValue
				case string:
					value = typedValue
				}
				//----------
				record[columns[index]] = value
				//------------------------------------------------------------
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
// ShowTables method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) ShowTables() ([]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return []string{}, errors.New("not connected")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT DISTINCT tbl_name FROM sqlite_master;")
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

func (conn *SQLiteDBStruct) ShowTablesMap() (map[string]map[string]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return map[string]map[string]string{}, errors.New("not connected")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT DISTINCT tbl_name FROM sqlite_master;")
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
// TableExists method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) TableExists(tableName string) (bool, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return false, errors.New("not connected")
	}
	//----------
	if tableName == "" {
		return false, errors.New("table name cannot be blank")
	}
	//----------
	if !CheckTableName(tableName) {
		return false, errors.New("invalid table name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query(fmt.Sprintf("SELECT COUNT(*) AS count FROM sqlite_master WHERE type='table' AND name='%s' LIMIT 1;", tableName))
	//------------------------------------------------------------
	if err != nil {
		return false, err
	}
	//------------------------------------------------------------
	defer rows.Close()
	//------------------------------------------------------------
	var count int
	//----------
	for rows.Next() {
		//----------
		err = rows.Scan(&count)
		//----------
	}
	//----------
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

func (conn *SQLiteDBStruct) GetSQLTableInfo(tableName string) (
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
	//----------
	if tableName == "" {
		return nil, nil, errors.New("table name cannot be blank")
	}
	//----------
	if !CheckTableName(tableName) {
		return nil, nil, errors.New("invalid table name")
	}
	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	rows, err = conn.DB.Query("SELECT IFNULL(cid, 0)+1, IFNULL(name, ''), IFNULL(type, '') FROM PRAGMA_TABLE_INFO(?);", tableName)
	//------------------------------------------------------------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	columnInfoMap := map[string]string{}
	//------------------------------------------------------------
	if err == nil {
		//------------------------------------------------------------
		defer rows.Close()
		//----------
		for rows.Next() {
			//----------
			var columInfoRow struct {
				Sequence int
				Name     string
				Type     string
			}
			//----------
			if err = rows.Scan(&columInfoRow.Sequence, &columInfoRow.Name, &columInfoRow.Type); err != nil {
				break
			}
			//----------
			columInfoRows = append(columInfoRows, columInfoRow)
			//----------
			columnInfoMap[columInfoRow.Name] = columInfoRow.Type
			//------------------------------------------------------------
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

func (conn *SQLiteDBStruct) GetTableInfo(tableName string) (
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
	//----------
	if tableName == "" {
		return nil, nil, errors.New("table name cannot be blank")
	}
	//----------
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

func (conn *SQLiteDBStruct) GetRowsInfo(rows *sql.Rows) (
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
	//----------
	columnInfoMap := map[string]string{}
	//------------------------------------------------------------
	colTypes, err = rows.ColumnTypes()
	//----------
	if err == nil {
		//------------------------------------------------------------
		for index, column := range colTypes {
			//----------
			Name := column.Name()
			Type := column.DatabaseTypeName()
			//----------
			columInfoRows = append(columInfoRows, struct {
				Sequence int
				Name     string
				Type     string
			}{index + 1, Name, Type})
			//----------
			columnInfoMap[Name] = Type
			//----------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return columInfoRows, columnInfoMap, err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Close method
//------------------------------------------------------------

func (conn *SQLiteDBStruct) Close() error {

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

//------------------------------------------------------------
// CheckDatabaseName
//------------------------------------------------------------

func CheckDatabaseName(databaseName string) bool {
	//------------------------------------------------------------
	return CheckTableName(databaseName)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// CheckTableName
//------------------------------------------------------------

func CheckTableName(tableName string) bool {
	//------------------------------------------------------------
	// checks table name (and database name if prepends table name)
	//------------------------------------------------------------
	var err error
	var match bool
	//------------------------------------------------------------
	if strings.Contains(tableName, ".") {
		//----------
		elements := strings.Split(tableName, ".")
		//----------
		if len(elements) != 2 {
			return false
		}
		//----------
		if !CheckDatabaseName(elements[0]) {
			return false
		}
		//----------
		tableName = elements[1]
		//----------
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
	//----------
	return err == nil && match
	//------------------------------------------------------------
}

//------------------------------------------------------------
// EscapeApostrophes
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
// EscapeDoubleQuotes
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
//############################################################
//------------------------------------------------------------
