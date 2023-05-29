import {request} from "../common/request";
import {message} from "antd";
import {
    CUSTOMER_ACCOUNT,
    CUSTOMER_ACCOUNT_MESSAGE,
    CUSTOMER_FOLDERS,
    CUSTOMER_MENU_OPEN_KEY,
    CUSTOMER_MENU_SELECT_KEY,
    CUSTOMER_OPEN_ACTIVE_FOLDERS,
    CUSTOMER_SEARCH,
    CUSTOMER_SHARE_ALL_PERSONNEL,
    CUSTOMER_SHARE_FOLDER,
    CUSTOMER_SHARE_FOLDERS,
    CUSTOMER_STATION,
    CUSTOMER_STATION_APPEND,
    CUSTOMER_STATION_CATEGORY,
    CUSTOMER_TEAMS,
    CUSTOMER_UPDATE_ACCOUNT,
    CUSTOMER_WEBLINKS,
    CUSTOMER_WORKSPACE,
    CUSTOMER_WORKSPACE_SWITCH,
    CUSTOMER_WORKSPACES,
    CUSTOMER_STATION_META,
    CUSTOMER_SHARE_FOLDER_CLEAR,
    CUSTOMER_SHARE_COLL,
    CUSTOMER_SHARE_INFO,
    CUSTOMER_STATION_DETAIL, CUSTOMER_STATION_AUDIT
} from "./index";
import {checkUserInfo, ErrCustomerWebLinkLimit, getUserInfo, HttpOK, isEmpty, isNotEmpty} from "./Base";

export const SetSubMenuOpenKeys = (openKeyID) => {
    return (dispatch) => {
        dispatch({
            type: CUSTOMER_MENU_OPEN_KEY,
            records: {
                openKey: openKeyID,
            }
        })
    }
}

export const SetMenuItemSelectKeys = (selectKeyID) => {
    return (dispatch) => {
        dispatch({
            type: CUSTOMER_MENU_SELECT_KEY,
            records: {
                selectKey: selectKeyID,
            }
        })
    }
}

function LoadSearchLinks(records) {
    return {
        type: CUSTOMER_SEARCH,
        records
    }
}

// 查看充值列表
export const SearchWebLinks = (keyword, type) => {
    return (dispatch) => {
        if (isEmpty(keyword)) {
            dispatch(LoadSearchLinks([]));
            return;
        }
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/search?keyword=" + keyword + "&type=" + type, {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadSearchLinks(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadTeams(records) {
    return {
        type: CUSTOMER_TEAMS,
        records
    }
}

// 获取用户团队列表
export const GetPersonalTeams = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/teams", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadTeams(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadWebLinkInfo(records) {
    return {
        type: CUSTOMER_WEBLINKS,
        records
    }
}

// 获取个人网址数量
export const GetPersonalWebLinkInfo = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/weblink/info", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadWebLinkInfo(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadAccountInfo(records) {
    return {
        type: CUSTOMER_ACCOUNT,
        records
    }
}

// 获取个人基本消息
export const GetPersonalAccountInfo = (callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/info", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadAccountInfo(res.data))
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback(res.data)
            }
        });
    }
}

function LoadAccountMessage(records) {
    return {
        type: CUSTOMER_ACCOUNT_MESSAGE,
        records
    }
}

