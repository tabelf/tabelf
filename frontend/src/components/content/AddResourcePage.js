import React, {Component} from 'react';
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {Button, Col, Form, Input, Menu, message, Modal, Row, Space, Upload} from "antd";
import {ELinkIcon, ELocalIcon, EMaxLocalIcon} from "../base/EIcon";
import {apiUrl, checkUserInfo, isEmpty} from "../../actions/Base";
import UpgradePage from "../recharge/UpgradePage";

class AddResourcePage extends Component {

    state = {
        showAddResourceType: 'url',
        showAddResourceModal: false,
        resourceAction: '',
        uploading: false,
        folderUID: '',
        workspaceUID: '',
        shareUID: '',
        hasColl: false
    }

    constructor(props) {
        super(props);
        this.webLinkForm = React.createRef();
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleChangeAddResource = ({key}) => {
        this.setState({showAddResourceType: key})
    }

    handleOpenAddResource = (folderUID, workspaceUID, shareUID, hasColl) => {
        const userinfo = checkUserInfo()
        const action = `${apiUrl}/customer/${userinfo.user_uid}/personal/folder/${folderUID}/workspace/${workspaceUID}/local`
        this.setState({
            showAddResourceModal: true,
            resourceAction: action,
            folderUID: folderUID,
            workspaceUID: workspaceUID,
            shareUID: shareUID,
            hasColl: hasColl
        })
    }

    handleCloseAddResource = () => {
        this.setState({
            showAddResourceModal: false,
            showAddResourceType: 'url'
        })
    }

    handleFileOnChange = (info) => {
        if (info.file.status === 'uploading') {
            this.setState({uploading: true})
            return;
        }
        if (info.file.status === 'done') {
            this.setState({
                uploading: false,
                showAddResourceModal: false
            })
            this.props.actions.AddLocalFileSuccessCallback(this.state.folderUID)
            message.success(`文件上传成功`);
            this.handleUpdateCommunityImage(this.state.folderUID)
        } else if (info.file.status === 'error') {
            this.setState({uploading: false})
            message.error(`文件上传失败, 请重新上传`);
        }
    }

    beforeUpload = (file) => {
        const isLt10M = file.size / 1024 / 1024 < 10;
        if (!isLt10M) {
            message.error('文件大小不能超过 10MB!');
        }
        return isLt10M;
    };

    handleAddResource = () => {
        const data = this.webLinkForm.current.getFieldsValue()
        if (isEmpty(data.url)) {
            message.error("URL不能为空")
            return;
        }
        if (this.state.workspaceUID === '') {
            message.error("未指定工作区域")
            return;
        }
        if (this.state.hasColl) {
            this.props.actions.AddSharePersonalWebLink(
                this.state.folderUID,
                this.state.shareUID,
                this.state.workspaceUID,
                data,
                () => {
                    this.handleCloseAddResource()
                    this.webLinkForm.current.resetFields()
                },
            );
        } else {
            this.props.actions.AddPersonalWebLink(
                this.state.folderUID,
                this.state.workspaceUID,
                data,
                () => {
                    this.handleCloseAddResource()
                    this.webLinkForm.current.resetFields()
                },
                this.handleUpgradeOpen
            );
        }
        this.handleUpdateCommunityImage(this.state.folderUID)
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    handleUpdateCommunityImage = (folderUID) => {
        this.props.handleUpdateCommunityImage(folderUID)
    }

    render() {
        const userinfo = checkUserInfo()
        const uploadProps = {
            name: "file",
            method: "POST",
            showUploadList: false,
            headers: {
                Authorization: userinfo.token,
                'X-Requested-With': null,
            }
        };

        return (
            <div>
                <Modal title={null}
                       open={this.state.showAddResourceModal}
                       footer={null}
                       forceRender={true}
                       onCancel={this.handleCloseAddResource}
                       width={510}
                       className={"add-resource"}>
                    <div className={"ant-modal-confirm-content"}>
                        <Row>
                            <Col span={7} className={"add-resource-sidebar"}>
                                <div className={"add-resource-modal-title"}>
                                    <span>
                                        <h2>添加资源</h2>
                                    </span>
                                </div>
                                <div>
                                    <Menu mode="vertical"
                                          selectedKeys={[this.state.showAddResourceType]}>
                                        <Menu.Item key="url" icon={<ELinkIcon/>} onClick={this.handleChangeAddResource}>
                                            URL链接
                                        </Menu.Item>
                                        <Menu.Item key="local" icon={<ELocalIcon/>}
                                                   onClick={this.handleChangeAddResource}>
                                            本地文件
                                        </Menu.Item>
                                    </Menu>
                                </div>
                            </Col>
                            <Col span={17}>
                                {
                                    this.state.showAddResourceType === 'url' ? (
                                        <Form labelCol={{offset: 2}}
                                              wrapperCol={{offset: 2, span: 20}}
                                              layout="vertical"
                                              className={"add-resource-form"}
                                              ref={this.webLinkForm}
                                              name="weblink-form">
                                            <Form.Item label="URL" name="url">
                                                <Input placeholder="https://www.baidu.com..."/>
                                            </Form.Item>
                                            <Form.Item label="标题" name="title">
                                                <Input placeholder="请输入标题"/>
                                            </Form.Item>
                                            <Form.Item wrapperCol={{offset: 2}}>
                                                <Space size={"large"} style={{marginTop: 18}}>
                                                    <Button type="primary" className={"add-resource-submit-btn"}
                                                            onClick={this.handleAddResource}>
                                                        保存
                                                    </Button>
                                                    <Button className={"add-resource-submit-btn"}
                                                            onClick={this.handleCloseAddResource}>
                                                        取消
                                                    </Button>
                                                </Space>
                                            </Form.Item>
                                        </Form>
                                    ) : (
                                        <div className={"add-resource-upload"}>
                                            <div>
                                                <EMaxLocalIcon/>
                                            </div>
                                            <div>
                                                <h3>从你的电脑上传文件</h3>
                                            </div>
                                            <div>
                                                <Upload {...uploadProps}
                                                        action={this.state.resourceAction}
                                                        onChange={this.handleFileOnChange}
                                                        beforeUpload={this.beforeUpload}>
                                                    <Button
                                                        shape="round"
                                                        type="primary"
                                                        loading={this.state.uploading}
                                                        disabled={this.state.uploading}>
                                                        {this.state.uploading ? "上传中..." : "选择文件"}
                                                    </Button>
                                                </Upload>
                                            </div>
                                        </div>
                                    )
                                }
                            </Col>
                        </Row>
                    </div>
                </Modal>

                <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(AddResourcePage);