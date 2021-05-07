// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 登录接口 */
export async function Signin(body: API.SigninRequest) {
  return request<API.CurrentUser>('/api/user/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    skipErrorHandler: true,
  });
}

/** 详情接口 */
export async function CurrentUser(body: API.SigninRequest) {
  return request<API.CurrentUser>('/api/user/me', {
    method: 'GET',
  });
}
