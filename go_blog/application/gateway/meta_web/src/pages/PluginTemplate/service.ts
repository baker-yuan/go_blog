import { request } from 'umi';

import { transformLabelList } from '@/helpers';

export const fetchList = ({ current = 1, pageSize = 10, ...res }) => {
  const { labels = [] } = res;

  return request('/plugin_configs', {
    params: {
      search: res.desc,
      label: labels.join(','),
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

export const remove = (rid: string) => request(`/plugin_configs/${rid}`, { method: 'DELETE' });

export const fetchItem = (id: string) =>
  request<{ data: PluginTemplateModule.ResEntity }>(`/plugin_configs/${id}`);

export const create = (data: PluginTemplateModule.Entity) =>
  request('/plugin_configs', {
    method: 'POST',
    data,
  });

export const update = (id: string, data: PluginTemplateModule.Entity) =>
  request(`/plugin_configs/${id}`, {
    method: 'PATCH',
    data,
  });

export const fetchLabelList = () =>
  request('/labels/plugin_config').then(({ data }) => transformLabelList(data.rows) as LabelList);
