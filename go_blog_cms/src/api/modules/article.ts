import { ResPage, Article } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加文章表
export const createArticleApi = (params: Article.Model) => {
  return http.post(BlogSrv + `/addOrUpdateArticle`, params);
};

// 修改文章表
export const updateArticleApi = (article: Article.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateArticle`, article);
};

// 分页查询文章表
export const getArticlesApi = (params: Article.ReqArticleListParams) => {
  // if (!params.pageSize) {
  //   const error = new Error("stack trace");
  //   console.log(error.stack);
  //   params.pageSize = 10;
  // }
  // if (!params.pageNum) {
  //   const error = new Error("stack trace");
  //   console.log(error.stack);
  //   params.pageNum = 1;
  // }
  return http.post<ResPage<Article.Model>>(BlogSrv + `/admin/searchArticle`, params);
};

// 根据id获取文章表
export const getArticleApi = (id: number) => {
  return http.post<Article.Model>(BlogSrv + `/admin/ArticleDetail`, {
    id: id
  });
};

// 根据id删除文章表
export const deleteArticleApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deleteArticle`, {
    id: id
  });
};
