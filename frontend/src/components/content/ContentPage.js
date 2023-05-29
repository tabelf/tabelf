import React, {Component} from 'react';
import {LinkOutlined} from '@ant-design/icons';
import {Affix, Button, Col, Divider, Layout, Row, Space, Tooltip} from 'antd';
import './style.css'
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import WorkspacePage from "./WorkspacePage";
import SharePage from "../share/SharePage";
import {withRouter} from "../base/Base";
import {isEmpty, isNotEmpty} from "../../actions/Base";
import IndexPage from "../home/IndexPage";
import {EGlobalAddIcon, ESendIcon} from "../base/EIcon";
import ReactDOM from "react-dom";
import PublicPage from "./PublicPage";
import html2canvas from 'html2canvas'

const {Content} = Layout;

class ContentPage extends Component {

    constructor(props) {
        super(props);
        this.sharePublicRef = React.createRef();
    }

    handleFolderShare = (folderUID) => {
        this.shareModalRef.handleOpenShare(folderUID);
    };

    componentDidMount() {
        const folderNumber = this.props.params.folder_number
        if (isNotEmpty(folderNumber)) {
            this.props.actions.GetFolderNumberWorkspaceContent(folderNumber)
        } else {
            this.props.actions.GetDefaultFolder()
        }
    }

    componentDidUpdate(prevProps) {
        const folderNumber = this.props.params.folder_number
        if (folderNumber !== prevProps.params.folder_number) {
            this.props.actions.GetFolderNumberWorkspaceContent(folderNumber)
        }
    }


    handlePublishShare = (folderUID) => {
        this.publicModalRef.handleOpenPublic(folderUID)
        this.generateWorkspaceSnapshot((canvas) => {
                this.publicModalRef.handleSetOuterHtml(canvas.toDataURL('image/png', 0.9), folderUID)
            }
        )
    }

    async generateWorkspaceSnapshot(callback) {
        return this.handleGenerateWorkspaceSnapshot(callback)
    }

    handleGenerateWorkspaceSnapshot = (callback) => {
        const componentNode = ReactDOM.findDOMNode(this.sharePublicRef.current);

        // 克隆组件节点
        const node = componentNode.cloneNode(true);

        // 删除分享按钮
        const elementRemove1 = node.querySelector("#resource-copy-link")
        elementRemove1.remove()

        // 删除工作内容额外内容
        const elementRemove2 = node.querySelectorAll(".workspace-extra-btn-group")
        elementRemove2.forEach(elem => elem.remove());

        // 删除空间名称前的icon
        const elementRemove3 = node.querySelectorAll(".icon-bold")
        elementRemove3.forEach(elem => elem.remove());

        // 删除添加工作内容按钮
        const elementRemove4 = node.querySelector(".add-workspace-btn")
        elementRemove4.remove()

        // 展开所有组件
        const element = node.querySelectorAll(".ant-collapse-content")
        element.forEach(elem => {
            if (elem && elem.style.display === "none") {
                elem.style.display = '';
            }
            if (elem.classList.contains('ant-collapse-content-hidden')) {
                elem.classList.remove('ant-collapse-content-hidden');
            }
        });

        const target = document.querySelector('#content-workspaces');
        const elementWidth = target.offsetWidth;

        // 获取组件内容
        const component = document.createElement('div');
        component.innerHTML = node.outerHTML;
        document.body.appendChild(component);

        html2canvas(component, {
            useCORS: true,
            windowWidth: elementWidth,
        }).then(canvas => {
            canvas.style.width = "100%"
            canvas.style.height = "100%"
            callback(canvas)
        });
        document.body.removeChild(component);
    }

    handleUpdateCommunityImage = (folderUID) => {
        setTimeout(
            () => this.handleGenerateWorkspaceSnapshot((canvas) => {
                    this.props.actions.UpdatePublicCommunityImage(folderUID, canvas.toDataURL('image/png', 0.9))
                },
            ),
            5000
        )
    }

    render() {
        const {personalWorkspaces} = this.props.state
        if (isEmpty(personalWorkspaces.folder_uid)) {
            return <IndexPage/>
        }
        return (
            <Content style={{margin: '20px 0px 0', overflow: 'initial'}}>
                <div ref={this.sharePublicRef} id={"content-workspaces"} className="site-content-background"
                     style={{padding: '16px 24px 10px 24px'}}>
                    <Row>
                        <Col span={18}>
                            <Space>
                                <div className="resource-title">{personalWorkspaces.folder_name}</div>
                            </Space>
                        </Col>
                        <Col span={5}>
                            <div className="resource-link" id={"resource-copy-link"}>
                                <Button type="primary"
                                        shape="round"
                                        onClick={() => this.handleFolderShare(personalWorkspaces.folder_uid)}
                                        icon={<LinkOutlined/>}>分享</Button>
                            </div>
                        </Col>
                    </Row>
                    <Divider style={{margin: '12px 0'}}/>
                    <Row>
                        <Col span={1}/>
                        <Col span={22} className="workspace-content">
                            <WorkspacePage handleUpdateCommunityImage={this.handleUpdateCommunityImage}/>
                        </Col>
                        <Col span={1}/>
                    </Row>
                </div>

                <Affix style={{position: 'fixed', bottom: 120, right: 33}}>
                    {/*<div>*/}
                    {/*    <Tooltip placement="leftTop" title={"添加资源"}>*/}
                    {/*        <Button size="large"*/}
                    {/*                type="primary"*/}
                    {/*                shape="circle"*/}
                    {/*                onClick={this.handleGoodStationOpen}*/}
                    {/*                icon={<EGlobalAddIcon/>}/>*/}
                    {/*    </Tooltip>*/}
                    {/*</div>*/}
                    <div style={{marginTop: 10}}>
                        <Tooltip placement="leftTop" title={<span style={{fontSize: 12}}>发布到分享社区</span>}>
                            <Button size="large"
                                    type="primary"
                                    shape="circle"
                                    onClick={() => this.handlePublishShare(personalWorkspaces.folder_uid)}
                                    icon={<ESendIcon/>}
                            />
                        </Tooltip>
                    </div>
                </Affix>

                <PublicPage bindRef={(ref) => this.publicModalRef = ref}/>

                <SharePage bindRef={(ref) => this.shareModalRef = ref}/>
            </Content>
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.customer}), dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    }))(ContentPage));