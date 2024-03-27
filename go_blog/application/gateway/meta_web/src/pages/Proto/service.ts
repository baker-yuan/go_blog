import { request } from 'umi';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) => {
  return request<Res<ResListData<ProtoModule.ResponseBody>>>('/proto', {
    params: {
      desc: res.desc,
      page: current,
      page_size: pageSize,
    },
  }).then(({ data }) => {
    return {
      data: data.rows,
      total: data.total_size,
    };
  });
};

export const create = (data: ProtoModule.ProtoData) =>
  request('/proto', {
    method: 'POST',
    data,
  });

export const update = (data: ProtoModule.ProtoData) => {
  request(`/proto/${data.id}`, {
    method: 'PUT',
    data,
  });
};
export const remove = (rid: string) => request(`/proto/${rid}`, { method: 'DELETE' });
