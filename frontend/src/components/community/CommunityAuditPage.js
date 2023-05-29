import React, {Component} from 'react';
import {Avatar, Button, Card, Col, Collapse, Empty, List, Row, Space, Tag} from "antd";
import {RightOutlined, ExclamationCircleOutlined, CheckCircleOutlined, CloseCircleOutlined} from "@ant-design/icons";
import {ECSVIcon, EExcelIcon, EJPGIcon, EMP3Icon, EMP4Icon, EPDFIcon, EPPTIcon, EWordIcon} from "../base/EIcon";
import DefaultNetwork from "../../assets/default_network.png";
import {downloadFile} from "../../actions/Base";
import {withRouter} from "../base/Base";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CommunityAction from "../../actions/CommunityAction";
import ReaderPage from "../content/ReaderPage";

const {Panel} = Collapse;

class CommunityAuditPage extends Component {

    componentDidMount() {
        this.props.bindRef(this)
        this.props.actions.GetPublicCommunityAudit()
    }

    handleOpenReader = (link) => {
        this.readerModalRef.handleOpenReader(link)
    }

    handleCommunityAudit = (communityUID, status) => {
        this.props.actions.UpdateAuditPublicCommunity(communityUID, status)
    }

    handleCommunityAuditNext = () => (
        this.props.actions.GetPublicCommunityAudit()
    )

    render() {
        const { audit } = this.props.state
        return (
            <div>
                <Row style={{alignItems: "center"}}>
                    <Col span={12}>
                        <div style={{textAlign: "left", opacity: 0.9}}>
                            <Space size={"large"}>
                                <div>待审核数量: {audit.audit_count}</div>
                                {audit.audit_count === 0 ? (<></>) : (
                                    <Space>
                                        <span>状态:</span>
                                        <div>
                                            {
                                                audit.status === "0" ? (
                                                    <Tag icon={<ExclamationCircleOutlined />} color="warning">
                                                        待审核
                                                    </Tag>
                                                ) : (
                                                    audit.status === "1" ?
                                                        (
                                                            <Tag icon={<CheckCircleOutlined />} color="success">
                                                                审核通过
                                                            </Tag>
                                                        ) : (
                                                            <Tag icon={<CloseCircleOutlined />} color="error">
                                                                审核拒绝
                                                            </Tag>
                                                        ))
                                            }
                                        </div>
                                    </Space>
                                )}
                            </Space>
                        </div>
                    </Col>
                    <Col span={12}>
                        <div style={{textAlign: "right"}}>
                            <Space>
                                <Button type={"primary"} onClick={() => this.handleCommunityAudit(audit.community_uid, "-1")} danger>审核拒绝</Button>
                                <Button type={"primary"} onClick={() => this.handleCommunityAudit(audit.community_uid, "1")}>审核通过</Button>
                                <Button onClick={this.handleCommunityAuditNext}>下一个</Button>
                            </Space>
                        </div>
                    </Col>
                    <Col span={24} style={{marginTop: 30}}>
                        {audit.audit_count === 0 ? (
                            <></>
                        ) : (
                            <Card>
                                <div>
                                    <h3>{audit.community_title}</h3>
                                </div>
                                <div dangerouslySetInnerHTML={{__html: audit.community_description}}></div>
                            </Card>
                        )}
                    </Col>
                </Row>

                <div>
                    <div style={{marginTop: 30}}>
                        <Collapse ghost
                                  expandIcon={({isActive}) => {
                                      return <RightOutlined style={{marginTop: "7px"}}
                                                            className="icon-bold"
                                                            rotate={isActive ? 90 : 0}/>}
                                  }>
                            {
                                audit.personal_workspaces.map(w => (
                                    <Panel header={<h3>{w.workspace_name}</h3>}
                                           collapsible={'header'}>
                                        <Card className="workspace">
                                            {
                                                w.web_links.length === 0 ? (
                                                    <div>
                                                        <EmptyWebLink/>
                                                    </div>
                                                ) : (
                                                    <div>
                                                        {w.web_links.map(link => (
                                                            <div key={link.link_uid} className="weblink-list">
                                                                <List.Item>
                                                                    <List.Item.Meta avatar={
                                                                        <div className="file-background">
                                                                            {
                                                                                (() => {
                                                                                    switch (link.file_type) {
                                                                                        case "pdf":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EPDFIcon/>}/>
                                                                                        case "docx":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EWordIcon/>}/>
                                                                                        case "ppt":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EPPTIcon/>}/>
                                                                                        case "xlsx":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EExcelIcon/>}/>
                                                                                        case "jpg":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EJPGIcon/>}/>
                                                                                        case "mp3":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EMP3Icon/>}/>
                                                                                        case "mp4":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<EMP4Icon/>}/>
                                                                                        case "csv":
                                                                                            return <Avatar
                                                                                                shape="square" size={16}
                                                                                                src={<ECSVIcon/>}/>
                                                                                        case "url":
                                                                                            return <img style={{
                                                                                                height: "16px",
                                                                                                width: "16px"
                                                                                            }}
                                                                                                        src={link.image}
                                                                                                        onError={(e) => {
                                                                                                            e.target.src = DefaultNetwork
                                                                                                            e.target.style.width = "20px"
                                                                                                            e.target.style.height = "20px"
                                                                                                            e.target.onerror = null
                                                                                                        }}/>
                                                                                        default:
                                                                                            return <img style={{
                                                                                                height: "20px",
                                                                                                width: "20px"
                                                                                            }} src={DefaultNetwork}/>
                                                                                    }
                                                                                })()
                                                                            }
                                                                        </div>}
                                                                                    title={
                                                                                        (() => {
                                                                                            switch (link.file_type) {
                                                                                                case "url":
                                                                                                    return <a href={link.link} target={"_blank"}>{link.title}</a>
                                                                                                case "pdf":
                                                                                                case "jpg":
                                                                                                case "csv":
                                                                                                case "mp3":
                                                                                                case "mp4":
                                                                                                case "docx":
                                                                                                    return <a onClick={() => this.handleOpenReader(link)}>{link.title}</a>
                                                                                                default:
                                                                                                    return <a onClick={() => downloadFile(link.link, link.title)}>{link.title}</a>
                                                                                            }
                                                                                        })()
                                                                                    }
                                                                                    description={link.description}/>
                                                                </List.Item>
                                                            </div>
                                                        ))
                                                        }
                                                    </div>
                                                )
                                            }
                                        </Card>
                                    </Panel>
                                ))
                            }
                        </Collapse>

                        <ReaderPage bindRef={(ref) => this.readerModalRef = ref}/>
                    </div>
                </div>
            </div>
        );
    }
}

const EmptyWebLink = () => (
    <Empty style={{margin: "10px 0"}}
           image={Empty.PRESENTED_IMAGE_SIMPLE}
           description={"在此处添加链接"}
    />
)

export default withRouter(connect(
    state => ({state: state.dataManage.community}),
    dispatch => ({
        actions: bindActionCreators(CommunityAction, dispatch)
    })
)(CommunityAuditPage));