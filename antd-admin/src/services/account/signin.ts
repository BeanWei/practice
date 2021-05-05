// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 账号登录 POST /api/user/signin */
export async function signIn(
  params: {
    /** 账号 */
    account: string;
    /** 密码 */
    password: string;
  },
  options?: { [key: string]: any },
) {
  return request<API.CurrentUser>('/api/user/signin', {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
