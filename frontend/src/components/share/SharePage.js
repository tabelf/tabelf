import React, {Component} from 'react';
import {Avatar, Button, Col, Divider, Dropdown, Input, List, message, Modal, Row, Space, Badge, Spin} from "antd";
import {CaretDownOutlined, CheckOutlined} from "@ant-design/icons";
import './style.css'
import copy from "copy-to-clipboard";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {EVipIcon} from "../base/EIcon";

const authority = [{
    label: '可查看',
    key: '0',
}, {
    label: '可编辑',
    key: '1',
}];

// 分享协作弹窗
class SharePage extends Component {

    state = {
        hasOpenShare: false,
        inviteAuthority: "0",
        emailValue: "",
        offset: 1,
        limit: 30,
        folderUID: "",
        loading: true
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleCancel = () => {
        this.props.actions.ClearSharePersonal()
        this.setState({hasOpenShare: false})
    };

    handleOpenShare = (folderUID) => {
        const {offset, limit} = this.state;
        this.props.actions.SharePersonalFolder(folderUID, (offset - 1) * limit, limit, () => {
            this.setState({hasOpenShare: true})
        })
        this.setState({folderUID: folderUID})
    }

    handleShareAuthorityUpdate = (key, share) => {
        this.props.actions.UpdateSharePersonalFolder(
            share.folder_uid,
            share.share_uid,
            key.key,
            share.expired_day,
        )
    }

    handleShareValidUpdate = (key, share) => {
        this.props.actions.UpdateSharePersonalFolder(
            share.folder_uid,
            share.share_uid,
            share.authority,
            parseInt(key.key),
        )
    }

    handleCopyShareURL = (shareURL) => {
        if (copy(shareURL)) {
            message.success("复制成功")
        }
    }

    handleShareCollEdit = (key, share, item) => {
        this.props.actions.UpdateCollSharePersonalFolder(share.folder_uid, share.share_uid, item.coll_uid, key.key)
    }

    handleInviteAuthority = (key) => {
        this.setState({inviteAuthority: key.key})
    }

    handleShareInvite = (shareUID) => {
        this.props.actions.ShareFriendFolder(
            shareUID,
            this.state.emailValue,
            this.state.inviteAuthority,
            () => {
                this.setState({emailValue: ""})
            }
        )
    }

    handleEmailChange = ({target: {value}}) => {
        this.setState({emailValue: value})
    }

    onLoadMore = () => {
        let {offset, limit, folderUID} = this.state;
        offset = offset + 1;
        this.props.actions.SharePersonalFolder(folderUID, (offset - 1) * limit, limit)
        this.setState({offset: offset});
    }

    render() {
        const {sharePersonalFolder} = this.props.state

        const loadMore = (sharePersonalFolder.share_personnel.length < sharePersonalFolder.share_num) && this.state.loading ? (
                <div style={{
                        textAlign: 'center',
                        marginTop: 12,
                        height: 32,
                        lineHeight: '32px',
                    }}>
                    <Button type={"text"} onClick={this.onLoadMore}>加载更多</Button>
                </div>
            ) : null;

        return (
            <div>
                <Modal title={null}
                       open={this.state.hasOpenShare}
                       footer={null}
                       onCancel={this.handleCancel}
                       width={610}
                       forceRender={true}
                       getContainer={false}
                       className={"share-folder"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>分享协作</h3>
                    </span>
                    <div className={"ant-modal-confirm-content"}>
                        <div>
                            <div>
                                <Row>
                                    <Col span={20}>
                                        <Input size="large"
                                               className="share-email-input"
                                               style={{width: 457}}
                                               placeholder="通过邮件邀请"
                                               value={this.state.emailValue}
                                               onChange={this.handleEmailChange}
                                               disabled={!sharePersonalFolder.has_owner}
                                               suffix={
                                                   <span>
                                                       <Divider type="vertical"/>
                                                       <Dropdown
                                                           overlayClassName={"share-email-input-dropdown share-edit-dropdown"}
                                                           overlayStyle={{
                                                               width: 95
                                                           }}
                                                           menu={{
                                                               items: authority,
                                                               onClick: (key) => this.handleInviteAuthority(key)
                                                           }}
                                                           trigger={['click']}>
                                                           <Space>
                                                               {this.state.inviteAuthority === "0" ? "可查看" : "可编辑"}
                                                               <CaretDownOutlined/>
                                                           </Space>
                                                       </Dropdown>
                                                   </span>
                                               }
                                        />
                                    </Col>
                                    <Col span={4}>
                                        <Button
                                            disabled={!sharePersonalFolder.has_owner}
                                            style={{height: 35, borderRadius: 4, marginTop: 15, marginLeft: 10}}
                                            type="primary"
                                            onClick={() => this.handleShareInvite(sharePersonalFolder.share_uid)}>
                                            发起邀请
                                        </Button>
                                    </Col>
                                </Row>
                            </div>
                            <div className="share-number">已有{sharePersonalFolder.share_num}人加入协作</div>
                            <div>
                                <List
                                    itemLayout="horizontal"
                                    className="share-user-list"
                                    split={false}
                                    loadMore={loadMore}>
                                    {sharePersonalFolder.share_personnel.map(item => (
                                        <List.Item actions={[
                                            <div>
                                                {
                                                    item.sequence === 0 ? (
                                                        <div>
                                                            <Space>
                                                                创建者
                                                                <CaretDownOutlined style={{opacity: 0}}/>
                                                            </Space>
                                                        </div>
                                                    ) : (
                                                        <div>
                                                            <Dropdown overlayClassName={"share-edit-dropdown"}
                                                                      overlayStyle={{
                                                                          width: 95
                                                                      }}
                                                                      menu={{
                                                                          items: [{
                                                                              label: '可查看',
                                                                              key: '0',
                                                                              icon: <CheckOutlined className={item.authority === "0" ? "" : "hidden-opacity"}/>
                                                                          }, {
                                                                              label: '可编辑',
                                                                              key: '1',
                                                                              icon: <CheckOutlined className={item.authority === "1" ? "" : "hidden-opacity"}/>
                                                                          }, {
                                                                              type: 'divider',
                                                                          }, {
                                                                              label: '移除',
                                                                              key: '3',
                                                                              icon: <CheckOutlined className={"hidden-opacity"}/>
                                                                          }],
                                                                          onClick: (key) => this.handleShareCollEdit(key, sharePersonalFolder, item)
                                                                      }}
                                                                      trigger={sharePersonalFolder.has_owner ? ['click'] : []}>
                                                                <Space>
                                                                    {item.authority === "0" ? "可查看" : "可编辑"}
                                                                    <CaretDownOutlined
                                                                        className={sharePersonalFolder.has_owner ? "" : "dropdown-fade"}
                                                                    />
                                                                </Space>
                                                            </Dropdown>
                                                        </div>
                                                    )
                                                }
                                            </div>
                                        ]}>
                                            <List.Item.Meta
                                                avatar={
                                                    <Badge count={item.has_membership ? <EVipIcon/> : 0} offset={[-6, 26]}>
                                                        <Avatar size={30} src={item.image}/>
                                                    </Badge>
                                                }
                                                title={<span>{item.user_name}</span>}
                                                description={item.email === "" ? <span>未填写邮件</span> : item.email}
                                            />
                                        </List.Item>
                                    ))}
                                </List>
                            </div>
                            <div>
                                <Divider style={{marginTop: 12, marginBottom: 14}}/>
                                <Row>
                                    <Col span={20}>
                                        <div>分享协作链接</div>
                                        <div style={{marginTop: 6, fontSize: 13, opacity: 0.9}}>
                                        <span>
                                            <Dropdown
                                                overlayClassName={"share-edit-dropdown"}
                                                overlayStyle={{width: 120}}
                                                menu={{
                                                    items: [{
                                                        label: '可查看权限',
                                                        key: '0',
                                                        icon: <CheckOutlined
                                                            className={sharePersonalFolder.authority === "0" ? "" : "hidden-opacity"}/>
                                                    }, {
                                                        label: '可编辑权限',
                                                        key: '1',
                                                        icon: <CheckOutlined
                                                            className={sharePersonalFolder.authority === "1" ? "" : "hidden-opacity"}/>
                                                    }],
                                                    onClick: (key) => this.handleShareAuthorityUpdate(key, sharePersonalFolder)
                                                }}
                                                trigger={sharePersonalFolder.has_owner ? ['click'] : []}>
                                                <Space>
                                                    {sharePersonalFolder.authority === "0" ? "可查看权限" : "可编辑权限"}
                                                    <CaretDownOutlined
                                                        className={sharePersonalFolder.has_owner ? "" : "dropdown-fade"}
                                                    />
                                                </Space>
                                            </Dropdown>
                                        </span>
                                            <span style={{marginLeft: 20}}>
                                            <Dropdown
                                                overlayClassName={"share-edit-dropdown"}
                                                overlayStyle={{
                                                    width: 120
                                                }}
                                                menu={{
                                                    items: [{
                                                        label: '7天有效',
                                                        key: '7',
                                                        icon: <CheckOutlined
                                                            className={sharePersonalFolder.expired_day === 7 ? "" : "hidden-opacity"}/>
                                                    }, {
                                                        label: '30天有效',
                                                        key: '30',
                                                        icon: <CheckOutlined
                                                            className={sharePersonalFolder.expired_day === 30 ? "" : "hidden-opacity"}/>
                                                    }, {
                                                        label: '永久有效',
                                                        key: '-1',
                                                        icon: <CheckOutlined
                                                            className={sharePersonalFolder.expired_day === -1 ? "" : "hidden-opacity"}/>
                                                    }],
                                                    onClick: (key) => this.handleShareValidUpdate(key, sharePersonalFolder)
                                                }}
                                                trigger={sharePersonalFolder.has_owner ? ['click'] : []}>
                                                <Space>
                                                    {sharePersonalFolder.expired_day === -1 ? "永久有效" : sharePersonalFolder.expired_day + "天有效"}
                                                    <CaretDownOutlined className={sharePersonalFolder.has_owner ? "" : "dropdown-fade"}/>
                                                </Space>
                                            </Dropdown>
                                        </span>
                                        </div>
                                    </Col>
                                    <Col span={4}>
                                        <Button
                                            style={{marginLeft: 10, borderRadius: 4}}
                                            onClick={() => this.handleCopyShareURL(sharePersonalFolder.share_link)}
                                            type="primary">复制链接</Button>
                                    </Col>
                                </Row>
                            </div>
                        </div>
                    </div>
                </Modal>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(SharePage);