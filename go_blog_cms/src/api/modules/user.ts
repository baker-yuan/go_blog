import { ResPage, User } from "@/api/interface/index";
import http from "@/api";
import { UserSrv } from "@/api/config/servicePort";

// 添加用户
export const createUserApi = (params: User.UserModel) => {
  return http.post(UserSrv + `/admin/addOrUpdateUser`, params);
};

// 修改用户
export const updateUserApi = (user: User.UserModel) => {
  return http.post(UserSrv + `/admin/addOrUpdateUser`, user);
};

// 分页查询用户
export const getUsersApi = (params: User.ReqUserListParams) => {
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
  return http.post<ResPage<User.UserModel>>(UserSrv + `/admin/searchUser`, params);
};

// 根据id获取用户
export const getUserApi = (id: number) => {
  return http.post<User.UserModel>(UserSrv + `/admin/userDetail`, {
    id: id
  });
};

// 根据id删除用户
export const deleteUserApi = (id: number) => {
  return http.post(UserSrv + `/admin/deleteUser`, {
    id: id
  });
};
