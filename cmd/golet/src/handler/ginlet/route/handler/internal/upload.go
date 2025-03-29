package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sinngae/golet/src/errcode"
)

// formDataCheck 文件类型检测 + 文件大小检测
//
//	引自https://www.jianshu.com/p/30d5138ceee6
//
// TODO 校验file格式为abc ?
// Done 校验file尺寸小于1M
// Done 仅读取指定长度数据
func formDataCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) (io.Reader, *pb.XlsCheckReqData, error) {
	var (
		// read limit
		bodyMaxLimit = int64(1 << 20) // body限制，1MB
		fileMaxLimit = int64(1 << 20) // file限制，1MB
		textMaxLimit = int64(-1)      // text限制，？

		// form表单的白名单
		formKeyFile = "file"
		formKeyData = "data"
	)

	files, texts, err := ParseForm(w, r, bodyMaxLimit, fileMaxLimit, textMaxLimit)
	if err != nil {
		return nil, nil, err
	}

	// file check
	fileBody, ok := files[formKeyFile]
	if !ok || fileBody == nil {
		err := fmt.Errorf("form-part[%s] file is expected", formKeyFile)
		err = errcode.Unknown.WithCause(err)
		return nil, nil, err
	}
	fileBuf := bytes.NewBuffer(fileBody)

	// param check
	textBody, ok := texts[formKeyData]
	if !ok {
		return fileBuf, nil, nil
	}
	pbReq := &pb.XlsCheckReqData{}
	err = json.Unmarshal(textBody, pbReq)
	if err != nil {
		err = fmt.Errorf("parse data[%v] failed, err=%v", textBody, err)
		err = bak.Gland(err, errcode.XlsBadRequest)
		return nil, nil, err
	}

	return fileBuf, pbReq, nil
}

func ParseForm(w http.ResponseWriter, r *http.Request, bodyMaxLimit, fileSumMaxLimit, textSumMaxLimit int64) (map[string][]byte, map[string][]byte, error) {
	r.Body = http.MaxBytesReader(w, r.Body, bodyMaxLimit)
	reader, err := r.MultipartReader()
	if err != nil { // http.StatusBadRequest
		err = fmt.Errorf("read form-multipart failed, err=%v", err)
		err = bak.Gland(err, errcode.XlsBadRequest)
		return nil, nil, err
	}

	files := make(map[string][]byte)
	texts := make(map[string][]byte)
	filesMaxLimit := int64(1 << 20) // default 1MB limit
	if fileSumMaxLimit > 0 {
		filesMaxLimit = fileSumMaxLimit
	}
	textsMaxLimit := int64(1 << 20) // default 1MB limit
	if textSumMaxLimit > 0 {
		textsMaxLimit = textSumMaxLimit
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF { // end of reader
			break
		}
		if err != nil {
			err = fmt.Errorf("read form-part failed, err=%v", err)
			err = stack.New()
			err = bak.Gland(err, errcode.XlsBadRequest)
			return nil, nil, err
		}

		var buf bytes.Buffer
		filename, formKey := part.FileName(), part.FormName()

		if filename == "" { // 非文件字段
			if _, ok := texts[formKey]; ok {
				err := fmt.Errorf("form-part[%s] text is duplicated", formKey)
				err = bak.Gland(err, errcode.XlsBadRequest)
				return nil, nil, err
			}

			n, err := io.CopyN(&buf, part, textsMaxLimit+1)
			if err != nil && err != io.EOF {
				err = fmt.Errorf("read form-part[%s] text failed, err=%v", formKey, err)
				err = bak.Gland(err, errcode.XlsBadRequest)
				return nil, nil, err
			}

			textsMaxLimit -= n
			if textsMaxLimit < 0 {
				err := fmt.Errorf("form-part[%s] text too large", formKey)
				err = bak.Gland(err, errcode.XlsSizeMaxExceed) // XXX ?
				return nil, nil, err
			}

			texts[formKey] = buf.Bytes()
		} else { // 文件字段
			if _, ok := files[formKey]; ok {
				err := fmt.Errorf("form-part[%s] file is duplicated", formKey)
				err = bak.Gland(err, errcode.XlsBadRequest)
				return nil, nil, err
			}

			n, err := io.CopyN(&buf, part, filesMaxLimit+1)
			if err != nil && err != io.EOF {
				err = fmt.Errorf("read form-part[%s] file failed, err=%v", formKey, err)
				err = bak.Gland(err, errcode.XlsSizeMaxExceed) // XXX always be
				return nil, nil, err
			}

			filesMaxLimit -= n
			if filesMaxLimit < 0 {
				err := fmt.Errorf("form-part[%s] file too large", formKey)
				err = bak.Gland(err, errcode.XlsSizeMaxExceed) // XXX ?
				return nil, nil, err
			}

			files[formKey] = buf.Bytes()
			contentType := http.DetectContentType(files[formKey])
			if contentType != constHttp.ContentTypeZip {
				err := fmt.Errorf("form-part[%s] file type is not allowed", formKey)
				err = bak.Gland(err, errcode.XlsBadRequest)
				return nil, nil, err
			}
		}
	}
	return files, texts, nil
}
