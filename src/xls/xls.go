package xls

import (
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

func ExportXls(data [][]string, opts ...OptionFunc) (*excelize.File, error) {
	options := parseOptionFunc(opts)

	xls := excelize.NewFile()
	if options.resetSheetName {
		xls.SetSheetName(DefaultSheetName, options.sheetName)
	}

	for rowIdx, row := range data {
		rowDest, axisFmtList := AxisFmtList(row)
		for colIdx, cell := range rowDest {
			if err := xls.SetCellValue(options.sheetName, fmt.Sprintf(axisFmtList[colIdx], rowIdx+1), cell); err != nil {
				return xls, err
			}
		}
	}
	return xls, nil
}

func ParseXls(reader io.Reader, opts ...OptionFunc) ([][]string, error) {
	xls, err := excelize.OpenReader(reader)
	if err != nil {
		err = fmt.Errorf("open xls failed, err=%w", err)
		return nil, err
	}
	// Get all the rows in the first Sheet1.
	firstSheetName := xls.GetSheetName(FirstSheetIdx)
	rows, err := xls.GetRows(firstSheetName) // must exist, empty xls?
	if err != nil {
		return nil, fmt.Errorf("get_rows failed, err=%w", err)
	}
	if len(rows) <= 1 {
		return nil, fmt.Errorf("no records")
	}
	return rows[1:], nil
}
