import {CUSTOMER_TEAMS, HISTORY_RECENT_WEBLINK} from "../actions";
import {combineReducers} from "redux";

const initHistoryWebLinks = {
    web_links: []
}

function LoadHistoryRecentWebLink(state = initHistoryWebLinks, action) {
    switch (action.type) {
        case HISTORY_RECENT_WEBLINK:
            return action.records
        default:
            return state
    }
}

export const historyReducers = combineReducers({
    historyWebLinks: LoadHistoryRecentWebLink,
})
