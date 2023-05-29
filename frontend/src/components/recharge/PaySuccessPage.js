import React, {Component} from 'react';
import {Button, Modal, Space} from 'antd';
import {EPaySuccessIcon} from "../base/EIcon";
import TransactionPage from "./TransactionPage";

class PaySuccessPage extends Component {

    state = {
        hasOpenRechargeSuccess: false,
        paymentAmount: ''
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleRechargeSuccessOpen = (paymentAmount) => {
        this.setState({
            hasOpenRechargeSuccess: true,
            paymentAmount: paymentAmount
        })
    }

    handleRechargeSuccessCancel = () => {
        this.setState({hasOpenRechargeSuccess: false})
    }

    handleOpenTransaction = () => {
        this.setState({hasOpenRechargeSuccess: false})
        this.transactionModalRef.handleOpenTransaction()
    };

    render() {
        return (
            <div>
                <Modal title={null}
                       open={this.state.hasOpenRechargeSuccess}
                       onCancel={this.handleRechargeSuccessCancel}
                       forceRender={true}
                       footer={null}
                       width={820}
                       height={500}
                       className={"share-folder"}>
                <span className={"ant-modal-confirm-title"}>
                </span>
                    <div className={"ant-modal-confirm-content"} style={{height: 410}}>
                        <div style={{marginTop: 30}}>
                            <div className={"recharge-success-container"}>
                                <div style={{marginTop: 10}}>
                                    <div className={"recharge-pay-success"}>
                                        <EPaySuccessIcon/>
                                        <span style={{marginLeft: 15, opacity: 0.8}}>支付成功</span>
                                    </div>
                                </div>
                                <div className="recharge-success-desc">
                                    <p>支付金额: ¥{this.state.paymentAmount}</p>
                                </div>
                                <div style={{marginTop: 50}}>
                                    <Space size={30}>
                                        <Button onClick={this.handleOpenTransaction}>查看订单</Button>
                                        <Button type="primary" href="/">返回首页</Button>
                                    </Space>
                                </div>
                            </div>
                        </div>
                    </div>
                </Modal>

                <TransactionPage bindRef={(ref) => this.transactionModalRef = ref}/>
            </div>
        );
    }
}

export default PaySuccessPage;