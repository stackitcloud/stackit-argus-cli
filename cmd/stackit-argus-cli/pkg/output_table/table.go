package output_table

/*
 * Implements output_table printer operations.
 */

import (
	"github.com/lensesio/tableprinter"
	"os"
)

// RemoveColumnsFromTable removes columns form the output_table
func RemoveColumnsFromTable(originalTable interface{}, fieldNames []string) interface{} {
	if len(fieldNames) == 0 {
		return originalTable
	}

	newTable := tableprinter.RemoveStructHeader(originalTable, fieldNames[0])
	for i := 1; i < len(fieldNames); i++ {
		newTable = tableprinter.RemoveStructHeader(newTable, fieldNames[i])
	}

	return newTable
}

// PrintTable prints a output_table
func PrintTable(in interface{}) {
	// init output_table printer
	printer := tableprinter.New(os.Stdout)

	// customize the output_table
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.RowLine = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.HeaderAlignment = tableprinter.AlignCenter
	printer.DefaultAlignment = tableprinter.AlignCenter
	printer.RowCharLimit = 30
	printer.RowTextWrap = true

	printer.Print(in)
}
