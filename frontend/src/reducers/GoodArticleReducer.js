import {
    CUSTOMER_ARTICLE,
    CUSTOMER_ARTICLE_APPEND,
    CUSTOMER_ARTICLE_CATEGORY, CUSTOMER_ARTICLE_COLLECTION, CUSTOMER_ARTICLE_HOT,
    CUSTOMER_ARTICLE_META, CUSTOMER_ARTICLE_UPDATE,
    CUSTOMER_ARTICLE_SUBMIT, CUSTOMER_ARTICLE_MENU, CUSTOMER_ARTICLE_AUDIT
} from "../actions";
import {combineReducers} from "redux";

const initArticleCategory = {
    categories: []
}

function LoadArticleCategory(state = initArticleCategory, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_CATEGORY:
            return action.records
        default:
            return state
    }
}

const initArticles = {
    total: 0,
    data: []
}

function LoadArticles(state = initArticles, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE:
            return action.records
        case CUSTOMER_ARTICLE_APPEND:
            const newData = {
                total: action.records.total,
                data: state.data,
            }
            newData.data.push(...action.records.data)
            return newData
        case CUSTOMER_ARTICLE_UPDATE:
        case CUSTOMER_ARTICLE_META:
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

const initArticleCollection = {
    data: []
}

function LoadArticleCollection(state = initArticleCollection, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_COLLECTION:
            return action.records
        default:
            return state
    }
}

const initArticleSubmit = {
    data: []
}

function LoadArticleSubmit(state = initArticleSubmit, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_SUBMIT:
            return action.records
        case CUSTOMER_ARTICLE_UPDATE:
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

const initArticleAudit = {
    data: []
}

function LoadArticleAudit(state = initArticleAudit, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_AUDIT:
            return action.records
        default:
            return state
    }
}

const initHotArticles = {
    data: []
}

function LoadHotGoodArticle(state = initHotArticles, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_HOT:
            return action.records
        default:
            return state
    }
}

const initArticleMenu = {}

function LoadMenuGoodArticle(state = initArticleMenu, action) {
    switch (action.type) {
        case CUSTOMER_ARTICLE_MENU:
            return action.records
        default:
            return state
    }
}

export const articleReducers = combineReducers({
    articleCategories: LoadArticleCategory,
    articles: LoadArticles,
    collections: LoadArticleCollection,
    hotArticle: LoadHotGoodArticle,
    submit: LoadArticleSubmit,
    menu: LoadMenuGoodArticle,
    audit: LoadArticleAudit
})