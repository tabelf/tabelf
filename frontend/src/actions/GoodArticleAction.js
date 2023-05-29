import {checkUserInfo, ErrCustomerWebLinkLimit, getUserInfo, HttpOK, isEmpty, isNotEmpty} from "./Base";
import {request} from "../common/request";
import {message} from "antd";
import {
    CUSTOMER_ARTICLE,
    CUSTOMER_ARTICLE_APPEND,
    CUSTOMER_ARTICLE_AUDIT,
    CUSTOMER_ARTICLE_CATEGORY,
    CUSTOMER_ARTICLE_COLLECTION,
    CUSTOMER_ARTICLE_HOT,
    CUSTOMER_ARTICLE_MENU,
    CUSTOMER_ARTICLE_META,
    CUSTOMER_ARTICLE_SUBMIT,
    CUSTOMER_ARTICLE_UPDATE,
} from "./index";

function LoadGoodArticleCategory(records) {
    return {
        type: CUSTOMER_ARTICLE_CATEGORY,
        records
    }
}

// 好文发现
export const GetGoodArticleCategory = () => {
    return (dispatch) => {
        let url = "/anonymous/article/category"
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "?user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleCategory(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticle(records) {
    return {
        type: CUSTOMER_ARTICLE,
        records
    }
}

function LoadAppendGoodArticle(records) {
    return {
        type: CUSTOMER_ARTICLE_APPEND,
        records
    }
}

// 好文推荐列表
export const GetGoodArticle = (categoryUID, sorted, offset, limit) => {
    return (dispatch) => {
        let url = "/anonymous/article/recommend?category_uid=" + categoryUID + "&sorted=" + sorted + "&offset=" + offset + "&limit=" + limit
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "&user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                if (offset === 0) {
                    dispatch(LoadGoodArticle(res.data))
                } else {
                    dispatch(LoadAppendGoodArticle(res.data))
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadHotGoodArticle(records) {
    return {
        type: CUSTOMER_ARTICLE_HOT,
        records
    }
}

// 热门好文推荐列表
export const GetHotGoodArticle = () => {
    return (dispatch) => {
        request("/anonymous/article/hot").then(res => {
            if (res.code === 200) {
                dispatch(LoadHotGoodArticle(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 保存好文发现
export const AddGoodArticleRecommend = (article, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/recommend", {
            method: 'post',
            body: article,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("创建成功");
                }
                callback()
                dispatch(GetGoodArticleMenuData())
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 收藏好文发现
export const GoodArticleStar = (articleUID, folderUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + articleUID + "/collection", {
            method: 'put',
            body: {
                "folder_uid": folderUID
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                message.success("添加成功")
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    callback()
                }
                message.error(res.data.message);
            }
        });
    }
}

// 好文精选浏览量
export const GoodArticleView = (articleUID) => {
    return (dispatch) => {
        request("/anonymous/article/" + articleUID + "/view", {
            method: 'put',
        }).then(res => {
            if (res.code === 200) {
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 获取用户文件夹列表
export const GetGoodArticleFolders = (callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                callback(res.data)
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticleMeta(records) {
    return {
        type: CUSTOMER_ARTICLE_META,
        records
    }
}

// 更新好站推荐元数据
export const UpdateGoodArticleMeta = (articleUID, metaType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + articleUID + "/meta", {
            method: 'put',
            body: {
                "meta_type": metaType
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleMeta(res.data))
                dispatch(GetGoodArticleMenuData())
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticleCollection(records) {
    return {
        type: CUSTOMER_ARTICLE_COLLECTION,
        records
    }
}

// 获取用户文件夹列表
export const GetGoodArticleCollection = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/collection", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleCollection(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticleSubmit(records) {
    return {
        type: CUSTOMER_ARTICLE_SUBMIT,
        records
    }
}

// 获取我的发布列表
export const GetGoodArticleSubmit = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/submit", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleSubmit(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticleAudit(records) {
    return {
        type: CUSTOMER_ARTICLE_AUDIT,
        records
    }
}

// 获取内容审核列表
export const GetGoodArticleAudit = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/audit", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleAudit(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

export const UpdateGoodArticleAudit = (articleUID, status) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + articleUID + "/audit", {
            method: 'put',
            body: {
                "status": status,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("操作成功");
                }
                dispatch(GetGoodArticleAudit())
                dispatch(GetGoodArticleMenuData())
            } else {
                message.error(res.data.message);
            }
        });
    }
}

export const GoodArticleUnStar = (articleUID, metaType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + articleUID + "/meta", {
            method: 'put',
            body: {
                "meta_type": metaType
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetGoodArticleCollection())
                dispatch(GetGoodArticleMenuData())
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodArticleMenu(records) {
    return {
        type: CUSTOMER_ARTICLE_MENU,
        records
    }
}

export const GetGoodArticleMenuData = () => {
    return (dispatch) => {
        let url = "/anonymous/article/norm"
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "?user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodArticleMenu(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 推荐好站详情
export const GetGoodArticleDetail = (articleUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + articleUID + "/detail", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (isNotEmpty(callback)) {
                    callback(res.data)
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadUpdateGoodArticle(records) {
    return {
        type: CUSTOMER_ARTICLE_UPDATE,
        records
    }
}

export const UpdateGoodArticleRecommend = (article, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/article/" + article.uid + "/detail", {
            method: 'put',
            body: article,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(LoadUpdateGoodArticle(res.data))
                dispatch(GetGoodArticleMenuData())
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}