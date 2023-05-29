import React, {Component} from 'react';
import {Avatar, Badge, Button, Dropdown, Empty, List, Popover, Tabs} from "antd";
import {BellOutlined} from "@ant-design/icons";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import './style.css'
import {EMessageAllDelIcon, EMessageAllReadIcon, EMessageSettingIcon} from "../base/EIcon";

class MessagePage extends Component {

    componentDidMount() {
        this.props.actions.GetMessages(1)
    }

    handleMessageChange = (key) => {
        this.props.actions.GetMessages(key)
    };

    handleReadMessage = (msgType, message) => {
        if (!message.has_read) {
            this.props.actions.ReadMessage(msgType, message.uid)
        }
    }

    handleReadAllMessage = () => {
        this.props.actions.ReadAllMessage(1)
    }

    handleDelMessage = (msgType, message) => {
        this.props.actions.DelMessage(msgType, message.uid)
    }

    handleDelAllMessage = () => {
        this.props.actions.DelAllMessage(1)
    }

    render() {

        const {messages} = this.props.state

        return (
            <div>
                <Popover
                    title={null}
                    overlayClassName={"user-message-popover"}
                    content={
                        <Tabs
                            className={"user-message-tabs"}
                            defaultActiveKey={"1"}
                            onChange={this.handleMessageChange}
                            tabBarExtraContent={{
                                right: <MessageMoreExtra
                                    handleReadAllMessage={this.handleReadAllMessage}
                                    handleDelAllMessage={this.handleDelAllMessage}
                                />
                            }}
                            items={[
                                {
                                    label: `未读消息`,
                                    key: "0",
                                    children: <List
                                        itemLayout="horizontal"
                                        dataSource={messages.user_messages}
                                        split={false}
                                        locale={{
                                            emptyText: <Empty description={"暂无消息"}
                                                              style={{marginTop: "50%"}}
                                                              image={Empty.PRESENTED_IMAGE_SIMPLE}/>
                                        }}
                                        renderItem={item => (
                                            <List.Item actions={[<a
                                                onClick={() => this.handleDelMessage(0, item)}
                                                className={"message-del-extra"}>删除</a>]}
                                                       onClick={() => this.handleReadMessage(0, item)}>
                                                <List.Item.Meta
                                                    avatar={<Avatar src={item.promoter_image}/>}
                                                    title={<span
                                                        className={item.has_read ? "message-read-title" : "message-unread-title"}>{item.promoter_name}</span>}
                                                    description={
                                                        <div
                                                            className={item.has_read ? "message-read-description" : "message-unread-description"}>
                                                            <div>
                                                                {item.description}
                                                            </div>
                                                            <div>{item.created_at}</div>
                                                        </div>
                                                    }
                                                />
                                            </List.Item>
                                        )}
                                    />,
                                },
                                {
                                    label: `全部消息`,
                                    key: "1",
                                    children: <List
                                        itemLayout="horizontal"
                                        dataSource={messages.user_messages}
                                        split={false}
                                        locale={{
                                            emptyText: <Empty description={"暂无消息"}
                                                              style={{marginTop: "50%"}}
                                                              image={Empty.PRESENTED_IMAGE_SIMPLE}/>
                                        }}
                                        renderItem={item => (
                                            <List.Item actions={[<a
                                                onClick={() => this.handleDelMessage(1, item)}
                                                className={"message-del-extra"}>删除</a>]}
                                                       onClick={() => this.handleReadMessage(1, item)}>
                                                <List.Item.Meta
                                                    avatar={<Avatar src={item.promoter_image}/>}
                                                    title={<span
                                                        className={item.has_read ? "message-read-title" : "message-unread-title"}>{item.promoter_name}</span>}
                                                    description={
                                                        <div
                                                            className={item.has_read ? "message-read-description" : "message-unread-description"}>
                                                            <div>
                                                                {item.description}
                                                            </div>
                                                            <div>{item.created_at}</div>
                                                        </div>
                                                    }
                                                />
                                            </List.Item>
                                        )}
                                    />,
                                }
                            ]}
                        />
                    }
                    placement="bottom"
                    trigger="click"
                    style={{width: "200px"}}>
                    <Badge dot={messages.unread !== 0} offset={[-10, 7]}>
                        <Avatar
                            style={{color: '#848B96', backgroundColor: '#ffffff'}}
                            icon={<BellOutlined/>}/>
                    </Badge>
                </Popover>
            </div>
        );
    }
}

const MessageMoreExtra = (props) => {

    const items = [
        {
            label: '全部已读',
            key: 'item-1',
            icon: <EMessageAllReadIcon/>,
            onClick: () => {
                props.handleReadAllMessage()
            }
        },
        {
            label: '全部删除',
            key: 'item-2',
            icon: <EMessageAllDelIcon/>,
            onClick: () => {
                props.handleDelAllMessage()
            }
        }
    ];

    return <Dropdown
        overlayClassName={"message-menu-extra"}
        placement="bottomRight"
        autoAdjustOverflow={false}
        overlayStyle={{
            width: 160,
            height: 100
        }}
        trigger={["click"]}
        menu={{items}}>
        <Button type="text"
                size={"small"}
                icon={<EMessageSettingIcon/>}/>
    </Dropdown>
}

export default connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(MessagePage);