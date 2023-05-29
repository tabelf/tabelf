import {
    CopyOutlined,
    DeleteOutlined,
    EditOutlined,
    FolderOutlined,
    HistoryOutlined,
    LinkOutlined,
    PlusOutlined,
    TeamOutlined,
    LikeOutlined,
    ShareAltOutlined,
    CompassOutlined
} from '@ant-design/icons';
import React, {Component} from 'react';
import {Breadcrumb, Button, Dropdown, Form, Input, Layout, Menu, Modal, Progress, message, Card, Image} from 'antd';
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {ECloseShareIcon, EHotIcon, EMoreIcon} from "../base/EIcon";
import {NavLink as Link} from 'react-router-dom';
import './style.css'
import UpgradePage from "../recharge/UpgradePage";
import SharePage from "../share/SharePage";
import LogoImage from '../../assets/logo.png'
import {isEmpty} from "../../actions/Base";

const {Sider} = Layout;

class SidebarPage extends Component {

    state = {
        isModalOpen: false,
        hasUpgradeOpen: false
    }

    handleUpgradeOpen = () => {
        this.upgradeModalRef.handleUpgradeOpen()
    };

    constructor(props) {
        super(props);
        this.folderRef = React.createRef();
    }

    handleCancel = () => {
        this.setState({isModalOpen: false})
    };

    componentDidMount() {
        this.props.actions.GetPersonalFolders()
        this.props.actions.GetSharePersonalFolders()
        this.props.actions.GetPersonalWebLinkInfo()
    }

    handleMenuClick = (item) => {
        this.props.actions.SetMenuItemSelectKeys(item.keyPath)
    }

    handleSubMenuOpenChange = (openKeys) => {
        this.props.actions.SetSubMenuOpenKeys(openKeys)
    }

    handleFolderShare = (folderUID) => {
        this.shareModalRef.handleOpenShare(folderUID);
    };

