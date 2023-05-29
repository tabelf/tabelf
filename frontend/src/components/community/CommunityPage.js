import React, {Component} from 'react';
import {Affix, Avatar, Card, Col, Divider, Image, Layout, Pagination, Row, Space, Tag, Badge} from "antd";
import {
    EPublicCommunityAuditIcon,
    EPublicCommunityIcon, EPublicCommunityUsedMiniIcon,
    ERecentCommunityIcon,
    EStarCommunityIcon,
} from "../base/EIcon";
import FooterPage from "../footer/FooterPage";
import './style.css'
import {EyeOutlined, LikeOutlined} from "@ant-design/icons";
import {Masonry} from "react-masonry-component2";
import NavHeader from "../base/NavHeader";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CommunityAction from "../../actions/CommunityAction";
import {isNotEmpty, CommunityView, CustomSEO, getUserInfo, isEmpty} from "../../actions/Base";
import {NavLink as Link} from "react-router-dom";

const {Content} = Layout;

const sortData = [
    {
        key: '推荐',
        value: 0,
    },
    {
        key: '热门🔥',
        value: 1,
    },
    {
        key: '使用最多',
        value: 4,
    },
    {
        key: '最新发布',
        value: 3,
    }
];

const defaultCategory = {
    name: "全部",
    uid: ''
}

const {CheckableTag} = Tag;

class CommunityPage extends Component {

    state = {
        categoryUID: '',
        categoryTag: defaultCategory.name,
        sortTag: 0,
        offset: 1,
        limit: 20,
    }

    componentDidMount() {
        CustomSEO('tab精灵 - 分享社区',
            'tab精灵,分享社区,在线书签',
            'tab精灵是一个国产的在线书签管理工具，提供分享社区功能，让您和其他用户分享和发现更多有价值的网站。快来试试吧！')
        const {offset, limit, categoryUID, sortTag} = this.state;
        this.props.actions.GetPublicCommunityCategory()
        this.props.actions.GetPublicCommunity(categoryUID, sortTag, (offset - 1) * limit, limit)
    }

    handleCategoryChange = (category) => {
        const {limit, sortTag} = this.state;
        let offset = 1;
        this.props.actions.GetPublicCommunity(category.uid, sortTag, (offset - 1) * limit, limit);
        this.setState({
            categoryTag: category.name,
            categoryUID: category.uid,
            offset: 1
        });
    };

    handleSortChange = (tag) => {
        const {categoryUID, limit} = this.state;
        let offset = 1;
        this.props.actions.GetPublicCommunity(categoryUID, tag, (offset - 1) * limit, limit);
        this.setState({
            sortTag: tag,
            offset: 1
        });
    };

    handlePageChange = (page) => {
        const {categoryUID, sortTag, limit} = this.state;
        let offset = page;
        this.props.actions.GetPublicCommunity(categoryUID, sortTag, (offset - 1) * limit, limit);
        this.setState({
            offset: offset
        });
    }

    handleCommunityMeta = (communityUID, metaType) => {
        if (metaType === CommunityView) {
            const userinfo = getUserInfo()
            if (isEmpty(userinfo)) {
                return
            }
        }
        this.props.actions.UpdatePublicCommunityMeta(communityUID, metaType)
    }


