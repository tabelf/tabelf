import React, {Component} from 'react';
import {Avatar, Badge, Col, Form, Input, message, Modal, Row, Upload, Typography} from "antd";
import {EMaxVipIcon} from "../base/EIcon";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {apiUrl, checkUserInfo, isEmpty} from "../../actions/Base";
import './style.css'

const {Paragraph} = Typography;

class UserInfoPage extends Component {
    state = {
        hasUserBaseInfo: false,
        imageURL: null,
    }

    constructor(props) {
        super(props);
        this.baseUserRef = React.createRef();
    }

    componentDidMount() {
        this.props.actions.GetPersonalAccountInfo()
        this.props.bindRef(this)
    }

    handleUserBaseOpen = () => {
        this.setState({hasUserBaseInfo: true})
        this.props.actions.GetPersonalAccountInfo((data) => {
            this.baseUserRef.current.setFieldsValue({
                "industry": data.industry,
                "email": data.email,
                "description": data.description,
            })
            this.setState({imageURL: data.image})
        })
    }

    handleUserBaseCancel = () => {
        this.setState({hasUserBaseInfo: false})
    }


    handleEditableStr = (value) => {
        const {accountInfo} = this.props.state
        if (!isEmpty(value) && accountInfo.user_name !== value) {
            this.props.actions.UpdatePersonalAccountInfo({"user_name": value})
        }
    }

    handleIndustryBlur = ({target: {value}}) => {
        const {accountInfo} = this.props.state
        if (accountInfo.industry !== value) {
            if (isEmpty(value)) {
                value = "null"
            }
            this.props.actions.UpdatePersonalAccountInfo({"industry": value})
        }
    }

    handleEmailBlur = ({target: {value}}) => {
        const {accountInfo} = this.props.state
        if (!isEmpty(value) && accountInfo.email !== value) {
            this.props.actions.UpdatePersonalAccountInfo({"email": value})
        }
    }

    handleDescriptionBlur = ({target: {value}}) => {
        const {accountInfo} = this.props.state
        if (accountInfo.description !== value) {
            if (isEmpty(value)) {
                value = "null"
            }
            this.props.actions.UpdatePersonalAccountInfo({"description": value})
        }
    }

    handleUserImageChange = (info) => {
        if (info.file.status === 'error') {
            message.error("上传失败")
            return;
        }
        if (info.file.status === 'done') {
            message.success("上传成功")
            this.setState({imageURL: info.file.response.image_url});
            this.props.actions.GetPersonalAccountInfo()
        }
    };

    beforeUpload = (file) => {
        const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/webp';
        if (!isJpgOrPng) {
            message.error('上传文件只能包含 JPG/PNG 格式!');
        }
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isLt2M) {
            message.error('文件大小不能超过 2MB!');
        }
        return isJpgOrPng && isLt2M;
    };

    render() {
        const userinfo = checkUserInfo()
        const uploadProps = {
            name: "file",
            method: "PUT",
            showUploadList: false,
            action: `${apiUrl}/customer/${userinfo.user_uid}/personal/upload`,
            headers: {
                Authorization: userinfo.token,
                'X-Requested-With': null,
            },
            listType: "picture-card",
        };

        const {accountInfo} = this.props.state

        return (
            <div>
                <Modal title={null}
                       open={this.state.hasUserBaseInfo}
                       onCancel={this.handleUserBaseCancel}
                       footer={null}
                       forceRender={true}
                       width={720}
                       height={172}
                       className={"share-folder"}>
                <span className={"ant-modal-confirm-title"}>
                    <h3>用户基本消息</h3>
                    {
                        accountInfo.has_entire ? (<></>) : (
                            <div style={{fontSize: 13, opacity: 0.4}}>完善基本信息，可获取5个链接数</div>
                        )
                    }
                </span>
                    <div className={"ant-modal-confirm-content"}>
                        <div className={"user-base-info-form"} style={{marginTop: 20}}>
                            <Row>
                                <Col span={18}>
                                    <Form name={"user-base-info"}
                                          ref={this.baseUserRef}
                                          wrapperCol={{offset: 1, span: 15}}
                                          labelCol={{span: 4}}>
                                        <Form.Item label={<span>昵称</span>}>
                                            <Paragraph editable={{onChange: this.handleEditableStr}}>
                                                {accountInfo.user_name}
                                            </Paragraph>
                                        </Form.Item>
                                        <Form.Item label={<span>行业</span>} name="industry">
                                            <Input placeholder="请输入行业" onBlur={this.handleIndustryBlur}/>
                                        </Form.Item>
                                        <Form.Item label={<span>邮件</span>} name="email">
                                            <Input placeholder="请输入邮件" onBlur={this.handleEmailBlur}/>
                                        </Form.Item>
                                        <Form.Item label={<span>个人简介</span>} name="description">
                                            <Input.TextArea
                                                style={{height: 120, resize: 'none'}}
                                                placeholder="请输入简介"
                                                onBlur={this.handleDescriptionBlur}
                                            />
                                        </Form.Item>
                                    </Form>
                                </Col>
                                <Col span={6}>
                                    <Upload {...uploadProps}
                                            onChange={this.handleUserImageChange}
                                            beforeUpload={this.beforeUpload}>
                                        <Badge count={accountInfo.has_membership ? <EMaxVipIcon/> : 0}
                                               offset={[-26, 94]}>
                                            <Avatar src={this.state.imageURL} size={104}/>
                                        </Badge>
                                    </Upload>
                                </Col>
                            </Row>
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
)(UserInfoPage);