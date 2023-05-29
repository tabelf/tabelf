import React, {Component} from 'react';
import NavHeader from "../base/NavHeader";
import {Affix, Avatar, Button, Card, Col, Divider, Image, Layout, List, message, Row, Space, Tag} from "antd";
import FooterPage from "../footer/FooterPage";
import {
    EyeOutlined,
    FieldTimeOutlined,
    HeartFilled,
    HeartOutlined,
    LikeFilled,
    LikeOutlined,
    PlusOutlined
} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CommunityAction from "../../actions/CommunityAction";
import {withRouter} from "../base/Base";
import {EPublicCommunityTagIcon, EUseCommunityIcon} from "../base/EIcon";
import {CommunityPraise, CommunityStar, CommunityUsed, CustomSEO} from "../../actions/Base";
import UpgradePage from "../recharge/UpgradePage";

const {Content} = Layout;

class CommunityDetailPage extends Component {

    componentDidMount() {
        const communityUID = this.props.params.community_uid
        this.props.actions.GetPublicCommunityDetail(communityUID, (data) => {
            CustomSEO('tab精灵 - ' + data.title, data.title, data.html_desc)
        })
    }

    handleCommunityMeta = (communityUID, metaType, callback) => {
        this.props.actions.UpdatePublicCommunityMeta(communityUID, metaType, callback)
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    handleFocus = (detail) => {
        if (detail.has_self) {
            message.warn("不能对自己进行操作")
            return;
        }
        this.props.actions.CommunityUserFocus(detail.uid, detail.user_uid, !detail.has_follow)
    }

    render() {
        const {detail} = this.props.state
        return (
            <div>
                <Layout>
                    <NavHeader/>

                    <Content style={{padding: 0, marginTop: 30, marginBottom: 30}}>
                        <div>
                            <Row>
                                <Col offset={4} span={16}>
                                    <Card>
                                        <div className={"detail-source-container"}>
                                            <h2 style={{marginBottom: 0}}>{detail.title}</h2>
                                            <div className="detail-source-meta">
                                                <Space>
                                                    <span><EyeOutlined/> {detail.view}</span>
                                                    <span><FieldTimeOutlined/> 2023-04-06 16:07:09</span>
                                                </Space>
                                            </div>
                                            <div className="detail-source-author">
                                                <List.Item actions={[
                                                    <div
                                                        className={detail.has_follow ? "community-detail-btn-focus" : ""}>
                                                        <Button type="primary"
                                                                style={{width: 120}}
                                                                onClick={() => this.handleFocus(detail)}
                                                                icon={detail.has_follow ? <></> : <PlusOutlined/>}>
                                                            {detail.has_follow ? "已关注" : "关注"}
                                                        </Button>
                                                    </div>
                                                ]}>
                                                    <List.Item.Meta
                                                        avatar={<Avatar size={40} src={detail.user_image}/>}
                                                        title={<span>{detail.user_name}</span>}
                                                        description={<span>
                                                        <Space>
                                                            <span>粉丝: {detail.fans}</span>
                                                            <span>分享: {detail.open}</span>
                                                        </Space>
                                                    </span>}
                                                    />
                                                </List.Item>
                                            </div>
                                            <Divider/>
                                            <div>
                                                <div dangerouslySetInnerHTML={{__html: detail.description}}></div>
                                            </div>
                                        </div>
                                    </Card>
                                </Col>
                            </Row>
                            <Row>
                                <Col offset={4} span={16}>
                                    <Card>
                                        <div id={"target-element"} className={"share-canvas"}>
                                            <Image preview={false} src={detail.image}/>
                                        </div>
                                    </Card>
                                </Col>
                            </Row>
                            {
                                detail.tags.length === 0 ? (
                                    <></>
                                ) : (
                                    <Row>
                                        <Col offset={4} span={16}>
                                            <Card>
                                                <div className={"detail-source-tag"}>
                                                    <div>
                                                        {
                                                            detail.tags.map(t => (
                                                                <Tag icon={<EPublicCommunityTagIcon/>}
                                                                     className={"community-detail-tags"}>{t}</Tag>
                                                            ))
                                                        }
                                                    </div>
                                                </div>
                                            </Card>
                                        </Col>
                                    </Row>
                                )
                            }
                        </div>
                    </Content>

                    <Affix style={{position: 'fixed', bottom: 140, right: "11%"}}>
                        <Card className={"community-detail-btn-group"}>
                            <Space direction="vertical" size={15} style={{fontSize: 13, textAlign: "center"}}>
                                <div>
                                    <div onClick={() => this.handleCommunityMeta(detail.uid, CommunityPraise)}
                                         className="community-detail-btn-style">
                                        {detail.has_praise ? (
                                            <LikeFilled style={{color: "#00aeec"}}/>
                                        ) : (
                                            <LikeOutlined/>
                                        )}
                                    </div>
                                    <span>{detail.praise}</span>
                                </div>
                                <div>
                                    <div onClick={() => this.handleCommunityMeta(detail.uid, CommunityStar)}
                                         className="community-detail-btn-style">
                                        {detail.has_star ? (
                                            <HeartFilled style={{color: "#00aeec"}}/>
                                        ) : (
                                            <HeartOutlined/>
                                        )}
                                    </div>
                                    <span>{detail.star}</span>
                                </div>
                                <div>
                                    <EUseCommunityIcon
                                        onClick={() => this.handleCommunityMeta(detail.uid, CommunityUsed, this.handleUpgradeOpen)}
                                        className="community-detail-btn-style"/>
                                    <div>
                                        <span>{detail.used}</span>
                                    </div>
                                </div>
                            </Space>
                        </Card>
                    </Affix>

                    <Divider style={{margin: 0}}/>
                    <FooterPage/>
                </Layout>

                <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>
            </div>
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.community}),
    dispatch => ({
        actions: bindActionCreators(CommunityAction, dispatch)
    })
)(CommunityDetailPage));