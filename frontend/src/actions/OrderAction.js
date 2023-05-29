import {
    CUSTOMER_ORDER,
    CUSTOMER_ORDER_CLEAR,
    CUSTOMER_ORDER_STATUS,
    CUSTOMER_ORDER_TRANSACTIONS,
    CUSTOMER_RECHARGE
} from "./index";
import {request} from "../common/request";
import {message} from "antd";
import {checkUserInfo, HttpOK, isEmpty, OrderCancelStateCode, OrderPaidStateCode} from "./Base";
import {GetPersonalAccountInfo, GetPersonalWebLinkInfo, GetMessages} from "./CustomerAction";

function LoadRecharges(records) {
    return {
        type: CUSTOMER_RECHARGE,
        records
    }
}

// 查看充值列表
export const GetUpgradeRecharges = (callback) => {
    return (dispatch) => {
        request("/anonymous/upgrade/recharge")
            .then(res => {
                if (res.code === 200) {
                    dispatch(LoadRecharges(res.data));
                } else {
                    message.error(res.data.message);
                }
                if (!isEmpty(callback)) {
                    callback(res.data)
                }
            });
    }
}

function LoadOrderPay(records) {
    return {
        type: CUSTOMER_ORDER,
        records
    }
}

// 创建订单
export const CreateOrder = (upgradeUID, totalPrice, paymentAmount, paymentType) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/order", {
            method: 'post',
            body: {
                "upgrade_uid": upgradeUID,
                "total_price": totalPrice,
                "payment_amount": paymentAmount,
                "payment_type": paymentType
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadOrderPay(res.data));
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadOrderStatus(records) {
    return {
        type: CUSTOMER_ORDER_STATUS,
        records
    }
}

// 查询订单状态
export const GetOrderStatus = (orderNumber, success, expired) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/order/" + orderNumber + "/status", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.status === OrderPaidStateCode) {
                    success()
                    dispatch(GetMessages(1));
                    dispatch(GetPersonalAccountInfo());
                    dispatch(GetPersonalWebLinkInfo());
                } else if (res.data.status === OrderCancelStateCode) {
                    expired()
                    dispatch(LoadOrderStatus(res.data));
                }
            } else {
                message.error(res.data.message);
            }
        });
    }
}

function LoadOrderTransactions(records) {
    return {
        type: CUSTOMER_ORDER_TRANSACTIONS,
        records
    }
}

// 查询订单交易记录
export const GetOrderTransactions = (callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/order/transactions", {
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                dispatch(LoadOrderTransactions(res.data));
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}

function LoadClearOrder(records) {
    return {
        type: CUSTOMER_ORDER_CLEAR,
        records
    }
}

// 清空订单消息
export const ClearOrder = () => {
    return (dispatch) => {
        dispatch(LoadClearOrder({}));
    }
}

// 问题反馈
export const CreateFeedback = (body, callback) => {
    return (dispatch) => {
        const userInfo = checkUserInfo();
        request("/customer/" + userInfo.user_uid + "/personal/feedback", {
            method: 'post',
            body: {
                ...body,
            },
            headers: {'Authorization': userInfo.token}
        }).then(res => {
            if (res.code === 200) {
                if (res.data.message === HttpOK) {
                    message.success("提交成功");
                }
            } else {
                message.error(res.data.message);
            }
            if (!isEmpty(callback)) {
                callback()
            }
        });
    }
}