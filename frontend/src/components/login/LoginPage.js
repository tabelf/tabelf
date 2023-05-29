import React, {Component} from 'react';
import {Button, Card, Divider, Image, Input, Layout, message, Popover, Avatar} from 'antd';
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as LoginAction from "../../actions/LoginAction";
import './style.css'
import {getUserInfo, isNotEmpty} from "../../actions/Base";
import {withRouter} from "../base/Base";
import WechatImage from '../../assets/wechat.png'
import HintImage from '../../assets/hint.jpeg'
import LoginImage from '../../assets/qrcode_login.jpg'
import FooterPage from "../footer/FooterPage";
import {EHintIcon} from "../base/EIcon";

const {Content} = Layout;

class LoginPage extends Component {

    state = {
        authCode: "",
    }

    componentDidMount() {
        let userInfo = getUserInfo();
        if (isNotEmpty(userInfo) && isNotEmpty(userInfo.user_uid)) {
            window.location.href = "/workspace"
        }
    }

    handleAuthCodeChange = ({target}) => {
        this.setState({authCode: target.value})
    }

    handleLoginClick = () => {
        if (this.state.authCode.length !== 6) {
            message.error("验证码错误");
            return;
        }
        const referralUID = this.props.params.referral_uid
        this.props.actions.AuthLogin(this.state.authCode, referralUID)
    }

    render() {
        return (
            <Layout className={"site-update-log-layout"}>
                <Content>
                    <Card bordered={false}>
                        <div className="login-layout">
                            <div style={{marginTop: "9%"}}>
                                <div>
                                    <h2>用户登录</h2>
                                </div>
                                <div className="wechat-scan">
                                    <img style={{
                                        width: "32px",
                                        height: "25px",
                                    }} src={WechatImage}/>
                                    <span>微信扫码关注公众号</span>
                                    <div>
                                        <span>输入「验证码」并发送进行获取验证码</span>
                                        <span>&nbsp;
                                            <Popover placement="right" content={<>
                                                <Image
                                                    width={271}
                                                    src={HintImage}
                                                    preview={false}
                                                />
                                            </>}>
                                                <EHintIcon/>
                                              </Popover>
                                        </span>
                                    </div>
                                </div>
                                <div>
                                    <Avatar
                                        shape="square"
                                        src={LoginImage}
                                        size={265}
                                    />
                                </div>
                                <div>
                                    <Input
                                        allowClear
                                        placeholder="输入6位验证码"
                                        style={{
                                            width: 230,
                                            height: 40,
                                            marginTop: 10
                                        }}
                                        onChange={this.handleAuthCodeChange}
                                    />
                                </div>
                                <div>
                                    <Button
                                        type="primary"
                                        block
                                        style={{
                                            width: 230,
                                            height: 40,
                                            marginTop: 10,
                                        }}
                                        onClick={this.handleLoginClick}
                                    >登录</Button>
                                </div>
                            </div>
                        </div>
                    </Card>
                </Content>

                <div className="container-footer">
                    <Divider style={{margin: 0}}/>
                    <FooterPage/>
                </div>
            </Layout>
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.login}),
    dispatch => ({
        actions: bindActionCreators(LoginAction, dispatch)
    })
)(LoginPage));