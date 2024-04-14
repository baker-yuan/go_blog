import { ResPage, Category } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加文章分类
export const createCategoryApi = (params: Category.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateCategory`, params);
};

// 修改文章分类
export const updateCategoryApi = (category: Category.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateCategory`, category);
};

// 分页查询文章分类
export const getCategorysApi = (params: Category.ReqCategoryListParams) => {
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
  return http.post<ResPage<Category.Model>>(BlogSrv + `/admin/searchCategory`, params);
};

// 根据id获取文章分类
export const getCategoryApi = (id: number) => {
  return http.post<Category.Model>(BlogSrv + `/admin/CategoryDetail`, {
    id: id
  });
};

// 根据id删除文章分类
export const deleteCategoryApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deleteCategory`, {
    id: id
  });
};
