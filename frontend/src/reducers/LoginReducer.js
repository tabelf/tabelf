import {combineReducers} from 'redux';
import {LOGIN_QR_CODE} from "../actions";

const initQrCode = {
    loading: true,
}

function LoginLoadQrCode(state = initQrCode, action) {
    switch (action.type) {
        case LOGIN_QR_CODE:
            return action.records
        default:
            return state
    }
}

export const loginReducers = combineReducers({
    qrcode: LoginLoadQrCode
})
