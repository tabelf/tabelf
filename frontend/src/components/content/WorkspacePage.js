import React, {Component} from "react";
import {ReactSortable} from "react-sortablejs";
import {Avatar, Button, Card, Collapse, Dropdown, Empty, Form, Input, List, message, Modal, Space, Tooltip} from 'antd';
import {DeleteOutlined, EditOutlined, RightOutlined} from '@ant-design/icons';
import './style.css'
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {
    EAddIcon, ECSVIcon,
    EDelIcon,
    EEditIcon,
    EExcelIcon,
    EJPGIcon,
    EMoreIcon, EMP3Icon, EMP4Icon,
    EPDFIcon,
    EPPTIcon,
    EWordIcon
} from "../base/EIcon";
import {downloadFile} from "../../actions/Base";
import DefaultNetwork from '../../assets/default_network.png'
import AddResourcePage from "./AddResourcePage";
import ReaderPage from "./ReaderPage";

const {Panel} = Collapse;

class WorkspacePage extends Component {

    state = {
        hasOpenWebLinkForm: false,
        workspaceUID: '',
        workspaceName: '',
        hasExpandActive: false,
        showAddResourceType: 'url',
        showAddResourceModal: false,
        resourceAction: ''
    }

    constructor(props) {
        super(props);
        this.webLinkForm = React.createRef();
        this.editWebLinkRef = React.createRef();
        this.editWorkspaceRef = React.createRef();
    }

    handleDeleteWorkspace = (workspaceUID) => {
        const {personalWorkspaces} = this.props.state
        this.props.actions.DeleteWorkspace(
            personalWorkspaces.folder_uid,
            workspaceUID,
        )

        this.handleUpdateCommunityImage(personalWorkspaces.folder_uid)
    }

    handleAddWorkspace = () => {
        const {personalWorkspaces} = this.props.state
        this.props.actions.AddWorkspace(
            personalWorkspaces.folder_uid,
            "未命名",
        )

        this.handleUpdateCommunityImage(personalWorkspaces.folder_uid)
    }

    genExtra = (folderUID, workspace) => {
        return <Space size={'small'} className={"workspace-extra-btn-group"}>
            <Tooltip placement="topLeft" title={<span style={{fontSize: 12}}>添加资源</span>}>
                <Button type="text"
                        shape="circle"
                        icon={<EAddIcon/>}
                        onClick={() => {
                            this.addResourceModalRef.handleOpenAddResource(folderUID, workspace.workspace_uid, '', false)
                        }}/>
            </Tooltip>
            <Tooltip placement="topLeft" title={<span style={{fontSize: 12}}>重命名</span>}>
                <Button type="text"
                        shape="circle"
                        onClick={() => EditWorkspace(this.props.actions, folderUID, workspace, this.editWorkspaceRef, this.handleUpdateCommunityImage)}
                        icon={<EEditIcon/>}
                />
            </Tooltip>
            <Tooltip placement="topLeft" title={<span style={{fontSize: 12}}>删除</span>}>
                <Button type="text"
                        shape="circle"
                        icon={<EDelIcon/>}
                        onClick={() => DeleteWorkspace(this.handleDeleteWorkspace, workspace.workspace_uid, folderUID, this.handleUpdateCommunityImage)}/>
            </Tooltip>
        </Space>
    }

    handlePanelClick = (newKey) => {
        const {personalWorkspaces} = this.props.state
        this.props.actions.UpdateWorkspaceSwitch(personalWorkspaces.folder_uid, newKey)
    }

    // 工作空间内的元素发生变动调用该方法
    handleWorkspaceItemInnerMove = (e) => {
        const {personalWorkspaces} = this.props.state
        const workspace = personalWorkspaces.personal_workspaces.find(w => w.workspace_uid === e.from.id)
        let webLinks = workspace.web_links.filter((link, index) => index !== e.oldIndex)
        webLinks.splice(e.newIndex, 0, workspace.web_links[e.oldIndex])
        webLinks.map((link, index) => {
            link.sequence = index
        })
        workspace.web_links = webLinks

        this.props.actions.UpdateWorkspaceWebLinks(
            personalWorkspaces.folder_uid,
            workspace,
            workspace
        )

        this.handleUpdateCommunityImage(personalWorkspaces.folder_uid)
    }

