import React, {Component} from 'react';
import {Spin} from 'antd';

class ESpinPage extends Component {
    render() {
        return (
            <div style={{
                height: "100vh",
                display: "grid",
                alignItems: "center",
                justifyContent: "center"
            }}>
                <Spin
                    size="large"
                    spinning={true}
                />
            </div>
        );
    }
}

export default ESpinPage;