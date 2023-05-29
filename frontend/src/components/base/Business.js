import React, {Component} from 'react';
import {Card, Divider, Alert, Image} from 'antd';
import FooterPage from "../footer/FooterPage";
import QrImage from '../../assets/img_v2_qrcode.jpg'

class Business extends Component {
    render() {
        return (
            <div style={{height: "100vh"}}>
                <div style={{
                    height: "88vh",
                    display: "grid",
                    alignItems: "center",
                    justifyContent: "center",
                    background: "#f0f2f5"
                }}>
                    <Card>
                        <Alert showIcon message="如果需要商务合作请添加微信详聊，备注：商务合作" type="warning" />
                        <div style={{textAlign: "center"}}>
                            <Image
                                src={QrImage}
                                width={210}
                                preview={false}
                            />
                        </div>
                    </Card>
                </div>
                <Divider style={{margin: 0}}/>
                <FooterPage/>
            </div>
        );
    }
}

export default Business;