// 获取个人基本消息
export const GetMessages = (msgType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/message?msg_type=" + msgType, {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadAccountMessage(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 用户关注
export const UserFocus = (followeeUID, status, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/focus", {
            method: 'post',
            body: {
                "followee_uid": followeeUID,
                "status": status
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
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

// 消息标记已读
export const ReadMessage = (msgType, messageUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/message/" + messageUID, {
            method: 'put',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetMessages(msgType))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 全部标记已读
export const ReadAllMessage = (msgType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/message", {
            method: 'put',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetMessages(msgType))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 删除消息
export const DelMessage = (msgType, messageUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/message/" + messageUID, {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetMessages(msgType))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 删除全部消息
export const DelAllMessage = (msgType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/message", {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetMessages(msgType))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadUpdateAccountInfo(records) {
    return {
        type: CUSTOMER_UPDATE_ACCOUNT,
        records
    }
}

// 更新个人消息
export const UpdatePersonalAccountInfo = (data) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/info", {
            method: 'put',
            body: data,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                // dispatch(LoadUpdateAccountInfo(data));
                dispatch(GetPersonalAccountInfo());
                dispatch(GetMessages(1));
                dispatch(GetPersonalWebLinkInfo());
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 获取用户文件夹列表
export const GetDefaultFolder = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadFolders(res.data));
                if (res.data.folders.length > 0) {
                    request("/customer/" + userInfo.user_uid + "/personal/folders/workspace/content?folder_number=" + res.data.folders[0].folder_number, {
                        headers: {'Authorization': userInfo.token}
                    }).then(res => {
                        if (res.code === 200) {
                            dispatch(LoadWorkspaces(res.data));
                        } else {
                            message.error(res.data.message);
                        }
                    });
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadFolders(records) {
    return {
        type: CUSTOMER_FOLDERS,
        records
    }
}

// 获取用户文件夹列表
export const GetPersonalFolders = (callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (isNotEmpty(callback)) {
                    callback(res.data)
                } else {
                    dispatch(LoadFolders(res.data));
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 添加个人文件
export const AddPersonalFolder = (data, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders", {
            method: 'post',
            body: data,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("添加成功");
                }
                dispatch(GetPersonalFolders());
                dispatch(LoadOpenActiveFolders({
                    openKey: ["my-files"],
                    selectKey: [res.data.folder_uid],
                }))
                window.location.replace("/content/0/" + res.data.folder_number);
            } else {
                message.error(res.data.message);
            }
            callback(res.data)
        });
    }
}

// 更新个人文件
export const UpdatePersonalFolder = (folderUID, data, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID, {
            method: 'put',
            body: data,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(GetPersonalFolders());
                dispatch(GetSharePersonalFolders());
                dispatch(GetWorkspaceContents(folderUID));
            } else {
                message.error(res.data.message);
            }
            callback(res.data)
        });
    }
}

// 删除个人文件
export const DeletePersonalFolder = (folderUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID, {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("删除成功");
                }
                window.location.href = "/"
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 复制个人文件
export const CopyPersonalFolder = (folderUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID + "/copy", {
            method: 'post',
            body: {
                "type": 0
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("复制成功");
                }
                dispatch(GetPersonalFolders());
                dispatch(GetPersonalWebLinkInfo())
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    callback()
                }
                message.error(res.data.message);
            }
        });
    }
}

function LoadShareFolders(records) {
    return {
        type: CUSTOMER_SHARE_FOLDERS,
        records
    }
}

// 获取协作个人文件
export const GetSharePersonalFolders = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadShareFolders(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadShareFolder(records) {
    return {
        type: CUSTOMER_SHARE_FOLDER,
        records
    }
}

export const ClearSharePersonal = () => {
    return (dispatch) => {
        dispatch({
            type: CUSTOMER_SHARE_FOLDER_CLEAR,
            data: []
        })
    }
}

// 分享个人文件链接
export const SharePersonalFolder = (folderUID, offset, limit, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID + "/share?offset="+offset + "&limit=" + limit , {
            method: 'put',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadShareFolder(res.data));
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback(res.data)
            }
        });
    }
}

function LoadUpdateShareColl(records) {
    return {
        type: CUSTOMER_SHARE_COLL,
        records
    }
}

// 分享个人文件链接
export const UpdateCollSharePersonalFolder = (folderUID, shareUID, collUID, type) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID + "/share/" + shareUID + "/coll/" + collUID, {
            method: 'put',
            body: {
                "type": type,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(LoadUpdateShareColl({
                    "coll_uid": collUID,
                    "type": type
                }));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 复制分享文件到我到个人文件
export const CopySharePersonalFolder = (shareUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/copy", {
            method: 'post',
            body: {
                "type": 0
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("复制成功");
                }
                dispatch(GetPersonalFolders());
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    callback()
                }
                message.error(res.data.message);
            }
        });
    }
}

function LoadUpdateShare(records) {
    return {
        type: CUSTOMER_SHARE_INFO,
        records
    }
}

