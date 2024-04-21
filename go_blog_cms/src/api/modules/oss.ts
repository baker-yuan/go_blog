import { UploadSign } from "@/api/interface/index";
import http from "@/api";
import { External } from "@/api/config/servicePort";

// 获取oss上传签名
export const getUploadSignApi = (req: UploadSign.Req) => {
  return http.post<UploadSign.Rsp>(External + `/admin/getUploadSign`, req);
};
