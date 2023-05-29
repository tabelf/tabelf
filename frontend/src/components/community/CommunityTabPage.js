import React, {Component} from 'react';
import {Layout, Tabs, Row, Col, Divider, Card, Image, Space, Avatar, Tag} from "antd";
import NavHeader from "../base/NavHeader";
import {withRouter} from "../base/Base";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CommunityAction from "../../actions/CommunityAction";
import FooterPage from "../footer/FooterPage";
import {
    CloseCircleOutlined,
    EyeOutlined,
    LikeOutlined,
} from "@ant-design/icons";
import {Masonry} from "react-masonry-component2";
import CommunityAuditPage from "./CommunityAuditPage";
import {EPublicCommunityUsedMiniIcon} from "../base/EIcon";

const {Content} = Layout;

class CommunityTabPage extends Component {

    state = {
        activeTab: "1"
    }

    componentDidMount() {
        const tabID = this.props.params.tab_id
        this.setState({activeTab: tabID})
        this.props.actions.GetPublicCommunitySelf(tabID)
    }

    onChange = (key) => {
        this.setState({activeTab: key})
        if (key === "4") {
            this.auditRef.handleCommunityAuditNext()
        } else {
            this.props.actions.GetPublicCommunitySelf(key)
        }
    };

    render() {
        const {selfCommunities} = this.props.state

        const elements = selfCommunities.data.map(item => (
            <Card className="community-element-container">
                <a href={"/detail/" + item.uid} target={"_blank"}>
                    <div style={{width: 197, height: 197, overflow: "hidden"}}>
                        <Image preview={false} src={item.image}/>
                    </div>
                    <div>
                        <div className="community-element-topic">
                            <h4>{item.title}</h4>
                        </div>
                        <div>
                            <Space size={15}>
                                <div>
                                    <Space size={5}>
                                        <span>
                                            <EyeOutlined className={"community-element-meta-basic"} style={{paddingTop: 6}}/>
                                        </span>
                                        <span className={"community-element-meta"}>{item.view}</span>
                                    </Space>
                                </div>
                                <div>
                                    <Space size={5}>
                                        <span>
                                            <EPublicCommunityUsedMiniIcon />
                                        </span>
                                        <span className={"community-element-meta"}>{item.used}</span>
                                    </Space>
                                </div>
                                <div>
                                    <Space size={5}>
                                        <span>
                                            <LikeOutlined className={"community-element-meta-basic"} style={{paddingTop: 4}}/>
                                        </span>
                                        <span className={"community-element-meta"}>{item.praise}</span>
                                    </Space>
                                </div>
                            </Space>
                        </div>
                        <div className="community-element-authority">
                            <Avatar size={24} src={item.user_image}/>
                            <span className={"community-detail-author"}>{item.user_name}</span>
                            {
                                item.status === "0" ? (
                                    <Tag color="warning">待审核</Tag>
                                ) : (item.status === "-1" ? <Tag icon={<CloseCircleOutlined />} color="error">不通过</Tag> : <></>)
                            }
                        </div>
                    </div>
                </a>
            </Card>
        ))

        const tabContainer = (
            <Masonry columnsCountBreakPoints={{
                1700: 6,
                1500: 5,
                1300: 5,
                1200: 5,
                1000: 4,
                960: 3,
                700: 3
            }}>
                {elements}
            </Masonry>
        )

        const items = [
            {
                label: `我的发布`,
                key: '1',
                children: tabContainer,
            },
            {
                label: `我的收藏`,
                key: '2',
                children: tabContainer,
            },
            {
                label: `最近使用`,
                key: '3',
                children: tabContainer,
            }
        ]

        if (selfCommunities.has_admin) {
            items.push({
                label: `我的审核`,
                key: '4',
                children: <CommunityAuditPage bindRef={(ref) => this.auditRef = ref}/>,
            })
        }

        return (
            <Layout>
                <NavHeader active={"my-community"}/>

                <Content className={"community-layout-content"}>
                    <Row>
                        <Col offset={3} span={18}>
                            <div style={{marginTop: 30, marginBottom: 40, minHeight: 580}}>
                                <Tabs
                                    activeKey={this.state.activeTab}
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
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.community}),
    dispatch => ({
        actions: bindActionCreators(CommunityAction, dispatch)
    })
)(CommunityTabPage));