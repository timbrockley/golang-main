//------------------------------------------------------------

package mysqldb

import (
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"
)

//------------------------------------------------------------

var conn1 MySQLdbStruct
var connX MySQLdbStruct

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// Connect
//------------------------------------------------------------

func TestConnect(t *testing.T) {

	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	conn1, err = Connect(MySQLdbStruct{Database: "MADE_UP_NAME_FDSDFDDVDHIFHDIH"}, true)
	//------------------------------------------------------------
	if err == nil {

		t.Error("Connect should fail if database does not exist")
	}
	//------------------------------------------------------------
	conn1, err = Connect(MySQLdbStruct{Database: "test", AutoCreate: true}, true)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec1
//------------------------------------------------------------

func TestExec1(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var result sql.Result
	//------------------------------------------------------------
	result, err = connX.Exec("CREATE DATABASE IF NOT EXISTS test;")
	//------------------------------------------------------------
	if result != nil {
		t.Error("result is not nil:", result)
	}
	//--------------------
	if err == nil {
		t.Errorf("err should = %q but = %q", "not connected", err)
	}
	//------------------------------------------------------------
	result, err = conn1.Exec("CREATE DATABASE IF NOT EXISTS test;")
	//------------------------------------------------------------
	if result == nil {
		t.Error("invalid result")
	}
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
	_, err = conn1.Exec(`
		USE test;
		DROP TABLE IF EXISTS cars;
		CREATE TABLE cars(id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL, PRIMARY KEY(id));
	`)
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec2
//------------------------------------------------------------

func TestExec2(t *testing.T) {

	//------------------------------------------------------------
	time.Sleep(1 * time.Second)
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	_, err = conn1.Exec(`
		USE test;
		DELETE FROM cars;
		INSERT INTO cars(name, price) VALUES('Skoda',9000);
		INSERT INTO cars(name, price) VALUES('Audi',52642);
		INSERT INTO cars(name, price) VALUES('Mercedes',57127);
		INSERT INTO cars(name, price) VALUES('Volvo',29000);
		INSERT INTO cars(name, price) VALUES('Bentley',350000);
		INSERT INTO cars(name, price) VALUES('Citroen',21000);
		INSERT INTO cars(name, price) VALUES('Hummer',41400);
		INSERT INTO cars(name, price) VALUES('Volkswagen', 21600);
	`)
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Query
//------------------------------------------------------------

func TestQuery(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var rows *sql.Rows
	//------------------------------------------------------------
	EXPECTED_count := 1
	//--------------------
	EXPECTED_id := 3
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	var id int
	var name string
	var price int
	//------------------------------------------------------------
	rows, err = conn1.Query("SELECT * FROM test.cars WHERE name = ?", EXPECTED_name)
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

//------------------------------------------------------------
// QueryRow
//------------------------------------------------------------

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
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	row1 = conn1.QueryRow("SELECT COUNT(*) AS count FROM test.cars")
	row2 = conn1.QueryRow("SELECT * FROM test.cars WHERE name = ?", EXPECTED_name)
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
	if name != EXPECTED_name {

		t.Errorf("name = %q but should = %q", name, EXPECTED_name)
	}
	//--------------------
	if price != EXPECTED_price {

		t.Errorf("price = %d but should = %d", price, EXPECTED_price)
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
	records, err = conn1.QueryRecords("SELECT * FROM test.cars WHERE name=?", "Skoda")
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		length := 1
		//--------------------
		if len(records) != length {
			t.Errorf("len(records) = %d but should = %d", len(records), length)
			//--------------------
		} else {
			//--------------------
			if fmt.Sprintf("%T", records[0]["id"]) != "int" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["id"]), "int")
			}
			//--------------------
			if fmt.Sprintf("%T", records[0]["name"]) != "string" {
				t.Errorf(`records[0]["id"] type = %q but should = %q`, fmt.Sprintf("%T", records[0]["name"]), "string")
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
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id int}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id int}")
			}
			//--------------------
			if fmt.Sprint(columnInfoMap) != "map[id:int name:varchar price:int]" {
				t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:int name:varchar price:int]")
			}
			//--------------------
		}
		//------------------------------------------------------------
		if fmt.Sprint(columnInfoMap) != "map[id:int name:varchar price:int]" {
			t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:int name:varchar price:int]")
		}
		//------------------------------------------------------------
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
	//--------------------
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
			if !strings.EqualFold(fmt.Sprint(columInfoRows[0]), "{1 id int}") {
				t.Errorf("columInfoRows[0] = %q but should = %q", fmt.Sprint(columInfoRows[0]), "{1 id int}")
			}
			//--------------------
		}
		//------------------------------------------------------------
		if fmt.Sprint(columnInfoMap) != "map[id:INT name:VARCHAR price:INT]" {
			t.Errorf("columnInfoMap = %q but should = %q", fmt.Sprint(columnInfoMap), "map[id:INT name:VARCHAR price:INT]")
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowDatabases
//------------------------------------------------------------

func TestShowDatabases(t *testing.T) {
	//------------------------------------------------------------
	databases, err := conn1.ShowDatabases()
	//--------------------
	if err != nil {
		t.Error(err)
	} else {
		//--------------------
		testExists := slices.Contains(databases, "test")
		//--------------------
		if !testExists {
			t.Errorf(`carsExists = %t but should = %t`, testExists, true)
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
		EXPECTED_result := "map[id:INT name:VARCHAR price:INT]"
		//--------------------
		if fmt.Sprint(tableInfoMap) != EXPECTED_result {
			t.Errorf(`tableInfoMap = %s but should = %s`, fmt.Sprint(tableInfoMap), EXPECTED_result)
		}
		//--------------------
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
	//--------------------
	if err == nil {

		t.Error("LockTables should report no tables defined")
	}
	//------------------------------------------------------------
	err = conn1.LockTables("")
	//--------------------
	if err == nil {

		t.Error("LockTables should report invalid table name")
	}
	//------------------------------------------------------------
	err = conn1.LockTables("1")
	//--------------------
	if err == nil {

		t.Error("LockTables should report invalid table name")
	}
	//------------------------------------------------------------
	err = conn1.LockTables("test.cars")
	//--------------------
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
	//--------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

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

func TestEscapeMySQLString(t *testing.T) {
	//------------------------------------------------------------
	result := EscapeMySQLString("\\_\x00_\r_\n_\x1A_\"_'")
	//--------------------
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

//------------------------------------------------------------
// Close
//------------------------------------------------------------

func TestClose(t *testing.T) {

	//------------------------------------------------------------
	conn1.Close()
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