// 修改分享个人文件链接
export const UpdateSharePersonalFolder = (folderUID, shareUID, authority, expiredDay) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID + "/share/" + shareUID, {
            method: 'put',
            body: {
                "authority": authority,
                "expired_day": expiredDay
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(LoadUpdateShare({
                    "authority": authority,
                    "expired_day": expiredDay
                }))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadOpenActiveFolders(records) {
    return {
        type: CUSTOMER_OPEN_ACTIVE_FOLDERS,
        records
    }
}

// 加入分享个人文件链接的协作
export const JoinSharePersonalFolder = (shareUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/join", {
            method: 'post',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetSharePersonalFolders());
                dispatch(LoadOpenActiveFolders({
                    openKey: ["my-share"],
                    selectKey: ["share" + res.data.folder_uid],
                }))
                window.location.replace("/collaboration/0/" + res.data.folder_number);
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

// 退出分享协作
export const ExitSharePersonalFolder = (shareUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/exit", {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("退出成功");
                }
                window.location.href = "/"
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 个人文件通过邀请分享好友
export const ShareFriendFolder = (shareUID, email, authority, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/friend", {
            method: 'post',
            body: {
                "email": email,
                "authority": authority,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("邀请成功");
                }
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

function LoadShareAllPersonnel(records) {
    return {
        type: CUSTOMER_SHARE_ALL_PERSONNEL,
        records
    }
}

// 查询协作管理的人员
export const GetShareAllPersonnel = (shareUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/personnel", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadShareAllPersonnel(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 创建者管理协作人员
export const UpdateManageSharePersonnel = (shareUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/personnel", {
            method: 'put',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(GetShareAllPersonnel(shareUID));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 创建者剔除协作人员
export const DeleteSharePersonnel = (shareUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/personnel", {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("删除成功");
                }
                dispatch(GetShareAllPersonnel(shareUID));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadWorkspaces(records) {
    return {
        type: CUSTOMER_WORKSPACES,
        records
    }
}

// 根据文件编号获取工作空间区域内容
export const GetFolderNumberWorkspaceContent = (folderNumber) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/workspace/content?folder_number=" + folderNumber, {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadWorkspaces(res.data));
                dispatch(LoadOpenActiveFolders({
                    openKey: ["my-files"],
                    selectKey: [res.data.folder_uid],
                }))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 获取文件工作区域内容
export const GetWorkspaceContents = (folderUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/" + folderUID + "/workspace/content", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadWorkspaces(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 获取协作文件工作区域内容
export const GetCollaborationWorkspaceContent = (folderNumber) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/workspace/content/share?folder_number=" + folderNumber, {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadWorkspaces({
                    ...res.data
                }));
                dispatch(LoadOpenActiveFolders({
                    openKey: ["my-share"],
                    selectKey: ["share" + res.data.folder_uid],
                }))
                if (res.data.authority === "0") {
                    message.warn("当前文件仅支持查看", 2)
                } else {
                    message.info("当前文件可进行编辑", 2)
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadActiveWorkspace(records) {
    return {
        type: CUSTOMER_WORKSPACE_SWITCH,
        records
    }
}

// 打开或关闭工作区域
export const UpdateWorkspaceSwitch = (folderUID, workspaceUIDs) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/switch", {
            method: 'put',
            body: {
                "active_workspace_uids": workspaceUIDs
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadActiveWorkspace(workspaceUIDs));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 添加工作空间
export const AddWorkspace = (folderUID, workspaceName) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace", {
            method: 'post',
            body: {
                "workspace_name": workspaceName
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetWorkspaceContents(folderUID))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 更新工作空间
export const UpdateWorkspace = (folderUID, workspaceUID, workspaceName) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + workspaceUID, {
            method: 'put',
            body: {
                "workspace_name": workspaceName
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetWorkspaceContents(folderUID))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 删除工作空间
export const DeleteWorkspace = (folderUID, workspaceUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + workspaceUID, {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(GetWorkspaceContents(folderUID))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadWorkspace(records) {
    return {
        type: CUSTOMER_WORKSPACE,
        records
    }
}

// 更新工作区域链接排序
export const UpdateWorkspaceWebLinks = (folderUID, newWorkspace, oldWorkspace) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + newWorkspace.workspace_uid + "/sort", {
            method: 'put',
            body: {
                "web_links": newWorkspace.web_links,
                "old_workspace_uid": oldWorkspace.workspace_uid,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(LoadWorkspace({
                    "new_workspace": newWorkspace,
                    "old_workspace": oldWorkspace,
                }))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 添加收藏链接
export const AddPersonalWebLink = (folderUID, workspaceUID, webLink, successCall, failedCall) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + workspaceUID + "/weblink", {
            method: 'post',
            body: webLink,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("添加成功");
                }
                dispatch(GetWorkspaceContents(folderUID))
                dispatch(GetPersonalWebLinkInfo())
                successCall()
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    failedCall()
                }
                message.error(res.data.message);
            }
        });
    }
}

// 添加资源上传成功后
export const AddLocalFileSuccessCallback = (folderUID) => {
    return (dispatch) => {
        dispatch(GetWorkspaceContents(folderUID))
        dispatch(GetPersonalWebLinkInfo())
    }
}

// 添加协作收藏链接
export const AddSharePersonalWebLink = (folderUID, shareUID, workspaceUID, webLink, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folders/share/" + shareUID + "/workspace/" + workspaceUID + "/weblink", {
            method: 'post',
            body: webLink,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("添加成功");
                }
                dispatch(GetWorkspaceContents(folderUID))
                dispatch(GetPersonalWebLinkInfo())
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    message.error("分享人链接数量超过限制");
                } else {
                    message.error(res.data.message);
                }
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

// 编辑收藏链接
export const UpdatePersonalWebLink = (folderUID, workspaceUID, webLink, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + workspaceUID + "/weblink/" + webLink.link_uid, {
            method: 'put',
            body: {
                "title": webLink.title,
                "description": webLink.description,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(GetWorkspaceContents(folderUID))
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

// 删除收藏链接
export const DeletePersonalWebLink = (folderUID, workspaceUID, linkUID) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/folder/" + folderUID + "/workspace/" + workspaceUID + "/weblink/" + linkUID, {
            method: 'delete',
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("删除成功");
                }
                dispatch(GetWorkspaceContents(folderUID))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodStationCategory(records) {
    return {
        type: CUSTOMER_STATION_CATEGORY,
        records
    }
}

// 推荐好站分类
export const GetGoodStationCategory = () => {
    return (dispatch) => {
        let url = "/anonymous/station/category"
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "?user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodStationCategory(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodStation(records) {
    return {
        type: CUSTOMER_STATION,
        records
    }
}

function LoadAppendGoodStation(records) {
    return {
        type: CUSTOMER_STATION_APPEND,
        records
    }
}

// 推荐好站列表
export const GetGoodStationRecommend = (categoryUID, sorted, offset, limit) => {
    return (dispatch) => {
        let url = "/anonymous/station/recommend?category_uid=" + categoryUID + "&sorted=" + sorted + "&offset=" + offset + "&limit=" + limit
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            url = url + "&user_uid=" + userInfo.user_uid
        }
        request(url).then(res => {
            if (res.code === 200) {
                if (offset === 0) {
                    dispatch(LoadGoodStation(res.data))
                } else {
                    dispatch(LoadAppendGoodStation(res.data))
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 保存好站推荐
export const AddGoodStationRecommend = (station, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/recommend", {
            method: 'post',
            body: station,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("创建成功");
                }
                if (!isEmpty(callback)) {
                    callback()
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

export const UpdateGoodStationRecommend = (station, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/" + station.station_uid + "/detail", {
            method: 'put',
            body: station,
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

function LoadGoodStationMeta(records) {
    return {
        type: CUSTOMER_STATION_META,
        records
    }
}

// 更新好站推荐元数据
export const UpdateGoodStationMeta = (stationUID, metaType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/" + stationUID + "/meta", {
            method: 'put',
            body: {
                "meta_type": metaType
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodStationMeta(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 收藏好站推荐
export const GoodStationStar = (stationUID, folderUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/" + stationUID + "/star", {
            method: 'put',
            body: {
                "folder_uid": folderUID
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                message.success("收藏成功")
                dispatch(LoadGoodStationMeta(res.data))
            } else {
                if (res.data.code === ErrCustomerWebLinkLimit) {
                    callback()
                }
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodStationDetail(records) {
    return {
        type: CUSTOMER_STATION_DETAIL,
        records
    }
}

// 推荐好站详情
export const GetGoodStationDetail = (stationUID, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/" + stationUID + "/detail", {
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

// 更新社区图片
export const UpdatePublicCommunityImage = (folderUID, image) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/community/image", {
            method: 'put',
            body: {
                "folder_uid": folderUID,
                "image": image
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadGoodStationAudit(records) {
    return {
        type: CUSTOMER_STATION_AUDIT,
        records
    }
}

// 审核推荐好站内容
export const GetGoodStationAudit = () => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/audit", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadGoodStationAudit(res.data))
            } else {
                message.error(res.data.message);
            }
        });
    }
}

// 审核推荐好站内容
export const UpdateGoodStationAudit = (stationUID, status) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/station/" + stationUID, {
            method: 'put',
            body: {
                "status": status
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("更新成功");
                }
                dispatch(GetGoodStationAudit())
            } else {
                message.error(res.data.message);
            }
        });
    }
}