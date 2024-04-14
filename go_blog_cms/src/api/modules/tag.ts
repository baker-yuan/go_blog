import { ResPage, Tag } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加文章标签
export const createTagApi = (params: Tag.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateTag`, params);
};

// 修改文章标签
export const updateTagApi = (tag: Tag.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateTag`, tag);
};

// 分页查询文章标签
export const getTagsApi = (params: Tag.ReqTagListParams) => {
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
  return http.post<ResPage<Tag.Model>>(BlogSrv + `/admin/searchTag`, params);
};

// 根据id获取文章标签
export const getTagApi = (id: number) => {
  return http.post<Tag.Model>(BlogSrv + `/admin/TagDetail`, {
    id: id
  });
};

// 根据id删除文章标签
export const deleteTagApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deleteTag`, {
    id: id
  });
};
