import { request } from 'umi';

import { transformData } from './transform';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) =>
  request('/services', {
    params: {
      id: res.id || '',
      desc: res.desc || '',
      name: res.name,
      page: current,
      page_size: pageSize,
    },
  }).then(({ data }) => ({
    data: data.rows,
    total: data.total_size,
  }));

export const create = (data: ServiceModule.Entity) =>
  request('/services', {
    method: 'POST',
    data: transformData(data),
  });

export const update = (serviceId: string, data: ServiceModule.Entity) =>
  request(`/services/${serviceId}`, {
    method: 'PUT',
    data: transformData(data),
  });

export const remove = (serviceId: string) =>
  request(`/services/${serviceId}`, { method: 'DELETE' });

export const fetchItem = (serviceId: number) =>
  request(`/services/${serviceId}`).then((data) => data);
