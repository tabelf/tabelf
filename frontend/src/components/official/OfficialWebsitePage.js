import React, {Component} from 'react';
import SidebarPage from "../sidebar/SidebarPage";
import {Divider, Layout} from "antd";
import HeaderPage from "../header/HeaderPage";
import {Route, Routes} from "react-router-dom";
import ContentPage from "../content/ContentPage";
import HistoryPage from "../history/HistoryPage";
import RecyclePage from "../recycle/RecyclePage";
import CollaborationPage from "../collaboration/CollaborationPage";
import FooterPage from "../footer/FooterPage";

class OfficialWebsitePage extends Component {
    render() {
        return (
            <Layout hasSider>
                <SidebarPage/>
                <Layout className="site-layout">
                    <HeaderPage/>
                    <Routes>
                        <Route path="/content/:q/:folder_number" element={<ContentPage/>}/>
                        <Route path="/history" exact={false} element={<HistoryPage/>}/>
                        <Route path="/recycle" exact={false} element={<RecyclePage/>}/>
                        <Route path="/collaboration/:q/:folder_number" exact={false} element={<CollaborationPage/>}/>
                    </Routes>
                    <div>
                        <div className="home-container-footer">
                            <Divider style={{margin: 0}}/>
                            <FooterPage/>
                        </div>
                    </div>
                </Layout>
            </Layout>
        );
    }
}

export default OfficialWebsitePage;