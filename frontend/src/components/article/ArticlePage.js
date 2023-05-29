import React, {Component} from 'react';
import {Avatar, Button, Form, Input, message, Modal, Select, Space, Upload} from "antd";
import {apiUrl, getUserInfo, isEmpty, isNotEmpty} from "../../actions/Base";
import {LoadingOutlined, PlusOutlined} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";

class ArticlePage extends Component {

    state = {
        hasGoodArticle: false,
        imageURL: null,
        loading: false,
        offset: 1,
        limit: 20,
    }

    constructor(props) {
        super(props);
        this.baseArticleRef = React.createRef();
    }

    componentDidMount() {
        this.props.bindRef(this)
        this.props.actions.GetGoodArticleCategory()
    }

    handleGoodArticleOpen = () => {
        this.setState({hasGoodArticle: true})
    }

    handleGoodArticleCancel = () => {
        this.baseArticleRef.current.resetFields()
        this.setState({hasGoodArticle: false})
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

    handleSaveArticle = () => {
        const data = this.baseArticleRef.current.getFieldsValue()
        let image = ""
        if (this.state.imageURL !== null) {
            image = this.state.imageURL
        }
        if (isNotEmpty(data.article_uid)) {
            this.props.actions.UpdateGoodArticleRecommend({...data, image}, this.handleGoodArticleCancel)
        } else {
            this.props.actions.AddGoodArticleRecommend({...data, image}, this.handleGoodArticleCancel)
        }
    }

    handleShowGoodArticle = (data) => {
        this.handleGoodArticleOpen()
        this.baseArticleRef.current.setFieldsValue(data)
        this.setState({imageURL: data.image})
    }

    render() {
        const userinfo = getUserInfo()
        if (isEmpty(userinfo)) {
            return <></>
        }

        const {articleCategories} = this.props.state

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
                       open={this.state.hasGoodArticle}
                       onCancel={this.handleGoodArticleCancel}
                       footer={null}
                       forceRender={true}
                       width={670}
                       height={172}
                       className={"share-folder"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>好文推荐</h3>
                    </span>
                    <div className={"ant-modal-confirm-content"}>
                        <div className={"station-base-info-form"} style={{marginTop: 20}}>
                            <Form name={"user-base-info"}
                                  ref={this.baseArticleRef}
                                  wrapperCol={{offset: 1, span: 13}}
                                  labelCol={{offset: 1, span: 4}}>
                                <Form.Item label={<span>分类</span>}
                                           rules={[{required: true, message: "分类不能为空"}]}
                                           name="category_uid">
                                    <Select placeholder="请选择推荐分类" allowClear>
                                        {
                                            articleCategories.categories.map(c => (
                                                <Select.Option value={c.uid}>{c.name}</Select.Option>
                                            ))
                                        }
                                    </Select>
                                </Form.Item>
                                <Form.Item label={<span>链接</span>}
                                           rules={[{required: true, message: "链接不能为空"}]}
                                           name="link">
                                    <Input placeholder="请输入链接"/>
                                </Form.Item>
                                <Form.Item label={<span>来源</span>}
                                           rules={[{required: true, message: "来源不能为空"}]}
                                           name="source">
                                    <Input placeholder="请输入来源"/>
                                </Form.Item>
                                <Form.Item label={<span>标题</span>}
                                           name="title">
                                    <Input placeholder="请输入自定义标题"/>
                                </Form.Item>
                                <Form.Item label={<span>头图</span>}>
                                    <Upload {...uploadProps}
                                            onChange={this.handleUserImageChange}
                                            beforeUpload={this.beforeUpload}>
                                        {this.state.imageURL ? <Avatar shape="square" src={this.state.imageURL} size={72}/> : uploadButton}
                                    </Upload>
                                </Form.Item>
                                <Form.Item noStyle name="article_uid"/>
                                <Form.Item wrapperCol={{offset: 6}}>
                                    <Space>
                                        <Button type="primary"
                                                className={"add-resource-submit-btn"}
                                                htmlType="submit"
                                                onClick={this.handleSaveArticle}>
                                            发布
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
    state => ({state: state.dataManage.article}),
    dispatch => ({
        actions: bindActionCreators(GoodArticleAction, dispatch)
    })
)(ArticlePage);