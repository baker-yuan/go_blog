// 请求响应参数（不包含data）
export interface Result {
  code: string;
  msg: string;
}

// 请求响应参数（包含data）
export interface ResultData<T = any> extends Result {
  data: T;
}

// 分页响应参数
export interface ResPage<T> {
  list: T[];
  pageNum: number;
  pageSize: number;
  total: number;
}

// 分页请求参数
export interface ReqPage {
  pageNum: number;
  pageSize: number;
}

// 文件上传模块
export namespace Upload {
  export interface ResFileUrl {
    fileUrl: string;
  }
}

// 登录模块
export namespace Login {
  export interface ReqLoginForm {
    username: string;
    password: string;
  }
  export interface ResLogin {
    access_token: string;
  }
  export interface ResAuthButtons {
    [key: string]: string[];
  }
}

// 用户管理模块
// export namespace User {
//   export interface ReqUserParams extends ReqPage {
//     username: string;
//     gender: number;
//     idCard: string;
//     email: string;
//     address: string;
//     createTime: string[];
//     status: number;
//   }
//   export interface ResUserList {
//     id: string;
//     username: string;
//     gender: number;
//     user: { detail: { age: number } };
//     idCard: string;
//     email: string;
//     address: string;
//     createTime: string;
//     status: number;
//     avatar: string;
//     photo: any[];
//     children?: ResUserList[];
//   }
//   export interface ResStatus {
//     userLabel: string;
//     userValue: number;
//   }
//   export interface ResGender {
//     genderLabel: string;
//     genderValue: number;
//   }
//   export interface ResDepartment {
//     id: string;
//     name: string;
//     children?: ResDepartment[];
//   }
//   export interface ResRole {
//     id: string;
//     name: string;
//     children?: ResDepartment[];
//   }
// }

// 树形结构
export namespace TreeNode {
  export interface Model {
    id: string;
    name: string;
    parentId: string;
    bizType: string;
    children?: Model[];
  }
}

// 接口资源
export namespace Resource {
  // 接口状态枚举
  export enum ResourceStatus {
    UNPUBLISHED = 0, // 未发布
    PUBLISHED = 1, // 已发布
    OFFLINE = 2 // 以下线
  }

  // 定义是否需要登录校验的枚举
  export enum NeedLogin {
    LOGIN_NOT_REQUIRED = 0, // 不需要登录校验
    LOGIN_REQUIRED = 1 // 需要登录校验
  }

  // 定义是否需要权限校验的枚举
  export enum NeedPermission {
    PERMISSION_NOT_REQUIRED = 0, // 不需要权限校验
    PERMISSION_REQUIRED = 1 // 需要权限校验
  }

  // http请求方式
  export enum HttpMethod {
    UNKNOWN = 0,
    GET = 1,
    POST = 2,
    PUT = 3,
    PATCH = 4,
    DELETE = 5
  }

  // 接口资源模型
  export interface ResourceModel {
    // 基本信息
    id: number; // ID
    treeNodeId: number; // 树节点ID
    name: string; // 接口名称
    describe: string; // 接口描述
    //
    status: ResourceStatus; // 接口状态
    // 路径定位
    url: string; // 接口URL（包含路径变量）
    httpMethod: string; // 请求方式
    // 权限校验
    isNeedLogin: NeedLogin; // 是否需要登录校验
    isNeedPermission: NeedPermission; // 是否需要权限校验
    // 下游服务信息
    service: string; // API服务名
    method: string; // API接口名
    // 公共字段
    createUserId: string; // 创建人id
    updateUserId: string; // 更新人id
    createTime: number; // 插入时间
    updateTime: number; // 修改时间
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
    treeNodePath: string; // 目录树路径
  }

  // 分页查询接口资源表
  export interface ReqResourceListParams extends ReqPage {
    name?: string; // 接口名称
  }
}

// 菜单资源
export namespace Menu {
  // 菜单资源表模型
  export interface Model {
    // 基本信息
    id: number; // 菜单id
    parentId: number; // 父级菜单id
    name: string; // 菜单名称
    icon: string; // 菜单图标
    describe: string; // 功能描述
    sortValue: number; // 排序
    // 前端
    path: string; // 对应路由path
    component: string; // 对应路由组件component
    // 状态
    isHidden: boolean; // 是否隐藏
    // 公共字段
    createUserId: number; // 创建人id
    updateUserId: number; // 更新人id
    createTime: string; // 插入时间
    updateTime: string; // 修改时间
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
  }
  // 分页查询菜单资源表
  export interface ReqMenuListParams extends ReqPage {}
}

