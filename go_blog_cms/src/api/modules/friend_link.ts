import { ResPage, FriendLink } from "@/api/interface/index";
import http from "@/api";
import { BlogSrv } from "@/api/config/servicePort";

// 添加友链表
export const createFriendLinkApi = (params: FriendLink.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateFriendLink`, params);
};

// 修改友链表
export const updateFriendLinkApi = (friendLink: FriendLink.Model) => {
  return http.post(BlogSrv + `/admin/addOrUpdateFriendLink`, friendLink);
};

// 分页查询友链表
export const getFriendLinksApi = (params: FriendLink.ReqFriendLinkListParams) => {
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
  return http.post<ResPage<FriendLink.Model>>(BlogSrv + `/admin/searchFriendLink`, params);
};

// 根据id获取友链表
export const getFriendLinkApi = (id: number) => {
  return http.post<FriendLink.Model>(BlogSrv + `/admin/friendLinkDetail`, {
    id: id
  });
};

// 根据id删除友链表
export const deleteFriendLinkApi = (id: number) => {
  return http.post(BlogSrv + `/admin/deleteFriendLink`, {
    id: id
  });
};
