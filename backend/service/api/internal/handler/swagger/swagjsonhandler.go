package swagger

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"tabelf/backend/service/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const swagFilePath = "gen/swagger/swagger.json"

// SwagByte 生成的swagger文档.
var SwagByte json.RawMessage

func init() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Panicf("no current folder")
	}
	index := strings.Index(curDir, "binghuang")
	root := "/app"
	if index > -1 {
		root = curDir[:index+3]
	}
	swagFile, err := os.Open(filepath.Join(root, swagFilePath))
	if err != nil {
		log.Println(err)
	}
	defer swagFile.Close()
	SwagByte, err = ioutil.ReadAll(swagFile)
	if err != nil {
		log.Println(err)
	}
}

func SwagJSONHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err := w.Write(SwagByte)
		if err != nil {
			httpx.Error(w, err)
		}
	}
}
