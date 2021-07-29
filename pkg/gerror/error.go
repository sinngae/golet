package gerror

import "encoding/json"

type GError struct {
	Err     error    `json:"error"`
	RetCode int64    `json:"ret_code"`
	Detail  []string `json:"detail"`
	Debug   string   `json:"debug"`
}

func (ge *GError) Error() string {
	data, err := json.Marshal(ge)
	if err != nil {
		return err.Error() // TODO
	}
	return string(data)
}
