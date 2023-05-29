import React, {Component} from 'react';
import {Button, Col, Divider, Layout, Row} from 'antd';
import './style.css'
import CollaborationWorkspacePage from "./CollaborationWorkspacePage";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {LinkOutlined} from "@ant-design/icons";
import SharePage from "../share/SharePage";
import {withRouter} from "../base/Base";
import {isEmpty} from "../../actions/Base";
import IndexPage from "../home/IndexPage";

const {Content} = Layout;


class CollaborationPage extends Component {

    state = {
        isModalOpen: false
    }

    handleCancel = () => {
        this.setState({isModalOpen: false})
    };

    handleFolderShare = (folderUID) => {
        this.shareModalRef.handleOpenShare(folderUID);
    };

    componentDidMount() {
        const folderNumber = this.props.params.folder_number
        this.props.actions.GetCollaborationWorkspaceContent(folderNumber)
    }

    componentDidUpdate(prevProps) {
        const folderNumber = this.props.params.folder_number
        if (folderNumber !== prevProps.params.folder_number) {
            this.props.actions.GetCollaborationWorkspaceContent(folderNumber)
        }
    }

    render() {
        const {personalWorkspaces} = this.props.state
        if (isEmpty(personalWorkspaces.folder_uid)) {
            return <IndexPage/>
        }
        return (
            <Content style={{margin: '20px 0px 0', overflow: 'initial'}}>
                <div className="site-collaboration-background" style={{padding: '16px 24px 10px 24px'}}>
                    <Row>
                        <Col span={18}>
                            <div className="resource-title">{personalWorkspaces.folder_name}</div>
                        </Col>
                        <Col span={5}>
                            <div className="resource-link">
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
                        <Col span={22} className="collaboration-content">
                            <CollaborationWorkspacePage/>
                        </Col>
                        <Col span={1}/>
                    </Row>
                </div>

                <SharePage bindRef={(ref) => this.shareModalRef = ref}/>
            </Content>
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(CollaborationPage));