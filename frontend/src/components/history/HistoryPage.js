import React, {Component} from 'react';
import {Avatar, Card, Col, Collapse, Divider, Empty, Layout, List, Row} from "antd";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as HistoryAction from "../../actions/HistoryAction";
import './style.css'
import {RightOutlined} from "@ant-design/icons";
import DefaultNetwork from "../../assets/default_network.png";
import {ECSVIcon, EExcelIcon, EJPGIcon, EMP3Icon, EMP4Icon, EPDFIcon, EPPTIcon, EWordIcon} from "../base/EIcon";
import {downloadFile} from "../../actions/Base";
import ReaderPage from "../content/ReaderPage";

const {Panel} = Collapse;
const {Content} = Layout;

class HistoryPage extends Component {

    componentDidMount() {
        this.props.actions.GetRecentPersonalWebLink()
    }

    handleOpenReader = (link) => {
        this.readerModalRef.handleOpenReader(link)
    }

    render() {
        const {historyWebLinks} = this.props.state
        return (
            <Content style={{margin: '20px 0px 0', overflow: 'initial'}}>
                <div className="site-history-background"
                     style={{padding: '16px 24px 10px 24px'}}>
                    <Row>
                        <Col span={18}>
                            <div className="history-title">最近</div>
                        </Col>
                    </Row>
                    <Divider style={{margin: '12px 0'}}/>
                    <Row>
                        <Col span={1}/>
                        <Col span={22} className="history-content">
                            <div>
                                <Collapse ghost activeKey={"history"}
                                          expandIcon={({isActive}) => {
                                              return <RightOutlined style={{marginTop: "7px"}}
                                                                    className="icon-bold"
                                                                    rotate={90}/>
                                          }}
                                >
                                    {
                                        <Panel header={<h3>最近更新</h3>} key={"history"} collapsible={'header'}>
                                            <Card className="history-workspace">
                                                {
                                                    historyWebLinks.web_links.length === 0 ? (
                                                        <div>
                                                            <EmptyWebLink/>
                                                        </div>
                                                    ) : (
                                                        historyWebLinks.web_links.map(link => (
                                                            <div key={link.link_uid} className="history-weblink-list">
                                                                <List.Item actions={[]}>
                                                                    <List.Item.Meta avatar={
                                                                        <div className="file-background">
                                                                            {
                                                                                (() => {
                                                                                    switch (link.file_type) {
                                                                                        case "pdf":
                                                                                            return <Avatar shape="square" size={16} src={<EPDFIcon/>}/>
                                                                                        case "docx":
                                                                                            return <Avatar shape="square" size={16} src={<EWordIcon/>}/>
                                                                                        case "ppt":
                                                                                            return <Avatar shape="square" size={16} src={<EPPTIcon/>}/>
                                                                                        case "xlsx":
                                                                                            return <Avatar shape="square" size={16} src={<EExcelIcon/>}/>
                                                                                        case "jpg":
                                                                                            return <Avatar shape="square" size={16} src={<EJPGIcon/>}/>
                                                                                        case "mp3":
                                                                                            return <Avatar shape="square" size={16} src={<EMP3Icon/>}/>
                                                                                        case "mp4":
                                                                                            return <Avatar shape="square" size={16} src={<EMP4Icon/>}/>
                                                                                        case "csv":
                                                                                            return <Avatar shape="square" size={16} src={<ECSVIcon/>}/>
                                                                                        case "url":
                                                                                            return <img style={{height: "16px", width: "16px"}}
                                                                                                        src={link.image}
                                                                                                        onError={(e) => {
                                                                                                            e.target.src = DefaultNetwork
                                                                                                            e.target.style.width = "20px"
                                                                                                            e.target.style.height = "20px"
                                                                                                            e.target.onerror = null
                                                                                                        }}/>
                                                                                        default: return <img style={{height: "20px", width: "20px"}} src={DefaultNetwork}/>
                                                                                    }
                                                                                })()
                                                                            }
                                                                        </div>}
                                                                                    title={
                                                                                        (() => {
                                                                                            switch (link.file_type) {
                                                                                                case "url": return <a href={link.link} target={"_blank"}>{link.title}</a>
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
                                                    )
                                                }
                                            </Card>
                                        </Panel>
                                    }
                                </Collapse>
                            </div>
                        </Col>
                        <Col span={1}/>
                    </Row>

                    <ReaderPage bindRef={(ref) => this.readerModalRef = ref}/>
                </div>
            </Content>
        );
    }
}

const EmptyWebLink = () => (
    <Empty style={{margin: "10px 0"}}
        image={Empty.PRESENTED_IMAGE_SIMPLE}
        description={"在此处添加链接"}
        onClick={() => {
            console.log("点击了区域")
        }}
    />
)


export default connect(
    state => ({state: state.dataManage.history}),
    dispatch => ({
        actions: bindActionCreators(HistoryAction, dispatch)
    })
)(HistoryPage);