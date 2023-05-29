import {combineReducers} from 'redux';
import {loginReducers} from "./LoginReducer";
import {customerReducers} from "./CustomerReducer";
import {historyReducers} from "./HistoryReducer";
import {recycleReducers} from "./RecycleReducer";
import {inviteReducers} from "./InviteReducer";
import {orderReducers} from "./OrderReducer";
import {communityReducers} from "./CommunityReducer";
import {articleReducers} from "./GoodArticleReducer";

// reducers 多个函数聚合管理
const reducersDataManagers = combineReducers({
    login: loginReducers,
    customer: customerReducers,
    history: historyReducers,
    recycle: recycleReducers,
    invite: inviteReducers,
    order: orderReducers,
    community: communityReducers,
    article: articleReducers
});

export default reducersDataManagers;