// 角色管理模块
export namespace Role {
  // 角色表模型
  export interface Model {
    // 基本信息
    id: number;
    name: string;
    code: string;
    describe: string;
    // 是否启用
    isEnable: boolean;
    // 公共字段
    createUserId: number;
    updateUserId: number;
    createTime: number;
    updateTime: number;
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
  }

  // 分页查询角色表
  export interface ReqRoleListParams extends ReqPage {}
}

// 用户管理模块
export namespace User {
  // 用户状态枚举
  export enum UserStatus {
    USER_STATUS_NORMAL = 0, // 正常
    USER_STATUS_DISABLED = 1 // 禁用
  }
  // 用户类型
  export const UserTypeAdmin: string = "ADMIN";
  export const UserTypeUser: string = "USER";

  // 用户模型
  export interface Model {
    // 账号信息
    uid: number; // 主键
    username: string; // 用户名(唯一)
    password: string; // 登录密码
    salt: string; // 盐
    userType: "ADMIN" | "USER"; // 用户类型 ADMIN-管理员 USER-普通用户
    // 基本信息
    email: string; // 邮箱号
    nickname: string; // 昵称
    avatar: string; // 头像地址
    intro: string; // 用户简介
    webSite: string; // 个人网站
    // 状态
    status: UserStatus; // 状态 0-正常 1-禁用
    // 三方登录
    loginType: number; // 登录方式 1-用户名 2-GitHub 3-码云 4-QQ 5-微博
    unionId: string; // 用户唯一标识（第三方网站）
    // 公共字段
    isDeleted: boolean; // 是否注销
    createUserId: number; // 创建人
    updateUserId: number; // 修改人
    createTime: number; // 创建时间
    updateTime: number; // 修改时间
  }

  // 分页查询用户
  export interface ReqUserListParams extends ReqPage {}
}

// 文章表管理模块
export namespace Article {
  // 文章表模型
  export interface Model {
    // 基本数据
    id: number; // 主键
    categoryID: number; // 文章分类表ID|biz_type表主键
    title: string; // 文章标题
    description: string; // 文章简介，最多200字
    content: string; // 文章内容
    coverImage: string; // 文章封面图片
    originalURL: string; // 原文链接
    //
    password: string; // 密码保护
    // 辅助
    words: number; // 文章字数
    readTime: number; // 文章阅读时长(分钟)
    // 类型描述
    type: ArticleType; // 文章类型 1-原创 2-转载 3-翻译
    status: ArticleStatus; // 文章状态 0-草稿 1-已发布
    format: ArticleFormat; // 文章格式 1-markdown 2-富文本
    visibility: ArticleVisibility; // 文章可见性 1-公开 2-私密 3-密码保护

    // 标志
    isTop: boolean; // 置顶开关
    isRecommend: boolean; // 推荐开关
    isAppreciation: boolean; // 赞赏开关
    isCommentEnabled: boolean; // 评论开关

    // 公共字段
    isDeleted: boolean; // 是否删除
    createUserId: number;
    updateUserId: number;
    createTime: number;
    updateTime: number;
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
  }

  // 分页查询文章表
  export interface ReqArticleListParams extends ReqPage {}
}

// 文章标签管理模块
export namespace Tag {
  // 文章标签模型
  export interface Model {
    id: number; // 主键
    name: string; // 标签名
    description: string; // 标签描述
    color: string; // 标签颜色
    // 公共字段
    createUserId: number;
    updateUserId: number;
    createTime: number;
    updateTime: number;
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
  }

  // 分页查询文章标签
  export interface ReqTagListParams extends ReqPage {}
}

// 文章分类管理模块
export namespace Category {
  // 文章分类模型
  export interface Model {
    id: number; // 主键
    name: string; // 文章类型名
    description: string; // 类型介绍
    sort: number; // 排序
    available: CategoryStatus; // 是否可用 0-不可用 1-可用
    // 公共字段
    createUserId: number;
    updateUserId: number;
    createTime: number;
    updateTime: number;
    // 查询补充
    createUser: string; // 创建人
    updateUser: string; // 更新人
  }

  // 分页查询文章分类
  export interface ReqCategoryListParams extends ReqPage {}
}
