//------------------------------------------------------------

package sqldb

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/timbrockley/golang-main/file"
)

//--------------------------------------------------------------------------------

var mysql_conn SQLdbStruct
var sqlite_conn SQLdbStruct

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

func TestConnect(t *testing.T) {
	//------------------------------------------------------------
	// Connect - tests connection function and method
	//------------------------------------------------------------
	var err error
	var db_type, filePath string
	//------------------------------------------------------------

	//------------------------------------------------------------
	db_type = "mysql"
	//------------------------------------------------------------
	mysql_conn, err = Connect(SQLdbStruct{DBType: db_type, Database: "test", AutoCreate: true}, true)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	db_type = "sqlite"
	//------------------------------------------------------------
	// runtime.Caller(0) = this script / runtime.Caller(1) = calling script
	_, filePath, _, _ = runtime.Caller(0)
	//----------
	filePath = file.FilePathBase(filePath) + ".db"
	//------------------------------------------------------------
	sqlite_conn, err = Connect(SQLdbStruct{DBType: db_type, FilePath: filePath, AutoCreate: true}, false)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//----------
		database := file.FilenameBase(filePath)
		database_ext := file.FilenameExt(filePath)
		//----------
		if sqlite_conn.FilePath != filePath {
			t.Errorf("filePath = %q but should = %q", sqlite_conn.FilePath, filePath)
		}
		//----------
		if sqlite_conn.Database != database {
			t.Errorf("Database = %q but should = %q", sqlite_conn.Database, database)
		}
		//----------
		if sqlite_conn.DatabaseExt != database_ext {
			t.Errorf("DatabaseExt = %q but should = %q", sqlite_conn.DatabaseExt, database_ext)
		}
		//----------
	}
	//------------------------------------------------------------
	// env_filePath := os.Getenv("SQLITE_FILEPATH")
	env_dataPath := os.Getenv("SQLITE_DATA_PATH")
	env_database := os.Getenv("SQLITE_DATABASE")
	//------------------------------------------------------------
	DATABASE := "_system"
	DATABASE_EXT := "db"
	//----------
	FILEPATH := file.FilePathJoin(env_dataPath, env_database+"."+DATABASE_EXT)
	//------------------------------------------------------------
	sqlite_conn.FilePath = ""
	//----------
	sqlite_conn.Database = ""
	sqlite_conn.DatabaseExt = ""
	//------------------------------------------------------------
	_ = sqlite_conn.Connect(true)
	//------------------------------------------------------------
	if FILEPATH != "/.db" {
		//----------
		if sqlite_conn.FilePath != FILEPATH {
			t.Errorf("filePath = %q but should = %q", sqlite_conn.FilePath, FILEPATH)
		}
		//----------
	}
	//----------
	if sqlite_conn.Database != DATABASE {
		t.Errorf("Database = %q but should = %q", sqlite_conn.Database, DATABASE)
	}
	//----------
	if sqlite_conn.DatabaseExt != DATABASE_EXT {
		t.Errorf("DatabaseExt = %q but should = %q", sqlite_conn.DatabaseExt, DATABASE_EXT)
	}
	//------------------------------------------------------------
	DATABASE = "test"
	DATABASE_EXT = "db"
	//----------
	FILEPATH = file.FilePathJoin(env_dataPath, DATABASE+"."+DATABASE_EXT)
	//------------------------------------------------------------
	sqlite_conn.FilePath = ""
	//----------
	sqlite_conn.Database = "test"
	sqlite_conn.DatabaseExt = ""
	//------------------------------------------------------------
	_ = sqlite_conn.Connect(true)
	//------------------------------------------------------------
	if FILEPATH != "/.db" {
		//----------
		if sqlite_conn.FilePath != FILEPATH {
			t.Errorf("filePath = %q but should = %q", sqlite_conn.FilePath, FILEPATH)
		}
		//----------
	}
	//----------
	if sqlite_conn.Database != DATABASE {
		t.Errorf("Database = %q but should = %q", sqlite_conn.Database, DATABASE)
	}
	//----------
	if sqlite_conn.DatabaseExt != DATABASE_EXT {
		t.Errorf("DatabaseExt = %q but should = %q", sqlite_conn.DatabaseExt, DATABASE_EXT)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestExec(t *testing.T) {
	//------------------------------------------------------------
	// Exec
	//------------------------------------------------------------
	var err error
	var testData []string
	//------------------------------------------------------------

	//------------------------------------------------------------
	if mysql_conn.DB == nil {
		t.Error("mysql database handle does not exist")
	} else {
		//------------------------------------------------------------
		// mysql
		//------------------------------------------------------------
		testData = []string{
			"CREATE DATABASE IF NOT EXISTS test;",
			"USE test;",
			"DROP TABLE IF EXISTS cars;",
			"CREATE TABLE cars(id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL, PRIMARY KEY(id));",
			"INSERT INTO cars(name,price) VALUES('Skoda',9000);",
		}
		//--------------------------------------------------
		for _, stmt := range testData {
			//------------------------------------------------------------
			_, err = mysql_conn.Exec(stmt)
			//------------------------------------------------------------
			if err != nil {
				t.Error(err)
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
	}

	//------------------------------------------------------------
	if sqlite_conn.DB == nil {
		t.Error("sqlite database handle does not exist")
	} else {
		//------------------------------------------------------------
		// sqlite
		//------------------------------------------------------------
		testData = []string{
			"DROP TABLE IF EXISTS cars;",
			"CREATE TABLE cars(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL);",
			"INSERT INTO cars(name, price) VALUES('Skoda',9000);",
		}
		//--------------------------------------------------
		for _, stmt := range testData {

			//------------------------------------------------------------
			_, err = sqlite_conn.Exec(stmt)
			//------------------------------------------------------------
			if err != nil {

				t.Error(err)
			}
			//------------------------------------------------------------
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestQueryRow(t *testing.T) {
	//------------------------------------------------------------
	// QueryRow
	//------------------------------------------------------------
	var err error
	var row1 *sql.Row
	var count int
	//------------------------------------------------------------
	EXPECTED_count := 1
	//------------------------------------------------------------

	//------------------------------------------------------------
	if mysql_conn.DB == nil {
		t.Error("mysql database handle does not exist")
	} else {
		//------------------------------------------------------------
		// mysql
		//------------------------------------------------------------
		row1 = mysql_conn.QueryRow("SELECT COUNT(*) AS count FROM cars")
		//------------------------------------------------------------
		err = row1.Scan(&count)
		//------------------------------------------------------------
		if err != nil {

			t.Error(err)

		} else {

			if count != EXPECTED_count {

				t.Errorf("count = %d but should = %d", count, EXPECTED_count)
			}
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	if sqlite_conn.DB == nil {
		t.Error("sqlite database handle does not exist")
	} else {
		//------------------------------------------------------------
		// sqlite
		//------------------------------------------------------------
		row1 = sqlite_conn.QueryRow("SELECT COUNT(*) AS count FROM cars")
		//------------------------------------------------------------
		err = row1.Scan(&count)
		//------------------------------------------------------------
		if err != nil {

			t.Error(err)

		} else {

			if count != EXPECTED_count {

				t.Errorf("count = %d but should = %d", count, EXPECTED_count)
			}
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestGetSQLTableInfo(t *testing.T) {
	//------------------------------------------------------------
	// GetSQLTableInfo
	//------------------------------------------------------------
	var err error
	//----------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	//----------
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = mysql_conn.GetSQLTableInfo("cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 3
		//----------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
			//----------
		} else {
			//----------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id int}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id int}")
			}
			//----------
			if fmt.Sprint(columnInfoMap) != "map[id:int name:varchar price:int]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:int name:varchar price:int]")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = sqlite_conn.GetSQLTableInfo("cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 3
		//----------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
			//----------
		} else {
			//----------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}")
			}
			//----------
			if fmt.Sprint(columnInfoMap) != "map[id:INTEGER name:VARCHAR(255) price:INT]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INTEGER name:VARCHAR(255) price:INT]")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------

func TestGetTableInfo(t *testing.T) {
	//------------------------------------------------------------
	// GetTableInfo
	//------------------------------------------------------------
	var err error
	//----------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	//----------
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = mysql_conn.GetTableInfo("cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 3
		//----------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
			//----------
		} else {
			//----------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id int}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id int}")
			}
			//----------
			if fmt.Sprint(columnInfoMap) != "map[id:INT name:VARCHAR price:INT]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INT name:VARCHAR price:INT]")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = sqlite_conn.GetTableInfo("cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 3
		//----------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
		} else {
			//----------
			//----------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}")
			}
			//----------
			if fmt.Sprint(columnInfoMap) != "map[id:INTEGER name:VARCHAR(255) price:INT]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INTEGER name:VARCHAR(255) price:INT]")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------

func TestQueryRecords(t *testing.T) {
	//------------------------------------------------------------
	// QueryRecords
	//------------------------------------------------------------
	var err error
	var records []map[string]any
	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	records, err = mysql_conn.QueryRecords("SELECT * FROM test.cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 1
		//----------
		if len(records) != length {
			t.Errorf("len(records) = %d but should = %d", len(records), length)
			//----------
		} else {
			//----------
			if !strings.EqualFold(fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]") {
				t.Errorf("records[0] = %q but should = %q", fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["id"]) != "int" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["id"]), "int")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["name"]) != "string" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["name"]), "string")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["price"]) != "int" {
				t.Errorf(`records[0]["price"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["price"]), "int")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	records, err = sqlite_conn.QueryRecords("SELECT * FROM cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 1
		//----------
		if len(records) != length {
			t.Errorf("len(records) = %d but should = %d", len(records), length)
			//----------
		} else {
			//----------
			if !strings.EqualFold(fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]") {
				t.Errorf("records[0] = %q but should = %q", fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["id"]) != "int" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["id"]), "int")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["name"]) != "string" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["name"]), "string")
			}
			//----------
			if fmt.Sprintf("%T", records[0]["price"]) != "int" {
				t.Errorf(`records[0]["price"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["price"]), "int")
			}
			//----------
		}
		//----------
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// LockTables
//--------------------------------------------------------------------------------

func TestLockTables(t *testing.T) {

	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------

	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	err = mysql_conn.LockTables("MADE_UP_TABLE_NAME_DFDFDFDFDSFDSFFD")
	//------------------------------------------------------------
	if err == nil {

		t.Error("LockTables should report no tables defined")
	}
	//------------------------------------------------------------
	err = mysql_conn.LockTables("test.cars")
	//----------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	err = sqlite_conn.LockTables()
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------

//--------------------------------------------------------------------------------
// UnlockTables
//--------------------------------------------------------------------------------

func TestUnlockTables(t *testing.T) {

	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------

	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	err = mysql_conn.UnlockTables()
	//----------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	err = sqlite_conn.UnlockTables()
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// TableExists
//--------------------------------------------------------------------------------

func TestTableExists(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var result bool
	//------------------------------------------------------------

	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	result, err = mysql_conn.TableExists("MADE_UP_TABLE_NAME_DFDFDFDFDSFDSFFD")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		if result != false {

			t.Errorf("result = %v but should = %v", result, !result)
		}
		//----------
	}
	//------------------------------------------------------------
	result, err = mysql_conn.TableExists("cars")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		if result != true {

			t.Errorf("result = %v but should = %v", result, !result)
		}
		//----------
	}
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	result, err = sqlite_conn.TableExists("MADE_UP_TABLE_NAME_DFDFDFDFDSFDSFFD")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		if result != false {

			t.Errorf("result = %v but should = %v", result, !result)
		}
		//----------
	}
	//------------------------------------------------------------
	result, err = sqlite_conn.TableExists("cars")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		if result != true {

			t.Errorf("result = %v but should = %v", result, !result)
		}
		//----------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------

func TestCheckTableName(t *testing.T) {
	//------------------------------------------------------------
	var result bool
	//------------------------------------------------------------
	if result = CheckTableName(""); result != false {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
	if result = CheckTableName("aa!!"); result != false {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
	if result = CheckTableName("1a"); result != false {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
	if result = CheckTableName("test1"); result != true {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
	if result = CheckTableName("_table_name"); result != true {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
	if result = CheckTableName("table_name"); result != true {

		t.Errorf("result = %v but should = %v", result, !result)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

func TestEscapeApostrophes(t *testing.T) {
	//------------------------------------------------------------
	result := EscapeApostrophes(`1'2''3`)
	//----------
	EXPECTED_result := `1''2''''3`
	//------------------------------------------------------------
	if result != EXPECTED_result {

		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestEscapeDoubleQuotes(t *testing.T) {
	//------------------------------------------------------------
	result := EscapeDoubleQuotes(`1"2""3`)
	//----------
	EXPECTED_result := `1""2""""3`
	//------------------------------------------------------------
	if result != EXPECTED_result {

		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestEscapeMySQLString(t *testing.T) {
	//------------------------------------------------------------
	result := EscapeMySQLString("\\_\x00_\r_\n_\x1A_\"_'")
	//----------
	EXPECTED_result := `\\_\` + "\x00" + `_\r_\n_\Z_\"_\'`
	//------------------------------------------------------------
	if result != EXPECTED_result {

		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

func TestClose(t *testing.T) {
	//------------------------------------------------------------
	// Close
	//------------------------------------------------------------

	//------------------------------------------------------------
	// mysql
	//------------------------------------------------------------
	mysql_conn.Close()
	//------------------------------------------------------------

	//------------------------------------------------------------
	// sqlite
	//------------------------------------------------------------
	sqlite_conn.Close()
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
