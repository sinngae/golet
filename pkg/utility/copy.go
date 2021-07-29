package utility

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func MustCopy(toValue interface{}, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		err = fmt.Errorf("must copy failed, err=%v", err)
		panic(err) // TODO CAT&SeaTalk report
	}
}
