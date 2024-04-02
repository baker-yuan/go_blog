import { ResPage, Menu } from "@/api/interface/index";
import http from "@/api";
import { AuthSrv } from "@/api/config/servicePort";

// 添加菜单资源表
export const createMenuApi = (menu: Menu.Model) => {
  return http.post(AuthSrv + `/admin/addOrUpdateMenu`, menu);
};

// 修改菜单资源表
export const updateMenuApi = (menu: Menu.Model) => {
  return http.post(AuthSrv + `/admin/addOrUpdateMenu`, menu);
};

// 分页查询菜单资源表
export const getMenusApi = (params: Menu.ReqMenuListParams) => {
  return http.post<ResPage<Menu.Model>>(AuthSrv + `/admin/searchMenu`, params);
};

// 根据id获取菜单资源表
export const getMenuApi = (id: number) => {
  return http.post<Menu.Model>(AuthSrv + `/admin/menuDetail`, {
    id: id
  });
};

// 根据id删除菜单资源表
export const deleteMenuApi = (id: number) => {
  return http.post(AuthSrv + `/admin/deleteMenu`, {
    id: id
  });
};

// 菜单树
export const getMenuTree = (params: Menu.ReqMenuListParams) => {
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
  return http.post<Menu.Model[]>(AuthSrv + `/admin/getMenuTree`, params);
};

// 菜单树
export const getMenuOption = () => {
  let params = {
    pageSize: 100000,
    pageNum: 1
  };
  return http.post<Menu.Model[]>(AuthSrv + `/admin/getMenuTree`, params).then(response => {
    let data = response.data.nodes;
    return {
      code: response.code,
      message: response.message,
      data: data
    };
  });
};
