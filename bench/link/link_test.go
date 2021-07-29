package lane_build

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	rscA1 = &Resource{
		name:        "A1",
		nextRscList: []string{"B2"},
	}
	rscA2 = &Resource{
		name:        "A2",
		nextRscList: []string{},
	}
	rscA3 = &Resource{
		name:        "A3",
		nextRscList: []string{"B2"},
	}
	rscB1 = &Resource{
		name:        "B1",
		nextRscList: []string{"C2"},
	}
	rscB2 = &Resource{
		name:        "B2",
		nextRscList: []string{"C2"},
	}
	rscB3 = &Resource{
		name:        "B3",
		nextRscList: []string{"C2"},
	}
	rscC1 = &Resource{
		name:        "C1",
		nextRscList: []string{},
	}
	rscC2 = &Resource{
		name:        "C2",
		nextRscList: []string{},
	}
	dataA = &[]*Resource{rscA1, rscA2, rscA3}
	dataB = &[]*Resource{rscB1, rscB2, rscB3}
	dataC = &[]*Resource{rscC1, rscC2}
	data  = []*[]*Resource{dataA, dataB, dataC}
)

func TestMapLink(t *testing.T) {
	t.Run("test map-link", func(t *testing.T) {
		got, fail := Link0(data)
		assert.NotNil(t, got)
		assert.NotNil(t, fail)
		for _, item := range got {
			val := Join(*item, "->")
			fmt.Printf("link %s ok\n", val)
		}

		for _, item := range fail {
			val := Join(*item, "->")
			fmt.Printf("link %s failed\n", val)
		}
	})
}
