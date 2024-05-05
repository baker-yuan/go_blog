import { ResPage, TreeNode } from "@/api/interface";
import { AuthSrv } from "@/api/config/servicePort";
import http from "@/api";

/**
 * 获取接口资源的树形结构
 */
export const getResourceTree = () => {
  let req = {
    bizType: TreeNode.BizType.BIZ_TYPE_RESOURCE
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

/**
 * 根据业务类型获取目录数
 */
export const getTreeByBizType = (bizType: TreeNode.BizType) => {
  let req = {
    bizType: bizType
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

// 添加树形结构表
export const createTreeNodeApi = (params: TreeNode.Model) => {
  return http.post(AuthSrv + `/admin/addOrUpdateTreeNode`, params);
};

// 修改树形结构表
export const updateTreeNodeApi = (treeNode: TreeNode.Model) => {
  return http.post(AuthSrv + `/admin/addOrUpdateTreeNode`, treeNode);
};

// 分页查询树形结构表
export const getTreeNodesApi = (params: TreeNode.ReqTreeNodeListParams) => {
  return http.post<ResPage<TreeNode.Model>>(AuthSrv + `/admin/searchTreeNode`, params);
};

// 根据id获取树形结构表
export const getTreeNodeApi = (id: number) => {
  return http.post<TreeNode.Model>(AuthSrv + `/admin/treeNodeDetail`, {
    id: id
  });
};

// 根据id删除树形结构表
export const deleteTreeNodeApi = (id: number) => {
  return http.post(AuthSrv + `/admin/deleteTreeNode`, {
    id: id
  });
};
