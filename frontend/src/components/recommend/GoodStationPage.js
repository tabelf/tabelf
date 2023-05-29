import React, {Component} from 'react';
import {Avatar, Button, Form, Input, message, Modal, Select, Space, Upload} from "antd";
import {apiUrl, getUserInfo, isEmpty, isNotEmpty} from "../../actions/Base";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {LoadingOutlined, PlusOutlined} from "@ant-design/icons";

class GoodStationPage extends Component {

    state = {
        hasGoodStation: false,
        imageURL: null,
        loading: false,
        offset: 1,
        limit: 20,
        callback: null,
        hasAudit: false
    }

    constructor(props) {
        super(props);
        this.baseStationRef = React.createRef();
    }

    componentDidMount() {
        this.props.bindRef(this)
        this.props.actions.GetGoodStationCategory()
    }

    handleGoodStationOpen = () => {
        this.setState({hasGoodStation: true})
    }

    handleGoodStationCancel = () => {
        this.baseStationRef.current.resetFields()
        this.setState({hasGoodStation: false, hasAudit: false})
    }

    handleUserImageChange = (info) => {
        if (info.file.status === 'uploading') {
            this.setState({loading: true});
            return;
        }
        if (info.file.status === 'error') {
            this.setState({loading: false});
            message.error("上传失败")
            return;
        }
        if (info.file.status === 'done') {
            this.setState({loading: false});
            message.success("上传成功")
            this.setState({imageURL: info.file.response.image_url});
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

    handleSaveStation = () => {
        const data = this.baseStationRef.current.getFieldsValue()
        let image = ""
        if (this.state.imageURL !== null) {
            image = this.state.imageURL
        }
        if (isNotEmpty(data.station_uid)) {
            this.props.actions.UpdateGoodStationRecommend({...data, image}, () => {
                if (this.state.hasAudit) {
                    this.props.actions.GetGoodStationAudit()
                } else {
                    this.props.actions.GetGoodStationRecommend("", "0", 0, this.state.limit)
                }
                this.handleGoodStationCancel()
            })
        } else {
            this.props.actions.AddGoodStationRecommend({...data, image}, () => {
                this.props.actions.GetGoodStationRecommend("", "0", 0, this.state.limit)
                this.handleGoodStationCancel()
            })
        }
    }

    handleShowGoodStation = (data, hasAudit) => {
        this.handleGoodStationOpen()
        this.baseStationRef.current.setFieldsValue(data)
        this.setState({
            imageURL: data.image,
            hasAudit: hasAudit
        })
    }

    render() {
        const userinfo = getUserInfo()
        if (isEmpty(userinfo)) {
            return <></>
        }

        const {stationCategories} = this.props.state

        const uploadButton = (
            <div>
                {this.state.loading ? <LoadingOutlined/> : <PlusOutlined/>}
                <div style={{marginTop: 8}}>上传图片</div>
            </div>
        );

        const uploadProps = {
            name: "file",
            method: "PUT",
            showUploadList: false,
            action: `${apiUrl}/customer/${userinfo.user_uid}/station/image`,
            headers: {
                Authorization: userinfo.token,
                'X-Requested-With': null,
            },
            listType: "picture-card"
        };
        return (
            <div>
                <Modal title={null}
                       open={this.state.hasGoodStation}
                       onCancel={this.handleGoodStationCancel}
                       footer={null}
                       forceRender={true}
                       width={670}
                       height={172}
                       className={"share-folder"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>好站推荐内容</h3>
                    </span>
                    <div className={"ant-modal-confirm-content"}>
                        <div className={"station-base-info-form"} style={{marginTop: 20}}>
                            <Form name={"user-base-info"}
                                  ref={this.baseStationRef}
                                  wrapperCol={{offset: 1, span: 13}}
                                  labelCol={{offset: 1, span: 4}}>
                                <Form.Item label={<span>分类</span>}
                                           rules={[{required: true}]}
                                           name="category_uid">
                                    <Select placeholder="请选择推荐分类" allowClear>
                                        {
                                            stationCategories.categories.map(c => (
                                                <Select.Option value={c.uid}>{c.name}</Select.Option>
                                            ))
                                        }
                                    </Select>
                                </Form.Item>
                                <Form.Item label={<span>链接</span>}
                                           rules={[{required: true}]}
                                           name="link">
                                    <Input placeholder="请输入链接"/>
                                </Form.Item>
                                <Form.Item label={<span>名称</span>}
                                           rules={[{required: true}]}
                                           name="site_name">
                                    <Input placeholder="请输入名称"/>
                                </Form.Item>
                                <Form.Item label={<span>标题</span>} name="title">
                                    <Input placeholder="请输入自定义标题"/>
                                </Form.Item>
                                <Form.Item label={<span>详情</span>} name="description">
                                    <Input.TextArea
                                        style={{height: 120, resize: 'none'}}
                                        placeholder="请输入自定义详情"
                                    />
                                </Form.Item>
                                <Form.Item label={<span>大图</span>}>
                                    <Upload {...uploadProps}
                                            onChange={this.handleUserImageChange}
                                            beforeUpload={this.beforeUpload}>
                                        {this.state.imageURL ? <Avatar shape="square" src={this.state.imageURL} size={72}/> : uploadButton}
                                    </Upload>
                                </Form.Item>
                                <Form.Item label={<span>标签</span>} name="tags">
                                    <Select mode="tags" open={false} placeholder="请输入自定义标签, 回车确认"/>
                                </Form.Item>
                                <Form.Item noStyle name="station_uid"/>
                                <Form.Item wrapperCol={{offset: 6}}>
                                    <Space>
                                        <Button type="primary"
                                                className={"add-resource-submit-btn"}
                                                onClick={this.handleSaveStation}>
                                            保存
                                        </Button>
                                        <Button className={"add-resource-submit-btn"}
                                                onClick={this.handleGoodStationCancel}>
                                            取消
                                        </Button>
                                    </Space>
                                </Form.Item>
                            </Form>
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
)(GoodStationPage);