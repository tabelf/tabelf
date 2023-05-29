import moment from 'moment';
import {apiUrl, clearUserInfo, ERR_CODE, ERR_FAILED, ERR_MESSAGE, isEmpty, toLogin} from "../actions/Base";

function Result(code, res) {
    return {
        code: code,
        data: res
    }
}

// 封装请求
export async function request(url, param) {
    let payload = {};
    if (param) {
        payload = {...param, body: JSON.stringify(param.body)};
    }

    let type = 'string';
    const {body} = payload;
    if (body != null) {
        type = typeof body;
    }

    let headers = isEmpty(payload.headers) ? {} : payload.headers;
    headers.Accept = 'application/json';
    let params = '';
    if (type === 'string') {
        headers['Content-Type'] = 'application/json';
        const timestamp = moment().format('x');
        if (url.indexOf("?") === -1) {
            params = `?timestamp=${timestamp}`;
        } else {
            params = `&timestamp=${timestamp}`;
        }
    }

    const options = {...payload, headers};
    let resp = null;
    try {
        resp = await fetch(`${apiUrl}` + url + params, options);
    } catch (e) {
        clearTimeout(window.timer);
    }
    if (resp === null) {
        return Result(ERR_FAILED, {
            code: ERR_CODE,
            message: ERR_MESSAGE
        })
    }
    // 验证过期
    if (resp.status === 401) {
        clearUserInfo();
        toLogin()
        return;
    }
    return Result(resp.status, await resp.json())
}