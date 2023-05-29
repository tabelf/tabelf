package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tabelf/backend/service/app"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type CreateTickerResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	URL           string `json:"url"`
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

type UserInfoResponse struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	Headimgurl     string `json:"headimgurl"`
	SubscribeTime  int    `json:"subscribe_time"`
	Remark         string `json:"remark"`
	Groupid        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
	ErrCode        int    `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
}

func GetWechatAccessToken(ctx context.Context) (token string, err error) {
	res, err := http.Get(fmt.Sprintf(app.AccessTokenURL, app.Wechat.AppID, app.Wechat.AppSecret))
	if err != nil {
		return "", err
	}
	defer func() {
		if e := res.Body.Close(); e != nil {
			return
		}
	}()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resp := &AccessTokenResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return "", err
	}
	if resp.ErrCode != 0 {
		return "", app.NewError(ctx, resp.ErrCode, resp.ErrMsg)
	}
	return resp.AccessToken, nil
}

func GetQrExpiredTicket(ctx context.Context, accessToken string, expire int) (resp *CreateTickerResponse, err error) {
	body := fmt.Sprintf(`{
			"expire_seconds": %d,
			"action_name": "%s",
			"action_info": {
				"scene": {
					"scene_str": "%s"
				}
			}
		}`, expire, "QR_STR_SCENE", app.QrScene)
	res, err := http.Post(fmt.Sprintf(app.TicketURL, accessToken),
		"application/json",
		strings.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resp = &CreateTickerResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	if resp.ErrCode != 0 {
		return nil, app.NewError(ctx, resp.ErrCode, resp.ErrMsg)
	}
	return resp, nil
}

func GetWechatUserInfo(ctx context.Context, accessToken, openID string) (resp *UserInfoResponse, err error) {
	res, err := http.Get(fmt.Sprintf(app.UserInfoURL, accessToken, openID))
	if err != nil {
		return nil, err
	}
	defer func() {
		if e := res.Body.Close(); e != nil {
			return
		}
	}()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resp = &UserInfoResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	if resp.ErrCode != 0 {
		return nil, app.NewError(ctx, resp.ErrCode, resp.ErrMsg)
	}
	return resp, nil
}
