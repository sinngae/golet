package xls

import "sync"

const (
	DefaultSheetName = "Sheet1"
	FirstSheetIdx    = 0
	ColumnSizeMax    = 700 // excel/xls column size limit is 16384, but 700 axis format is enough
)

var (
	alphabetList = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	axisFormatList []string
	axisFmtBatch   int
	axisFormatOnce sync.Once
	axisFormatLock sync.Mutex
)

// AxisFmtList PAY ATTENTION: support 700 columns only
func AxisFmtList(row []string) ([]string, []string) {
	axisFormatOnce.Do(
		func() {
			for _, item := range alphabetList {
				axisFormatList = append(axisFormatList, string(item)+"%d")
			}
			axisFmtBatch = 1
		},
	)

	columnCnt := len(row)
	if columnCnt > ColumnSizeMax {
		// panic("reach xls column size limit") // almost never happen
		row = row[:ColumnSizeMax]
	}

	axisFormatLock.Lock()
	defer axisFormatLock.Unlock()
	if columnCnt > len(axisFormatList) {
		batch := (columnCnt + len(alphabetList) - 1) / len(alphabetList)
		for i := axisFmtBatch; i < batch; i++ {
			idx := i - 1
			for _, item := range alphabetList {
				axisFormatList = append(axisFormatList, string(alphabetList[idx])+string(item)+"%d")
			}
		}
		axisFmtBatch = batch
	}

	return row, axisFormatList
}
