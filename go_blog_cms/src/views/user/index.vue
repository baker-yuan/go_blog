<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      :columns="columns"
      :request-api="getTableList"
      :init-param="initParam"
      :data-callback="dataCallback"
      @darg-sort="sortTable"
      :border="false"
    >
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增</el-button>
      </template>
      <!-- 头像 -->
      <template #avatar="scope">
        <!-- eslint-disable-next-line prettier/prettier -->
        <img style="height: 80px" :src="scope.row.avatar" alt="">
      </template>
      <!-- 禁用 -->
      <template #status="scope">
        <el-switch
          @change="handleStatusChange(scope.row)"
          :active-value="User.UserStatus.USER_STATUS_DISABLED"
          :inactive-value="User.UserStatus.USER_STATUS_NORMAL"
          v-model="scope.row.status"
        >
        </el-switch>
      </template>
      <!-- 创建时间 -->
      <template #createTime="scope">
        <div>{{ formatTimestamp(scope.row.createTime) }}</div>
      </template>
      <!-- 修改时间 -->
      <template #updateTime="scope">
        <div>{{ formatTimestamp(scope.row.updateTime) }}</div>
      </template>
      <!-- 表格操作 -->
      <template #operation="scope">
        <div v-if="scope.row.userType == User.UserTypeAdmin">
          <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
          <el-button type="primary" link :icon="Delete" @click="deleteUser(scope.row)">删除</el-button>
        </div>
      </template>
    </ProTable>
    <UserDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="UserManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { User } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import UserDrawer from "@/views/user/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { createUserApi, updateUserApi, deleteUserApi, getUsersApi } from "@/api/modules/user.ts";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<User.Model>[]>([
  {
    prop: "uid",
    label: "uid",
    width: 100,
    search: { el: "input" }
  },
  {
    prop: "username",
    label: "用户名",
    width: 150,
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left" // 设置列标题左对齐
  },
  {
    prop: "nickname",
    label: "昵称",
    width: 150,
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left" // 设置列标题左对齐
  },
  {
    prop: "avatar",
    label: "头像",
    width: 150
  },
  {
    prop: "status",
    label: "禁用",
    width: 150
  },
  {
    prop: "userType",
    label: "用户类型",
    width: 150
  },
  { prop: "createUser", label: "创建人", minWidth: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", minWidth: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", minWidth: 170 }
]);

// 如果表格需要初始化请求参数，直接定义传给 ProTable (之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({});

// dataCallback 是对于返回的表格数据做处理，如果你后台返回的数据不是 list && total && pageNum && pageSize 这些字段，可以在这里进行处理成这些字段
// 或者直接去 hooks/useTable.ts 文件中把字段改为你后端对应的就行
const dataCallback = (data: any) => {
  console.log("dataCallback", data);
  return {
    list: data.data,
    total: data.pageTotal,
    pageNum: data.pageNum,
    pageSize: data.pageSize
  };
};

// 如果你想在请求之前对当前请求参数做一些操作，可以自定义如下函数：params 为当前所有的请求参数（包括分页），最后返回请求列表接口
// 默认不做操作就直接在 ProTable 组件上绑定	:requestApi="getTableList"
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  console.log("getTableList", newParams);
  return getUsersApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除用户
const deleteUser = async (row: User.Model) => {
  await useHandleData(deleteUserApi, row.id, `删除【${row.name}】用户`);
  proTable.value?.getTableList();
};

// 是否禁用改变
const handleStatusChange = (row: User.Model) => {
  if (row.uid == 0 || row.uid == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateUserApi(newParams);
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof UserDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<User.Model> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createUserApi : title === "编辑" ? updateUserApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
