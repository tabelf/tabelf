import {combineReducers} from 'redux';
import {
    CUSTOMER_ACCOUNT,
    CUSTOMER_ACCOUNT_MESSAGE,
    CUSTOMER_FOLDERS,
    CUSTOMER_MENU_OPEN_KEY,
    CUSTOMER_MENU_SELECT_KEY,
    CUSTOMER_OPEN_ACTIVE_FOLDERS,
    CUSTOMER_SEARCH,
    CUSTOMER_SHARE_ALL_PERSONNEL, CUSTOMER_SHARE_COLL,
    CUSTOMER_SHARE_FOLDER, CUSTOMER_SHARE_FOLDER_CLEAR,
    CUSTOMER_SHARE_FOLDERS, CUSTOMER_SHARE_INFO,
    CUSTOMER_STATION, CUSTOMER_STATION_APPEND, CUSTOMER_STATION_AUDIT,
    CUSTOMER_STATION_CATEGORY, CUSTOMER_STATION_DETAIL, CUSTOMER_STATION_META,
    CUSTOMER_TEAMS,
    CUSTOMER_UPDATE_ACCOUNT,
    CUSTOMER_WEBLINKS,
    CUSTOMER_WORKSPACE,
    CUSTOMER_WORKSPACE_SWITCH,
    CUSTOMER_WORKSPACES
} from "../actions";
import {GetGoodStationAudit} from "../actions/CustomerAction";

const initSearchLinks = {
    search_web_links: []
}

function LoadSearchLinks(state = initSearchLinks, action) {
    switch (action.type) {
        case CUSTOMER_SEARCH:
            return action.records
        default:
            return state
    }
}

const initTeams = {
    personal_teams: []
}

function LoadPersonalTeams(state = initTeams, action) {
    switch (action.type) {
        case CUSTOMER_TEAMS:
            return action.records
        default:
            return state
    }
}

const initFolders = {
    folders: []
}

function LoadPersonalFolders(state = initFolders, action) {
    switch (action.type) {
        case CUSTOMER_FOLDERS:
            return action.records
        default:
            return state
    }
}

const initAccountInfo = {}

function LoadPersonalAccountInfo(state = initAccountInfo, action) {
    switch (action.type) {
        case CUSTOMER_ACCOUNT:
            return action.records
        case CUSTOMER_UPDATE_ACCOUNT:
            return {
                ...state,
                ...action.records
            }
        default:
            return state
    }
}

const initWebLink = {
    has_limit: false,
    used_quantity: 0,
    total_quantity: 0
}

function LoadWebLinkInfo(state = initFolders, action) {
    switch (action.type) {
        case CUSTOMER_WEBLINKS:
            return action.records
        default:
            return state
    }
}


const initShareFolders = {
    openKey: "",
    selectKey: "",
    folders: []
}

function LoadSharePersonalFolders(state = initShareFolders, action) {
    switch (action.type) {
        case CUSTOMER_SHARE_FOLDERS:
            return action.records
        default:
            return state
    }
}

const initMenuActive = {
    openKey: [],
    selectKey: []
}

function LoadMenuActive(state = initMenuActive, action) {
    switch (action.type) {
        case CUSTOMER_OPEN_ACTIVE_FOLDERS:
            return action.records
        case CUSTOMER_MENU_OPEN_KEY:
            return {
                ...state,
                openKey: action.records.openKey
            }
        case CUSTOMER_MENU_SELECT_KEY:
            return {
                ...state,
                selectKey: action.records.selectKey
            }
        default:
            return state
    }
}

const initShareFolder = {
    share_num: 0,
    share_personnel: [],
}

function LoadSharePersonalFolder(state = initShareFolder, action) {
    switch (action.type) {
        case CUSTOMER_SHARE_FOLDER:
            const newData = {
                ...action.records,
                share_num: action.records.share_num,
                share_personnel: state.share_personnel,
            }
            newData.share_personnel.push(...action.records.share_personnel)
            return newData
        case CUSTOMER_SHARE_FOLDER_CLEAR:
            return {
                share_num: 0,
                share_personnel: [],
            }
        case CUSTOMER_SHARE_COLL:
            const newState = state.share_personnel.filter(p => {
                if (p.coll_uid === action.records.coll_uid) {
                    if (action.records.type === "3") { //移除
                        return false
                    }
                    p.authority = action.records.type
                }
                return true
            })
            return {
                ...state,
                "share_personnel": newState
            }
        case CUSTOMER_SHARE_INFO:
            return {
                ...state,
                "authority": action.records.authority,
                "expired_day": action.records.expired_day,
            }
        default:
            return state
    }
}

