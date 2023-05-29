import React, {Component} from 'react';
import {Card, Space, Tag} from "antd";
import {NavLink as Link} from "react-router-dom";
import {StarFilled} from "@ant-design/icons";
import {EGoodArticleAuditIcon, EGoodArticlePublishIcon, EGoodArticleSubmitIcon} from "../base/EIcon";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as GoodArticleAction from "../../actions/GoodArticleAction";
import FeedbackPage from "../feedback/FeedbackPage";
import ArticlePage from "./ArticlePage";

class ArticleMenuPage extends Component {

    componentDidMount() {
        this.props.actions.GetGoodArticleMenuData()
    }

    handleOpenFeedback = () => {
        this.feedbackModalRef.handleOpenFeedback()
    }

    handleGoodArticleOpen = () => {
        this.articleModalRef.handleGoodArticleOpen()
    }

    render() {

        const {menu} = this.props.state

        return (
            <div>
                <div>
                    <Card className={"good-article-collection"}>
                        <Space direction={"vertical"}>
                            <div>
                                <Link to={"/good/article/collections"} style={{color: "#5a5a5a"}}>
                                    <Space>
                                        <span>
                                            <StarFilled
                                                className={"good-stations-unstar"}
                                                style={{paddingTop: 4, fontSize: 16}}/>
                                        </span>
                                        <span>我的收藏</span>
                                        <span>
                                            {
                                                menu.has_anonymous ? (
                                                    <></>
                                                ) : (<Tag style={{border: 0}}>{menu.collection}</Tag>)
                                            }
                                        </span>
                                    </Space>
                                </Link>
                            </div>
                            <div>
                                <Link to={"/good/article/submit"} style={{color: "#5a5a5a"}}>
                                    <Space>
                                        <span>
                                            <EGoodArticleSubmitIcon style={{paddingTop: 4}}/>
                                        </span>
                                        <span>我的发布</span>
                                        <span>
                                            {
                                                menu.has_anonymous ? (
                                                    <></>
                                                ) : (<Tag style={{border: 0}}>{menu.publish}</Tag>)
                                            }
                                        </span>
                                    </Space>
                                </Link>
                            </div>

                            {
                                menu.has_authority ? (
                                    <div>
                                        <Link to={"/good/article/audit"} style={{color: "#5a5a5a"}}>
                                            <Space>
                                                <span>
                                                    <EGoodArticleAuditIcon style={{paddingTop: 4}}/>
                                                </span>
                                                <span>内容审核</span>
                                                <span>
                                                    {
                                                        <Tag style={{border: 0}}>{menu.audit}</Tag>
                                                    }
                                                </span>
                                            </Space>
                                        </Link>
                                    </div>
                                ) : null
                            }
                            <div>
                                <Link onClick={this.handleGoodArticleOpen} style={{color: "#5a5a5a"}}>
                                    <Space>
                                        <span>
                                            <EGoodArticlePublishIcon style={{paddingTop: 4}}/>
                                        </span>
                                        <span>好文推荐</span>
                                    </Space>
                                </Link>
                            </div>
                        </Space>
                    </Card>
                </div>

                <div>
                    <div style={{marginTop: 10, marginLeft: 15, marginBottom: 200}}>
                        <Space direction={"vertical"} size={5} style={{fontSize: 13}}>
                            <span>
                                <a href="/" style={{color: "#8590a6"}}>Tab 精灵 ©2022-2023</a>
                            </span>
                            <span>
                                <a href="/business" style={{color: "#8590a6"}} target="_blank">商务合作</a>
                            </span>
                            <span>
                                <a href="/update/log" style={{color: "#8590a6"}} target="_blank">更新日志</a>
                            </span>
                            <span>
                                <a href="/disclaimer" style={{color: "#8590a6"}} target="_blank">免责声明与条款 ‧ </a>
                                <a onClick={this.handleOpenFeedback} style={{color: "#8590a6"}}>问题反馈</a>
                            </span>
                            <span>
                                <a style={{color: "#8590a6"}} href="https://beian.miit.gov.cn"
                                   target="_blank">豫ICP备案号2023002238</a>
                            </span>
                        </Space>
                    </div>
                </div>

                <FeedbackPage bindRef={(ref) => this.feedbackModalRef = ref}/>

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
)(ArticleMenuPage);