import { request } from 'umi';

import { convertToFormData } from '@/components/Upstream/service';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) => {
  return request<Res<ResListData<UpstreamModule.RequestBody>>>('/upstreams', {
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
};

export const fetchOne = (id: string) =>
  request<Res<any>>(`/upstreams/${id}`).then(({ data }) => convertToFormData(data));

export const create = (data: UpstreamModule.RequestBody) =>
  request('/upstreams', {
    method: 'POST',
    data,
  });

export const update = (id: string, data: UpstreamModule.RequestBody) =>
  request(`/upstreams/${id}`, {
    method: 'PUT',
    data,
  });

export const remove = (id: string) => request(`/upstreams/${id}`, { method: 'DELETE' });
