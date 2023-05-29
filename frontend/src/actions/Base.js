import {message} from "antd";
import axios from 'axios';

// 通用错误码
export const ERR_FAILED = 400
export const ERR_CODE = 999999;
export const ERR_MESSAGE = "服务调用失败"

export const ErrCustomerWebLinkLimit = 1000403

// order.
export const OrderUnpaidStateCode = "001"; // 待支付.
export const OrderCancelStateCode = "040"; // 订单取消
export const OrderPayFailStateCode = "050"; // 支付失败.
export const OrderPaidStateCode = "100"; // 已支付.

// community.
export const CommunityView = 0
export const CommunityPraise = 1
export const CommunityStar = 2
export const CommunityNew = 3
export const CommunityUsed = 4

// cache
export const USER_TOKEN_EXPIRED = 60 * 24 * 30; // 用户信息 30 天有效期

// token
export const USER_TOKEN_KEY = "USER_TOKEN_KEY"

export const HttpOK = "OK"

export const apiUrl = `${process.env.REACT_APP_API}`;

export function checkUserInfo() {
    const token = Cache.get(USER_TOKEN_KEY);
    if (isEmpty(token)) {
        toLogin()
    } else {
        return token;
    }
}

export function getUserInfo() {
    return Cache.get(USER_TOKEN_KEY);
}

// 清空用户登录信息
export function clearUserInfo() {
    Cache.remove(USER_TOKEN_KEY);
    return true;
}

export function toLogin() {
    window.location.href = "/login";
}

export function CustomSEO(title, keywords, description) {
    document.title = title;
    // 获取 head 元素，即 <head> 标签
    const head = document.querySelector('head');
    // 获取当前页面中的 meta 标签
    const metaKeywords = head.querySelector('meta[name="keywords"]');
    const metaDescription = head.querySelector('meta[name="description"]');
    // 修改 metaKeywords 的 content 属性值
    metaKeywords.content = keywords;
    // 修改 metaDescription 的 content 属性值
    metaDescription.content = description;
}

export function isEmpty(data) {
    return data === null || data === undefined || data === '';
}

export function isNotEmpty(data) {
    return !isEmpty(data)
}

export const downloadFile = (url, filename) => {
    axios({
        url: url,
        method: 'get',
        responseType: 'blob', // 将响应数据以二进制格式处理
    }).then(function (response) {
        const fileName = filename; // 定义文件名
        const blob = new Blob([response.data]); // 创建一个Blob对象
        const url = window.URL.createObjectURL(blob); // 创建一个临时URL
        const link = document.createElement('a'); // 创建一个a标签
        link.href = url; // 设置a标签的href属性为临时URL
        link.setAttribute('download', fileName); // 设置a标签的download属性为文件名
        document.body.appendChild(link); // 将a标签添加到body中
        link.click(); // 模拟点击链接进行下载
        document.body.removeChild(link); // 下载完成后删除a标签
    }).catch((error) => {
        console.log(error);
    });
};

export const Cache = {
    /*
    * set 存储方法
    * @ param {String}     key 键
    * @ param {String}     value 值，
    * @ param {String}     expired 过期时间，以分钟为单位，非必须
    */
    set(key, val, expired) {
        if (typeof val !== 'string') {
            val = JSON.stringify(val);
        }
        window.localStorage.setItem(key, val);
        if (expired) {
            window.localStorage.setItem(`${key}__expires__`, `${Date.now() + 1000 * 60 * expired}`);
        }
    },

    /*
   * get 获取方法
   * @ param {String}     key 键
   * @ param {String}     expired 存储时为非必须字段，所以有可能取不到，默认为 Date.now+1
   */
    get(key) {
        const expired = window.localStorage.getItem(`${key}__expires__`) || Date.now + 1;
        const now = Date.now();

        if (now >= expired) {
            this.remove(key);
            return;
        }
        let val = window.localStorage.getItem(key);
        try {
            val = JSON.parse(val);
        } catch (e) {
            console.log(`${val} Unexpected token H in JSON at position 0`);
        }
        return val;
    },

    clear() {
        window.localStorage.clear();
    },
    /*
    * remove 移除
    * */
    remove(key) {
        if (window.localStorage.getItem(`${key}__expires__`)) {
            window.localStorage.removeItem(`${key}__expires__`);
        }

        if (window.localStorage.getItem(key)) {
            window.localStorage.removeItem(key);
        }
    }
}



