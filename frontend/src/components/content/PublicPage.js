import React, {Component} from 'react';
import {Button, Form, Input, Modal, Select, Space} from "antd";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CommunityAction from "../../actions/CommunityAction";
import '@wangeditor/editor/dist/css/style.css'
import {Editor, Toolbar} from '@wangeditor/editor-for-react'
import {GetPublicCommunityCategory} from "../../actions/CommunityAction";

const toolbarConfig = {
    excludeKeys: [ // 剔除
        "group-more-style",
        "|",
        "bgColor",
        "fontSize",
        "fontFamily",
        "lineHeight",
        "emotion",
        "group-image",
        "group-video",
        "insertTable",
        "codeBlock",
        "undo",
        "redo",
        "fullScreen"
    ]
}

const editorConfig = {
    placeholder: '请输入内容...',
}

class PublicPage extends Component {

    state = {
        hasOpenPublic: false,
        outerHtml: null,
        editor: null,
        html: "",
        folderUID: ''
    }

    constructor(props) {
        super(props);
        this.basePublicRef = React.createRef();
    }

    componentDidMount() {
        this.props.bindRef(this)
        this.props.actions.GetPublicCommunityCategory()
    }

    handleOpenPublic = (folderUID) => {
        console.log("打开")
        this.setState({
            hasOpenPublic: true,
            folderUID: folderUID
        })
    }

    handleSetOuterHtml = (outerHtml) => {
        console.log("设置2")
        this.setState({outerHtml: outerHtml})
    }

    handleClosePublic = () => {
        this.basePublicRef.current.resetFields()
        this.setState({html: ''})
        this.setState({hasOpenPublic: false})
    }

    handleSetEditor = (editor) => {
        this.setState({editor: editor})
    }

    handleSetHtml = (html) => {
        this.setState({html: html})
    }

    handleSavePublic = () => {
        const data = this.basePublicRef.current.getFieldsValue()
        const body = {
            ...data,
            "description": this.state.html,
            "image": this.state.outerHtml,
            "folder_uid": this.state.folderUID,
        }
        this.props.actions.CreatePublicCommunity(body, () => {
            this.handleClosePublic()
        })
    }

    render() {
        const {communityCategories} = this.props.state

        return (
            <div>
                <Modal title={null}
                    open={this.state.hasOpenPublic}
                       onCancel={this.handleClosePublic}
                       forceRender={true}
                       footer={null}
                       width={720}
                       className={"public-resource"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>公开到资源社区</h3>
                    </span>
                    <div className={"ant-modal-confirm-content"} style={{height: 450}}>
                        <div style={{marginTop: 30}}>
                            <Form name={"public-base-info"}
                                  ref={this.basePublicRef}
                                  wrapperCol={{offset: 1}}>
                                <Form.Item label={<span>标题</span>} name="title">
                                    <Input placeholder="请输入发布标题"/>
                                </Form.Item>
                                <Form.Item label={<span>详情</span>}>
                                    <div style={{border: '1px solid #ccc', zIndex: 100}}>
                                        <Toolbar
                                            editor={this.state.editor}
                                            defaultConfig={toolbarConfig}
                                            mode="default"
                                            style={{borderBottom: '1px solid #ccc'}}
                                        />
                                        <Editor
                                            defaultConfig={editorConfig}
                                            value={this.state.html}
                                            onCreated={this.handleSetEditor}
                                            onChange={editor => this.handleSetHtml(editor.getHtml())}
                                            mode="default"
                                            style={{height: '180px', overflowY: 'hidden'}}
                                        />
                                    </div>
                                </Form.Item>
                                <Form.Item label={<span>分类</span>} wrapperCol={{offset: 1, span: 8}}
                                           name="category_uid">
                                    <Select placeholder="请选择推荐分类" allowClear>
                                        {
                                            communityCategories.categories.map(c => (
                                                <Select.Option value={c.uid}>{c.name}</Select.Option>
                                            ))
                                        }
                                    </Select>
                                </Form.Item>
                                <Form.Item label={<span>标签</span>} name="tags">
                                    <Select mode="tags" open={false} placeholder="请输入自定义标签, 回车确认"/>
                                </Form.Item>
                                <Form.Item style={{textAlign: "right"}}>
                                    <Space>
                                        <Button className={"add-resource-submit-btn"}
                                                onClick={this.handleGoodStationCancel}>
                                            取消
                                        </Button>
                                        <Button type="primary"
                                                className={"add-resource-submit-btn"}
                                                onClick={this.handleSavePublic}>
                                            保存
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
    state => ({state: state.dataManage.community}),
    dispatch => ({
        actions: bindActionCreators(CommunityAction, dispatch)
    })
)(PublicPage);