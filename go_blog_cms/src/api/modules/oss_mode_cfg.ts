import { ResPage, OssModeCfg } from "@/api/interface/index";
import http from "@/api";
import { External } from "@/api/config/servicePort";

// 添加文件路径
export const createOssModeCfgApi = (params: OssModeCfg.Model) => {
  return http.post(External + `/admin/addOrUpdateOssModeCfg`, params);
};

// 修改文件路径
export const updateOssModeCfgApi = (OssModeCfg: OssModeCfg.Model) => {
  return http.post(External + `/admin/addOrUpdateOssModeCfg`, OssModeCfg);
};

// 分页查询文件路径
export const getOssModeCfgsApi = (params: OssModeCfg.ReqOssModeCfgListParams) => {
  // if (!params.pageSize) {
  //   const error = new Error("stack trace");
  //   console.log(error.stack);
  //   params.pageSize = 10;
  // }
  // if (!params.pageNum) {
  //   const error = new Error("stack trace");
  //   console.log(error.stack);
  //   params.pageNum = 1;
  // }
  return http.post<ResPage<OssModeCfg.Model>>(External + `/admin/searchOssModeCfg`, params);
};

// 根据id获取文件路径
export const getOssModeCfgApi = (id: number) => {
  return http.post<OssModeCfg.Model>(External + `/admin/OssModeCfgDetail`, {
    id: id
  });
};

// 根据id删除文件路径
export const deleteOssModeCfgApi = (id: number) => {
  return http.post(External + `/admin/deleteOssModeCfg`, {
    id: id
  });
};
