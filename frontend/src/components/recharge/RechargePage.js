import React, {Component} from 'react';
import {Card, Image, Modal, PageHeader} from 'antd';
import {EWechatPayIcon} from "../base/EIcon";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as OrderAction from "../../actions/OrderAction";
import {isEmpty, isNotEmpty, OrderCancelStateCode, OrderUnpaidStateCode} from "../../actions/Base";
import PaySuccessPage from "./PaySuccessPage";
import FeedbackPage from "../feedback/FeedbackPage";

class RechargePage extends Component {

    state = {
        hasOpenRecharge: false,
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleRechargeOpen = () => {
        this.setState({hasOpenRecharge: true})
    }

    handleRechargeCancel = () => {
        this.setState({hasOpenRecharge: false})
        this.clearOrderTimer()
    }

    startOrderTimer = (order) => {
        this.orderTimer = setInterval(() => {
            this.props.actions.GetOrderStatus(order.order_number, () => {
                this.setState({hasOpenRecharge: false})
                this.successModalRef.handleRechargeSuccessOpen(order.payment_amount)
                this.clearOrderTimer()
            }, this.clearOrderTimer)
        }, 2000)
    }

    clearOrderTimer = () => {
        clearInterval(this.orderTimer)
        this.orderTimer = null
        this.props.actions.ClearOrder()
    }

    handlePrevBack = () => {
        this.setState({hasOpenRecharge: false})
        this.props.handleUpgradeOpen()
        this.clearOrderTimer()
    }


    componentDidUpdate(prevProps) {
        const {order, orderStatus} = this.props.state
        if (this.state.hasOpenRecharge &&
            isNotEmpty(order.order_number) &&
            isEmpty(this.orderTimer) &&
            (isEmpty(orderStatus.status) || orderStatus.status === OrderUnpaidStateCode)
        ) {
            this.startOrderTimer(order)
        }
    }

    handleRefreshOrder = () => {

    }

    handleOpenFeedback = (orderNumber) => {
        this.feedbackModalRef.handleOpenFeedback(orderNumber)
    }


    render() {
        const {order, orderStatus} = this.props.state
        return (
            <div>
                <Modal title={null}
                       open={this.state.hasOpenRecharge}
                       onCancel={this.handleRechargeCancel}
                       forceRender={true}
                       footer={null}
                       width={820}
                       height={500}
                       className={"share-folder"}>
                <span className={"ant-modal-confirm-title"}>
                    <div className={"return-previous-step"}>
                        <PageHeader onBack={this.handlePrevBack} title="返回上一步"/>
                    </div>
                </span>
                    <div className={"ant-modal-confirm-content"} style={{height: 410}}>
                        <div style={{marginTop: 30}}>
                            <div className={"recharge-image-container"}>
                                <h3 style={{fontSize: 19}}>支付金额：¥{order.payment_amount}</h3>
                                <div className="upgrade-pay">
                                    <EWechatPayIcon/>
                                    <span style={{fontSize: 12, opacity: 0.8}}>&nbsp;使用微信扫码支付</span>
                                </div>
                                <div style={{marginTop: 10}}>
                                    <Card className={"recharge-qrcode-image"}>
                                        <Image
                                            width={120}
                                            src={order.payment_link}
                                            preview={false}
                                        />
                                        {
                                            orderStatus.status === OrderCancelStateCode ? (
                                                <div className={"recharge-qrcode-mark"}>
                                                    二维码过期，<a onClick={this.handlePrevBack}>返回</a>
                                                </div>
                                            ) : (
                                                <></>
                                            )
                                        }
                                    </Card>
                                </div>
                                <div className="recharge-problem-desc">
                                    <div>支付后，如果会员未开通</div>
                                    <div>请进行<a onClick={() => this.handleOpenFeedback(order.order_number)}>问题反馈</a>，我们将快速处理</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </Modal>

                <PaySuccessPage bindRef={(ref) => this.successModalRef = ref}/>

                <FeedbackPage bindRef={(ref) => this.feedbackModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.order}), dispatch => ({
        actions: bindActionCreators(OrderAction, dispatch)
    }))(RechargePage);