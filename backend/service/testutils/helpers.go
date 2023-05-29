package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"tabelf/backend/service/api/pkg/utils"
	"tabelf/backend/service/app"

	"github.com/stretchr/testify/suite"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Helper struct {
	Suite  suite.Suite
	Route  httpx.Router
	Config app.Config

	method      string
	url         string
	contentType string
	body        io.Reader
	jwt         bool
	header      map[string]string
}

type HelperResult struct {
	helper Helper

	response *httptest.ResponseRecorder
}

func (h *Helper) Get(url string, args ...interface{}) *Helper {
	h.method = http.MethodGet
	h.url = fmt.Sprintf(url, args...)
	return h
}

func (h *Helper) Post(url string, args ...interface{}) *Helper {
	h.method = http.MethodPost
	h.url = fmt.Sprintf(url, args...)
	return h
}

func (h *Helper) Put(url string, args ...interface{}) *Helper {
	h.method = http.MethodPut
	h.url = fmt.Sprintf(url, args...)
	return h
}

func (h *Helper) Delete(url string, args ...interface{}) *Helper {
	h.method = http.MethodDelete
	h.url = fmt.Sprintf(url, args...)
	return h
}

func (h *Helper) JWT() *Helper {
	h.jwt = true
	return h
}

func (h *Helper) Body(data interface{}) *Helper {
	h.contentType = "application/json;charset=utf-8"
	content, err := json.Marshal(data)
	h.Suite.Nil(err)
	h.body = bytes.NewBuffer(content)
	return h
}

func (h *Helper) Header(date map[string]string) *Helper {
	h.header = date
	return h
}

func (h *Helper) Excel(data [][]string) *Helper {
	excel := excelize.NewFile()
	w, err := excel.NewStreamWriter("Sheet1")
	h.Suite.Nil(err)
	for i, row := range data {
		r := make([]interface{}, len(row))
		for j, value := range row {
			r[j] = value
		}
		cell, err2 := excelize.CoordinatesToCellName(1, i+1)
		h.Suite.Nil(err2)
		err2 = w.SetRow(cell, r)
		h.Suite.Nil(err2)
	}
	h.Suite.Nil(w.Flush())
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "test.xlsx")
	h.Suite.Nil(err)
	h.Suite.Nil(excel.Write(file))
	h.Suite.Nil(writer.Close())
	h.body = body
	h.contentType = writer.FormDataContentType()
	return h
}

func (h *Helper) do() *HelperResult {
	if h.contentType == "" {
		h.contentType = "application/json;charset=utf-8"
		if h.body == nil {
			h.body = bytes.NewBuffer([]byte("{}"))
		}
	}
	request, err := http.NewRequestWithContext(context.Background(), h.method, h.url, h.body)
	h.Suite.Nil(err)
	request.Header.Set("content-type", h.contentType)
	if h.jwt {
		token, err := utils.GetJwtToken(
			h.Config.Jwt.JwtKey,
			h.Config.Jwt.JwtIssuer,
			h.Config.Jwt.JwtExpire,
			utils.JWTClaims{UID: "test"},
		)
		h.Suite.Nil(err)
		request.Header.Set("Authorization", token)
	}
	for key, value := range h.header {
		request.Header.Set(key, value)
	}
	hr := HelperResult{helper: *h}
	hr.response = httptest.NewRecorder()
	h.Route.ServeHTTP(hr.response, request)
	return &hr
}

func (h *Helper) OK(vs ...interface{}) *HelperResult {
	hr := h.do()
	hr.helper.Suite.Equal(http.StatusOK, hr.response.Code)
	return hr.result(vs...)
}

func (h *Helper) Bad(vs ...interface{}) *HelperResult {
	hr := h.do()
	hr.helper.Suite.Equal(http.StatusBadRequest, hr.response.Code)
	return hr.result(vs...)
}

// result 会把 response.Body 中的参数解析到 vs 中，目前应该用 .OK/.Bad 来调用，所以设为 private 的.
func (hr *HelperResult) result(vs ...interface{}) *HelperResult {
	if len(vs) == 0 {
		return hr
	}
	v := vs[0]
	err := json.Unmarshal(hr.response.Body.Bytes(), v)
	hr.helper.Suite.Nil(err)
	return hr
}
