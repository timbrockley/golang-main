//------------------------------------------------------------

package sqldb

import (
	"database/sql"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/timbrockley/golang-main/database/mysqldb"
	"github.com/timbrockley/golang-main/database/sqlitedb"
	"github.com/timbrockley/golang-main/file"
	"github.com/timbrockley/golang-main/system"
)

//------------------------------------------------------------

type SQLdbStruct struct {
	//----------
	DBType string
	//----------
	Host string
	//----------
	User     string
	Password string
	//----------
	AllowNativePasswords bool
	//----------
	Database string
	//----------
	DatabaseExt string
	//----------
	FilePath string
	//----------
	AutoCreate bool
	//----------
	DB *sql.DB
	//----------
	connMySQL  mysqldb.MySQLdbStruct
	connSQLite sqlitedb.SQLiteDBStruct
	//----------
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

func (conn *SQLdbStruct) Connect(checkENV ...bool) error {
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	conn.DB = nil
	//------------------------------------------------------------
	if conn.DBType == "" && checkENV != nil && checkENV[0] {
		//----------
		conn.DBType = os.Getenv("DB_TYPE")
		//----------
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {

		//------------------------------------------------------------
		if checkENV != nil && checkENV[0] {
			//------------------------------------------------------------
			conn.Host = os.Getenv("MYSQL_HOST")
			//----------
			conn.User = os.Getenv("MYSQL_USER")
			conn.Password = os.Getenv("MYSQL_PWD")
			//----------
			conn.AllowNativePasswords = os.Getenv("MYSQL_ALLOW_NATIVE_PASSWORDS") == "true"
			//----------
			if conn.Database == "" {
				conn.Database = os.Getenv("MYSQL_DATABASE")
			}
			//----------
			if conn.Database == "" {
				conn.Database = "_system"
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
		conn.connMySQL = mysqldb.MySQLdbStruct{Host: conn.Host, User: conn.User, Password: conn.Password, AllowNativePasswords: conn.AllowNativePasswords, Database: conn.Database, AutoCreate: conn.AutoCreate}
		//----------
		err = conn.connMySQL.Connect()
		//------------------------------------------------------------
		if err == nil {
			//----------
			conn.DB = conn.connMySQL.DB
			//----------
		}
		//------------------------------------------------------------

	} else {

		//------------------------------------------------------------
		// if file path is not passed then check environment variables
		//------------------------------------------------------------
		if conn.FilePath != "" {
			//----------
			conn.Database = ""
			conn.DatabaseExt = ""
			//----------

		} else {

			//------------------------------------------------------------
			var filePath, dataPath, database string
			//------------------------------------------------------------
			if checkENV != nil && checkENV[0] {
				//----------
				filePath = os.Getenv("SQLITE_FILEPATH")
				dataPath = os.Getenv("SQLITE_DATA_PATH")
				//----------
				if conn.Database == "" {
					database = os.Getenv("SQLITE_DATABASE")
				}
				//----------
			}
			//------------------------------------------------------------
			if filePath != "" {
				//----------
				conn.FilePath = filePath
				conn.Database = ""
				conn.DatabaseExt = ""
				//----------
			} else {
				//----------
				if conn.Database == "" {
					//----------
					if database != "" {
						//----------
						conn.Database = database
						conn.DatabaseExt = "db"
						//----------
					} else {
						//----------
						conn.Database = "_system"
						conn.DatabaseExt = "db"
						//----------
					}
					//----------
				}
				//----------
				if conn.DatabaseExt == "" {
					//----------
					conn.DatabaseExt = "db"
					//----------
				}
				//----------
				if dataPath != "" {
					//----------
					conn.FilePath = file.FilePathJoin(dataPath, conn.Database+"."+conn.DatabaseExt)
					//----------
					conn.Database = ""
					conn.DatabaseExt = ""
					//----------
				}
				//----------
			}
			//----------
		}
		//------------------------------------------------------------
		conn.connSQLite = sqlitedb.SQLiteDBStruct{FilePath: conn.FilePath, Database: conn.Database, DatabaseExt: conn.DatabaseExt, AutoCreate: conn.AutoCreate}
		//----------
		err = conn.connSQLite.Connect()
		//------------------------------------------------------------
		conn.FilePath = conn.connSQLite.FilePath
		//----------
		conn.Database = conn.connSQLite.Database
		conn.DatabaseExt = conn.connSQLite.DatabaseExt
		//----------
		if err == nil {
			//----------
			conn.DB = conn.connSQLite.DB
			//----------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	return err
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Connect - connect to database and return passed SQLdbStruct Object
//------------------------------------------------------------
// conn, err = Connect(SQLdbStruct{ }, checkENV)
//------------------------------------------------------------

func Connect(conn SQLdbStruct, checkENV ...bool) (SQLdbStruct, error) {
	//------------------------------------------------------------
	return conn, conn.Connect(checkENV...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec method
//------------------------------------------------------------

func (conn *SQLdbStruct) Exec(query string, args ...any) (sql.Result, error) {
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

func (conn *SQLdbStruct) Query(query string, args ...any) (*sql.Rows, error) {
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

func (conn *SQLdbStruct) QueryRow(query string, args ...any) *sql.Row {
	//------------------------------------------------------------
	return conn.DB.QueryRow(strings.TrimSpace(query), args...)
	//------------------------------------------------------------
}

//------------------------------------------------------------
// LockTables method
//------------------------------------------------------------

func (conn *SQLdbStruct) LockTables(Tables ...string) error {
	//------------------------------------------------------------
	if conn.DB == nil {
		return errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.LockTables(Tables...)
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.LockTables()
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// UnlockTables method
//------------------------------------------------------------

func (conn *SQLdbStruct) UnlockTables() error {
	//------------------------------------------------------------
	if conn.DB == nil {
		return errors.New("not connected")
	}
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.UnlockTables()
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.UnlockTables()
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Close method
//------------------------------------------------------------

func (conn *SQLdbStruct) Close() error {
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

func (conn *SQLdbStruct) GetSQLTableInfo(tableName string) (
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
	if strings.ToLower(conn.DBType) != "sqlite" {
		//----------
		if conn.Database == "" {
			return nil, nil, errors.New("database name cannot be blank")
		}
		//----------
		if !CheckDatabaseName(conn.Database) {
			return nil, nil, errors.New("invalid database name")
		}
		//----------
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
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.GetSQLTableInfo(tableName)
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

func (conn *SQLdbStruct) GetTableInfo(tableName string) (
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
	if strings.ToLower(conn.DBType) != "sqlite" {
		//----------
		if conn.Database == "" {
			return nil, nil, errors.New("database name cannot be blank")
		}
		//----------
		if !CheckDatabaseName(conn.Database) {
			return nil, nil, errors.New("invalid database name")
		}
		//----------
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
	if strings.ToLower(conn.DBType) == "mysql" {
		//------------------------------------------------------------
		return conn.connMySQL.GetTableInfo(tableName)
		//------------------------------------------------------------
	} else {
		//------------------------------------------------------------
		return conn.connSQLite.GetTableInfo(tableName)
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords method
//------------------------------------------------------------

func (conn *SQLdbStruct) QueryRecords(query string, args ...any) ([]map[string]any, error) {
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//----------
		return conn.connMySQL.QueryRecords(strings.TrimSpace(query), args...)
		//----------
	} else {
		//----------
		return conn.connSQLite.QueryRecords(strings.TrimSpace(query), args...)
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// TableExists method
//------------------------------------------------------------

func (conn *SQLdbStruct) TableExists(tableName string) (bool, error) {
	//------------------------------------------------------------
	if strings.ToLower(conn.DBType) == "mysql" {
		//----------
		return conn.connMySQL.TableExists(tableName)
		//----------
	} else {
		//----------
		return conn.connSQLite.TableExists(tableName)
		//----------
	}
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
	return mysqldb.EscapeMySQLString(dataString)
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
//############################################################
//------------------------------------------------------------
