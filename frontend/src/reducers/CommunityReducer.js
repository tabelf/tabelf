import {combineReducers} from "redux";
import {
    CUSTOMER_COMMUNITY, CUSTOMER_COMMUNITY_AUDIT, CUSTOMER_COMMUNITY_AUDIT_STATUS,
    CUSTOMER_COMMUNITY_CATEGORY,
    CUSTOMER_COMMUNITY_DETAIL,
    CUSTOMER_COMMUNITY_META,
    CUSTOMER_COMMUNITY_SELF
} from "../actions";

const initCommunityCategory = {
    categories: []
}

function LoadCommunityCategory(state = initCommunityCategory, action) {
    switch (action.type) {
        case CUSTOMER_COMMUNITY_CATEGORY:
            return action.records
        default:
            return state
    }
}

const initCommunities = {
    total: 0,
    data: []
}

function LoadCommunities(state = initCommunities, action) {
    switch (action.type) {
        case CUSTOMER_COMMUNITY:
            return action.records
        default:
            return state
    }
}

const initCommunityDetail = {
    tags: []
}

function LoadCommunityDetail(state = initCommunityDetail, action) {
    switch (action.type) {
        case CUSTOMER_COMMUNITY_DETAIL:
            return action.records
        case CUSTOMER_COMMUNITY_META:
            return {
                ...state,
                ...action.records,
            }
        default:
            return state
    }
}

const initSelfCommunities = {
    data: []
}

function LoadSelfCommunities(state = initSelfCommunities, action) {
    switch (action.type) {
        case CUSTOMER_COMMUNITY_SELF:
            return action.records
        default:
            return state
    }
}

const initAuditCommunity = {
    personal_workspaces: []
}

function LoadAuditCommunity(state = initAuditCommunity, action) {
    switch (action.type) {
        case CUSTOMER_COMMUNITY_AUDIT:
            return action.records
        case CUSTOMER_COMMUNITY_AUDIT_STATUS:
            return {
                ...state,
                ...action.records,
            }
        default:
            return state
    }
}

export const communityReducers = combineReducers({
    communityCategories: LoadCommunityCategory,
    communities: LoadCommunities,
    detail: LoadCommunityDetail,
    selfCommunities: LoadSelfCommunities,
    audit: LoadAuditCommunity
})