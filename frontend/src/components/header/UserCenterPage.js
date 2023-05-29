import React, {Component} from 'react';
import {Avatar, Badge, Dropdown, Menu, Space} from "antd";
import {CaretDownOutlined, LikeOutlined, UserOutlined, ShareAltOutlined, CompassOutlined} from "@ant-design/icons";
import {EHotIcon, EInviteIcon, ELogoutIcon, EOrderMembershipIcon, EOrderTransactionIcon, EVipIcon} from "../base/EIcon";
import {NavLink as Link} from "react-router-dom";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import UserInfoPage from "./UserInfoPage";
import InvitePage from "./InvitePage";
import UpgradePage from "../recharge/UpgradePage";
import TransactionPage from "../recharge/TransactionPage";
import {clearUserInfo, toLogin} from "../../actions/Base";
import './style.css'

class UserCenterPage extends Component {


    handleUserBaseOpen = () => {
        this.userInfoModalRef.handleUserBaseOpen()
    }

    handleLogout = () => {
        if (clearUserInfo()) {
            toLogin()
        }
    }

    handleInviteOpen = () => {
        this.inviteModalRef.handleOpenInvite();
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    handleOpenTransaction = () => {
        this.transactionModalRef.handleOpenTransaction()
    };


    render() {

        const {accountInfo} = this.props.state

        return (
            <div>
                <Space size={"large"}>
                    <Dropdown
                        overlayClassName={"account-image-menu-extra"}
                        overlay={
                            <Menu>
                                <Menu.Item key="my-user"
                                           icon={<UserOutlined/>}
                                           onClick={this.handleUserBaseOpen}>个人消息</Menu.Item>
                                <Menu.Item key="my-invite"
                                           icon={<EInviteIcon/>}
                                           onClick={this.handleInviteOpen}>
                                    邀请有礼
                                </Menu.Item>
                                <Menu.Item key="my-good-article" icon={<CompassOutlined />}>
                                    <Link to="/good/article" target={"_blank"}>好文精选&nbsp;<EHotIcon/></Link>
                                </Menu.Item>
                                <Menu.Item key="my-community" icon={<ShareAltOutlined />}>
                                    <Link to="/community" target={"_blank"}>分享社区</Link>
                                </Menu.Item>
                                <Menu.Item key="my-good-station" icon={<LikeOutlined />}>
                                    <Link to="/recommend" target={"_blank"}>好站推荐</Link>
                                </Menu.Item>
                                <Menu.Divider/>
                                <Menu.Item key="my-logout"
                                           icon={<ELogoutIcon/>}
                                           onClick={this.handleLogout}>退出登录</Menu.Item>
                            </Menu>
                        }
                        overlayStyle={{width: "150px"}}
                        placement="bottom"
                        arrow>
                        <Badge count={accountInfo.has_membership ? <EVipIcon/> : 0} offset={[-6, 26]}>
                            <Avatar src={accountInfo.image}/>
                        </Badge>
                    </Dropdown>

                    <div>
                        <Link onClick={this.handleUserBaseOpen}>
                            <Space>
                                <div title={accountInfo.user_name} className={"account-username"}>
                                    {accountInfo.user_name}
                                </div>
                                <div>
                                    <CaretDownOutlined style={{color: "rgba(0, 0, 0, 0.45)"}}/>
                                </div>
                            </Space>
                        </Link>
                    </div>
                </Space>

                <UserInfoPage bindRef={(ref) => this.userInfoModalRef = ref}/>

                <InvitePage bindRef={(ref) => this.inviteModalRef = ref}/>

                <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>

                <TransactionPage bindRef={(ref) => this.transactionModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(UserCenterPage);