import {CUSTOMER_INVITE} from "./index";
import {request} from "../common/request";
import {message} from "antd";
import {checkUserInfo, isEmpty} from "./Base";

function LoadPersonalInvite(records) {
    return {
        type: CUSTOMER_INVITE,
        records
    }
}

// 查询邀请记录信息
export const GetPersonalInvite = (callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/invite", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
                if (res.code === 200) {
                    dispatch(LoadPersonalInvite(res.data));
                } else {
                    message.error(res.data.message);
                }
                if (!isEmpty(callback)) {
                    callback()
                }
            });
    }
}