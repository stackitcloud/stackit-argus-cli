package output_table

/*
 * Implements output_table printer operations.
 */

import (
	"github.com/lensesio/tableprinter"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
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

// customizeTabs customizes a table
func customizeTabs(printer *tableprinter.Printer) {
	printer.HeaderLine = false
	printer.RowCharLimit = 30
	printer.RowTextWrap = true
	printer.NumbersAlignment = tableprinter.AlignLeft
	printer.DefaultAlignment = tableprinter.AlignLeft
}

// customizeTable customizes a table
func customizeTable(printer *tableprinter.Printer) {
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.RowLine = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.HeaderLine = false
	printer.RowCharLimit = 30
	printer.RowTextWrap = true
	printer.HeaderAlignment = tableprinter.AlignCenter
	printer.NumbersAlignment = tableprinter.AlignLeft
	printer.DefaultAlignment = tableprinter.AlignLeft
}

// PrintTable prints a output table
func PrintTable(in interface{}, outputType config.OutputType) {
	// init output_table printer
	printer := tableprinter.New(os.Stdout)

	if outputType == "table" || outputType == "wide-table" {
		customizeTable(printer)
	} else {
		customizeTabs(printer)
	}

	printer.Print(in)
}
