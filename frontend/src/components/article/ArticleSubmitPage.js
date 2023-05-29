import React, {Component} from 'react';
import NavHeader from "../base/NavHeader";
import {Card, Col, Divider, Empty, Image, Layout, List, Row, Space, Tag} from "antd";
import {NavLink as Link} from "react-router-dom";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";
import ArticlePage from "./ArticlePage";
import ArticleMenuPage from "./ArticleMenuPage";

const {Content} = Layout;

class ArticleSubmitPage extends Component {

    componentDidMount() {
        this.props.actions.GetGoodArticleSubmit()
    }

    handleEditGoodArticle = (article) => {
        this.props.actions.GetGoodArticleDetail(article.uid, (data) => {
            this.articleModalRef.handleShowGoodArticle(data)
        })
    }

    render() {
        const {submit} = this.props.state
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
                                            <h3>我的发布</h3>
                                        </div>
                                        <Divider style={{marginBottom: 0, marginTop: 16}}/>
                                    </Card>

                                    <div className={"good-article-list"}>
                                        <div>
                                            <List
                                                itemLayout="horizontal"
                                                dataSource={submit.data}
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
                                                            title={
                                                                <div style={{marginBottom: 14}}>
                                                                    <a href={item.link} onClick={() => this.handleGoodArticleView(item)} style={{color: "rgb(0 0 0 / 85%)"}} target={"_blank"}>{item.title}</a>
                                                                    <span style={{marginLeft: 30}}>
                                                                        {
                                                                            item.status === "0" ? (
                                                                                <Tag color="warning">
                                                                                    待审核
                                                                                </Tag>
                                                                            ) : (
                                                                                item.status === "1" ? (
                                                                                    <Tag color="success">
                                                                                        审核通过
                                                                                    </Tag>
                                                                                ) : (
                                                                                    <Tag color="error">
                                                                                        审核拒绝
                                                                                    </Tag>
                                                                                )
                                                                            )
                                                                        }
                                                                    </span>
                                                                </div>
                                                            }
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
                                                                            <div style={{
                                                                                textAlign: "right",
                                                                                alignItems: "center"
                                                                            }}>
                                                                                <Space size={15}>
                                                                                    <div>
                                                                                        <Link onClick={() => this.handleEditGoodArticle(item)}>
                                                                                            <Space>
                                                                                                <div style={{paddingTop: 8}}>
                                                                                                    <span className={"good-article-element-source-data"}>重新编辑</span>
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

                <ArticlePage bindRef={(ref) => this.articleModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.article}),
    dispatch => ({
        actions: bindActionCreators(GoodArticleAction, dispatch)
    })
)(ArticleSubmitPage);