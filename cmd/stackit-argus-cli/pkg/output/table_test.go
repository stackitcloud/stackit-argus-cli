package output

import (
	"reflect"
	"testing"
)

type testTable struct {
	Column1 string `header:"column1"`
	Column2 string `header:"column2"`
}

func TestRemoveColumnsFromTable(t *testing.T) {
	table := testTable{
		Column1: "1",
		Column2: "2",
	}

	// test not existed header
	notChanged := RemoveColumnsFromTable(table, []string{"unexportedField"})
	if !reflect.DeepEqual(notChanged, table) {
		t.Fatalf("expected the whole value to be exactly the same, original value(%#+v) should be returned instead of: %#+v", table, notChanged)
	}

	// test with one header
	newTable := RemoveColumnsFromTable(table, []string{"Column2"})
	typ := reflect.TypeOf(newTable)
	f, ok := typ.FieldByName("Column2")
	if !ok {
		t.Fatalf("'Column2' not found")
	}
	if got := string(f.Tag); got != "" {
		t.Fatalf("expected the whole field tag of 'Column2' to be removed, but got: '%s'", got)
	}

	// test with multiple headers
	newTable = RemoveColumnsFromTable(table, []string{"Column1", "Column2"})
	typ = reflect.TypeOf(newTable)
	f, ok = typ.FieldByName("Column1")
	if !ok {
		t.Fatalf("'Column1' not found")
	}
	if got := string(f.Tag); got != "" {
		t.Fatalf("expected the whole field tag of 'Column2' to be removed, but got: '%s'", got)
	}
	f, ok = typ.FieldByName("Column2")
	if !ok {
		t.Fatalf("'Column2' not found")
	}
	if got := string(f.Tag); got != "" {
		t.Fatalf("expected the whole field tag of 'Column2' to be removed, but got: '%s'", got)
	}
}
