import React, {Component} from 'react';
import {
    Affix,
    Avatar, Badge,
    Button,
    Card,
    Col,
    Divider,
    Form,
    Layout,
    Modal,
    Row,
    Select,
    Space,
    Spin,
    Tag,
} from 'antd';
import './style.css'
import {EyeFilled, LikeFilled, StarFilled, EllipsisOutlined} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {Masonry} from "react-masonry-component2";
import FooterPage from "../footer/FooterPage";
import {
    EPublicCommunityAuditIcon,
    EStationPublishIcon
} from "../base/EIcon";
import {CustomSEO, getUserInfo, isEmpty} from "../../actions/Base";
import InfiniteScroll from 'react-infinite-scroll-component';
import UpgradePage from "../recharge/UpgradePage";
import NavHeader from "../base/NavHeader";
import {NavLink as Link} from "react-router-dom";
import GoodStationPage from "./GoodStationPage";

const {Content} = Layout;

const {CheckableTag} = Tag;

const sortData = [
    {
        key: '默认',
        value: 0,
    },
    {
        key: '点赞',
        value: 1,
    },
    {
        key: '收藏',
        value: 2,
    },
    {
        key: '最新',
        value: 3,
    }
];

const defaultCategory = {
    name: "全部",
    uid: ''
}

const stationView = 0
const stationPraise = 1
const stationStar = 2

class RecommendPage extends Component {

    state = {
        hasGoodStation: false,
        categoryUID: '',
        categoryTag: defaultCategory.name,
        sortTag: 0,
        imageURL: null,
        offset: 1,
        limit: 20,
        loading: false
    }

    constructor(props) {
        super(props);
        this.baseStarRef = React.createRef();
    }

    componentDidMount() {
        CustomSEO('tab精灵 - 好站推荐',
            'tab精灵,好站推荐,在线书签',
           'tab精灵是一个国产的在线书签管理工具，提供好站推荐功能，让您收藏、分享、发现更多有价值的网站。快来试试吧！')
        const {offset, limit, categoryUID, sortTag} = this.state;
        this.props.actions.GetGoodStationCategory()
        this.props.actions.GetGoodStationRecommend(categoryUID, sortTag, (offset - 1) * limit, limit);
    }

    handleGoodStationOpen = () => {
        this.stationModalRef.handleGoodStationOpen()
    }

    handleCategoryChange = (category) => {
        const {limit, sortTag} = this.state;
        let offset = 1;
        this.props.actions.GetGoodStationRecommend(category.uid, sortTag, (offset - 1) * limit, limit);
        this.setState({
            categoryTag: category.name,
            categoryUID: category.uid,
            offset: 1
        });
    };

    handleSortChange = (tag) => {
        const {categoryUID, limit} = this.state;
        let offset = 1;
        this.props.actions.GetGoodStationRecommend(categoryUID, tag, (offset - 1) * limit, limit);
        this.setState({
            sortTag: tag,
            offset: 1
        });
    };

    loadMoreData = () => {
        let {offset, limit, categoryUID, sortTag} = this.state;
        offset = offset + 1
        this.props.actions.GetGoodStationRecommend(categoryUID, sortTag, (offset - 1) * limit, limit);
        this.setState({offset: offset});
    }

