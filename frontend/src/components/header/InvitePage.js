import React, {Component} from 'react';
import {Avatar, Card, Col, Divider, Input, List, Modal, Row, Steps, Button, message, Empty} from 'antd';
import {UserOutlined} from "@ant-design/icons";
import {EInviteLinkIcon, EInviteLoginIcon, EInviteSendIcon} from "../base/EIcon";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as InviteAction from "../../actions/InviteAction";
import copy from "copy-to-clipboard";

class InvitePage extends Component {

    state = {
        hasOpenInvite: false
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleOpenInvite = () => {
        this.props.actions.GetPersonalInvite(() => {
            this.setState({hasOpenInvite: true})
        })
    }

    handleCloseInvite = () => {
        this.setState({hasOpenInvite: false})
    }

    handleCopyShareURL = (shareURL) => {
        if (copy(shareURL)) {
            message.success("复制成功")
        }
    }

    render() {
        const {invites} = this.props.state;
        return (
            <Modal title={null}
                   open={this.state.hasOpenInvite}
                   onCancel={this.handleCloseInvite}
                   forceRender={true}
                   footer={null}
                   width={820}
                   className={"invite-modal"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>邀请好友获得奖励</h3>
                    </span>
                <div className={"ant-modal-confirm-content"} style={{height: 410}}>
                    <div style={{marginTop: 30}}>
                        <div>
                            <Steps
                                labelPlacement="vertical"
                                items={[
                                    {
                                        icon: <UserOutlined style={{opacity: 0}}/>
                                    },
                                    {
                                        title: <span style={{fontSize: "14px"}}>1. 复制专属链接</span>,
                                        icon: <EInviteLinkIcon/>
                                    },
                                    {
                                        title: <span style={{fontSize: "14px"}}>2. 将链接分享到 QQ / 微信等</span>,
                                        icon: <EInviteSendIcon/>
                                    },
                                    {
                                        title: <span style={{fontSize: "14px"}}>3. 好友注册, 奖励5个链接数</span>,
                                        icon: <EInviteLoginIcon/>
                                    },
                                    {
                                        icon: <UserOutlined style={{opacity: 0}}/>
                                    }
                                ]}
                            />
                        </div>
                        <div style={{marginTop: 40}}>
                            <Row>
                                <Col offset={5} span={15}>
                                    <Input size="large"
                                           disabled
                                           value={invites.invite_url}
                                           style={{ width: 376 }}
                                    />
                                    <Button
                                        type="primary"
                                        size="large"
                                        onClick={() => this.handleCopyShareURL(invites.invite_url)}
                                    >
                                        复制链接
                                    </Button>
                                </Col>
                            </Row>
                        </div>
                        <div style={{marginTop: 40}}>
                            <Row>
                                <Col offset={5} span={15}>
                                    <Card title={"我的成就"}>
                                        <Row>
                                            <Col offset={3} span={6}>
                                                <div style={{textAlign: "center"}}>
                                                    <h2>{invites.invited}</h2>
                                                    <p>邀请人数</p>
                                                </div>
                                            </Col>
                                            <Col offset={3} span={3}>
                                                <Divider style={{height: 70}} type={"vertical"}/>
                                            </Col>
                                            <Col span={6}>
                                                <div style={{textAlign: "center"}}>
                                                    <h2>{invites.earned}</h2>
                                                    <p>获取链接数</p>
                                                </div>
                                            </Col>
                                        </Row>
                                    </Card>
                                </Col>
                            </Row>
                        </div>
                        <div style={{marginTop: 40}}>
                            <Row>
                                <Col offset={5} span={15}>
                                    <Card title={"我的邀请记录"}>
                                        <List
                                            itemLayout="horizontal"
                                            split={false}
                                            dataSource={invites.achievements}
                                            locale={{emptyText: <Empty description={"邀请记录为空"}
                                                                       style={{marginTop: "20%"}}
                                                                       image={Empty.PRESENTED_IMAGE_SIMPLE}/>}}
                                            className="invite-user-list"
                                            renderItem={item => (
                                                <List.Item>
                                                    <List.Item.Meta
                                                        avatar={<Avatar size={30} src={item.image}/>}
                                                        title={<span>{item.user_name}</span>}
                                                        description={item.created_at}
                                                    />
                                                </List.Item>
                                            )}
                                        />
                                    </Card>
                                </Col>
                            </Row>
                        </div>
                        <div style={{height: 40}}/>
                    </div>
                </div>
            </Modal>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.invite}),
    dispatch => ({
        actions: bindActionCreators(InviteAction, dispatch)
    })
)(InvitePage);