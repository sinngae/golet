package utility

import (
	"encoding/json"
	"fmt"
)

func MustMarshal(src interface{}) []byte {
	dest, err := json.Marshal(src)
	if err != nil {
		err = fmt.Errorf("must-marshal failed, err=%s", err.Error())
		panic(err)
	}
	return dest
}

func MustUnmarshal(src []byte, dest interface{}) {
	err := json.Unmarshal(src, dest)
	if err != nil {
		err = fmt.Errorf("must-unmarshal failed, err=%s", err.Error())
		panic(err)
	}
}
