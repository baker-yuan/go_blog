import { TreeNode } from "@/api/interface";
import { AuthSrv } from "@/api/config/servicePort";
import http from "@/api";

/**
 * @name 获取接口资源的树形结构
 */
export const getResourceTree = () => {
  let req = {
    bizType: 1
  };
  return http.post<TreeNode.Model[]>(AuthSrv + `/admin/getTreeByBizType`, req).then(response => {
    let data = response.data.nodes;
    return {
      code: response.code,
      message: response.message,
      data: data
    };
  });
};