    // 跨工作空间的元素发生变动调用该方法
    handleWorkspaceItemOuterMove = (e) => {
        const {personalWorkspaces} = this.props.state
        const workspaces = personalWorkspaces.personal_workspaces.filter(w => true)
        const fromWorkspace = workspaces.find(w => w.workspace_uid === e.from.id)
        const toWorkspace = workspaces.find(w => w.workspace_uid === e.to.id)
        let webLinks = toWorkspace.web_links.filter(link => true)
        webLinks.splice(e.newIndex, 0, fromWorkspace.web_links[e.oldIndex])
        webLinks.map((link, index) => (
            link.sequence = index
        ))
        toWorkspace.web_links = webLinks
        fromWorkspace.web_links.splice(e.oldIndex, 1)

        this.props.actions.UpdateWorkspaceWebLinks(
            personalWorkspaces.folder_uid,
            toWorkspace,
            fromWorkspace
        )

        this.handleUpdateCommunityImage(personalWorkspaces.folder_uid)
    }

    handleDelWebLink = (folderUID, workspaceUID, linkUID) => {
        this.props.actions.DeletePersonalWebLink(folderUID, workspaceUID, linkUID)
        this.handleUpdateCommunityImage(folderUID)
    }

    handleOpenReader = (link) => {
        this.readerModalRef.handleOpenReader(link)
    }

    handleUpdateCommunityImage = (folderUID) => {
        this.props.handleUpdateCommunityImage(folderUID)
    }