const initShareAllPersonnel = []

function LoadShareAllPersonnel(state = initShareAllPersonnel, action) {
    switch (action.type) {
        case CUSTOMER_SHARE_ALL_PERSONNEL:
            return action.records
        default:
            return state
    }
}

const initWorkspaces = {
    loading: true,
    personal_workspaces: []
}

function LoadWorkspaceContents(state = initWorkspaces, action) {
    switch (action.type) {
        case CUSTOMER_WORKSPACES:
            return action.records
        case CUSTOMER_WORKSPACE:
            let idx1, idx2 = 0
            const newWorkspaces = [...state.personal_workspaces]
            newWorkspaces.map((s, index) => {
                if (s.workspace_uid === action.records.new_workspace.workspace_uid) {
                    idx1 = index
                }
                if (s.workspace_uid === action.records.old_workspace.workspace_uid) {
                    idx2 = index
                }
            })
            newWorkspaces[idx1] = action.records.new_workspace
            newWorkspaces[idx2] = action.records.old_workspace
            return {
                ...state,
                personal_workspaces: newWorkspaces
            }
        case CUSTOMER_WORKSPACE_SWITCH:
            const sourceWorkspaces = [...state.personal_workspaces]
            sourceWorkspaces.map(s => {
                if (action.records.indexOf(s.workspace_uid) !== -1) {
                    s.is_open = true
                } else {
                    s.is_open = false
                }
            })
            return {
                ...state,
                personal_workspaces: sourceWorkspaces,
                active_workspace_uids: action.records
            }
        default:
            return state
    }
}

const initAccountMessage = {
    unread: 0,
    user_messages: []
}

function LoadAccountMessages(state = initAccountMessage, action) {
    switch (action.type) {
        case CUSTOMER_ACCOUNT_MESSAGE:
            return action.records
        default:
            return state
    }
}

const initStationCategory = {
    categories: []
}

function LoadStationCategory(state = initStationCategory, action) {
    switch (action.type) {
        case CUSTOMER_STATION_CATEGORY:
            return action.records
        default:
            return state
    }
}

const initStations = {
    total: 0,
    data: []
}

function LoadStations(state = initStations, action) {
    switch (action.type) {
        case CUSTOMER_STATION:
            return action.records
        case CUSTOMER_STATION_APPEND:
            const newData = {
                total: action.records.total,
                data: state.data,
            }
            newData.data.push(...action.records.data)
            return newData
        case CUSTOMER_STATION_META:
            const newMetaData = state.data.map(item => {
                if (item.uid === action.records.uid) {
                    return {
                        ...item,
                        ...action.records,
                    }
                }
                return item
            })
            return {
                ...state,
                "data": newMetaData
            }
        default:
            return state
    }
}


const initStationDetail = {}

function LoadStationDetail(state = initStationDetail, action) {
    switch (action.type) {
        case CUSTOMER_STATION_DETAIL:
            return action.records
        default:
            return state
    }
}

const initStationAudit = {
    data: []
}

function LoadGoodStationAudit(state = initStationAudit, action) {
    switch (action.type) {
        case CUSTOMER_STATION_AUDIT:
            return action.records
        default:
            return state
    }
}

export const customerReducers = combineReducers({
    personalTeams: LoadPersonalTeams,
    personalFolders: LoadPersonalFolders,
    sharePersonalFolders: LoadSharePersonalFolders,
    personalWorkspaces: LoadWorkspaceContents,
    sharePersonalFolder: LoadSharePersonalFolder,
    menuActive: LoadMenuActive,
    personalWebLink: LoadWebLinkInfo,
    shareAllPersonnel: LoadShareAllPersonnel,
    accountInfo: LoadPersonalAccountInfo,
    searchLinks: LoadSearchLinks,
    messages: LoadAccountMessages,
    stationCategories: LoadStationCategory,
    stations: LoadStations,
    stationDetail: LoadStationDetail,
    stationAudits: LoadGoodStationAudit
})
