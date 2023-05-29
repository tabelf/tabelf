package account

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"github.com/clbanning/mxj"
	"tabelf/backend/service/api/internal/logic/anonymous/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
)

func WechatMessageCallbackHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WechatMessageCallbackRequest
		m, err := mxj.NewMapXmlReader(r.Body)
		if err != nil {
			return
		}
		if _, ok := m["xml"]; !ok {
			return
		}
		bytes, err := json.Marshal(m["xml"])
		if err != nil {
			return
		}
		if err = json.Unmarshal(bytes, &req); err != nil {
			return
		}
		l := account.NewWechatMessageCallbackLogic(r.Context(), ctx)
		resp, err := l.WechatMessageCallback(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			if _, err = w.Write(TextMessageReply(
				req.ToUserName,
				req.FromUserName,
				resp.Message,
			)); err != nil {
				return
			}
		}
	}
}

func TextMessageReply(fromUser, toUser, content string) []byte {
	if content == "" {
		return []byte("")
	}
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
	}
	return msg
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string   // 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName      xml.Name `xml:"xml"`
}