    handleGoodStationMeta = (stationUID, metaType) => {
        if (metaType === stationView) {
            const userinfo = getUserInfo()
            if (isEmpty(userinfo)) {
                return
            }
        }
        this.props.actions.UpdateGoodStationMeta(stationUID, metaType)
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    handleGoodStationCollection = (station) => {
        this.props.actions.GetPersonalFolders((data) => {
            StationStar(
                station.uid,
                data,
                this.baseStarRef,
                this.props.actions,
                this.handleUpgradeOpen
            )
        })
    }

    handleEditGoodStation = (station) => {
        this.props.actions.GetGoodStationDetail(station.uid, (data) => {
            this.stationModalRef.handleShowGoodStation(data, false)
        })
    }

    render() {
        const {stationCategories, stations} = this.props.state

        const elements = stations.data.map((station) => {
            return (
                <Card className="good-stations-element-container">
                    <a onClick={() => this.handleGoodStationMeta(station.uid, stationView)}
                       href={station.link}
                       target="_blank"
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
                                            <Tag color="#eee" style={{color: "#999", marginTop: 5}}># {tag}</Tag>
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
                                        <span onClick={() => this.handleGoodStationMeta(station.uid, stationPraise)}>
                                            <LikeFilled className={station.has_praise ? "good-stations-star" : "good-stations-unstar"} style={{paddingTop: 4}}/>
                                        </span>
                                        <span className={"good-stations-element-source-data"}>{station.praise}</span>
                                    </Space>
                                </div>
                                <div>
                                    <Space>
                                        <span onClick={() => this.handleGoodStationCollection(station)}>
                                            <StarFilled className={station.has_star ? "good-stations-star" : "good-stations-unstar"} style={{paddingTop: 6}}/>
                                        </span>
                                        <span className={"good-stations-element-source-data"}>{station.star}</span>
                                    </Space>
                                </div>
                                <div>
                                    <Space>
                                        <span>
                                            <EyeFilled style={{color: "#aaa", paddingTop: 6}}/>
                                        </span>
                                        <span className={"good-stations-element-source-data"}>{station.view}</span>
                                    </Space>
                                </div>
                                <div>
                                    {
                                        stationCategories.has_authority ? (
                                            <Space>
                                                <span>
                                                    <Button type={"text"}
                                                            shape="circle"
                                                            size={"small"}
                                                            onClick={() => this.handleEditGoodStation(station)}
                                                            icon={<EllipsisOutlined style={{color: "#aaa"}}/>} />
                                                </span>
                                            </Space>
                                        ) : (<></>)
                                    }
                                </div>
                            </Space>
                        </Space>
                    </div>
                </Card>
            );
        })

        return (
            <div>
                <Layout>
                    <NavHeader active={"my-good-station"}/>

                    <Content className={"recommend-layout-content"}>
                        <Row>
                            <Col offset={2} span={20}>
                                <div style={{marginTop: 20}}>
                                    <h3>全部内容</h3>
                                </div>
                                <div style={{marginTop: 20}}>
                                    <span style={{marginRight: 8, fontWeight: 500}}>分类:</span>
                                    <CheckableTag
                                        key={defaultCategory.name}
                                        checked={this.state.categoryTag === defaultCategory.name}
                                        onChange={() => this.handleCategoryChange(defaultCategory)}>
                                        {defaultCategory.name}
                                    </CheckableTag>
                                    {stationCategories.categories.map((tag) => (
                                        <CheckableTag
                                            key={tag.name}
                                            checked={this.state.categoryTag === tag.name}
                                            onChange={() => this.handleCategoryChange(tag)}>
                                            {tag.name}
                                        </CheckableTag>
                                    ))}
                                </div>
                                <div style={{marginTop: 20, marginBottom: 20}}>
                                    <span style={{marginRight: 8, fontWeight: 500}}>排序:</span>
                                    {sortData.map((tag) => (
                                        <CheckableTag
                                            key={tag.value}
                                            checked={this.state.sortTag === tag.value}
                                            onChange={() => this.handleSortChange(tag.value)}>
                                            {tag.key}
                                        </CheckableTag>
                                    ))}
                                </div>
                            </Col>
                            <Col offset={2} span={20}>
                                <div className={"recommend-layout-content-container"}>
                                    <InfiniteScroll
                                        dataLength={stations.data.length}
                                        style={{overflow: "unset"}}
                                        next={this.loadMoreData}
                                        hasMore={stations.data.length < stations.total}
                                        loader={<div className="loading-container"><Spin/></div>}
                                        endMessage={<Divider plain></Divider>}
                                        scrollableTarget="scrollableDiv">
                                        <Masonry columnsCountBreakPoints={{
                                            1700: 5,
                                            1300: 4,
                                            1200: 3,
                                            960: 3,
                                            700: 2
                                        }}>
                                            {elements}
                                        </Masonry>
                                    </InfiniteScroll>
                                </div>
                            </Col>
                        </Row>

                        <Affix style={{position: 'fixed', bottom: 140, right: 0}}>
                            {
                                stationCategories.has_authority ? (
                                        <Card className={"community-btn-group"}>
                                            <Space direction="vertical" size={15} style={{fontSize: 10, textAlign: "center"}}>
                                                <div>
                                                    <Link onClick={this.handleGoodStationOpen}>
                                                        <EStationPublishIcon/>
                                                        <div>
                                                            <span>好站发布</span>
                                                        </div>
                                                    </Link>
                                                </div>
                                                <div>
                                                    <Link to="/recommend/audit">
                                                        <div>
                                                            <Badge dot={stationCategories.wait !== 0}>
                                                                <EPublicCommunityAuditIcon />
                                                            </Badge>
                                                        </div>
                                                        <span>好站审核</span>
                                                    </Link>
                                                </div>
                                            </Space>
                                        </Card>
                                ) : <></>
                            }
                        </Affix>

                    </Content>
                    <Divider style={{margin: 0}}/>
                    <FooterPage/>

                    <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>

                    <GoodStationPage bindRef={(ref) => this.stationModalRef = ref}/>
                </Layout>
            </div>
        );
    }
}

const StationStar = (stationUID, personalFolders, ref, actions, handleUpgradeOpen) => {
    return Modal.confirm({
        icon: null,
        title: <h4>好站收藏</h4>,
        okText: '确认',
        cancelText: '取消',
        width: 406,
        height: 172,
        bodyStyle: {
            padding: 24,
        },
        className: "station-star-folder",
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
            actions.GoodStationStar(stationUID, data.folder_uid, handleUpgradeOpen)
            ref.current.resetFields()
        },
        onCancel: () => {}
    });
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(RecommendPage);