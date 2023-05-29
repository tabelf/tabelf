import React, {Component} from 'react';
import NavHeader from "../base/NavHeader";
import {
    Avatar,
    Button,
    Card,
    Col,
    Divider,
    Layout,
    Row,
    Space,
    Tabs,
    Tag,
} from "antd";
import FooterPage from "../footer/FooterPage";
import {EllipsisOutlined} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {Masonry} from "react-masonry-component2";
import GoodStationPage from "./GoodStationPage";
import {UpdateGoodStationAudit} from "../../actions/CustomerAction";

const {Content} = Layout;

class RecommendAuditPage extends Component {

    componentDidMount() {
        this.props.actions.GetGoodStationAudit()
    }

    handleEditGoodStation = (station) => {
        this.props.actions.GetGoodStationDetail(station.uid, (data) => {
            this.stationModalRef.handleShowGoodStation(data, true)
        })
    }

    handleAuditGoodStation = (stationUID, status) => {
        this.props.actions.UpdateGoodStationAudit(stationUID, status)
    }

    render() {
        const {stationAudits} = this.props.state

        const elements = stationAudits.data.map((station) => {
            return (
                <Card className="good-stations-element-container">
                    <a href={station.link} target="_blank"
                       style={{color: "rgba(0, 0, 0, 0.85)"}}>
                        <div className={"good-stations-element-content"}>
                            <div>
                                <div className={"good-stations-element-title"}>
                                    <span>{station.title}</span>
                                </div>
                                <div
                                    className={"good-stations-element-description"}>
                                    <span>{station.description}</span>
                                </div>
                                <div className={"good-stations-element-tags"}>
                                    {
                                        station.tags.map(tag => (
                                            <Tag color="#eee" style={{color: "#999"}}># {tag}</Tag>
                                        ))
                                    }
                                </div>
                            </div>
                            <div className={"good-stations-element-image"}>
                                <Avatar shape="square" size={72} src={station.image}/>
                            </div>
                        </div>
                    </a>
                    <div className={"good-stations-element-source"}>
                        <Space>
                            <div>
                                <Space>
                                    <span>
                                        <Avatar shape="square" size={16} src={station.icon}/>
                                    </span>
                                    <span className={"good-stations-element-source-title"}>{station.source}</span>
                                </Space>
                            </div>
                            <Divider type="vertical"/>
                            <Space size={15}>
                                <div>
                                    <Space>
                                        <a onClick={() => this.handleAuditGoodStation(station.uid, true)} className={"good-stations-element-source-data"}>通过</a>
                                    </Space>
                                </div>
                                <div>
                                    <Space>
                                        <a onClick={() => this.handleAuditGoodStation(station.uid, false)} className={"good-stations-element-source-data"}>拒绝</a>
                                    </Space>
                                </div>
                                <div>
                                    <Space>
                                        <span>
                                            <Button type={"text"}
                                                    shape="circle"
                                                    size={"small"}
                                                    onClick={() => this.handleEditGoodStation(station)}
                                                    icon={<EllipsisOutlined style={{color: "#aaa"}}/>} />
                                        </span>
                                    </Space>
                                </div>
                            </Space>
                        </Space>
                    </div>
                </Card>
            );
        })

        const items = [
            {
                label: `我的审核`,
                key: '1',
                children: <Row style={{alignItems: "center"}}>
                    <Col span={12}>
                        <div style={{textAlign: "left", opacity: 0.9}}>
                            <Space size={"large"}>
                                <div>待审核数量: {stationAudits.data.length}</div>
                            </Space>
                        </div>
                    </Col>
                    <Col span={24}>
                        <Masonry columnsCountBreakPoints={{
                            1700: 4,
                            1500: 4,
                            1300: 3,
                            960: 3,
                            700: 2
                        }}>
                            {elements}
                        </Masonry>
                    </Col>
                </Row>,
            },
        ]
        return (
            <div>
                <Layout>
                    <NavHeader active={"my-good-station"}/>

                    <Content className={"community-layout-content"}>
                        <Row>
                            <Col offset={3} span={18}>
                                <div style={{marginTop: 30, marginBottom: 40, minHeight: 580}}>
                                    <Tabs
                                        onChange={this.onChange}
                                        items={items}
                                    />
                                </div>
                            </Col>
                        </Row>
                    </Content>

                    <Divider style={{margin: 0}}/>
                    <FooterPage/>
                </Layout>

                <GoodStationPage bindRef={(ref) => this.stationModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(RecommendAuditPage);