    render() {
        const {
            personalFolders,
            sharePersonalFolders,
            menuActive,
            personalWebLink,
        } = this.props.state
        return (
                <Sider style={{
                            overflow: 'auto',
                            height: '100vh',
                            position: 'fixed',
                            left: 0,
                            top: 0,
                            bottom: 0,
                            borderRight: '1px solid #e9edf2',
                        }}
                       width={320}
                       theme="light">
                    <div className="logo">
                        <img src={LogoImage} alt="logo" className="logo-image"/>
                    </div>
                    <div className="tab-create">
                        <div className="tab-create-button">
                            <Button
                                icon={<PlusOutlined/>}
                                block={true}
                                onClick={() => NewAddFolder(this.props.actions, this.folderRef)}
                                type="primary">新建</Button>
                        </div>
                    </div>
                    <div className="custom-space">
                        <Breadcrumb className="breadcrumb-space">
                            <Breadcrumb.Item>
                                <span>切换空间</span>
                            </Breadcrumb.Item>
                            <Breadcrumb.Item
                                // overlay={
                                //     <Menu>
                                //         <Menu.Item>
                                //             <span>个人空间</span>
                                //         </Menu.Item>
                                //         {
                                //             personalTeams.personal_teams.map(t => {
                                //                 return <Menu.Item>
                                //                     <span>{t.team_name}</span>
                                //                 </Menu.Item>
                                //             })
                                //         }
                                //         <Menu.Divider/>
                                //         <Menu.Item icon={<PlusOutlined/>}>创建新团队</Menu.Item>
                                //     </Menu>
                                // }
                            >
                                <span>个人空间</span>
                            </Breadcrumb.Item>
                        </Breadcrumb>
                    </div>
                    <div className="user-menu-inline">
                        <Menu mode="inline"
                              forceSubMenuRender={true}
                              openKeys={menuActive.openKey}
                              selectedKeys={menuActive.selectKey}
                              onClick={this.handleMenuClick}
                              onOpenChange={this.handleSubMenuOpenChange}
                        >
                            <Menu.Item key="my-history" icon={<HistoryOutlined/>}>
                                <Link to="/history">最近</Link>
                            </Menu.Item>
                            <Menu.SubMenu key="my-files" icon={<FolderOutlined/>} title="我的文件">
                                {
                                    personalFolders.folders.map(f => {
                                        return <Menu.Item key={f.folder_uid} icon={<FolderOutlined/>}>
                                        <span>
                                           <Link to={"/content/0/" + f.folder_number}>
                                               {f.folder_name}
                                           </Link>
                                        </span>
                                            <span className="sidebar-menu-suffix">
                                            <FolderMoreExtra
                                                actions={this.props.actions}
                                                folder={f}
                                                handleFolderShare={this.handleFolderShare}
                                                handleUpgradeOpen={this.handleUpgradeOpen}
                                            />
                                        </span>
                                        </Menu.Item>
                                    })
                                }
                            </Menu.SubMenu>
                            <Menu.SubMenu key="my-share" icon={<TeamOutlined/>} title="与我协作">
                                {
                                    sharePersonalFolders.folders.map(f => {
                                        return <Menu.Item icon={<FolderOutlined/>} key={"share" + f.folder_uid}>
                                        <span>
                                            <Link to={"/collaboration/0/"+f.folder_number}>{f.folder_name}</Link>
                                        </span>
                                            <span className="sidebar-menu-suffix">
                                            <CollaborationFolderMoreExtra
                                                actions={this.props.actions}
                                                folder={f}
                                                handleFolderShare={this.handleFolderShare}
                                                handleUpgradeOpen={this.handleUpgradeOpen}
                                            />
                                        </span>
                                        </Menu.Item>
                                    })
                                }
                            </Menu.SubMenu>
                            <Menu.Divider/>
                            <Menu.Item key="my-good-article" icon={<CompassOutlined />}>
                                <Link to="/good/article" target={"_blank"}>好文精选&nbsp;<EHotIcon/></Link>
                            </Menu.Item>
                            <Menu.Item key="my-community" icon={<ShareAltOutlined />}>
                                <Link to="/community" target={"_blank"}>分享社区</Link>
                            </Menu.Item>
                            <Menu.Item key="my-good-station" icon={<LikeOutlined />}>
                                <Link to="/recommend" target={"_blank"}>好站推荐</Link>
                            </Menu.Item>
                            <Menu.Divider/>
                            <Menu.Item key="my-recycle" icon={<DeleteOutlined/>}>
                                <Link to="/recycle">回收站</Link>
                            </Menu.Item>
                        </Menu>
                    </div>

                    <SharePage bindRef={(ref) => this.shareModalRef = ref}/>

                    <UpgradePage bindRef={(ref) => this.upgradeModalRef = ref}/>
                </Sider>
        );
    }
}

const NewAddFolder = (actions, ref) => {
    return Modal.confirm({
        icon: null,
        title: <h4>新建文件夹</h4>,
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
                <Form style={{marginTop: 20}} ref={ref}>
                    <Form.Item label={<span>名称</span>} name="folder_name">
                        <Input placeholder="请输入名称"/>
                    </Form.Item>
                </Form>
            </div>
        ),
        onOk: () => {
            const data = ref.current.getFieldsValue()
            if (isEmpty(data.folder_name) || data.folder_name.trim() === '') {
                message.error("请输入名称")
                return;
            }
            actions.AddPersonalFolder(data, () => ref.current.resetFields())
        },
        onCancel: () => {
        }
    });
}

const DeleteFolder = (actions, folderUID) => {
    return Modal.confirm({
        title: "删除操作",
        okText: '确认',
        cancelText: '取消',
        width: 386,
        height: 172,
        bodyStyle: {padding: 24},
        className: "add-folder",
        content: "确认要删除该文件下的所有内容吗?",
        onOk: () => {
            actions.DeletePersonalFolder(folderUID)
        }
    });
}

const ExitCollFolder = (actions, shareUID) => {
    return Modal.confirm({
        title: "退出操作",
        okText: '确认',
        cancelText: '取消',
        width: 386,
        height: 172,
        bodyStyle: {padding: 24},
        className: "add-folder",
        content: "确认要退出当前文件协作的内容吗?",
        onOk: () => {
            actions.ExitSharePersonalFolder(shareUID)
        }
    });
}


