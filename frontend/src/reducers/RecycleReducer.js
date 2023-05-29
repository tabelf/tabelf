import {RECYCLE_WEBLINK} from "../actions";
import {combineReducers} from "redux";

const initRecycleWebLinks = {
    web_links: []
}

function LoadHistoryRecentWebLink(state = initRecycleWebLinks, action) {
    switch (action.type) {
        case RECYCLE_WEBLINK:
            return action.records
        default:
            return state
    }
}

export const recycleReducers = combineReducers({
    recycleWebLinks: LoadHistoryRecentWebLink,
})
