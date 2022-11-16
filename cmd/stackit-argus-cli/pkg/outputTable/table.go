package output_table

/*
 * Implements outputTable printer operations.
 */

import (
	"github.com/lensesio/tableprinter"
	"os"
)

// RemoveColumnsFromTable removes columns form the outputTable
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

// PrintTable prints a outputTable
func PrintTable(in interface{}) {
	// init outputTable printer
	printer := tableprinter.New(os.Stdout)

	// customize the outputTable
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
