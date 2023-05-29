import {checkUserInfo, HttpOK} from "./Base";
import {request} from "../common/request";
import {message} from "antd";
import {RECYCLE_WEBLINK} from "./index";
import {GetPersonalFolders, GetPersonalWebLinkInfo} from "./CustomerAction";

function LoadRecycleWebLink(records) {
    return {
        type: RECYCLE_WEBLINK,
        records
    }
}

// 查看删除的收藏链接
export const GetRecyclingPersonalWebLink = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/weblink/recycling", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadRecycleWebLink(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 恢复删除的收藏链接
export const RestoreDeletePersonalWebLink = (linkUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/weblink/" + linkUID + "/restore", {
            method: 'put',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(GetRecyclingPersonalWebLink());
                dispatch(GetPersonalFolders());
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 回收站彻底删除的收藏链接
export const DeleteForeverPersonalWebLink = (linkUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/weblink/" + linkUID + "/forever", {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("删除成功");
                }
                dispatch(GetRecyclingPersonalWebLink());
                dispatch(GetPersonalWebLinkInfo())
            } else {
                message.error(res.data.message);
            }
        });
    }
}