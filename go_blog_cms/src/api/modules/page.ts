import { ResPage, Page } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加页面
export const createPageApi = (params: Page.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdatePage`, params);
};

// 修改页面
export const updatePageApi = (page: Page.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdatePage`, page);
};

// 分页查询页面
export const getPagesApi = (params: Page.ReqPageListParams) => {
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
  return http.post<ResPage<Page.Model>>(BlogSrv + `/admin/searchPage`, params);
};

// 根据id获取页面
export const getPageApi = (id: number) => {
  return http.post<Page.Model>(BlogSrv + `/admin/pageDetail`, {
    id: id
  });
};

// 根据id删除页面
export const deletePageApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deletePage`, {
    id: id
  });
};
