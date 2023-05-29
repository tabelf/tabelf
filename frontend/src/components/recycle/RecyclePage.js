import React, {Component} from 'react';
import {Avatar, Button, Card, Col, Collapse, Divider, Dropdown, Empty, Layout, List, Row} from "antd";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as RecycleAction from "../../actions/RecycleAction";
import './style.css'
import {
    ECSVIcon,
    EExcelIcon,
    EJPGIcon,
    EMoreIcon,
    EMP3Icon,
    EMP4Icon,
    EPDFIcon,
    EPPTIcon,
    EWordIcon
} from "../base/EIcon";
import {DeleteOutlined, ReloadOutlined, RightOutlined} from "@ant-design/icons";
import DefaultNetwork from "../../assets/default_network.png";
import {downloadFile} from "../../actions/Base";
import ReaderPage from "../content/ReaderPage";

const {Panel} = Collapse;
const {Content} = Layout;

class RecyclePage extends Component {

    componentDidMount() {
        this.props.actions.GetRecyclingPersonalWebLink()
    }

    handleOpenReader = (link) => {
        this.readerModalRef.handleOpenReader(link)
    }

    render() {
        const {recycleWebLinks} = this.props.state
        return (
            <Content style={{margin: '20px 0px 0', overflow: 'initial'}}>
                <div className="site-recycle-background"
                     style={{padding: '16px 24px 10px 24px'}}>
                    <Row>
                        <Col span={18}>
                            <div className="recycle-title">回收站</div>
                        </Col>
                    </Row>
                    <Divider style={{margin: '12px 0'}}/>
                    <Row>
                        <Col span={1}/>
                        <Col span={22} className="recycle-content">
                            <div>
                                <Collapse ghost activeKey={"recycle"}
                                          expandIcon={({isActive}) => {
                                              return <RightOutlined style={{marginTop: "7px"}}
                                                                    className="icon-bold"
                                                                    rotate={90}/>
                                          }}
                                >
                                    {
                                        <Panel header={<h3>最近删除</h3>} key={"recycle"} collapsible={'header'}>
                                            <Card className="recycle-workspace">
                                                {
                                                    recycleWebLinks.web_links.length === 0 ? (
                                                        <div>
                                                            <EmptyRecycleWebLink/>
                                                        </div>
                                                    ) : (
                                                        recycleWebLinks.web_links.map(link => (
                                                            <div key={link.link_uid} className="recycle-list">
                                                                <List.Item actions={[
                                                                    <RecycleMoreExtra
                                                                        actions={this.props.actions}
                                                                        link={link}
                                                                    />
                                                                ]}>
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

const EmptyRecycleWebLink = () => (
    <Empty style={{margin: "10px 0"}}
        image={Empty.PRESENTED_IMAGE_SIMPLE}
        description={"回收站内为空"}
        onClick={() => {
            console.log("点击了区域")
        }}
    />
)

const RecycleMoreExtra = ({actions, link}) => {

    const items = [
        {
            label: '恢复',
            key: 'item-1',
            icon: <ReloadOutlined />,
            onClick: () => {
                actions.RestoreDeletePersonalWebLink(link.link_uid)
            }
        },
        {
            label: '彻底删除',
            key: 'item-2',
            icon: <DeleteOutlined />,
            onClick: () => {
                actions.DeleteForeverPersonalWebLink(link.link_uid)
            }
        }
    ];

    return <Dropdown
        overlayClassName={"recycle-menu-extra"}
        placement="bottomRight"
        autoAdjustOverflow={false}
        overlayStyle={{
            width: 160,
            height: 100
        }}
        trigger={["click"]}
        menu={{items}}>
        <Button type="text"
                shape="circle"
                className={"recycle-menu-btn"}
                size={"small"}
                icon={<EMoreIcon/>}/>
    </Dropdown>
}



export default connect(
    state => ({state: state.dataManage.recycle}),
    dispatch => ({
        actions: bindActionCreators(RecycleAction, dispatch)
    })
)(RecyclePage);