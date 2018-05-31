package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	. "github.com/book-library-seat-system/go-server/util"
)

func init() {
	ser := NewServer()
	ser.Run(":8899")
}

func TesttestGET(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 发送http get请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8899/v1/test", strings.NewReader(""))
	CheckErr(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

	// 接收响应并且读取body信息
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}
}
