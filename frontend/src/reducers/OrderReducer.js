import {
    CUSTOMER_ORDER,
    CUSTOMER_ORDER_CLEAR,
    CUSTOMER_ORDER_STATUS,
    CUSTOMER_ORDER_TRANSACTIONS,
    CUSTOMER_RECHARGE
} from "../actions";
import {combineReducers} from "redux";

const initRecharges = {
    upgrade_recharges: []
}

function LoadRecharges(state = initRecharges, action) {
    switch (action.type) {
        case CUSTOMER_RECHARGE:
            return action.records
        default:
            return state
    }
}

const initOrder = {}

function LoadOrder(state = initOrder, action) {
    switch (action.type) {
        case CUSTOMER_ORDER:
            return action.records
        case CUSTOMER_ORDER_CLEAR:
            return action.records
        default:
            return state
    }
}

const initOrderStatus = {}

function LoadOrderStatus(state = initOrderStatus, action) {
    switch (action.type) {
        case CUSTOMER_ORDER_STATUS:
            return action.records
        default:
            return state
    }
}

const initOrderTransactions = {
    transactions: []
}

function LoadOrderTransactions(state = initOrderTransactions, action) {
    switch (action.type) {
        case CUSTOMER_ORDER_TRANSACTIONS:
            return action.records
        default:
            return state
    }
}


export const orderReducers = combineReducers({
    recharges: LoadRecharges,
    order: LoadOrder,
    orderStatus: LoadOrderStatus,
    transactions: LoadOrderTransactions
})
