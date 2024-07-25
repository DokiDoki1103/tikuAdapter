import request from '../utils/request'


export function login(params) {
  return request({
    url: `/adapter-service/user/login`,
    method: 'get',
    params
  });
}


export function getPlat() {
  return request({
    url: `/adapter-service/plat`,
    method: 'get',
  })
}

export function getCourses(params) {
  return request({
    url: `/adapter-service/courses`,
    method: 'get',
    params
  })
}

export function getQuestions(data) {
  return request({
    url: `/adapter-service/questions/search`,
    method: 'post',
    data
  })
}

export function updateQuestions(data) {
  return request({
    url: `/adapter-service/questions/${data.id}`,
    method: 'put',
    data
  })
}

export function delQuestions(data) {
  return request({
    url: `/adapter-service/questions/${data.id}`,
    method: 'delete'
  })
}

export function createQuestions(data) {
  return request({
    url: `/adapter-service/questions`,
    method: 'post',
    data
  })
}
// 获取用户列表
export function getUserList() {
  return request({
    url: `/adapter-service/user`,
    method: 'get'
  })
}
// 删除用户
export function delUser(data) {
  return request({
    url: `/adapter-service/user/${data.id}`,
    method: 'delete'
  })
}

// 新增用户
export function addUser(data) {
  return request({
    url: `/adapter-service/user`,
    method: 'post',
    data
  })
}

export function getLogList() {
  return request({
    url: `/adapter-service/logs`,
    method: 'get',
  })
}




