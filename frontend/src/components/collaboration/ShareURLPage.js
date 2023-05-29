import React, {Component} from 'react';
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import * as CustomerAction from "../../actions/CustomerAction";
import {withRouter} from "../base/Base";
import {checkUserInfo, isEmpty} from "../../actions/Base";
import HomePage from "../home/HomePage";

class ShareURLPage extends Component {

    componentDidMount() {
        const token = checkUserInfo()
        if (isEmpty(token)) {
            return;
        }
        const shareUID = this.props.params.share_uid
        if (isEmpty(shareUID)) {
            window.location.href = "/workspace"
            return;
        }
        this.props.actions.JoinSharePersonalFolder(shareUID)
    }

    render() {
        return (
            <div>
                <HomePage/>
            </div>
        );
    }
}

export default withRouter(connect(
    state => ({state: state.dataManage.customer}),
    dispatch => ({
        actions: bindActionCreators(CustomerAction, dispatch)
    })
)(ShareURLPage));