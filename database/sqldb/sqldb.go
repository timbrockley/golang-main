/*

Copyright 2023-2024, Tim Brockley. All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

*/

package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/timbrockley/golang-main/database/mysqldb"
	"github.com/timbrockley/golang-main/database/postgresdb"
	"github.com/timbrockley/golang-main/database/sqlitedb"
	"github.com/timbrockley/golang-main/file"
	"github.com/timbrockley/golang-main/system"
)

//------------------------------------------------------------

type SQLdb struct {
	//--------------------
	DBType string
	//--------------------
	Host string
	//--------------------
	User     string
	Password string
	//--------------------
	AllowNativePasswords bool
	//--------------------
	FilePath    string
	DataPath    string
	Database    string
	DatabaseExt string
	//--------------------
	AutoCreate bool
	//--------------------
	DB *sql.DB
	//--------------------
	connMySQL    mysqldb.MySQLdb
	connPostgres postgresdb.PostgresDBStruct
	connSQLite   sqlitedb.SQLiteDB
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

func (conn *SQLdb) Connect(checkENV ...bool) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if conn.DB != nil {
		conn.DB.Close()
	}
	//------------------------------------------------------------
	if conn.DBType == "" && checkENV != nil && checkENV[0] {
		//--------------------
		conn.DBType = os.Getenv("DB_TYPE")
		//--------------------
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {

		//------------------------------------------------------------
		if len(checkENV) > 0 && checkENV[0] {
			//------------------------------------------------------------
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
			if conn.Database == "" {
				conn.Database = "_system"
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
		conn.connMySQL = mysqldb.MySQLdb{Host: conn.Host, User: conn.User, Password: conn.Password, AllowNativePasswords: conn.AllowNativePasswords, Database: conn.Database, AutoCreate: conn.AutoCreate}
		//--------------------
		err = conn.connMySQL.Connect()
		//------------------------------------------------------------
		if err == nil {
			//--------------------
			conn.DB = conn.connMySQL.DB
			//--------------------
		}
		//------------------------------------------------------------

	} else if strings.ToLower(conn.DBType) == "postgres" {

		//------------------------------------------------------------
		if len(checkENV) > 0 && checkENV[0] {
			//------------------------------------------------------------
			conn.Host = os.Getenv("POSTGRES_HOST")
			//--------------------
			conn.User = os.Getenv("POSTGRES_USER")
			conn.Password = os.Getenv("POSTGRES_PWD")
			//--------------------
			if conn.Database == "" {
				conn.Database = os.Getenv("POSTGRES_DATABASE")
			}
			//--------------------
			if conn.Database == "" {
				conn.Database = "_system"
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
		conn.connPostgres = postgresdb.PostgresDBStruct{Host: conn.Host, User: conn.User, Password: conn.Password, Database: conn.Database, AutoCreate: conn.AutoCreate}
		//--------------------
		err = conn.connPostgres.Connect()
		//------------------------------------------------------------
		if err == nil {
			//--------------------
			conn.DB = conn.connPostgres.DB
			//--------------------
		}
		//------------------------------------------------------------

	} else {

		//------------------------------------------------------------
		// first check for :memory: database otherwise
		// if file path is not passed then check environment variables
		//------------------------------------------------------------
		if conn.FilePath != ":memory:" && conn.Database != ":memory:" {
			//------------------------------------------------------------

			if conn.FilePath != "" {
				//--------------------
				conn.DataPath = ""
				conn.Database = ""
				conn.DatabaseExt = ""
				//--------------------
			} else {
				//------------------------------------------------------------
				var filePath, database string
				//------------------------------------------------------------
				if len(checkENV) > 0 && checkENV[0] {
					//--------------------
					filePath = os.Getenv("SQLITE_FILEPATH")
					//--------------------
					if conn.DataPath == "" {
						conn.DataPath = os.Getenv("SQLITE_DATA_PATH")
					}
					//--------------------
					if conn.Database == "" {
						database = os.Getenv("SQLITE_DATABASE")
					}
					//--------------------
					if conn.DatabaseExt == "" {
						conn.DatabaseExt = os.Getenv("SQLITE_DATABASE_EXT")
					}
					//--------------------
				}
				//------------------------------------------------------------
				if filePath != "" {
					//--------------------
					conn.FilePath = filePath
					conn.DataPath = ""
					conn.Database = ""
					conn.DatabaseExt = ""
					//--------------------
				} else {
					//--------------------
					if conn.DatabaseExt == "" {
						//--------------------
						conn.DatabaseExt = "db"
						//--------------------
					}
					//--------------------
					if conn.Database == "" {
						//--------------------
						if database != "" {
							//--------------------
							conn.Database = database
							//--------------------
						} else {
							//--------------------
							conn.Database = "_system"
							//--------------------
						}
						//--------------------
					}
					//--------------------
					if conn.DataPath != "" {
						//--------------------
						conn.FilePath = file.FilePathJoin(conn.DataPath, conn.Database+"."+conn.DatabaseExt)
						//--------------------
						conn.DataPath = ""
						conn.Database = ""
						conn.DatabaseExt = ""
						//--------------------
					} else {
						return fmt.Errorf("invalid dataPath")
					}
					//--------------------
				}
				//--------------------
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
		conn.connSQLite = sqlitedb.SQLiteDB{FilePath: conn.FilePath, DataPath: conn.DataPath, Database: conn.Database, DatabaseExt: conn.DatabaseExt, AutoCreate: conn.AutoCreate}
		//--------------------
		err = conn.connSQLite.Connect()
		//------------------------------------------------------------
		conn.FilePath = conn.connSQLite.FilePath
		//--------------------
		conn.Database = conn.connSQLite.Database
		conn.DatabaseExt = conn.connSQLite.DatabaseExt
		//--------------------
		if err == nil {
			//--------------------
			conn.DB = conn.connSQLite.DB
			//--------------------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect - connect to database and return passed SQLdb Object
//------------------------------------------------------------
// conn, err = Connect(SQLdb{ }, checkENV)
//------------------------------------------------------------

func Connect(conn SQLdb, checkENV ...bool) (SQLdb, error) {
	//------------------------------------------------------------
	return conn, conn.Connect(checkENV...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec method
//------------------------------------------------------------

func (conn *SQLdb) Exec(query string, args ...any) (sql.Result, error) {
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

func (conn *SQLdb) Query(query string, args ...any) (*sql.Rows, error) {
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

func (conn *SQLdb) QueryRow(query string, args ...any) *sql.Row {
	//------------------------------------------------------------
	return conn.DB.QueryRow(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords method
//------------------------------------------------------------

func (conn *SQLdb) QueryRecords(query string, args ...any) ([]map[string]any, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//--------------------
		return conn.connMySQL.QueryRecords(strings.TrimSpace(query), args...)
		//--------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//--------------------
		return conn.connPostgres.QueryRecords(strings.TrimSpace(query), args...)
		//--------------------
	} else {
		//--------------------
		return conn.connSQLite.QueryRecords(strings.TrimSpace(query), args...)
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Close method
//------------------------------------------------------------

func (conn *SQLdb) Close() error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	if conn.DB != nil {
		err = conn.DB.Close()
	}
	//------------------------------------------------------------
	conn.DB = nil
	//------------------------------------------------------------
	_ = conn.connMySQL.Close()
	_ = conn.connPostgres.Close()
	_ = conn.connSQLite.Close()
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// GetSQLTableInfo method
//------------------------------------------------------------

func (conn *SQLdb) GetSQLTableInfo(tableName string) (
	[]struct {
		Sequence int
		Name     string
		Type     string
	},
	map[string]string,
	error,
) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, nil, errors.New("not connected")
	}
	//--------------------
	if strings.ToLower(conn.DBType) != "sqlite" {
		//--------------------
		if conn.Database == "" {
			return nil, nil, errors.New("database name cannot be blank")
		}
		//--------------------
		if !CheckDatabaseName(conn.Database) {
			return nil, nil, errors.New("invalid database name")
		}
		//--------------------
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
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.GetSQLTableInfo(tableName)
		//------------------------------------------------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//------------------------------------------------------------
		return conn.connPostgres.GetSQLTableInfo(tableName)
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.GetSQLTableInfo(tableName)
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetTableInfo method
//------------------------------------------------------------

func (conn *SQLdb) GetTableInfo(tableName string) (
	[]struct {
		Sequence int
		Name     string
		Type     string
	},
	map[string]string,
	error,
) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, nil, errors.New("not connected")
	}
	//--------------------
	if strings.ToLower(conn.DBType) != "sqlite" {
		//--------------------
		if conn.Database == "" {
			return nil, nil, errors.New("database name cannot be blank")
		}
		//--------------------
		if !CheckDatabaseName(conn.Database) {
			return nil, nil, errors.New("invalid database name")
		}
		//--------------------
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
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.GetTableInfo(tableName)
		//------------------------------------------------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//------------------------------------------------------------
		return conn.connPostgres.GetTableInfo(tableName)
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.GetTableInfo(tableName)
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowDatabases method
//------------------------------------------------------------

func (conn *SQLdb) ShowDatabases() ([]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.ShowDatabases()
		//------------------------------------------------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//------------------------------------------------------------
		return conn.connPostgres.ShowDatabases()
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return []string{}, nil
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTables method
//------------------------------------------------------------

func (conn *SQLdb) ShowTables() ([]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.ShowTables()
		//------------------------------------------------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//------------------------------------------------------------
		return conn.connPostgres.ShowTables()
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.ShowTables()
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTablesMap method
//------------------------------------------------------------

func (conn *SQLdb) ShowTablesMap() (map[string]map[string]string, error) {
	//------------------------------------------------------------
	if conn.DB == nil {
		return nil, errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.ShowTablesMap()
		//------------------------------------------------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//------------------------------------------------------------
		return conn.connPostgres.ShowTablesMap()
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.ShowTablesMap()
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// TableExists method
//------------------------------------------------------------

func (conn *SQLdb) TableExists(tableName string) (bool, error) {
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//--------------------
		return conn.connMySQL.TableExists(tableName)
		//--------------------
	} else if strings.ToLower(conn.DBType) == "postgres" {
		//--------------------
		return conn.connPostgres.TableExists(tableName)
		//--------------------
	} else {
		//--------------------
		return conn.connSQLite.TableExists(tableName)
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// NullStringToString
//------------------------------------------------------------

func NullStringToString(value sql.NullString) string {
	//------------------------------------------------------------
	if value.Valid {
		return value.String
	}
	return ""
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

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
// EscapeMySQLString
//------------------------------------------------------------

func EscapeMySQLString(dataString string) string {
	//------------------------------------------------------------
	return mysqldb.EscapeMySQLString(dataString)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// EscapePostgreSQLString
//------------------------------------------------------------

func EscapePostgreSQLString(dataString string) string {
	//------------------------------------------------------------
	return postgresdb.EscapePostgreSQLString(dataString)
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
