import React, {Component} from 'react';
import {Button, Form, Input, Modal, Select} from "antd";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as OrderAction from "../../actions/OrderAction";
import './style.css'

class FeedbackPage extends Component {

    state = {
        hasOpenFeedback: false,
        orderNumber: ''
    }

    constructor(props) {
        super(props);
        this.feedbackRef = React.createRef();
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleOpenFeedback = (orderNumber) => {
        this.setState({
            hasOpenFeedback: true,
            orderNumber: orderNumber
        })
    }

    handleCloseFeedback = () => {
        this.setState({hasOpenFeedback: false, orderNumber: ''})
    }

    handleFeedback = () => {
        const data = this.feedbackRef.current.getFieldsValue()
        this.props.actions.CreateFeedback({
            "order_number": this.state.orderNumber,
            ...data,
        }, () => {
            this.feedbackRef.current.resetFields()
            this.setState({hasOpenFeedback: false, orderNumber: ''})
        })
    }

    render() {
        return (
            <Modal title={null}
                   open={this.state.hasOpenFeedback}
                   onCancel={this.handleCloseFeedback}
                   forceRender={true}
                   footer={null}
                   width={470}
                   className={"invite-modal"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>问题反馈</h3>
                    </span>
                <div className={"ant-modal-confirm-content"} style={{height: 230}}>
                    <div style={{marginTop: 30}}>
                        <div className={"add-feedback"}>
                            <Form style={{marginTop: 20}}
                                  labelCol={{span: 3}}
                                  wrapperCol={{span: 20}}
                                  ref={this.feedbackRef}
                            >
                                <Form.Item label={<span>类型</span>}
                                           rules={[{required: true, message: '请选择问题类型'}]}
                                           name="category">
                                    <Select placeholder="请选择问题类型"
                                            allowClear>
                                        <Select.Option value="支付问题">支付问题</Select.Option>
                                        <Select.Option value="Bug问题">Bug问题</Select.Option>
                                        <Select.Option value="系统优化">系统优化</Select.Option>
                                        <Select.Option value="其他问题">其他问题</Select.Option>
                                    </Select>
                                </Form.Item>
                                <Form.Item label={<span>描述</span>} name="description">
                                    <Input.TextArea rows={5} placeholder="请输入问题描述"/>
                                </Form.Item>
                                <Form.Item wrapperCol={{offset: 3, span: 8}}>
                                    <Button type="primary" onClick={this.handleFeedback}>
                                        提交
                                    </Button>
                                </Form.Item>
                            </Form>
                        </div>
                    </div>
                </div>
            </Modal>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.order}),
    dispatch => ({
        actions: bindActionCreators(OrderAction, dispatch)
    })
)(FeedbackPage);