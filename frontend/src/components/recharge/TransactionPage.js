import React, {Component} from 'react';
import {Empty, Modal, Table} from 'antd';
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as OrderAction from "../../actions/OrderAction";

const columns = [
    {
        title: '订单号',
        dataIndex: 'order_number',
        key: 'order_number',
        align: 'center'
    },
    {
        title: '交易时间',
        dataIndex: 'created_at',
        key: 'created_at',
        align: 'center'
    },
    {
        title: '交易金额',
        dataIndex: 'payment_amount',
        key: 'payment_amount',
        align: 'center'
    },
    {
        title: '支付方式',
        dataIndex: 'payment_type',
        key: 'payment_type',
        align: 'center',
        render: (text) => {
            return text === "wechat" ? "微信支付" : ""
        }
    },
    {
        title: '购买内容',
        dataIndex: 'payment_content',
        key: 'payment_content',
        align: 'center'
    },
    {
        title: '有效期',
        dataIndex: 'membership_expired',
        key: 'membership_expired',
        align: 'center'
    }
]

class TransactionPage extends Component {

    state = {
        hasOpenTransaction: false
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleOpenTransaction = () => {
        this.props.actions.GetOrderTransactions(() => {
            this.setState({hasOpenTransaction: true})
        })
    }

    handleCloseTransaction = () => {
        this.setState({hasOpenTransaction: false})
    }

    render() {
        const {transactions} = this.props.state
        return (
            <Modal title={null}
                   open={this.state.hasOpenTransaction}
                   onCancel={this.handleCloseTransaction}
                   forceRender={true}
                   footer={null}
                   width={820}
                   className={"invite-modal"}>
                    <span className={"ant-modal-confirm-title"}>
                        <h3>交易记录</h3>
                    </span>
                <div className={"ant-modal-confirm-content"} style={{height: 410}}>
                    <div style={{marginTop: 30}}>
                        <div>
                            <Table bordered
                                   locale={{emptyText: <Empty description={"交易记录为空"}
                                                              image={Empty.PRESENTED_IMAGE_SIMPLE}/>}}
                                   size={"small"}
                                   pagination={false}
                                   columns={columns}
                                   className={"order-transaction-table"}
                                   rowClassName={"order-transaction-table-row"}
                                   dataSource={transactions.transactions}/>
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
)(TransactionPage);