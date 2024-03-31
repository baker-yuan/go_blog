import { ResPage, Resource } from "@/api/interface/index";
import http from "@/api";
import { AuthSrv } from "@/api/config/servicePort";

// 添加接口资源表
export const createResourceApi = (params: Resource.ResourceModel) => {
  return http.post(AuthSrv + `/admin/addOrUpdateResource`, params);
};

// 修改接口资源表
export const updateResourceApi = (resource: Resource.ResourceModel) => {
  return http.post(AuthSrv + `/admin/addOrUpdateResource`, resource);
};

// 分页查询接口资源表
export const getResourcesApi = (params: Resource.ReqResourceListParams) => {
  if (!params.pageSize) {
    const error = new Error("stack trace");
    console.log(error.stack);
    params.pageSize = 10;
  }
  if (!params.pageNum) {
    const error = new Error("stack trace");
    console.log(error.stack);
    params.pageNum = 1;
  }
  return http.post<ResPage<Resource.ResourceModel>>(AuthSrv + `/admin/searchResource`, params);
};

// 根据id获取接口资源表
export const getResourceApi = (id: number) => {
  return http.post<Resource.ResourceModel>(AuthSrv + `/admin/resourceDetail`, {
    id: id
  });
};

// 根据id删除接口资源表
export const deleteResourceApi = (id: number) => {
  return http.post(AuthSrv + `/admin/deleteResource`, {
    id: id
  });
};
