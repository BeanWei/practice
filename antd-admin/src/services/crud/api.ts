// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 查询列表 */
export async function List(module: string, model: string, options?: { [key: string]: any }) {
  return request<Record<string, any>>(`/api/${module}/${model}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取详情 */
export async function Get(module: string, model: string, id: string) {
  return request<Record<string, any>>(`/api/${module}/${model}/${id}`, {
    method: 'GET',
  });
}

/** 新增数据 */
export async function Create(module: string, model: string, body: Record<string, any>) {
  return request<Record<string, any>>(`/api/${module}/${model}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
  });
}

/** 更新数据 */
export async function Update(module: string, model: string, id: string, body: Record<string, any>) {
  return request<Record<string, any>>(`/api/${module}/${model}/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
  });
}

/** 删除数据, 批量删除时：多个id使用逗号拼接 */
export async function Delete(module: string, model: string, id: string) {
  return request<Record<string, any>>(`/api/${module}/${model}/${id}`, {
    method: 'DELETE',
  });
}
