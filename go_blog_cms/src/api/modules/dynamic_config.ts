import { ResPage, DynamicConfig } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加动态配置
export const createDynamicConfigApi = (params: DynamicConfig.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateDynamicConfig`, params);
};

// 修改动态配置
export const updateDynamicConfigApi = (dynamicConfig: DynamicConfig.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateDynamicConfig`, dynamicConfig);
};

// 分页查询动态配置
export const getDynamicConfigsApi = (params: DynamicConfig.ReqDynamicConfigListParams) => {
  return http.post<ResPage<DynamicConfig.Model>>(BlogSrv + `/admin/searchDynamicConfig`, params);
};

// 根据id获取动态配置
export const getDynamicConfigApi = (id: number) => {
  return http.post<DynamicConfig.Model>(BlogSrv + `/admin/dynamicConfigDetail`, {
    id: id
  });
};

// 根据id删除动态配置
export const deleteDynamicConfigApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deleteDynamicConfig`, {
    id: id
  });
};
