import React, {Component} from 'react';
import NavHeader from "../base/NavHeader";
import {Layout, Row, Col, Card, Divider, List, Image, Space, Empty} from "antd";
import {NavLink as Link} from "react-router-dom";
import {CloseCircleFilled, CheckCircleFilled} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";
import ArticleMenuPage from "./ArticleMenuPage";

const {Content} = Layout;

class ArticleAuditPage extends Component {

    componentDidMount() {
        this.props.actions.GetGoodArticleAudit()
    }

    handleGoodArticleAudit = (articleUID, status) => {
        this.props.actions.UpdateGoodArticleAudit(articleUID, status)
    }

    render() {
        const {audit} = this.props.state
        return (
            <div>
                <Layout>
                    <NavHeader active={"my-good-article"}/>

                    <Content className={"good-article-layout-content"}>
                        <Row gutter={[14, 8]} style={{marginTop: 30}}>
                            <Col offset={2} span={15}>
                                <Card className={"good-article-content"}>
                                    <Card className={"good-article-affix-header"}>
                                        <div className={"good-article-collection-title"}>
                                            <h3>内容审核</h3>
                                        </div>
                                        <Divider style={{marginBottom: 0, marginTop: 16}}/>
                                    </Card>

                                    <div className={"good-article-list"}>
                                        <div>
                                            <List
                                                itemLayout="horizontal"
                                                dataSource={audit.data}
                                                locale={{
                                                    emptyText: <Empty description={"暂无内容"} image={Empty.PRESENTED_IMAGE_SIMPLE}/>
                                                }}
                                                renderItem={item => (
                                                    <List.Item>
                                                        <List.Item.Meta
                                                            avatar={
                                                                item.image !== "" ? (
                                                                    <Image style={{
                                                                        width: 123,
                                                                        height: 77,
                                                                        borderRadius: 6
                                                                    }} preview={false} src={item.image}/>
                                                                ) : null
                                                            }
                                                            title={<div style={{marginBottom: 14}}>
                                                                <a href={item.link} onClick={() => this.handleGoodArticleView(item)} style={{color: "rgb(0 0 0 / 85%)"}} target={"_blank"}>{item.title}</a>
                                                            </div>}
                                                            description={
                                                                <div>
                                                                    <Row style={{alignItems: "flex-end"}}>
                                                                        <Col span={12}>
                                                                            <Space style={{fontSize: 12}}>
                                                                                <span>{item.source}</span>
                                                                                <span>{item.created_at}</span>
                                                                            </Space>
                                                                        </Col>
                                                                        <Col span={12}>
                                                                            <div style={{textAlign: "right", alignItems: "center"}}>
                                                                                <Space size={15}>
                                                                                    <div>
                                                                                        <Link onClick={() => this.handleGoodArticleAudit(item.uid, true)}>
                                                                                            <Space>
                                                                                                <span>
                                                                                                    <CheckCircleFilled className={"good-stations-unstar"} style={{paddingTop: 12}}/>
                                                                                                </span>
                                                                                                <div style={{paddingTop: 6}}>
                                                                                                    <span className={"good-article-element-source-data"}>通过</span>
                                                                                                </div>
                                                                                            </Space>
                                                                                        </Link>
                                                                                    </div>
                                                                                    <div>
                                                                                        <Link onClick={() => this.handleGoodArticleAudit(item.uid, false)}>
                                                                                            <Space>
                                                                                                <span>
                                                                                                    <CloseCircleFilled className={"good-stations-unstar"} style={{paddingTop: 12}}/>
                                                                                                </span>
                                                                                                <div style={{paddingTop: 6}}>
                                                                                                    <span className={"good-article-element-source-data"}>拒绝</span>
                                                                                                </div>
                                                                                            </Space>
                                                                                        </Link>
                                                                                    </div>
                                                                                </Space>
                                                                            </div>
                                                                        </Col>
                                                                    </Row>
                                                                </div>
                                                            }
                                                        />
                                                    </List.Item>
                                                )}
                                            />
                                        </div>
                                    </div>
                                </Card>
                            </Col>

                            <Col span={5}>
                                <ArticleMenuPage />
                            </Col>

                        </Row>
                    </Content>
                </Layout>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.article}),
    dispatch => ({
        actions: bindActionCreators(GoodArticleAction, dispatch)
    })
)(ArticleAuditPage);