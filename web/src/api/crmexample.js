import request from '@/utils/request'

export function listExample(query) {
  return request({
    url: '/api/v1/crmexample',
    method: 'get',
    params: query
  })
}

export function getExample(id) {
  return request({
    url: '/api/v1/crmexample',
    method: 'get'
  })
}

export function addExample(data) {
  return request({
    url: '/api/v1/crmexample',
    method: 'post',
    data: data
  })
}

export function updateExample(data) {
  return request({
    url: '/api/v1/crmexample',
    method: 'put',
    data: data
  })
}

export function deleteExample(id) {
  return request({
    url: '/api/v1/crmexample',
    method: 'delete',
  })
}