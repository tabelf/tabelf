import {Divider, Layout} from 'antd';
import React from 'react';
import './style.css'
import SidebarPage from "../sidebar/SidebarPage";
import HeaderPage from "../header/HeaderPage";
import ContentPage from "../content/ContentPage";
import FooterPage from "../footer/FooterPage";
import {checkUserInfo} from "../../actions/Base";

class HomePage extends React.Component {

    componentDidMount() {
        checkUserInfo()
    }

    render() {
        return (
            <Layout hasSider>
                <SidebarPage/>
                <Layout className="site-layout">
                    <HeaderPage/>
                    <ContentPage/>
                    <div>
                        <div className="home-container-footer">
                            <Divider style={{margin: 0}}/>
                            <FooterPage/>
                        </div>
                    </div>
                </Layout>
            </Layout>
        )
    }
}

export default HomePage;