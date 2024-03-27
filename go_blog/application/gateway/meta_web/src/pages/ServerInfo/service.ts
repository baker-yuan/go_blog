import { request } from 'umi';

export const fetchInfoList = () => {
  return request<Res<ResListData<ServerInfoModule.Node>>>('/server_info').then(
    ({ data }) => data.rows,
  );
};

export const fetchVersion = () => {
  return request<Res<ServerInfoModule.DashboardInfo>>('/tool/version').then(({ data }) => data);
};
