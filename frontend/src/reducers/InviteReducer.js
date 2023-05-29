import {combineReducers} from 'redux';
import {CUSTOMER_INVITE} from "../actions";

const initPersonalInvite = {}

function LoadPersonalInvite(state = initPersonalInvite, action) {
    switch (action.type) {
        case CUSTOMER_INVITE:
            return action.records
        default:
            return state
    }
}

export const inviteReducers = combineReducers({
    invites: LoadPersonalInvite
})