    render() {
        const userinfo = getUserInfo()
        const {communityCategories, communities} = this.props.state

        const elements = communities.data.map(item => (
            <Card className="community-element-container">
                <a href={"/detail/" + item.uid}
                   target={"_blank"}
                   onClick={() => this.handleCommunityMeta(item.uid, CommunityView)}>
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
                            <span>{item.user_name}</span>
                        </div>
                    </div>
                </a>
            </Card>
        ))

        return (
            <div>
                <Layout>
                    <NavHeader active={"my-community"}/>

                    <Content className={"community-layout-content"}>
                        <Row>
                            <Col offset={4} span={16}>
                                <div style={{marginTop: 20}}>
                                    <h3>全部分享</h3>
                                </div>
                                <div style={{marginTop: 20, display: "flex"}}>
                                    <div style={{marginRight: 8, fontWeight: 500, width: 50}}>专题:</div>
                                    <div style={{width: "100%"}}>
                                        <CheckableTag
                                            key={defaultCategory.name}
                                            style={{marginBottom: 10}}
                                            checked={this.state.categoryTag === defaultCategory.name}
                                            onChange={() => this.handleCategoryChange(defaultCategory)}>
                                            {defaultCategory.name}
                                        </CheckableTag>
                                        {communityCategories.categories.map((tag) => (
                                            <CheckableTag
                                                key={tag.name}
                                                style={{marginBottom: 10}}
                                                checked={this.state.categoryTag === tag.name}
                                                onChange={() => this.handleCategoryChange(tag)}>
                                                {tag.name}
                                            </CheckableTag>
                                        ))}
                                    </div>
                                </div>
                                <div style={{marginTop: 20, marginBottom: 20, display: "flex"}}>
                                    <div style={{marginRight: 8, fontWeight: 500, width: 50}}>排序:</div>
                                    <div style={{width: "100%"}}>
                                        {sortData.map((tag) => (
                                            <CheckableTag
                                                key={tag.value}
                                                checked={this.state.sortTag === tag.value}
                                                onChange={() => this.handleSortChange(tag.value)}>
                                                {tag.key}
                                            </CheckableTag>
                                        ))}
                                    </div>
                                </div>
                                <Divider/>
                            </Col>


                            <Col offset={4} span={16}>
                                <div className={"community-layout-content-container"}>
                                    <Masonry columnsCountBreakPoints={{
                                        1700: 5,
                                        1500: 5,
                                        1300: 4,
                                        1100: 4,
                                        900: 3,
                                        700: 2
                                    }}>
                                        {elements}
                                    </Masonry>
                                </div>
                            </Col>

                            <Col offset={8} span={8} style={{
                                textAlign: "center",
                                marginBottom: 50,
                                marginTop: 50
                            }}>
                                <Pagination hideOnSinglePage={true}
                                            onChange={this.handlePageChange}
                                            current={this.state.offset}
                                            total={communities.total}
                                            pageSize={this.state.limit}
                                            showSizeChanger={false}/>
                            </Col>
                        </Row>

                        <Affix style={{position: 'fixed', bottom: 140, right: "5%"}}>
                            {
                                isNotEmpty(userinfo) ? (
                                    <Card className={"community-btn-group"}>
                                        <Space direction="vertical" size={15} style={{fontSize: 10, textAlign: "center"}}>
                                            <div>
                                                <Link to="/community/tab/1">
                                                    <EPublicCommunityIcon/>
                                                    <div>
                                                        <span>我的发布</span>
                                                    </div>
                                                </Link>
                                            </div>
                                            <div>
                                                <Link to="/community/tab/2">
                                                    <div>
                                                        <EStarCommunityIcon />
                                                    </div>
                                                    <span>我的收藏</span>
                                                </Link>
                                            </div>
                                            <div>
                                                <Link to="/community/tab/3">
                                                    <div>
                                                        <ERecentCommunityIcon />
                                                    </div>
                                                    <span>最近使用</span>
                                                </Link>
                                            </div>
                                            {
                                                communityCategories.has_admin ? (
                                                    <div>
                                                        <Link to="/community/tab/4">
                                                            <div>
                                                                <Badge dot={communityCategories.wait !== 0} >
                                                                    <EPublicCommunityAuditIcon />
                                                                </Badge>
                                                            </div>
                                                            <span>我的审核</span>
                                                        </Link>
                                                    </div>
                                                ) : (<></>)
                                            }
                                        </Space>
                                    </Card>
                                ) : (
                                    <></>
                                )
                            }
                        </Affix>

                    </Content>
                    <Divider style={{margin: 0}}/>
                    <FooterPage/>
                </Layout>
            </div>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.community}),
    dispatch => ({
        actions: bindActionCreators(CommunityAction, dispatch)
    })
)(CommunityPage);