package invoke

type FastResponse struct {
	StatusCode  int    `json:"status_code"`
	Body        []byte `json:"body"`
	ContentType string `json:"content_type"`
	Server      string `json:"server"`
}

func (resp *FastResponse) GetStatusCode() int {
	if resp == nil {
		return 200
	}
	return resp.StatusCode
}

func (resp *FastResponse) GetBody() []byte {
	if resp == nil {
		return []byte{}
	}
	return resp.Body
}

func (resp *FastResponse) GetContentType() string {
	if resp == nil {
		return ""
	}
	return resp.ContentType
}

func (resp *FastResponse) GetServer() string {
	if resp == nil {
		return ""
	}
	return resp.Server
}
