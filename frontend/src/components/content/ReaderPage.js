import React, {Component, lazy, Suspense} from 'react';
import {Button, Col, Modal, Row} from "antd";
import {CloudDownloadOutlined} from '@ant-design/icons';
import {downloadFile} from "../../actions/Base";

const LazyFileViewer = lazy(() => import('react-file-viewer'));

class ReaderPage extends Component {

    state = {
        hasOpenReader: false,
        link: {}
    }

    componentDidMount() {
        this.props.bindRef(this)
    }

    handleOpenReader = (link) => {
        this.setState({
            hasOpenReader: true,
            link: link
        })
    }

    handleCloseReader = () => {
        this.setState({hasOpenReader: false})
    }

    render() {
        return (
            <div>
                <Modal title={
                            <Row style={{alignItems: "center"}}>
                                <Col span={18}>
                                    {this.state.link.title}
                                </Col>
                                <Col>
                                    <Button type="primary"
                                            shape="round"
                                            onClick={() => downloadFile(this.state.link.link, this.state.link.title)}
                                            icon={<CloudDownloadOutlined />}>下载文件</Button>
                                </Col>
                            </Row>
                        }
                       open={this.state.hasOpenReader}
                       footer={null}
                       onCancel={this.handleCloseReader}
                       className={"reader-modal"}
                       width={710}>
                    <Suspense fallback={<div>加载中...</div>}>
                        {
                            this.state.link.file_type === "pdf" ? (
                                <iframe
                                    src={this.state.link.link}
                                    style={{ width: '100%', height: '100%' }}
                                    frameBorder="0"
                                />
                            ) : (
                                <LazyFileViewer fileType={this.state.link.file_type}
                                                errorComponent={<div>加载失败</div>}
                                                filePath={this.state.link.link}
                                                unsupportedComponent={() => {
                                                    return <div className='not_support'>
                                                        加载失败, 不支持
                                                    </div>
                                                }}
                                />
                            )
                        }
                    </Suspense>
                </Modal>
            </div>
        );
    }
}

export default ReaderPage;