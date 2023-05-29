import React, {Component} from 'react';
import {Timeline, Row, Col, Card, Typography, Layout, Divider, Button} from 'antd';
import FooterPage from "../footer/FooterPage";
import './style.css'

const { Title, Paragraph } = Typography;

const { Content } = Layout;

class UpdateLogPage extends Component {
    render() {
        return (
            <div>
                <Layout className={"site-update-log-layout"}>
                    <Content>
                        <Row>
                            <Col offset={6} span={2} style={{marginTop: 20}}>
                                <h2>
                                    更新日志
                                </h2>
                            </Col>
                            <Col offset={1} style={{marginTop: 20, marginLeft: 0}}>
                                <Button href="/" type="primary">返回首页</Button>
                            </Col>
                            <Col offset={6} span={12} style={{marginTop: 20}}>
                                <Timeline>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/05/05</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.2.0</Title>
                                                <Paragraph>
                                                    1. 好文精选新功能;<br/>
                                                    2. 手机端页面显示缩放;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/04/04</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.1.0</Title>
                                                <Paragraph>
                                                    1. 分享社区新功能;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/04/04</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.0.3</Title>
                                                <Paragraph>
                                                    1. 好站推荐栏目;<br/>
                                                    2. 优化首页侧边栏、分享权限Bug问题;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/03/29</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.0.2</Title>
                                                <Paragraph>
                                                    1. 支持本地文件资源上传功能;<br/>
                                                    2. 优化首页页面显示;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/03/08</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.0.1</Title>
                                                <Paragraph>
                                                    1. 支持好友邀请有礼功能;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item>
                                        <div style={{opacity: 0.5, marginBottom: 10}}>2023/02/22</div>
                                        <Card>
                                            <Typography>
                                                <Title level={5}>v1.0.0</Title>
                                                <Paragraph>
                                                    1. 支持新建工作空间功能;<br/>
                                                    2. 支持网页链接的收藏、分类等功能;<br/>
                                                    3. 支持工作空间链接分享进行好友共同协作功能;
                                                </Paragraph>
                                            </Typography>
                                        </Card>
                                    </Timeline.Item>
                                    <Timeline.Item color={"gray"}>
                                    </Timeline.Item>
                                </Timeline>
                            </Col>
                        </Row>
                    </Content>
                    <Divider style={{margin: 0}}/>
                    <FooterPage/>
                </Layout>
            </div>
        );
    }
}

export default UpdateLogPage;