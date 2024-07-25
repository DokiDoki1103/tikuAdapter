
import Cookies from 'js-cookie'

/** 
  * 设置cookies
  * */
export function getCookies(key) {
  return Cookies.get(key)
}
/** 
  * 设置Cookies
  * */
export function setCookies(key, value) {
  let seconds = 180 * 24 * 60 * 60; // 6个月的秒数
  let expires = new Date(new Date() * 1 + seconds * 1000)
  return Cookies.set(key, value, { expires: expires })
}
/** 
  * 移除Cookies
  * */
export function removeCookies(key) {
  return Cookies.remove(key)
}