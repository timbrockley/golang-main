//------------------------------------------------------------

package mysqldb

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
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
	conn1, err = Connect(MySQLdbStruct{Database: "test"}, true)
	//------------------------------------------------------------
	if err != nil {

		t.Error(err)
	}
	//------------------------------------------------------------
}

//------------------------------------------------------------
// Exec
//------------------------------------------------------------

func TestExec(t *testing.T) {

	//------------------------------------------------------------
	var err error
	var result sql.Result
	//------------------------------------------------------------
	result, err = connX.Exec("CREATE DATABASE IF NOT EXISTS test;")
	//------------------------------------------------------------
	if result != nil {
		t.Error("result is not nil:", result)
	}
	//----------
	if err == nil {
		t.Errorf("err should = %q but = %q", "not connected", err)
	}
	//------------------------------------------------------------
	result, err = conn1.Exec("CREATE DATABASE IF NOT EXISTS test;")
	//------------------------------------------------------------
	if result == nil {
		t.Error("invalid result")
	}
	//----------
	if err != nil {
		t.Error(err)
	}
	//------------------------------------------------------------
	testData := []string{
		"USE test",
		"DROP TABLE IF EXISTS cars;",
		"CREATE TABLE cars(id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255), price INT DEFAULT 0 NOT NULL, PRIMARY KEY(id));",
		"INSERT INTO cars(name, price) VALUES('Skoda',9000);",
		"INSERT INTO cars(name, price) VALUES('Audi',52642);",
		"INSERT INTO cars(name, price) VALUES('Mercedes',57127);",
		"INSERT INTO cars(name, price) VALUES('Volvo',29000);",
		"INSERT INTO cars(name, price) VALUES('Bentley',350000);",
		"INSERT INTO cars(name, price) VALUES('Citroen',21000);",
		"INSERT INTO cars(name, price) VALUES('Hummer',41400);",
	}
	//--------------------------------------------------
	for _, stmt := range testData {
		//------------------------------------------------------------
		_, err = conn1.Exec(stmt)
		//------------------------------------------------------------
		if err != nil {
			t.Error(err)
		}
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	result, err = conn1.Exec("INSERT INTO cars(name, price) VALUES(?,?);", "Volkswagen", 21600)
	//------------------------------------------------------------
	if result == nil {
		t.Error("invalid result")
	}
	//----------
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
	rows, err = conn1.Query("SELECT * FROM test.cars")
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
	//----------
	EXPECTED_id := 3
	EXPECTED_name := "Mercedes"
	EXPECTED_price := 57127
	//------------------------------------------------------------
	row1 = conn1.QueryRow("SELECT COUNT(*) AS count FROM test.cars")
	row2 = conn1.QueryRow("SELECT * FROM test.cars WHERE id = ?", EXPECTED_id)
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
	columnInfoMap := map[string]string{}
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
	//----------
	columnInfoMap := map[string]string{}
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
}

//------------------------------------------------------------
// QueryRecords
//------------------------------------------------------------

func TestQueryRecords(t *testing.T) {
	//------------------------------------------------------------
	var err error
	var records []map[string]any
	//------------------------------------------------------------
	records, err = conn1.QueryRecords("SELECT * FROM test.cars")
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

//------------------------------------------------------------
// Close
//------------------------------------------------------------

func TestClose(t *testing.T) {

	//------------------------------------------------------------
	conn1.Close()
	//------------------------------------------------------------
}

//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
