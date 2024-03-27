import { request } from 'umi';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) =>
  request('/consumers', {
    params: {
      username: res.username,
      page: current,
      page_size: pageSize,
    },
  }).then(({ data }) => ({
    data: data.rows,
    total: data.total_size,
  }));

export const fetchItem = (username: string) =>
  request<{ data: ConsumerModule.ResEntity }>(`/consumers/${username}`);

export const create = (data: ConsumerModule.Entity) =>
  request('/consumers', {
    method: 'PUT',
    data,
  });

export const update = (username: string, data: ConsumerModule.Entity) =>
  request(`/consumers/${username}`, {
    method: 'PUT',
    data,
  });

export const remove = (username: string) => request(`/consumers/${username}`, { method: 'DELETE' });
