
import {useCookies} from 'vue3-cookies'
const {cookies} = useCookies()

const TOKEN_KEY = 'token'

export function setToken(value) {
    cookies.set(TOKEN_KEY, value)
}

export function getToken() {
    return cookies.get(TOKEN_KEY)
}

export function removeCookie() {
    cookies.remove(TOKEN_KEY)
}


export function setCookie(key, value) {
    cookies.set(key, value)
}

export function getCookie(key) {
    return cookies.get(key)
}