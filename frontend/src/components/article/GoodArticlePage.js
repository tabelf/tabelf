import React, {Component} from 'react';
import {
    Affix,
    Button,
    Card,
    Col,
    Divider, Empty,
    Form,
    Image,
    Layout,
    List,
    Modal,
    Row,
    Select,
    Space,
    Spin,
    Tag
} from "antd";
import NavHeader from "../base/NavHeader";
import {NavLink as Link} from "react-router-dom";
import {EGoodArticleAddIcon} from "../base/EIcon";
import UpgradePage from "../recharge/UpgradePage";
import './style.css'
import {CaretRightOutlined, EllipsisOutlined, StarFilled} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";
import ArticlePage from "./ArticlePage";
import InfiniteScroll from 'react-infinite-scroll-component';
import ArticleMenuPage from "./ArticleMenuPage";

const {CheckableTag} = Tag;

const defaultCategory = {
    name: "全部",
    uid: ''
}

const {Content} = Layout;

const ArticleView = 0;
const ArticleUsed = 1;
const ArticleStar = 2;
const ArticleNew  = 3;

class GoodArticlePage extends Component {

    state = {
        categoryTag: defaultCategory.name,
        categoryUID: '',
        sortTag: 0,
        offset: 1,
        limit: 20,
    }

    constructor(props) {
        super(props);
        this.baseStarRef = React.createRef();
    }

    componentDidMount() {
        const {offset, limit, categoryUID} = this.state;
        this.props.actions.GetGoodArticleCategory()
        this.props.actions.GetGoodArticle(categoryUID, ArticleNew, (offset - 1) * limit, limit)
        this.props.actions.GetHotGoodArticle()
    }

