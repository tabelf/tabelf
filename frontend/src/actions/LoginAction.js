import {request} from "../common/request";
import {message} from "antd";
import {LOGIN_QR_CODE} from "./index";
import {Cache, USER_TOKEN_EXPIRED, USER_TOKEN_KEY, isEmpty} from "./Base";

function LoadQrCode(records) {
    return {
        type: LOGIN_QR_CODE,
        records
    }
}

// 获取登录的二维码
export const GetLoginQrCode = () => {
    return (dispatch) => {
        request("/anonymous/account/auth/login/qr_code")
            .then(res => {
                if (res.code === 200) {
                    dispatch(LoadQrCode(res.data));
                } else {
                    message.error(res.data.message);
                }
            });
    }
}

// 用户登录
export const AuthLogin = (authCode, referralUID) => {
    return (dispatch) => {
        request("/anonymous/account/auth/login", {
            method: 'post',
            body: {
                "auth_code": authCode,
                "referral_uid": referralUID,
            }
        }).then(res => {
            if (res.code === 200) {
                Cache.set(USER_TOKEN_KEY, res.data, USER_TOKEN_EXPIRED);
                let beforePage = document.referrer;
                if (isEmpty(beforePage) ||
                    beforePage === "https://www.tabelf.com/" ||
                    beforePage === "https://tabelf.com/" ||
                    beforePage === "http://www.tabelf.com/" ||
                    beforePage === "http://tabelf.com/" ||
                    beforePage === "www.tabelf.com/" ||
                    beforePage === "tabelf.com/" ||
                    beforePage === "http://localhost:3000/" ||
                    beforePage === "localhost:3000/") {
                    beforePage = "/workspace";
                }
                window.location.href = beforePage; // 返回进行登录页的前一个页面
            } else {
                message.error(res.data.message);
            }
        });
    }
}