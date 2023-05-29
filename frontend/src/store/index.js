import { legacy_createStore as createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';
import reducersDataManagers from '../reducers'

// 引入中间件 redux-thunk, 异步处理中间件
// 中间件一但加上，dispatch 就会被中间件拦截下来
const middleware = applyMiddleware(thunk);

// reduces 整合
const reducers = combineReducers({
    dataManage: reducersDataManagers
});

const store = createStore(reducers, middleware);

export default store;