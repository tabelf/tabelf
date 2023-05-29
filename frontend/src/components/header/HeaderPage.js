import React, {Component} from 'react';
import {SearchOutlined, UserOutlined} from '@ant-design/icons';
import {
    Avatar,
    Button,
    Col,
    Empty,
    Input,
    Layout,
    List,
    Modal,
    Row,
    Space,
    Tabs,
} from 'antd';
import './style.css'
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import MessagePage from "./MessagePage";
import UserCenterPage from "./UserCenterPage";

const {Header} = Layout;

class HeaderPage extends Component {

    state = {
        hasSearchInfo: false,
        searchKeyword: '',
        activeTab: 0,
    }

    handleOpenSearch = () => {
        this.setState({hasSearchInfo: true})
    }

    handleCloseSearch = () => {
        this.props.actions.SearchWebLinks('', this.state.activeTab)
        this.setState({
            hasSearchInfo: false,
            searchKeyword: '',
            activeTab: 0
        })
    }

    handleSearchChange = ({target: {value}}) => {
        this.setState({searchKeyword: value})
        this.props.actions.SearchWebLinks(value, this.state.activeTab)
    }

    handleTabsChange = (key) => {
        this.setState({activeTab: key})
        this.props.actions.SearchWebLinks(this.state.searchKeyword, key)
    };

    render() {
        const {searchLinks} = this.props.state
        const tabList = (
            <List
                itemLayout="horizontal"
                dataSource={searchLinks.search_web_links}
                locale={{
                    emptyText: <Empty description={"内容为空"}
                                      image={Empty.PRESENTED_IMAGE_SIMPLE}/>
                }}
                split={false}
                renderItem={item => (
                    <List.Item>
                        <List.Item.Meta
                            avatar={<Avatar shape="square" size="small" icon={<UserOutlined/>}/>}
                            title={<a href={"/content/0/" + item.folder_number}>
                                <span>{item.title}</span>
                                <span style={{opacity: 0.5, fontWeight: 400}}>&nbsp;-&nbsp;{item.description}</span>
                            </a>}
                        />
                    </List.Item>
                )}
            />
        )
        return (
            <Header className="site-layout-background">
                <Row>
                    <Col span={18}>
                        <div className="header-search">
                            <Button
                                onClick={this.handleOpenSearch}
                                shape="round"
                                block>
                                <SearchOutlined/>
                                搜索书签页
                            </Button>
                        </div>
                    </Col>
                    <Col span={6}>
                        <div className="header-user">
                            <Space size={"large"}>
                                <MessagePage />
                                <UserCenterPage/>
                            </Space>

                            {/*<Button type="text"*/}
                            {/*        shape="round"*/}
                            {/*        href={"/login"}>*/}
                            {/*    登录 / 注册*/}
                            {/*</Button>*/}
                        </div>
                    </Col>
                </Row>

                <Modal title={null}
                       open={this.state.hasSearchInfo}
                       footer={null}
                       forceRender={true}
                       width={720}
                       onCancel={this.handleCloseSearch}
                       className={"search-content"}>
                    <div className={"ant-modal-confirm-content"}>
                        <Input prefix={<SearchOutlined/>}
                               onChange={this.handleSearchChange}
                               placeholder="搜索文件"
                               size="large"
                               value={this.state.searchKeyword}
                               bordered={false}/>
                        <div className={"search-content-tab"}>
                            <Tabs
                                activeKey={this.state.activeTab}
                                onChange={this.handleTabsChange}
                                items={[
                                    {
                                        label: `全部文件`,
                                        key: 0,
                                        children: tabList,
                                    },
                                    {
                                        label: `我的文件`,
                                        key: 1,
                                        children: tabList,
                                    },
                                    {
                                        label: `与我协作`,
                                        key: 2,
                                        children: tabList,
                                    },
                                ]}
                            />
                        </div>
                    </div>
                </Modal>
            </Header>
        );
    }
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(HeaderPage);