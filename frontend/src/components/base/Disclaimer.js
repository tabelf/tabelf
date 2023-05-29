import React, {Component} from 'react';
import {Typography, Row, Col, Divider} from 'antd';
import FooterPage from "../footer/FooterPage";

const { Title, Paragraph } = Typography;

class Disclaimer extends Component {
    render() {
        return (
            <div>
                <Row>
                    <Col offset={6} span={12} style={{marginTop: 50, marginBottom: 100}}>
                        <Typography>
                            <Title level={2}>免责声明与条款</Title>
                            <Paragraph style={{fontSize: 15}}>
                                欢迎访问本网站，我们是一家提供网页地址收藏服务的网站。在使用本网站之前，请您仔细阅读并理解下列条款和条件，这些条款和条件适用于所有用户，并构成您与本网站之间的协议。
                            </Paragraph>
                            <Title level={3}>网站服务说明</Title>
                            <Paragraph style={{fontSize: 15}}>
                                1. 本网站作为一个网址收藏类型网站，提供链接并引导用户访问第三方网站。我们不直接提供任何产品或服务，也不对这些第三方网站的内容、服务、及其可用性等方面承担任何明示或暗示的保证或责任。在使用本网站时，您应当自行判断和评估所有第三方网站的信息和内容，并自行承担因使用或无法使用该网站所产生的风险和后果。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                2. 我们致力于提供最好的网址收藏服务，但我们没有义务监控用户发布的内容或行为，也不能保证所有内容都是准确、完整、真实、合法和适当的。我们保留随时修改、暂停、中止或停止本网站任何部分的权利，包括但不限于对其进行更新和对其内容进行更改、删除、添加甚至关闭本网站。
                            </Paragraph>
                            <Title level={3}>免责声明</Title>
                            <Paragraph style={{fontSize: 15}}>
                                1. 本网站收集和整理的所有链接均来源于网络, 并不表示本站观点。本站仅提供基于搜索引擎类的推荐服务,所有详细信息均跳转到原始网页地址访问,不做任何转码或存储操作。
                                服务器仅存储文章标题和链接,不做任何正文内容抓取或存储。如果本站的某些内容侵犯您的权益,请及时与我们联系,我们会尽快处理。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                2. 本站不对链接的错误、不实、争议性、违法或侵权的信息内容负责。一切由此引起的争议,本站不承担任何法律责任。
                                相关法律责任归原信息来源网站所有。如果您对本站推荐的某些内容持有异议,请直接联系原信息来源网站。
                                本站长期以来一直秉承为读者查找和推荐网络上有价值的信息资源为宗旨。我们会尽量推荐真实、合法和有价值的内容,但鉴于网络信息的开放性和信息量的庞大性,我们无法保证所推荐内容的绝对真实性、准确性和权威性。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                3. 本网站不代表或担保链接到的任何第三方网站或在这些网站上提供的产品、服务或信息的准确性、完整性、合法性、可靠性、适用性、安全性等方面。任何与第三方网站相关的问题，包括但不限于产品、服务、信息、广告或其他内容等，均由该第三方网站或内容提供者对齐负责。在任何情况下，我们对于任何第三方网站或内容提供者所造成的任何损失或索赔不承担任何责任。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                4. 另外，本网站提供的内容仅供参考，不构成任何投资、法律或其他专业建议。在任何情况下，我们不对网站上提供的任何信息或建议的准确性、合法性、适用性或完整性做出任何保证或承担任何责任。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                5. 您应当遵守有关的法律法规和政策规定，不得利用本网站从事任何违法或有悖公序良俗的活动。如果您违反了任何法律法规或政策规定，应自行承担全部责任，且我们有权停止您使用本网站的权利。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                6. 如果您发现任何第三方网站链接存在问题，欢迎随时联系我们并告知，或者通过问题反馈与我们取得联系。我们会尽快处理您的请求，并在必要时立即删除链接。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                最后，本网站的免责声明可能随时更新，一旦更新将立即生效。请您定期访问本页面以查看最新版本的免责声明。
                            </Paragraph>
                            <Paragraph style={{fontSize: 15}}>
                                感谢您阅读本网站的免责声明。如果您有任何疑问或意见，请随时联系我们。
                            </Paragraph>
                        </Typography>
                    </Col>
                </Row>

                <Divider style={{margin: 0}}/>
                <FooterPage/>
            </div>
        );
    }
}

export default Disclaimer;