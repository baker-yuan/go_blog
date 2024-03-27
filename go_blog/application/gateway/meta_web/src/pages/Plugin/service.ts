import { request } from 'umi';

import { DEFAULT_GLOBAL_RULE_ID } from '@/constants';

export const fetchList = (): Promise<{
  data: PluginModule.TransformedPlugin[];
  total: number;
}> =>
  request<{
    data: {
      plugins: Record<string, any>;
    };
  }>(`/global_rules/${DEFAULT_GLOBAL_RULE_ID}`).then(({ data }) => {
    const plugins = Object.entries(data.plugins || {})
      .filter(([, value]) => !value.disable)
      .map(([name, value]) => ({
        id: name,
        name,
        value,
      }));

    return {
      data: plugins,
      total: plugins.length,
    };
  });

export const createOrUpdate = (data: Partial<Omit<PluginModule.GlobalRule, 'id'>>) =>
  request(`/global_rules/${DEFAULT_GLOBAL_RULE_ID}`, {
    method: 'PUT',
    data: { id: DEFAULT_GLOBAL_RULE_ID, ...data },
  });

export const fetchPluginList = () => {
  return request<Res<PluginComponent.Meta[]>>('/plugins?all=true').then((data) => {
    return data.data;
  });
};
