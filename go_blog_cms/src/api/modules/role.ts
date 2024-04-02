import { ResPage, Role } from "@/api/interface/index";
import http from "@/api";
import { AuthSrv } from "@/api/config/servicePort";

// 添加角色表
export const createRoleApi = (params: Role.RoleModel) => {
  return http.post(AuthSrv + `/admin/addOrUpdateRole`, params);
};

// 修改角色表
export const updateRoleApi = (role: Role.RoleModel) => {
  return http.post(AuthSrv + `/admin/addOrUpdateRole`, role);
};

// 分页查询角色表
export const getRolesApi = (params: Role.ReqRoleListParams) => {
  return http.post<ResPage<Role.RoleModel>>(AuthSrv + `/admin/searchRole`, params);
};

// 根据id获取角色表
export const getRoleApi = (id: number) => {
  return http.post<Role.RoleModel>(AuthSrv + `/admin/RoleDetail`, {
    id: id
  });
};

// 根据id删除角色表
export const deleteRoleApi = (id: number) => {
  return http.post(AuthSrv + `/admin/deleteRole`, {
    id: id
  });
};
