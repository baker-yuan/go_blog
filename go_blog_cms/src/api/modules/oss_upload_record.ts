import { ResPage, OssUploadRecord } from "@/api/interface/index";
import http from "@/api";
import { External } from "@/api/config/servicePort";

// 分页查询OSS文件上传记录
export const getOssUploadRecordsApi = (params: OssUploadRecord.ReqOssUploadRecordListParams) => {
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
  return http.post<ResPage<OssUploadRecord.Model>>(External + `/admin/searchOssUploadRecord`, params);
};

// 根据id获取OSS文件上传记录
export const getOssUploadRecordApi = (id: number) => {
  return http.post<OssUploadRecord.Model>(External + `/admin/ossUploadRecordDetail`, {
    id: id
  });
};
