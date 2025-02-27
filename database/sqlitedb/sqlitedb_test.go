//------------------------------------------------------------

package sqlitedb

import (
	"database/sql"
	"fmt"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/timbrockley/golang-main/file"
)

//--------------------------------------------------------------------------------

var (
	conn1 SQLiteDB
	conn2 SQLiteDB
	conn3 SQLiteDB
)

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
	conn1, err = Connect(SQLiteDB{FilePath: tempFilePath, AutoCreate: false})
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
	conn1, _ = Connect(SQLiteDB{AutoCreate: false})
	//------------------------------------------------------------
	if conn1.FilePath != filePath {
		t.Errorf("filePath = %q but should = %q", conn1.FilePath, filePath)
	}
	//--------------------
	if conn1.Database != filenameBase {
		t.Errorf("database = %q but should = %q", conn1.Database, filenameBase)
	}
	//--------------------
	if conn1.DatabaseExt != "db" {
		t.Errorf("databaseExt = %q but should = %q", conn1.DatabaseExt, "db")
	}
	//------------------------------------------------------------
	// runtime.Caller(0) => this script / runtime.Caller(1) => calling script
	_, filePath, _, _ = runtime.Caller(0)
	filePath = file.FilePathBase(filePath) + ".db"
	filenameBase = file.FilenameBase(filePath)
	//------------------------------------------------------------
	conn1, err = Connect(SQLiteDB{AutoCreate: true, FilePath: filePath})
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {

		//--------------------
		if conn1.FilePath != filePath {
			t.Errorf("filePath = %q but should = %q", conn1.FilePath, filePath)
		}
		//--------------------
		if conn1.Database != filenameBase {
			t.Errorf("database = %q but should = %q", conn1.Database, filenameBase)
		}
		//--------------------
		if conn1.DatabaseExt != "db" {
			t.Errorf("databaseExt = %q but should = %q", conn1.DatabaseExt, "db")
		}
		//--------------------
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
	//--------------------
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
		//--------------------
		for rows.Next() {
			//--------------------
			err := rows.Scan(&id, &name, &price)
			//--------------------
			if err != nil {
				t.Error(err)
			} else {

				//--------------------
				result_count += 1
				row_found = true
				//--------------------
				if id == EXPECTED_id {
					//--------------------
					if name != EXPECTED_name {
						t.Errorf("name = %q but should = %q", name, EXPECTED_name)
					}
					//--------------------
					if price != EXPECTED_price {
						t.Errorf("price = %d but should = %d", price, EXPECTED_price)
					}
					//--------------------
				}
				//--------------------
			}
			//--------------------
		}
		//--------------------
		if result_count != EXPECTED_count {
			t.Errorf("count = %d but should = %d", result_count, EXPECTED_count)
		}
		//--------------------
		if !row_found {
			t.Error("results row not found")
		}
		//--------------------
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
	//--------------------
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
	//--------------------
	if err2 != nil {
		t.Error(err2)
	}
	//------------------------------------------------------------
	if count != EXPECTED_count {
		t.Errorf("count = %d but should = %d", count, EXPECTED_count)
	}
	//--------------------
	if id != EXPECTED_id {
		t.Errorf("id = %d but should = %d", id, EXPECTED_id)
	}
	//--------------------
	if name != EXPECTED_name {
		t.Errorf("name = %q but should = %q", name, EXPECTED_name)
	}
	//--------------------
	if price != EXPECTED_price {
		t.Errorf("price = %d but should = %d", price, EXPECTED_price)
	}
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//------------------------------------------------------------
// QueryRecords
//------------------------------------------------------------

func TestQueryRecords(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var records []map[string]any
	//------------------------------------------------------------
	records, err = conn1.QueryRecords("SELECT * FROM cars")
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		length := 8
		//--------------------
		if len(records) != length {
			t.Errorf("len(records) = %d but should = %d", len(records), length)
			//--------------------
		} else {
			//--------------------
			if !strings.EqualFold(fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]") {
				t.Errorf("records[0] = %q but should = %q", fmt.Sprint(records[0]), "map[id:1 name:Skoda price:9000]")
			}
			//--------------------
			if fmt.Sprintf("%T", records[0]["id"]) != "int" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["id"]), "int")
			}
			//--------------------
			if fmt.Sprintf("%T", records[0]["name"]) != "string" {
				t.Errorf(`records[0]["name"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["name"]), "string")
			}
			//--------------------
			if fmt.Sprintf("%T", records[0]["price"]) != "int" {
				t.Errorf(`records[0]["price"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["price"]), "int")
			}
			//--------------------
		}
		//--------------------
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
	//--------------------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	//--------------------
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = conn1.GetSQLTableInfo("cars")
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		length := 3
		//--------------------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
			//--------------------
		} else {
			//--------------------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}")
			}
			//--------------------
			if fmt.Sprint(columnInfoMap) != "map[id:INTEGER name:VARCHAR(255) price:INT]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INTEGER name:VARCHAR(255) price:INT]")
			}
			//--------------------
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTables
//------------------------------------------------------------

