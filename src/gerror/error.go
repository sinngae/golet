package gerror

import (
	"encoding/json"
)

type GError struct {
	Err     error       `json:"error,omitempty"`
	RetCode int64       `json:"ret_code"`
	Message string      `json:"detail"`
	Debug   string      `json:"debug"`
	Data    interface{} `json:"data"`
}

func (ge *GError) Error() string {
	data, err := json.Marshal(ge)
	if err != nil {
		return err.Error() // TODO
	}
	return string(data)
}