    render() {
        const {personalWorkspaces} = this.props.state
        return (
            <div>
                <Collapse ghost
                          activeKey={personalWorkspaces.active_workspace_uids}
                          onChange={this.handlePanelClick}
                          expandIcon={({isActive}) => {
                              return <RightOutlined style={{marginTop: "7px"}}
                                                    className="icon-bold"
                                                    rotate={isActive ? 90 : 0}/>
                          }}>
                    {
                        personalWorkspaces.personal_workspaces.map(w => (
                            <Panel header={<h3>{w.workspace_name}</h3>}
                                   key={w.workspace_uid}
                                   forceRender={true}
                                   extra={w.is_open ? this.genExtra(personalWorkspaces.folder_uid, w) : <></>}
                                   collapsible={'header'}>
                                <Card className="workspace">
                                    {
                                        <ReactSortable
                                            id={w.workspace_uid}
                                            group={"shared"}
                                            list={w.web_links}
                                            setList={(newList) => {
                                                w.web_links = newList
                                            }}
                                            filter=".empty-web-link-filter"
                                            onUpdate={this.handleWorkspaceItemInnerMove}
                                            onAdd={this.handleWorkspaceItemOuterMove}>
                                            {
                                                w.web_links.length === 0 ? (
                                                    <div className="empty-web-link-filter">
                                                        <EmptyWebLink/>
                                                    </div>
                                                ) : (
                                                    w.web_links.map(link => (
                                                        <div key={link.link_uid} className="weblink-list">
                                                            <List.Item actions={[
                                                                <WebLinkMoreExtra
                                                                    handleEditWebLink={() => EditWebLink(this.props, personalWorkspaces.folder_uid, w.workspace_uid, link, this.editWebLinkRef, this.handleUpdateCommunityImage)}
                                                                    handleDelWebLink={() => this.handleDelWebLink(personalWorkspaces.folder_uid, w.workspace_uid, link.link_uid)}
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
                                                                                    default: return <img style={{height: "20px", width: "20px"}}  crossOrigin="anonymous" src={DefaultNetwork}/>
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
                                        </ReactSortable>
                                    }
                                </Card>
                            </Panel>
                        ))
                    }
                </Collapse>

                <div className="add-workspace-btn">
                    <Button block onClick={this.handleAddWorkspace}>✨ 创建新的书签栏 ✨</Button>
                </div>

                <AddResourcePage handleUpdateCommunityImage={this.handleUpdateCommunityImage}
                                 bindRef={(ref) => this.addResourceModalRef = ref}/>

                <ReaderPage bindRef={(ref) => this.readerModalRef = ref}/>
            </div>
        );
    }
}

const DeleteWorkspace = (handleDeleteWorkspace, workspaceUID, folderUID, updateCommunityImage) => {
    return Modal.confirm({
        title: "删除操作",
        okText: '确认',
        cancelText: '取消',
        width: 386,
        height: 172,
        bodyStyle: {padding: 24},
        className: "add-folder",
        content: "确认要删除该书签拦下的所有内容吗?",
        onOk: () => {
            handleDeleteWorkspace(workspaceUID)
            updateCommunityImage(folderUID)
        }
    });
}


const WebLinkMoreExtra = ({handleEditWebLink, handleDelWebLink}) => {

    const items = [
        {
            label: '编辑链接',
            key: 'item-1',
            icon: <EditOutlined/>,
            onClick: () => {
                handleEditWebLink()
            }
        },
        {
            label: '删除链接',
            key: 'item-2',
            icon: <DeleteOutlined/>,
            onClick: () => {
                handleDelWebLink()
            }
        }
    ];

    return <Dropdown
        overlayClassName={"workspace-menu-extra"}
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
                className={"weblink-menu-btn"}
                size={"small"}
                icon={<EMoreIcon/>}/>
    </Dropdown>
}

const EmptyWebLink = () => (
    <Empty style={{margin: "10px 0"}}
           image={Empty.PRESENTED_IMAGE_SIMPLE}
           description={"在此处添加链接"}
    />
)

const EditWebLink = (props, folderUID, workspaceUID, link, ref, updateCommunityImage) => {
    return Modal.confirm({
        icon: null,
        title: <h4>编辑链接</h4>,
        okText: '确认',
        cancelText: '取消',
        width: 470,
        height: 200,
        className: "add-weblink",
        content: (
            <div>
                <Form style={{marginTop: 20}} name="edit-weblink-form" ref={ref} initialValues={{
                    "title": link.title,
                    "description": link.description,
                }}>
                    <Form.Item label={<span>标题</span>} name="title">
                        <Input/>
                    </Form.Item>
                    <Form.Item label={<span>描述</span>} name="description">
                        <Input/>
                    </Form.Item>
                </Form>
            </div>
        ),
        onOk: () => {
            const data = ref.current.getFieldsValue()
            props.actions.UpdatePersonalWebLink(folderUID, workspaceUID, {
                ...data,
                "link_uid": link.link_uid,
            }, () => ref.current.resetFields())
            updateCommunityImage(folderUID)
        },
        onCancel: () => {
            ref.current.resetFields()
        }
    });
}

const EditWorkspace = (actions, folderUID, workspace, ref, updateCommunityImage) => {
    return Modal.confirm({
        icon: null,
        title: <h4>修改书签栏名称</h4>,
        okText: '确认',
        cancelText: '取消',
        width: 386,
        height: 172,
        bodyStyle: {
            padding: 24,
        },
        className: "add-folder",
        content: (
            <div>
                <Form style={{marginTop: 20}} initialValues={{
                    workspace_name: workspace.workspace_name,
                }} ref={ref}>
                    <Form.Item label={<span>名称</span>} name="workspace_name">
                        <Input placeholder="请输入名称"/>
                    </Form.Item>
                </Form>
            </div>
        ),
        onOk: () => {
            const data = ref.current.getFieldsValue()
            actions.UpdateWorkspace(
                folderUID,
                workspace.workspace_uid,
                data.workspace_name,
            )
            updateCommunityImage(folderUID)
        },
        onCancel: () => {
        }
    });
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(WorkspacePage);