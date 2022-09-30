package utility

import (
	"encoding/json"
	"fmt"
)

func MustMarshal(src interface{}) []byte {
	dest, err := json.Marshal(src)
	if err != nil {
		err = fmt.Errorf("must copy failed, err=%v", err)
		panic(err) // TODO CAT&SeaTalk report
	}
	return dest
}

func MustUnmarshal(src []byte, dest interface{}) {
	err := json.Unmarshal(src, dest)
	if err != nil {
		err = fmt.Errorf("must copy failed, err=%v", err)
		panic(err) // TODO CAT&SeaTalk report
	}
}
