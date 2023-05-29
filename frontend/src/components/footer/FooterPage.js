import React, {Component} from 'react';
import {Layout, Space} from 'antd';
import FeedbackPage from "../feedback/FeedbackPage";

const {Footer} = Layout;

class FooterPage extends Component {

    handleOpenFeedback = () => {
        this.feedbackModalRef.handleOpenFeedback()
    }

    render() {
        return (
            <Footer style={{textAlign: 'center', background: '#ffffff', opacity: 0.5, fontSize: 13}}>
                <div>
                    <Space size={13}>
                        <a href="/" style={{color: "black"}}>Tab 精灵 ©2022-2023</a>
                        <a href="/update/log" style={{color: "black"}} target="_blank">更新日志</a>
                        <a onClick={this.handleOpenFeedback} style={{color: "black"}}>问题反馈</a>
                        <a href="/disclaimer" style={{color: "black"}} target="_blank">免责声明与条款</a>
                        <a href="/business" style={{color: "black"}} target="_blank">商务合作</a>
                    </Space>
                </div>
                <div style={{marginTop: 3}}>
                    <a style={{color: "black"}} href="https://beian.miit.gov.cn" target="_blank">豫ICP备案号2023002238</a>
                </div>

                <FeedbackPage bindRef={(ref) => this.feedbackModalRef = ref}/>
            </Footer>
        );
    }
}

export default FooterPage;