func TestShowTables(t *testing.T) {
	//------------------------------------------------------------
	tables, err := conn1.ShowTables()
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		carsExists := slices.Contains(tables, "cars")
		//--------------------
		if !carsExists {
			t.Errorf(`carsExists = %t but should = %t`, carsExists, true)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTablesMap
//------------------------------------------------------------

func TestShowTablesMap(t *testing.T) {
	//------------------------------------------------------------
	tablesMap, err := conn1.ShowTablesMap()
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		tableInfoMap := tablesMap["cars"]
		//--------------------
		EXPECTED_result := "map[id:INTEGER name:VARCHAR(255) price:INT]"
		//--------------------
		if fmt.Sprint(tableInfoMap) != EXPECTED_result {
			t.Errorf(`tableInfoMap = %s but should = %s`, fmt.Sprint(tableInfoMap), EXPECTED_result)
		}
		//--------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetTableInfo
//------------------------------------------------------------

func TestGetTableInfo(t *testing.T) {
	//------------------------------------------------------------
	var err error
	//--------------------
	var columInfoRows []struct {
		Sequence int
		Name     string
		Type     string
	}
	var columnInfoMap map[string]string
	//------------------------------------------------------------
	columInfoRows, columnInfoMap, err = conn1.GetTableInfo("cars")
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		length := 3
		//--------------------
		if len(columInfoRows) != length {
			t.Errorf("len(columInfoRows) = %d but should = %d", len(columInfoRows), length)
			//--------------------
		} else {
			//--------------------
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id INTEGER}")
			}
			//--------------------
			if fmt.Sprint(columnInfoMap) != "map[id:INTEGER name:VARCHAR(255) price:INT]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INTEGER name:VARCHAR(255) price:INT]")
			}
			//--------------------
		}
		//--------------------
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
		//--------------------
		if result != false {
			t.Errorf("result = %v but should = %v", result, !result)
		}
		//--------------------
	}
	//------------------------------------------------------------
	result, err = conn1.TableExists("cars")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		if result != true {
			t.Errorf("result = %v but should = %v", result, !result)
		}
		//--------------------
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

func TestNullStringToString(t *testing.T) {
	//------------------------------------------------------------
	var data1, data2 sql.NullString
	//------------------------------------------------------------
	data1 = sql.NullString{String: "", Valid: false}     // Represents a null value
	data2 = sql.NullString{String: "test1", Valid: true} // Represents a non-null value
	//------------------------------------------------------------
	result1 := NullStringToString(data1)
	result2 := NullStringToString(data2)
	//--------------------
	EXPECTED_result1 := ""
	EXPECTED_result2 := "test1"
	//------------------------------------------------------------
	if result1 != EXPECTED_result1 {
		t.Errorf("result1 = %s but should = %s", result1, EXPECTED_result1)
	}
	if result2 != EXPECTED_result2 {
		t.Errorf("result2 = %s but should = %s", result2, EXPECTED_result2)
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
	//--------------------
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
	//--------------------
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

//--------------------------------------------------------------------------------
// :memory: database
//--------------------------------------------------------------------------------

func TestMemoryDatabase1(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var count int
	//------------------------------------------------------------
	conn2, err = Connect(SQLiteDB{Database: ":memory:", AutoCreate: false})
	//--------------------------------------------
	if err != nil {
		t.Error(err)
	} else {

		//--------------------------------------------
		stmt := `
			DROP TABLE IF EXISTS cars;
			CREATE TABLE cars(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL);
			INSERT INTO cars(name, price) VALUES('Skoda',9000);
			`
		//--------------------------------------------
		_, err = conn2.Exec(stmt)
		//--------------------------------------------
		if err != nil {
			t.Error(err)
		} else {

			//--------------------------------------------
			row := conn2.QueryRow("SELECT COUNT(*) AS count FROM cars")
			//--------------------------------------------
			err = row.Scan(&count)
			//--------------------------------------------
			if err != nil {
				t.Error(err)
			} else {

				//--------------------------------------------
				EXPECTED_count := 1
				//--------------------------------------------
				if count != EXPECTED_count {
					t.Errorf("count = %d but should = %d", count, EXPECTED_count)
				}
				//--------------------------------------------

			}
			//--------------------------------------------
		}
		//--------------------------------------------
	}
	//--------------------------------------------
	conn2.Close()
	//------------------------------------------------------------
}

func TestMemoryDatabase2(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var count int
	//------------------------------------------------------------
	conn3, err = Connect(SQLiteDB{FilePath: ":memory:", AutoCreate: false})
	//--------------------------------------------
	if err != nil {
		t.Error(err)
	} else {

		//--------------------------------------------
		stmt := `
			DROP TABLE IF EXISTS cars;
			CREATE TABLE cars(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL);
			INSERT INTO cars(name, price) VALUES('Skoda',9000);
			INSERT INTO cars(name, price) VALUES('Audi',52642);
			`
		//--------------------------------------------
		_, err = conn3.Exec(stmt)
		//--------------------------------------------
		if err != nil {
			t.Error(err)
		} else {

			//--------------------------------------------
			row := conn3.QueryRow("SELECT COUNT(*) AS count FROM cars")
			//--------------------------------------------
			err = row.Scan(&count)
			//--------------------------------------------
			if err != nil {
				t.Error(err)
			} else {

				//--------------------------------------------
				EXPECTED_count := 2
				//--------------------------------------------
				if count != EXPECTED_count {
					t.Errorf("count = %d but should = %d", count, EXPECTED_count)
				}
				//--------------------------------------------

			}
			//--------------------------------------------
		}
		//--------------------------------------------
	}
	//--------------------------------------------
	conn3.Close()
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
