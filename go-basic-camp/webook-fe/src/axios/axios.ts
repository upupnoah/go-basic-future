import axios from "axios";
import router from "next/router";
const instance = axios.create({
    // 修改对应的配置文件
    baseURL: "http://192.168.31.38:8099", // 若部署到云服务器上，需要修改为云服务器的地址
    withCredentials: true
})

export interface Result<T> {
    code: number,
    msg: string,
    data: T,
}


instance.interceptors.response.use(function (resp) {
    const newToken = resp.headers["x-jwt-token"]
    const newRefreshToken = resp.headers["x-refresh-token"]
    if (newToken) {
        localStorage.setItem("token", newToken)
    }
    if (newRefreshToken) {
        localStorage.setItem("refresh_token", newRefreshToken)
    }
    if (resp.status == 401) {
        window.location.href = "/users/login"
    }
    return resp
}, (err) => {
    console.log(err)
    if (err.response.status == 401) {
        window.location.href = "/users/login"
    }
    return err
})

// 在这里让每一个请求都加上 authorization 的头部
instance.interceptors.request.use((req) => {
    const token = localStorage.getItem("token")
    req.headers.setAuthorization("Bearer " + token, true)
    return req
}, (err) => {
    console.log(err)
})

export default instance