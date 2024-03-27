import { request } from 'umi';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) => {
  return request<Res<ResListData<SSLModule.ResponseBody>>>('/ssl', {
    params: {
      page: current,
      page_size: pageSize,
      ...res,
    },
  }).then(({ data }) => {
    return {
      total: data.total_size,
      data: data.rows,
    };
  });
};

export const remove = (id: string) =>
  request(`/ssl/${id}`, {
    method: 'DELETE',
  });

export const create = (data: SSLModule.SSL) =>
  request('/ssl', {
    method: 'POST',
    data,
  });

export const verifyKeyPaire = (cert = '', key = ''): Promise<SSLModule.VerifyKeyPaireProps> =>
  request('/check_ssl_cert', {
    method: 'POST',
    data: { cert, key },
  });

export const update = (id: string, data: SSLModule.SSL) =>
  request(`/ssl/${id}`, {
    method: 'PUT',
    data,
  });
