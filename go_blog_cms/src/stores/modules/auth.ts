import { defineStore } from "pinia";
import { AuthState } from "@/stores/interface";
import { getAuthButtonListApi, getAuthMenuListApi } from "@/api/modules/login";
import { getFlatMenuList, getShowMenuList, getAllBreadcrumbList } from "@/utils";

export const useAuthStore = defineStore({
  id: "geeker-auth",
  state: (): AuthState => ({
    // 按钮权限列表
    authButtonList: {},
    // 菜单权限列表
    authMenuList: [],
    // 当前页面的 router name，用来做按钮权限筛选
    routeName: ""
  }),
  getters: {
    // 按钮权限列表
    authButtonListGet: state => state.authButtonList,
    // 菜单权限列表 ==> 这里的菜单没有经过任何处理
    authMenuListGet: state => state.authMenuList,
    // 菜单权限列表 ==> 左侧菜单栏渲染，需要剔除 isHide == true
    showMenuListGet: state => getShowMenuList(state.authMenuList),
    // 菜单权限列表 ==> 扁平化之后的一维数组菜单，主要用来添加动态路由
    flatMenuListGet: state => getFlatMenuList(state.authMenuList),
    // 递归处理后的所有面包屑导航列表
    breadcrumbListGet: state => getAllBreadcrumbList(state.authMenuList)
  },
  actions: {
    // Get AuthButtonList
    async getAuthButtonList() {
      const { data } = await getAuthButtonListApi();
      this.authButtonList = data;
    },
    // Get AuthMenuList
    async getAuthMenuList() {
      const { data } = await getAuthMenuListApi();
      // 类型转换
      let authMenuList = transformItems(data.nodes);
      console.log(JSON.stringify(authMenuList));
      this.authMenuList = authMenuList;
    },
    // Set RouteName
    async setRouteName(name: string) {
      this.routeName = name;
    }
  }
});

// 数据转换
// {"id":90,"parent_id":0,"name":"图片管理","icon":"el-icon-mydaohanglantubiao_quanxianguanli","describe":"","sort_value":4,"path":"","component":"Layout","is_hidden":false,"create_user":0,"update_user":0,"create_time":1644081726,"update_time":1644081726,"children":[{"id":91,"parent_id":90,"name":"OSS图片管理","icon":"el-icon-myyonghuliebiao","describe":"","sort_value":1,"path":"/img-manage","component":"/imgManage/ImgList.vue","is_hidden":false,"create_user":0,"update_user":0,"create_time":1644081726,"update_time":1644081726,"children":[]}]}
// {"path":"/auth","name":"auth","redirect":"/auth/menu","meta":{"icon":"Lock","title":"权限管理","isLink":"","isHide":false,"isFull":false,"isAffix":false,"isKeepAlive":true},"children":[{"path":"/auth/menu","name":"authMenu","component":"/auth/menu/index","meta":{"icon":"Menu","title":"菜单权限","isLink":"","isHide":false,"isFull":false,"isAffix":false,"isKeepAlive":true}}]}
function transformItems(items) {
  return items.map(item => transformItem(item));
}

function transformItem(item) {
  // 基础转换
  const transformedItem = {
    path: item.path,
    name: item.name,
    component: item.component,
    meta: {
      icon: item.icon,
      title: item.name,
      isLink: "",
      isHide: false,
      isFull: false,
      isAffix: false,
      isKeepAlive: true
    },
    children: item.children ? item.children.map(transformItem) : []
  };

  // 特殊处理：如果没有子项，则不包含children属性
  if (transformedItem.children.length === 0) {
    delete transformedItem.children;
  }

  // 特殊处理：如果有子项并且路径为空，则添加redirect到第一个子项
  if (transformedItem.children && transformedItem.children.length > 0 && !item.path) {
    transformedItem.redirect = transformedItem.children[0].path;
  }

  return transformedItem;
}
