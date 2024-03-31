import { Resource } from "@/api/interface";

// 接口状态下拉框
export const statusOption = [
  {
    value: Resource.ResourceStatus.UNPUBLISHED,
    label: "未发布"
  },
  {
    value: Resource.ResourceStatus.PUBLISHED,
    label: "已发布"
  },
  {
    value: Resource.ResourceStatus.OFFLINE,
    label: "以下线"
  }
];

// http请求方式下拉框
export const httpMethodOption = [
  {
    value: Resource.HttpMethod.UNKNOWN,
    label: "未知"
  },
  {
    value: Resource.HttpMethod.GET,
    label: "GET"
  },
  {
    value: Resource.HttpMethod.POST,
    label: "POST"
  },
  {
    value: Resource.HttpMethod.PUT,
    label: "PUT"
  },
  {
    value: Resource.HttpMethod.PATCH,
    label: "PATCH"
  },
  {
    value: Resource.HttpMethod.DELETE,
    label: "DELETE"
  }
];
