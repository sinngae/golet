package fasthttp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestUploadFileServer(t *testing.T) {
	err := fasthttp.ListenAndServe(":8080", UploadHandler)
	fmt.Println(err)
}
func UploadHandler(ctx *fasthttp.RequestCtx) {
	//根据参数名获取上传的文件
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	//打开上传的文件
	file, err := fileHeader.Open()
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	//使用完关闭文件
	defer file.Close()
	newFile := "d:\\bbb.txt" //windows下需要转移反斜线
	//新建一个文件，此处使用默认的txt格式
	nf, err := os.OpenFile(newFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	//使用完需要关闭
	defer nf.Close()
	//复制文件内容
	_, err = io.Copy(nf, file)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	ctx.WriteString("success")
}

func TestUploadFile(t *testing.T) {
	//待上传文件
	uploadFile := "d:\\aaa.txt"
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	//从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("file", path.Base(uploadFile))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file, err := os.Open(uploadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	//不要忘记关闭打开的文件
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err.Error())
	}

	//关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	request.SetBody(bodyBufer.Bytes())
	request.Header.SetMethod("POST")
	request.SetRequestURI("http://localhost:8080/")
	response := fasthttp.AcquireResponse()
	err = fasthttp.Do(request, response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(response.Body()))
}
