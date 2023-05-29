import {checkUserInfo} from "./Base";
import {request} from "../common/request";
import {message} from "antd";
import {HISTORY_RECENT_WEBLINK} from "./index";

function LoadRecentWebLink(records) {
    return {
        type: HISTORY_RECENT_WEBLINK,
        records
    }
}

// 查看最近更新的收藏链接
export const GetRecentPersonalWebLink = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/weblink/recent", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadRecentWebLink(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}