    handleCategoryChange = (category) => {
        const {limit} = this.state;
        let offset = 1;
        this.props.actions.GetGoodArticle(category.uid, ArticleNew, (offset - 1) * limit, limit);
        this.setState({
            categoryTag: category.name,
            categoryUID: category.uid,
            offset: 1
        });
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

    handleGoodArticleMeta = (articleUID, metaType) => {
        this.props.actions.UpdateGoodArticleMeta(articleUID, metaType)
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    loadMoreData = () => {
        let {offset, limit, categoryUID, sortTag} = this.state;
        offset = offset + 1
        this.props.actions.GetGoodArticle(categoryUID, sortTag, (offset - 1) * limit, limit);
        this.setState({offset: offset});
    }

    handleEditGoodArticle = (article) => {
        this.props.actions.GetGoodArticleDetail(article.uid, (data) => {
            this.articleModalRef.handleShowGoodArticle(data)
        })
    }

    handleGoodArticleView = (article) => {
        this.props.actions.GoodArticleView(article.uid)
    }

    render() {
        const {articles, articleCategories, hotArticle} = this.props.state

        return (
            <div>
                <Layout>
                    <NavHeader active={"my-good-article"}/>

                    <Content className={"good-article-layout-content"}>
                        <Row gutter={[14, 8]} style={{marginTop: 30}}>
                            <Col offset={2} span={15}>
                                <Card className={"good-article-content"}>
                                    <Affix offsetTop={0}>
                                        <Card className={"good-article-affix-header"}>
                                            <div className={"good-article-topic"}>
                                                <div style={{marginRight: 8, fontWeight: 500, width: 50}}>专题:</div>
                                                <div style={{width: "100%"}}>
                                                    <CheckableTag
                                                        key={defaultCategory.name}
                                                        checked={this.state.categoryTag === defaultCategory.name}
                                                        onChange={() => this.handleCategoryChange(defaultCategory)}>
                                                        {defaultCategory.name}
                                                    </CheckableTag>
                                                    {articleCategories.categories.map((tag) => (
                                                        <CheckableTag
                                                            key={tag.name}
                                                            checked={this.state.categoryTag === tag.name}
                                                            onChange={() => this.handleCategoryChange(tag)}>
                                                            {tag.name}
                                                        </CheckableTag>
                                                    ))}
                                                </div>
                                            </div>
                                            <Divider style={{marginBottom: 0}}/>
                                        </Card>
                                    </Affix>

                                    <div className={"good-article-list"}>
                                        <div>
                                            <InfiniteScroll
                                                dataLength={articles.data.length}
                                                style={{overflow: "unset"}}
                                                next={this.loadMoreData}
                                                hasMore={articles.data.length < articles.total}
                                                loader={<div className="loading-container"><Spin/></div>}
                                                endMessage={<Divider plain></Divider>}
                                                scrollableTarget="scrollableDiv">
                                                <List
                                                    itemLayout="horizontal"
                                                    dataSource={articles.data}
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
                                                                                                        <span className={"good-article-element-source-data"} style={{paddingTop: 12}}>添加</span>
                                                                                                    </div>
                                                                                                </Space>
                                                                                            </Link>
                                                                                        </div>
                                                                                        <div>
                                                                                            <Link onClick={() => this.handleGoodArticleMeta(item.uid, ArticleStar)}>
                                                                                                <Space>
                                                                                                    <span>
                                                                                                        <StarFilled className={item.has_star ? "good-stations-star" : "good-stations-unstar"} style={{paddingTop: 12}}/>
                                                                                                    </span>
                                                                                                    <div style={{paddingTop: 8}}>
                                                                                                        <span className={"good-article-element-source-data"}>收藏</span>
                                                                                                    </div>
                                                                                                </Space>
                                                                                            </Link>
                                                                                        </div>
                                                                                        <div>
                                                                                            <Space>
                                                                                                <div style={{paddingTop: 8}}>
                                                                                                    <Button type={"text"}
                                                                                                            shape="circle"
                                                                                                            size={"small"}
                                                                                                            onClick={() => this.handleEditGoodArticle(item)}
                                                                                                            icon={<EllipsisOutlined style={{color: "#aaa"}}/>}/>
                                                                                                </div>
                                                                                            </Space>
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
                                            </InfiniteScroll>
                                        </div>
                                    </div>
                                </Card>

                                <div style={{textAlign: "center", marginTop: 15, fontSize: 12, color: "#8590a6"}}>
                                    <span>已经到底了，不要在往下拉了～</span>
                                </div>
                            </Col>

                            <Col span={5}>
                                <div style={{marginBottom: 10}}>
                                    <Card className={"good-article-hot-top"}>
                                        <div className={"good-article-hot-title"}>
                                            <h3>热点TOP</h3>
                                        </div>
                                        <Divider style={{margin: 0}}/>
                                        <div className={"good-article-hot-container"}>
                                            {
                                                hotArticle.data.map(h => (
                                                    <div className={"good-article-hot-article"}>
                                                        <div>
                                                            <CaretRightOutlined style={{color: "#ccc"}}/>
                                                        </div>
                                                        <div className={"good-article-hot-article-title"}>
                                                            <a href={h.link} style={{color: "rgb(0 0 0 / 85%)"}} target={"_blank"}>{h.title}</a>
                                                        </div>
                                                    </div>
                                                ))
                                            }
                                        </div>
                                    </Card>
                                </div>

                                <ArticleMenuPage />
                            </Col>

                        </Row>

                    </Content>

                    <ArticlePage bindRef={(ref) => this.articleModalRef = ref}/>

                    <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>
                </Layout>
            </div>
        )
    }
}

export const ArticleCollection = (articleUID, personalFolders, ref, actions, handleUpgradeOpen) => {
    return Modal.confirm({
        icon: null,
        title: <h4>好文添加</h4>,
        okText: '确认',
        cancelText: '取消',
        width: 406,
        height: 172,
        bodyStyle: {
            padding: 24,
        },
        className: "article-star-folder",
        content: (
            <div>
                <Form style={{marginTop: 20}} ref={ref}>
                    <Form.Item label={<span>我的文件夹</span>} name="folder_uid">
                        <Select placeholder="请选择文件夹" allowClear>
                            {
                                personalFolders.folders.map(f => (
                                    <Select.Option value={f.folder_uid}>{f.folder_name}</Select.Option>
                                ))
                            }
                        </Select>
                    </Form.Item>
                </Form>
            </div>
        ),
        onOk: () => {
            const data = ref.current.getFieldsValue()
            actions.GoodArticleStar(articleUID, data.folder_uid, handleUpgradeOpen)
            ref.current.resetFields()
        },
        onCancel: () => {}
    });
}

export default connect(
    state => ({state: state.dataManage.article}),
    dispatch => ({
        actions: bindActionCreators(GoodArticleAction, dispatch)
    })
)(GoodArticlePage);