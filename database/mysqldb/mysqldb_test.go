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

var (
	conn1 MySQLdb
	connX MySQLdb

	signalChannel = map[string]chan bool{}
)

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------

//------------------------------------------------------------
// setup channel helpers
//------------------------------------------------------------

func init() {
	sendSignal("TestConnect")
}

func sendSignal(name string) {
	signalChannel[name] = make(chan bool)
	go func() { signalChannel[name] <- true }()
}

func receiveSignal(name string) {
	<-signalChannel[name]
	// delete(signalChannel, name)
}

//------------------------------------------------------------
// Connect
//------------------------------------------------------------

func TestConnect(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestConnect")
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	conn1, err = Connect(MySQLdb{Database: "MADE_UP_NAME_FDSDFDDVDHIFHDIH"}, true)
	//------------------------------------------------------------
	if err == nil {
		t.Error("Connect should fail if database does not exist")
	}
	//------------------------------------------------------------
	conn1, err = Connect(MySQLdb{Database: "test", AutoCreate: true}, true)
	//------------------------------------------------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
	sendSignal("TestExec1")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec1
//------------------------------------------------------------

func TestExec1(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestExec1")
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
	time.Sleep(1 * time.Second)
	//------------------------------------------------------------
	sendSignal("TestExec2")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec2
//------------------------------------------------------------

func TestExec2(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestExec2")
	//------------------------------------------------------------
	var err error
	//------------------------------------------------------------
	_, err = conn1.Exec(`
		USE test;
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
	sendSignal("TestQuery")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Query
//------------------------------------------------------------

func TestQuery(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestQuery")
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
	sendSignal("TestQueryRow")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRow
//------------------------------------------------------------

func TestQueryRow(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestQueryRow")
	//------------------------------------------------------------
	var err1 error
	var row1 *sql.Row
	//------------------------------------------------------------
	var id int
	var name string
	var price int
	//------------------------------------------------------------
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	row1 = conn1.QueryRow("SELECT * FROM test.cars WHERE name = ?", EXPECTED_name)
	//------------------------------------------------------------
	err1 = row1.Scan(&id, &name, &price)
	//------------------------------------------------------------
	if err1 != nil {
		t.Error(err1)
	}
	//------------------------------------------------------------
	if name != EXPECTED_name {
		t.Errorf("name = %q but should = %q", name, EXPECTED_name)
	}
	//--------------------
	if price != EXPECTED_price {
		t.Errorf("price = %d but should = %d", price, EXPECTED_price)
	}
	//------------------------------------------------------------
	sendSignal("TestQueryRecords")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// QueryRecords
//------------------------------------------------------------

func TestQueryRecords(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestQueryRecords")
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
	sendSignal("TestGetSQLTableInfo")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetSQLTableInfo
//------------------------------------------------------------

func TestGetSQLTableInfo(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestGetSQLTableInfo")
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
	sendSignal("TestGetTableInfo")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// GetTableInfo
//------------------------------------------------------------

func TestGetTableInfo(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestGetTableInfo")
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
	sendSignal("TestShowDatabases")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowDatabases
//------------------------------------------------------------

func TestShowDatabases(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestShowDatabases")
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
	sendSignal("TestShowTables")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTables
//------------------------------------------------------------

func TestShowTables(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestShowTables")
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
	sendSignal("TestShowTablesMap")
	//------------------------------------------------------------
}

//------------------------------------------------------------
// ShowTablesMap
//------------------------------------------------------------

func TestShowTablesMap(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestShowTablesMap")
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
	sendSignal("TestLockTables")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// LockTables
//--------------------------------------------------------------------------------

func TestLockTables(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestLockTables")
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
	sendSignal("TestUnlockTables")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// UnlockTables
//--------------------------------------------------------------------------------

func TestUnlockTables(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestUnlockTables")
	//------------------------------------------------------------
	err := conn1.UnlockTables()
	//--------------------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
	sendSignal("TestTableExists")
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
	receiveSignal("TestTableExists")
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
	sendSignal("TestCheckTableName")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// CheckTableName
//--------------------------------------------------------------------------------

func TestCheckTableName(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestCheckTableName")
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
	sendSignal("TestNullStringToString")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------

//--------------------------------------------------------------------------------
// TestNullStringToString
//--------------------------------------------------------------------------------

func TestNullStringToString(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestNullStringToString")
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
	sendSignal("TestEscapeApostrophes")
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
	receiveSignal("TestEscapeApostrophes")
	//------------------------------------------------------------
	result := EscapeApostrophes(`1'2''3`)
	//--------------------
	EXPECTED_result := `1''2''''3`
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	sendSignal("TestEscapeDoubleQuotes")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
// EscapeDoubleQuotes
//--------------------------------------------------------------------------------

func TestEscapeDoubleQuotes(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestEscapeDoubleQuotes")
	//------------------------------------------------------------
	result := EscapeDoubleQuotes(`1"2""3`)
	//--------------------
	EXPECTED_result := `1""2""""3`
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	sendSignal("TestEscapeMySQLString")
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------

func TestEscapeMySQLString(t *testing.T) {
	//------------------------------------------------------------
	receiveSignal("TestEscapeMySQLString")
	//------------------------------------------------------------
	result := EscapeMySQLString("\\_\x00_\r_\n_\x1A_\"_'")
	//--------------------
	EXPECTED_result := `\\_\` + "\x00" + `_\r_\n_\Z_\"_\'`
	//------------------------------------------------------------
	if result != EXPECTED_result {
		t.Errorf("result = %s but should = %s", result, EXPECTED_result)
	}
	//------------------------------------------------------------
	sendSignal("TestClose")
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
	receiveSignal("TestClose")
	//------------------------------------------------------------
	conn1.Close()
	//------------------------------------------------------------
}

//--------------------------------------------------------------------------------
//################################################################################
//--------------------------------------------------------------------------------
