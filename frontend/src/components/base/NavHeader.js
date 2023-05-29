import React, {Component} from 'react';
import {Col, Layout, Menu, Row, Space, Button} from "antd";
import LogoImage from "../../assets/logo.png";
import {NavLink as Link} from "react-router-dom";
import MessagePage from "../header/MessagePage";
import UserCenterPage from "../header/UserCenterPage";
import './style.css'
import {getUserInfo, isEmpty} from "../../actions/Base";

const {Header} = Layout;

class NavHeader extends Component {

    render() {
        const userinfo = getUserInfo()
        return (
            <div>
                <Header className={"base-layout-header"}>
                    <Row>
                        <Col span={7}>
                            <div>
                                <img src={LogoImage} alt="logo" className="base-header-logo"/>
                            </div>
                        </Col>
                        <Col span={9}>
                            <Menu mode="horizontal"
                                  defaultSelectedKeys={[this.props.active]}
                            >
                                <Menu.Item key="my-personal">
                                    <Link to="/workspace">个人</Link>
                                </Menu.Item>
                                <Menu.Item key="my-good-station">
                                    <Link to="/recommend">好站推荐</Link>
                                </Menu.Item>
                                <Menu.Item key="my-community">
                                    <Link to="/community">分享社区</Link>
                                </Menu.Item>
                                <Menu.Item key="my-good-article">
                                    <Link to="/good/article">好文精选</Link>
                                </Menu.Item>
                            </Menu>
                        </Col>
                        <Col span={8}>
                            <div className="header-user">
                                {
                                    isEmpty(userinfo) ? (
                                        <Button href={"/login"} type={"primary"}>登录</Button>
                                    ) : (
                                        <Space size={"large"}>
                                            <MessagePage/>
                                            <UserCenterPage/>
                                        </Space>
                                    )
                                }
                            </div>
                        </Col>
                    </Row>
                </Header>
            </div>
        );
    }
}

export default NavHeader;