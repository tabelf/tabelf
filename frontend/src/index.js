import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import {Provider} from "react-redux";
import store from "./store";

const root = ReactDOM.createRoot(document.getElementById('root'));
// React.StrictMode 主要是检查提示, 代码中存在不规范操作就会提示, 坑!!! 会导致你的代码执行两次
root.render(
    <Provider store={store}>
        <App/>
    </Provider>
);
