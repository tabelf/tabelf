import React, {Component} from 'react';
import NavHeader from "../base/NavHeader";
import {Layout, Row, Col, Card, Divider, List, Image, Space, Empty} from "antd";
import {NavLink as Link} from "react-router-dom";
import {EGoodArticleAddIcon} from "../base/EIcon";
import {StarFilled} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";
import UpgradePage from "../recharge/UpgradePage";
import ArticleMenuPage from "./ArticleMenuPage";
import {ArticleCollection} from "./GoodArticlePage";

const {Content} = Layout;

const ArticleStar = 2;

class ArticleCollectionPage extends Component {

    componentDidMount() {
        this.props.actions.GetGoodArticleCollection()
    }

    constructor(props) {
        super(props);
        this.baseStarRef = React.createRef();
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    handleGoodArticleAdd = (article) => {
        this.props.actions.GetGoodArticleFolders((data) => {
            ArticleCollection(
                article.uid,
                data,
                this.baseStarRef,
                this.props.actions,
                this.handleUpgradeOpen
            )
        })
    }

    handleGoodArticleOpen = () => {
        this.articleModalRef.handleGoodArticleOpen()
    }

    handleGoodArticleMeta = (articleUID, metaType) => {
        this.props.actions.GoodArticleUnStar(articleUID, metaType)
    }

    render() {
        const {collections} = this.props.state
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
                                            <h3>我的收藏</h3>
                                        </div>
                                        <Divider style={{marginBottom: 0, marginTop: 16}}/>
                                    </Card>

                                    <div className={"good-article-list"}>
                                        <div>
                                            <List
                                                itemLayout="horizontal"
                                                dataSource={collections.data}
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
                                                                                        <Link onClick={() => this.handleGoodArticleAdd(item)}>
                                                                                            <Space>
                                                                                                <span>
                                                                                                    <EGoodArticleAddIcon style={{paddingTop: 12}}/>
                                                                                                </span>
                                                                                                <div style={{paddingTop: 8}}>
                                                                                                    <span className={"good-article-element-source-data"}>添加</span>
                                                                                                </div>
                                                                                            </Space>
                                                                                        </Link>
                                                                                    </div>
                                                                                    <div>
                                                                                        <Link onClick={() => this.handleGoodArticleMeta(item.uid, ArticleStar)}>
                                                                                            <Space>
                                                                                                <span>
                                                                                                    <StarFilled className={"good-stations-unstar"} style={{paddingTop: 12}}/>
                                                                                                </span>
                                                                                                <div style={{paddingTop: 8}}>
                                                                                                    <span className={"good-article-element-source-data"}>取消收藏</span>
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

                <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.article}),
    dispatch => ({
        actions: bindActionCreators(GoodArticleAction, dispatch)
    })
)(ArticleCollectionPage);