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

      <!-- 是否启用 -->
      <template #isEnable="scope">
        <el-switch
          @change="handleIsEnableChange(scope.row)"
          :active-value="true"
          :inactive-value="false"
          v-model="scope.row.isEnable"
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
        <el-button type="primary" link :icon="EditPen">菜单权限</el-button>
        <el-button type="primary" link :icon="EditPen">资源权限</el-button>
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteRole(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <RoleDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="RoleManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { Role } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import RoleDrawer from "@/views/auth/role/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { createRoleApi, updateRoleApi, deleteRoleApi, getRolesApi } from "@/api/modules/role.ts";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Role.Model>[]>([
  {
    prop: "name",
    label: "名称",
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 150,
    search: { el: "input" }
  },
  {
    prop: "code",
    label: "编码",
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 150,
    search: { el: "input" }
  },
  {
    prop: "describe",
    label: "功能描述",
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 150,
    search: { el: "input" }
  },
  {
    prop: "isEnable",
    label: "是否启用",
    width: 200
  },
  { prop: "createUser", label: "创建人", minWidth: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", minWidth: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", width: 350 }
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
  return getRolesApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除角色表
const deleteRole = async (row: Role.Model) => {
  await useHandleData(deleteRoleApi, row.id, `删除【${row.name}】角色表`);
  proTable.value?.getTableList();
};

// 是否启用改变
const handleIsEnableChange = (row: Menu.Model) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateRoleApi(newParams);
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof RoleDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Role.Model> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createRoleApi : title === "编辑" ? updateRoleApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
