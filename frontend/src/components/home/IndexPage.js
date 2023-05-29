import React, {Component} from 'react';
import {Layout} from "antd";

const {Content} = Layout;

class IndexPage extends Component {
    render() {
        return (
            <div>
                <Content style={{margin: '20px 0px 0', overflow: 'initial'}}>
                    <div className="site-content-background" style={{padding: '16px 24px 10px 24px'}}>
                    </div>
                </Content>
            </div>
        );
    }
}

export default IndexPage;