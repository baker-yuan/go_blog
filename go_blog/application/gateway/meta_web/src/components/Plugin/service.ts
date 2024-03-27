import { omit } from 'lodash';
import { request } from 'umi';

import { PLUGIN_LIST, PluginType } from './data';

const cached: {
  list: PluginComponent.Meta[];
} = {
  list: [],
};

export const fetchList = () => {
  if (cached.list.length) {
    return Promise.resolve(cached.list);
  }

  return request<Res<PluginComponent.Meta[]>>('/plugins?all=true').then((data) => {
    const typedData = data.data.map((item) => ({
      ...item,
      type: PLUGIN_LIST[item.name]?.type || 'other',
      originType: item.type,
      hidden: PLUGIN_LIST[item.name]?.hidden || false,
    }));

    let finalList: PluginComponent.Meta[] = [];

    Object.values(PluginType).forEach((type) => {
      finalList = finalList.concat(typedData.filter((item) => item.type === type));
    });

    if (cached.list.length === 0) {
      cached.list = finalList;
    }

    return finalList;
  });
};

/**
 * cache plugin schema by schemaType
 * default schema is route for plugins in route
 * support schema: consumer for plugins in consumer
 */
const cachedPluginSchema: Record<string, any> = {
  route: {},
  consumer: {},
};
export const fetchSchema = async (
  name: string,
  schemaType: PluginComponent.Schema,
): Promise<any> => {
  if (!cachedPluginSchema[schemaType][name]) {
    const queryString = schemaType !== 'route' ? `?schema_type=${schemaType}` : '';
    cachedPluginSchema[schemaType][name] = (
      await request(`/schema/plugins/${name}${queryString}`)
    ).data;
    // for plugins schema returned with properties: [], which will cause parse error
    if (JSON.stringify(cachedPluginSchema[schemaType][name].properties) === '[]') {
      cachedPluginSchema[schemaType][name] = omit(
        cachedPluginSchema[schemaType][name],
        'properties',
      );
    }
  }
  return cachedPluginSchema[schemaType][name];
};

export const fetchPluginTemplateList = () => {
  return request<Res<ResListData<PluginTemplateModule.ResEntity>>>('/plugin_configs').then(
    (data) => {
      return data.data.rows;
    },
  );
};