const RenameFolder = (actions, folder, ref) => {
    return Modal.confirm({
        icon: null,
        title: <h4>重命名</h4>,
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
                    folder_name: folder.folder_name,
                }} ref={ref}>
                    <Form.Item label={<span>名称</span>} name="folder_name">
                        <Input placeholder="请输入名称"/>
                    </Form.Item>
                </Form>
            </div>
        ),
        onOk: () => {
            const data = ref.current.getFieldsValue()
            actions.UpdatePersonalFolder(
                folder.folder_uid,
                data,
                () => ref.current.resetFields(),
            )
        },
        onCancel: () => {
        }
    });
}

const CollaborationFolderMoreExtra = (props) => {
    const renameFolderRef = React.createRef();

    const items = [];

    if (props.folder.authority === "1") {
        items.push(
            {
                label: '分享协作',
                key: 'item-3',
                icon: <LinkOutlined/>,
                onClick: () => {
                    props.handleFolderShare(props.folder.folder_uid)
                }
            },
            {
                label: '重命名',
                key: 'item-1',
                icon: <EditOutlined/>,
                onClick: () => RenameFolder(
                    props.actions,
                    props.folder,
                    renameFolderRef
                ),
            },
            {
                label: '复制到我的文件',
                key: 'item-2',
                icon: <CopyOutlined/>,
                onClick: () => {
                    props.actions.CopySharePersonalFolder(
                        props.folder.share_uid,
                        props.handleUpgradeOpen
                    )
                }
            },
            {
                label: '退出协作',
                key: 'item-4',
                icon: <ECloseShareIcon/>,
                onClick: () => {
                    ExitCollFolder(props.actions, props.folder.share_uid)
                }
            }
        )
    } else if (props.folder.authority === "0")  {
        items.push(
            {
                label: '分享协作',
                key: 'item-3',
                icon: <LinkOutlined/>,
                onClick: () => {
                    props.handleFolderShare(props.folder.folder_uid)
                }
            },
            {
                label: '复制到我的文件',
                key: 'item-2',
                icon: <CopyOutlined/>,
                onClick: () => {
                    props.actions.CopySharePersonalFolder(
                        props.folder.share_uid,
                        props.handleUpgradeOpen
                    )
                }
            },
            {
                label: '退出协作',
                key: 'item-4',
                icon: <ECloseShareIcon/>,
                onClick: () => {
                    ExitCollFolder(props.actions, props.folder.share_uid)
                }
            }
        )
    }

    return <Dropdown
        overlayClassName={"folder-menu-extra"}
        autoAdjustOverflow={false}
        overlayStyle={{
            width: 160,
            height: 280
        }}
        trigger={["click"]}
        menu={{items}}>
        <Button type="text"
                shape="circle"
                className={"sidebar-menu-btn"}
                size={"small"}
                icon={<EMoreIcon/>}/>
    </Dropdown>
}


const FolderMoreExtra = (props) => {
    const renameFolderRef = React.createRef();

    const items = [
        {
            label: '重命名',
            key: 'item-1',
            icon: <EditOutlined/>,
            onClick: () => RenameFolder(
                props.actions,
                props.folder,
                renameFolderRef
            ),
        },
        {
            label: '复制',
            key: 'item-2',
            icon: <CopyOutlined/>,
            onClick: () => {
                props.actions.CopyPersonalFolder(props.folder.folder_uid, props.handleUpgradeOpen)
            }
        },
        {
            label: '分享协作',
            key: 'item-3',
            icon: <LinkOutlined/>,
            onClick: () => {
                props.handleFolderShare(props.folder.folder_uid)
            }
        },
        {
            label: '删除',
            key: 'item-4',
            icon: <DeleteOutlined/>,
            onClick: () => DeleteFolder(props.actions, props.folder.folder_uid)
        },
    ];

    return <Dropdown
        overlayClassName={"folder-menu-extra"}
        autoAdjustOverflow={false}
        overlayStyle={{
            width: 160,
            height: 280
        }}
        trigger={["click"]}
        menu={{items}}>
        <Button type="text"
                shape="circle"
                className={"sidebar-menu-btn"}
                size={"small"}
                icon={<EMoreIcon/>}/>
    </Dropdown>
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(SidebarPage);