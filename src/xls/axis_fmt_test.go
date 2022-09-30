package xls

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAxisFmtList(t *testing.T) {
	var rowEmpty []string
	rowT10 := []string{"", "", "", "", "", "", "", "", "", ""}
	rowT20 := append(rowT10, rowT10...)
	rowT26 := append(rowT20, []string{"", "", "", "", "", ""}...)
	rowT27 := append(rowT26, "")
	rowT60 := append(rowT20, append(rowT20, rowT20...)...)
	rowT100 := append(rowT60, append(rowT20, rowT20...)...)
	rowT200 := append(rowT100, rowT100...)
	rowT400 := append(rowT200, rowT200...)
	rowT700 := append(rowT400, append(rowT200, rowT100...)...)
	rowT701 := append(rowT700, "")
	tests := []struct {
		name string
		row  []string
		want int
	}{
		{"case 0", rowEmpty, 26},
		{"case 0", rowT10, 26},
		{"case 0", rowT26, 26},
		{"case 0", rowT27, 52},
		{"case 0", rowT60, 78},
		{"case 0", rowT100, 104},
		{"case 0", rowT700, 700},
		{"case 0", rowT701, 700}, // should panic
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, got := AxisFmtList(tt.row)
			data := strings.Join(got, ", ")
			fmt.Println(data)
			fmt.Println(row)
			assert.Equal(t, tt.want, len(got))
		})
	}
}
