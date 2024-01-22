//------------------------------------------------------------

package sqlitedb

import (
	"database/sql"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/timbrockley/golang-main/file"
)

//--------------------------------------------------------------------------------

var conn1 SQLiteDBStruct

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// Connect
//--------------------------------------------------------------------------------

func TestConnect(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var tempFilePath, filePath, filenameBase string
	//------------------------------------------------------------
	tempFilePath, _ = file.TempFilePath()
	//------------------------------------------------------------
	conn1, err = Connect(SQLiteDBStruct{FilePath: tempFilePath, AutoCreate: false})
	//------------------------------------------------------------
	if err == nil {

		t.Error("error not returned for made up database name using autoCreate=false")
	}
	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ = runtime.Caller(0)
	filePath = file.FilePathBase(filePath) + ".db"
	filePath = strings.Replace(filePath, `_test`, "", -1)
	filenameBase = file.FilenameBase(filePath)
	//------------------------------------------------------------
	conn1, _ = Connect(SQLiteDBStruct{AutoCreate: false})
	//------------------------------------------------------------
	if conn1.FilePath != filePath {

		t.Errorf("filePath = %q but should = %q", conn1.FilePath, filePath)
	}
	//----------
	if conn1.Database != filenameBase {

		t.Errorf("database = %q but should = %q", conn1.Database, filenameBase)
	}
	//----------
	if conn1.DatabaseExt != "db" {

		t.Errorf("databaseExt = %q but should = %q", conn1.DatabaseExt, "db")
	}
	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ = runtime.Caller(0)
	filePath = file.FilePathBase(filePath) + ".db"
	filenameBase = file.FilenameBase(filePath)
	//------------------------------------------------------------
	conn1, err = Connect(SQLiteDBStruct{AutoCreate: true, FilePath: filePath})
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//----------
		if conn1.FilePath != filePath {

			t.Errorf("filePath = %q but should = %q", conn1.FilePath, filePath)
		}
		//----------
		if conn1.Database != filenameBase {

			t.Errorf("database = %q but should = %q", conn1.Database, filenameBase)
		}
		//----------
		if conn1.DatabaseExt != "db" {

			t.Errorf("databaseExt = %q but should = %q", conn1.DatabaseExt, "db")
		}
		//----------
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// Exec
//--------------------------------------------------------------------------------

func TestExec(t *testing.T) {

	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	stmt := `
	DROP TABLE IF EXISTS cars;
	CREATE TABLE cars(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL);
	INSERT INTO cars(name, price) VALUES('Skoda',9000);
	INSERT INTO cars(name, price) VALUES('Audi',52642);
	INSERT INTO cars(name, price) VALUES('Mercedes',57127);
	INSERT INTO cars(name, price) VALUES('Volvo',29000);
	INSERT INTO cars(name, price) VALUES('Bentley',350000);
	INSERT INTO cars(name, price) VALUES('Citroen',21000);
	INSERT INTO cars(name, price) VALUES('Hummer',41400);
	INSERT INTO cars(name, price) VALUES('Volkswagen',21600);
	`
	//------------------------------------------------------------
	_, err = conn1.Exec(stmt)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// Query
//--------------------------------------------------------------------------------

func TestQuery(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	EXPECTED_count := 8
	//----------
	EXPECTED_id := 3
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	var id int
	var name string
	var price int
	//------------------------------------------------------------
	rows, err = conn1.Query("SELECT * FROM cars")
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)

	} else {

		//------------------------------------------------------------
		defer rows.Close()
		//------------------------------------------------------------
		result_count := 0
		row_found := false
		//----------
		for rows.Next() {
			//----------
			err := rows.Scan(&id, &name, &price)
			//----------
			if err != nil {

				t.Error(err)

			} else {

				//----------
				result_count += 1
				row_found = true
				//----------
				if id == EXPECTED_id {
					//----------
					if name != EXPECTED_name {

						t.Errorf("name = %q but should = %q", name, EXPECTED_name)
					}
					//----------
					if price != EXPECTED_price {

						t.Errorf("price = %d but should = %d", price, EXPECTED_price)
					}
					//----------
				}
				//----------
			}
			//----------
		}
		//----------
		if result_count != EXPECTED_count {

			t.Errorf("count = %d but should = %d", result_count, EXPECTED_count)
		}
		//----------
		if !row_found {

			t.Error("results row not found")
		}
		//----------
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// QueryRow
//--------------------------------------------------------------------------------

func TestQueryRow(t *testing.T) {

	//------------------------------------------------------------
	var err1, err2 error
	var row1, row2 *sql.Row
	var count int
	//------------------------------------------------------------
	var id int
	var name string
	var price int
	//------------------------------------------------------------
	EXPECTED_count := 8
	//----------
	EXPECTED_id := 3
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	row1 = conn1.QueryRow("SELECT COUNT(*) AS count FROM cars")
	row2 = conn1.QueryRow("SELECT * FROM cars WHERE id = ?", EXPECTED_id)
	//------------------------------------------------------------
	err1 = row1.Scan(&count)
	err2 = row2.Scan(&id, &name, &price)
	//------------------------------------------------------------
	if err1 != nil {

		t.Error(err1)
	}
	//----------
	if err2 != nil {

		t.Error(err2)
	}
	//------------------------------------------------------------
	if count != EXPECTED_count {

		t.Errorf("count = %d but should = %d", count, EXPECTED_count)
	}
	//----------
	if id != EXPECTED_id {

		t.Errorf("id = %d but should = %d", id, EXPECTED_id)
	}
	//----------
	if name != EXPECTED_name {

		t.Errorf("name = %q but should = %q", name, EXPECTED_name)
	}
	//----------
	if price != EXPECTED_price {

		t.Errorf("price = %d but should = %d", price, EXPECTED_price)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//------------------------------------------------------------
// GetSQLTableInfo
//------------------------------------------------------------

func TestGetSQLTableInfo(t *testing.T) {
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
	columInfoRows, columnInfoMap, err = conn1.GetSQLTableInfo("cars")
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
// GetTableInfo
//------------------------------------------------------------

func TestGetTableInfo(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//----------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = conn1.GetTableInfo("cars")
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
// QueryRecords
//------------------------------------------------------------

func TestQueryRecords(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var records []map[string]any
	//------------------------------------------------------------
	records, err = conn1.QueryRecords("SELECT * FROM cars")
	//----------
	if err != nil {
		t.Error(err)
	} else {
		//----------
		length := 8
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
	err = conn1.LockTables()
	//----------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// UnlockTables
//--------------------------------------------------------------------------------

func TestUnlockTables(t *testing.T) {
	//------------------------------------------------------------
	err := conn1.UnlockTables()
	//----------
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
	result, err = conn1.TableExists("MADE_UP_TABLE_NAME_DFDFDFDFDSFDSFFD")
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
	result, err = conn1.TableExists("cars")
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

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// CheckTableName
//--------------------------------------------------------------------------------

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

//--------------------------------------------------------------------------------
// EscapeApostrophes
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
// EscapeDoubleQuotes
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
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// Close
//--------------------------------------------------------------------------------

func TestClose(t *testing.T) {

	//------------------------------------------------------------
	conn1.Close()
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
