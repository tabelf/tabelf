import React, {Component} from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom'
import HomePage from "./components/home/HomePage";
import LoginPage from "./components/login/LoginPage";
import ShareURLPage from "./components/collaboration/ShareURLPage";
import UpdateLogPage from "./components/updatelog/UpdateLogPage";
import Disclaimer from "./components/base/Disclaimer";
import RecommendPage from "./components/recommend/RecommendPage";
import Business from "./components/base/Business";
import CommunityPage from "./components/community/CommunityPage";
import CommunityDetailPage from "./components/community/CommunityDetailPage";
import CommunityTabPage from "./components/community/CommunityTabPage";
import RecommendAuditPage from "./components/recommend/RecommendAuditPage";
import GoodArticlePage from "./components/article/GoodArticlePage";
import ArticleCollectionPage from "./components/article/ArticleCollectionPage";
import ArticleSubmitPage from "./components/article/ArticleSubmitPage";
import ArticleAuditPage from "./components/article/ArticleAuditPage";
import OfficialWebsitePage from "./components/official/OfficialWebsitePage";

import indexHTML from './components/official/template/index.html';
import loginIndexHTML from './components/official/template/login_index.html';
import './components/official/template/static/css/animate.min.css'
import './components/official/template/static/css/cf.errors.css'
import './components/official/template/static/css/css2.css'
import './components/official/template/static/css/slick.css'
import './components/official/template/static/css/tailwind-built.css'
import {getUserInfo, isNotEmpty} from "./actions/Base";

// Routes 替换 Switch
// Route 必须作为 Routes 的子节点存在
//  如果不使用/*，该组件二级页面不会生效
class App extends Component {
    render() {
        let innerHtml = loginIndexHTML
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo)) {
            innerHtml = indexHTML
        }
        return (
            <BrowserRouter>
                <Routes>
                    <Route path="/" exact={false} element={<div dangerouslySetInnerHTML={{ __html: innerHtml }} />}/>
                    <Route path="/*" exact={false} element={<OfficialWebsitePage />}/>
                    <Route path="/workspace" exact={false} element={<HomePage/>}/>
                    <Route path="/v/:share_uid" exact={false} element={<ShareURLPage/>}/>
                    <Route path="/login" exact={false} element={<LoginPage/>}/>
                    <Route path="/r/:referral_uid" exact={false} element={<LoginPage/>}/>
                    <Route path="/update/log" exact={false} element={<UpdateLogPage/>}/>
                    <Route path="/disclaimer" exact={false} element={<Disclaimer/>}/>
                    <Route path="/recommend" exact={false} element={<RecommendPage/>}/>
                    <Route path="/recommend/audit" exact={false} element={<RecommendAuditPage/>}/>
                    <Route path="/business" exact={false} element={<Business/>}/>
                    <Route path="/community" exact={false} element={<CommunityPage/>}/>
                    <Route path="/community/tab/:tab_id" exact={false} element={<CommunityTabPage/>}/>
                    <Route path="/detail/:community_uid" exact={false} element={<CommunityDetailPage/>}/>
                    <Route path="/good/article" exact={false} element={<GoodArticlePage/>}/>
                    <Route path="/good/article/collections" exact={false} element={<ArticleCollectionPage/>}/>
                    <Route path="/good/article/submit" exact={false} element={<ArticleSubmitPage/>}/>
                    <Route path="/good/article/audit" exact={false} element={<ArticleAuditPage/>}/>
                </Routes>
            </BrowserRouter>
        );
    }
}

export default App;