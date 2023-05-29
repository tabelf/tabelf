import {checkUserInfo, CommunityUsed, ErrCustomerWebLinkLimit, getUserInfo, HttpOK, isEmpty, isNotEmpty} from "./Base";
import {request} from "../common/request";
import {message} from "antd";
import {
    CUSTOMER_COMMUNITY, CUSTOMER_COMMUNITY_AUDIT, CUSTOMER_COMMUNITY_AUDIT_STATUS,
    CUSTOMER_COMMUNITY_CATEGORY,
    CUSTOMER_COMMUNITY_DETAIL,
    CUSTOMER_COMMUNITY_META,
    CUSTOMER_COMMUNITY_SELF
} from "./index";
import {UserFocus} from "./CustomerAction";

// 发布内容到分享社区
export const CreatePublicCommunity = (data, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/public", {
            method: 'post',
            body: data,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("发布成功, 等待审核");
                }
                // 审核中
                if (!isEmpty(callback)) {
                    callback()
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadPublicCommunityCategory(records) {
    return {
        type: CUSTOMER_COMMUNITY_CATEGORY,
        records
    }
}

// 分享社区分类
export const GetPublicCommunityCategory = () => {
    return (dispatch) => {
        let url = "/anonymous/community/category"
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "?user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunityCategory(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadPublicCommunity(records) {
    return {
        type: CUSTOMER_COMMUNITY,
        records
    }
}

// 分享社区列表
export const GetPublicCommunity = (categoryUID, sorted, offset, limit) => {
    return (dispatch) => {
        request("/anonymous/community/public?category_uid=" + categoryUID + "&sorted=" + sorted + "&offset=" + offset + "&limit=" + limit)
        .then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunity(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadPublicCommunitySelf(records) {
    return {
        type: CUSTOMER_COMMUNITY_SELF,
        records
    }
}

// 分享社区列表
export const GetPublicCommunitySelf = (category) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/self?category=" + category, {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunitySelf(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}


function LoadPublicCommunityDetail(records) {
    return {
        type: CUSTOMER_COMMUNITY_DETAIL,
        records
    }
}

// 分享社区列表
export const GetPublicCommunityDetail = (detailUID, callback) => {
    return (dispatch) => {
        let url = "/anonymous/community/" + detailUID + "/detail"
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "?user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunityDetail(res.data))
                if (isNotEmpty(callback)) {
                    callback(res.data)
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}


function LoadPublicCommunityMeta(records) {
    return {
        type: CUSTOMER_COMMUNITY_META,
        records
    }
}

// 更新好站推荐元数据
export const UpdatePublicCommunityMeta = (communityUID, metaType, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/" + communityUID + "/meta", {
            method: 'put',
            body: {
                "meta_type": metaType
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (metaType === CommunityUsed) {
                    window.location.replace("/content/0/" + res.data.folder_number);
                } else {
                    dispatch(LoadPublicCommunityMeta(res.data))
                }
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit && !isEmpty(callback)) {
                    callback()
                }
                message.error(res.data.message);
            }
        });
    }
}

// 分享社区用户关注
export const CommunityUserFocus = (detailUID, followeeUID, status) => {
    return (dispatch) => {
        dispatch(UserFocus(followeeUID, status, () => {
            dispatch(GetPublicCommunityDetail(detailUID))
        }))
    }
}

function LoadPublicCommunityAudit(records) {
    return {
        type: CUSTOMER_COMMUNITY_AUDIT,
        records
    }
}

// 审核分享社区内容
export const GetPublicCommunityAudit = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/audit", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunityAudit(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadPublicCommunityAuditStatus(records) {
    return {
        type: CUSTOMER_COMMUNITY_AUDIT_STATUS,
        records
    }
}

export const UpdateAuditPublicCommunity = (communityUID, status) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/" + communityUID + "/audit", {
            method: 'put',
            body: {
                "status": status
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadPublicCommunityAuditStatus(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}