<<<<<<< HEAD
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
=======
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
>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511
