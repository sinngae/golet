package utility

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
)

func MustCopy(to interface{}, from interface{}) {
	if reflect.TypeOf(to) != reflect.TypeOf(from) {
		err := fmt.Errorf("different type copy is forbidden")
		panic(err)
	}

	err := copier.Copy(to, from)
	if err != nil {
		err = fmt.Errorf("must copy failed, err=%v", err)
		panic(err) // TODO CAT&SeaTalk report
	}
}
