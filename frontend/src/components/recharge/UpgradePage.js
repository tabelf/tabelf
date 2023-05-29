import React, {Component} from 'react';
import {Button, Card, Col, Modal, Row, Space, Typography} from 'antd';
import './style.css'
import {CheckOutlined} from "@ant-design/icons";
import {EWechatIcon} from "../base/EIcon";
import {withRouter} from "../base/Base";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as OrderAction from "../../actions/OrderAction";
import RechargePage from "./RechargePage";

const {Title, Text} = Typography;

class UpgradePage extends Component {

    state = {
        upgradeActive: 0,
        payMethodActive: 'wechat',
        paymentAmount: '0.00',
        hasOpenUpgrade: false,
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handlePayMethodActive = (key) => {
        this.setState({payMethodActive: key})
    }

    handleUpgradeActive = (key, payment) => {
        this.setState({
            upgradeActive: key,
            paymentAmount: payment
        })
    }

    handlePayOrder = () => {
        const {recharges} = this.props.state
        const recharge = recharges.upgrade_recharges[this.state.upgradeActive]
        this.props.actions.CreateOrder(
            recharge.uid,
            recharge.origin_amount,
            this.state.paymentAmount,
            this.state.payMethodActive
        )
        this.setState({hasOpenUpgrade: false})
        this.rechargeModalRef.handleRechargeOpen()
    }

    handleUpgradeOpen = () => {
        this.props.actions.GetUpgradeRecharges((data) => {
            this.setState({
                hasOpenUpgrade: true,
                paymentAmount: data.amount,
            })
        })
    };

    handleUpgradeCancel = () => {
        this.setState({hasOpenUpgrade: false})
    };

    render() {
        const {recharges} = this.props.state
        return (
            <div>
                <Modal title={null}
                       open={this.state.hasOpenUpgrade}
                       onCancel={this.handleUpgradeCancel}
                       forceRender={true}
                       footer={null}
                       width={820}
                       height={500}
                       className={"share-folder"}>
                <span className={"ant-modal-confirm-title"}>
                    <h3>升级享全部会员权益</h3>
                </span>
                    <div className={"ant-modal-confirm-content"} style={{height: 410}}>
                        <div style={{marginTop: 30}}>
                            <div>
                                <Row>
                                    <Col span={4}>
                                        <div style={{marginTop: 6}}>请选择购买方案</div>
                                    </Col>
                                    <Col offset={15} span={4}>
                                        <Button block className="upgrade-pay-btn"
                                                type="primary"
                                                onClick={this.handlePayOrder}
                                        >
                                            <span style={{fontWeight: 500}}>立即支付¥</span>
                                            <span style={{fontSize: 16, fontWeight: 500}}>
                                            {this.state.paymentAmount}
                                        </span>
                                        </Button>
                                    </Col>
                                </Row>
                            </div>
                            <div style={{marginTop: 25}}>
                                <Row>
                                    {
                                        recharges.upgrade_recharges.map((r, idx) => (
                                            <Col span={6}>
                                                <Card
                                                    className={this.state.upgradeActive === idx ? "upgrade-user-declare upgrade-user-declare-active" : "upgrade-user-declare"}
                                                    key={idx}
                                                    style={{width: 175}}
                                                    onClick={() => this.handleUpgradeActive(idx, r.amount)}
                                                >
                                                    <Typography style={{textAlign: "center"}}>
                                                        <Title level={5}>{r.title}</Title>
                                                        <div className="upgrade-price-desc">
                                                            <Text style={{fontSize: 15}}>¥{r.amount}</Text>
                                                            <Text
                                                                style={{marginLeft: 0, fontSize: 13, color: "#9aa5b8"}}
                                                                delete>¥{r.origin_amount}</Text>
                                                        </div>
                                                        <Space className="upgrade-space-desc" direction="vertical">
                                                            {
                                                                r.descriptions.map(d => (
                                                                    <Text>
                                                                        <CheckOutlined/>
                                                                        <Text className="upgrade-desc">{d}</Text>
                                                                    </Text>
                                                                ))
                                                            }
                                                        </Space>
                                                    </Typography>
                                                </Card>
                                            </Col>
                                        ))
                                    }
                                </Row>
                            </div>
                            {/*<Row>*/}
                            {/*    <Col span={6}>*/}
                            {/*        <Card className={"upgrade-user-declare"} style={{width: 190}}>*/}
                            {/*            <Typography style={{textAlign: "center"}}>*/}
                            {/*                <Title level={5}>1年会员</Title>*/}
                            {/*                <div className="upgrade-price-desc">*/}
                            {/*                    <Text style={{fontSize: 15}}>¥88</Text>*/}
                            {/*                    <Text style={{marginLeft: 50, fontSize: 13, color: "#9aa5b8"}} delete>¥99</Text>*/}
                            {/*                </div>*/}
                            {/*                <Space className="upgrade-space-desc" direction="vertical">*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">收藏数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">分享数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">协作人数不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">会员用户专属客服</Text>*/}
                            {/*                    </Text>*/}
                            {/*                </Space>*/}
                            {/*            </Typography>*/}
                            {/*        </Card>*/}
                            {/*    </Col>*/}
                            {/*    <Col span={6}>*/}
                            {/*        <Card className={"upgrade-user-declare"} style={{width: 190}}>*/}
                            {/*            <Typography style={{textAlign: "center"}}>*/}
                            {/*                <Title level={5}>2年会员</Title>*/}
                            {/*                <div className="upgrade-price-desc">*/}
                            {/*                    <Text style={{fontSize: 15}}>¥168</Text>*/}
                            {/*                    <Text style={{marginLeft: 50, fontSize: 13, color: "#9aa5b8"}} delete>¥198</Text>*/}
                            {/*                </div>*/}
                            {/*                <Space className="upgrade-space-desc" direction="vertical">*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">收藏数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">分享数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">协作人数不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">会员用户专属客服</Text>*/}
                            {/*                    </Text>*/}
                            {/*                </Space>*/}
                            {/*            </Typography>*/}
                            {/*        </Card>*/}
                            {/*    </Col>*/}
                            {/*    <Col span={6}>*/}
                            {/*        <Card className={"upgrade-user-declare"} style={{width: 190}}>*/}
                            {/*            <Typography style={{textAlign: "center"}}>*/}
                            {/*                <Title level={5}>3年会员</Title>*/}
                            {/*                <div className="upgrade-price-desc">*/}
                            {/*                    <Text style={{fontSize: 15}}>¥258</Text>*/}
                            {/*                    <Text style={{marginLeft: 50, fontSize: 13, color: "#9aa5b8"}} delete>¥297</Text>*/}
                            {/*                </div>*/}
                            {/*                <Space className="upgrade-space-desc" direction="vertical">*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">收藏数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">分享数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">协作人数不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">会员用户专属客服</Text>*/}
                            {/*                    </Text>*/}
                            {/*                </Space>*/}
                            {/*            </Typography>*/}
                            {/*        </Card>*/}
                            {/*    </Col>*/}
                            {/*    <Col span={6}>*/}
                            {/*        <Card className={"upgrade-user-declare"} style={{width: 190}}>*/}
                            {/*            <Typography style={{textAlign: "center"}}>*/}
                            {/*                <Title level={5}>终身会员</Title>*/}
                            {/*                <div className="upgrade-price-desc">*/}
                            {/*                    <Text style={{fontSize: 15}}>¥588</Text>*/}
                            {/*                    <Text style={{marginLeft: 50, fontSize: 13, color: "#9aa5b8"}} delete>¥998</Text>*/}
                            {/*                </div>*/}
                            {/*                <Space className="upgrade-space-desc" direction="vertical">*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">收藏数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">分享数量不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">协作人数不做限制</Text>*/}
                            {/*                    </Text>*/}
                            {/*                    <Text>*/}
                            {/*                        <CheckOutlined />*/}
                            {/*                        <Text className="upgrade-desc">会员用户专属客服</Text>*/}
                            {/*                    </Text>*/}
                            {/*                </Space>*/}
                            {/*            </Typography>*/}
                            {/*        </Card>*/}
                            {/*    </Col>*/}
                            {/*</Row>*/}
                            <div style={{marginTop: 30}}>请选择支付方式:</div>
                            <div style={{marginTop: 20}}>
                                <Row>
                                    <Col span={6}>
                                        <Card
                                            className={this.state.payMethodActive === 'wechat' ? "upgrade-pay-declare upgrade-pay-declare-active" : "upgrade-pay-declare"}
                                            key={0}
                                            style={{width: 180}}
                                            onClick={() => this.handlePayMethodActive('wechat')}>
                                            <div className="upgrade-pay">
                                                <EWechatIcon/>
                                                <span>&nbsp;&nbsp;微信支付</span>
                                            </div>
                                        </Card>
                                    </Col>
                                    {/*<Col span={6}>*/}
                                    {/*    <Card*/}
                                    {/*        className={this.state.payMethodActive === 'alipay' ? "upgrade-pay-declare upgrade-pay-declare-active" : "upgrade-pay-declare"}*/}
                                    {/*        key={1}*/}
                                    {/*        style={{width: 180}}*/}
                                    {/*        onClick={() => this.handlePayMethodActive('alipay')}*/}
                                    {/*    >*/}
                                    {/*        <div className="upgrade-pay">*/}
                                    {/*            <EAlipayIcon/>*/}
                                    {/*            <span>&nbsp;&nbsp;支付宝支付</span>*/}
                                    {/*        </div>*/}
                                    {/*    </Card>*/}
                                    {/*</Col>*/}
                                </Row>
                            </div>
                        </div>
                    </div>
                </Modal>

                <RechargePage
                    bindRef={(ref) => this.rechargeModalRef = ref}
                    handleUpgradeOpen={this.handleUpgradeOpen}
                />
            </div>
        );
    }
}

export default withRouter(connect(
state => ({state: state.dataManage.order}), dispatch => ({
    actions: bindActionCreators(OrderAction, dispatch)
}))(UpgradePage));