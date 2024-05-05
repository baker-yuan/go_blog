<template>
  <div class="main-box">
    <TreeFilter multiple label="name" :request-api="getResourceTree" @change="changeTreeFilter" :expand-all="false" />
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
        <!-- 接口状态 -->
        <template #status="scope">
          <el-switch
            @change="handleStatusChange(scope.row)"
            :active-value="Resource.ResourceStatus.PUBLISHED"
            :inactive-value="Resource.ResourceStatus.OFFLINE"
            v-model="scope.row.status"
          >
          </el-switch>
        </template>

        <!-- 登录校验 -->
        <template #isNeedLogin="scope">
          <el-switch
            @change="handleIsNeedLoginChange(scope.row)"
            :active-value="Resource.NeedLogin.LOGIN_REQUIRED"
            :inactive-value="Resource.NeedLogin.LOGIN_NOT_REQUIRED"
            v-model="scope.row.isNeedLogin"
          >
          </el-switch>
        </template>
        <!-- 权限校验 -->
        <template #isNeedPermission="scope">
          <el-switch
            @change="handleIsNeedPermissionChange(scope.row)"
            :active-value="Resource.NeedPermission.PERMISSION_REQUIRED"
            :inactive-value="Resource.NeedPermission.PERMISSION_NOT_REQUIRED"
            v-model="scope.row.isNeedPermission"
          >
          </el-switch>
        </template>
        <!-- 标准输出 -->
        <template #isNeedStdRes="scope">
          <el-switch
            @change="handleIsNeedStdResChange(scope.row)"
            :active-value="Resource.NeedStdRes.RES_REQUIRED"
            :inactive-value="Resource.NeedStdRes.RES_NOT_REQUIRED"
            v-model="scope.row.isNeedStdRes"
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
          <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
          <el-button type="primary" link :icon="Delete" @click="deleteBakerAuthResource(scope.row)">删除</el-button>
        </template>
      </ProTable>
      <ResourceDrawer ref="drawerRef" />
    </div>
  </div>
</template>
<script setup lang="ts" name="ResourceList">
import { ref, reactive } from "vue";
import TreeFilter from "@/components/TreeFilter/index.vue";
import { getResourceTree } from "@/api/modules/tree_node";
import ResourceDrawer from "@/views/auth/resource/detail.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { Resource } from "@/api/interface";
import { useHandleData } from "@/hooks/useHandleData";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { createResourceApi, updateResourceApi, deleteResourceApi, getResourcesApi } from "@/api/modules/resource.ts";
import ProTable from "@/components/ProTable/index.vue";
import { statusOption, httpMethodOption, protocolOption } from "@/views/auth/resource/index";
import { formatTimestamp } from "@/utils/time";

// TreeFilter选中的节点
const treeFilterValues = reactive({ treeNodeId: [] });
const changeTreeFilter = (val: string[]) => {
  treeFilterValues.treeNodeId = val;
  // 重置列表页面pageNum
  proTable.value!.pageable.pageNum = 1;
  proTable.value?.getTableList();
};

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Resource.ResourceModel>[]>([
  {
    prop: "name",
    label: "接口名称",
    showOverflowTooltip: false,
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 150,
    search: { el: "input" }
  },
  {
    prop: "describe",
    label: "接口描述",
    showOverflowTooltip: false,
    width: 150
  },
  {
    prop: "treeNodePath",
    label: "路径",
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 200,
    showOverflowTooltip: false
  },
  {
    prop: "url",
    label: "接口URL",
    align: "left", // 设置列内容左对齐
    headerAlign: "left", // 设置列标题左对齐
    width: 300,
    showOverflowTooltip: false
  },
  {
    prop: "status",
    label: "是否生效",
    width: 90,
    enum: statusOption
  },
  { prop: "isNeedLogin", label: "登录校验", width: 90 },
  { prop: "isNeedPermission", label: "权限校验", width: 90 },
  { prop: "isNeedStdRes", label: "标准输出", width: 90 },
  {
    prop: "httpMethod",
    label: "请求方式",
    width: 100,
    enum: httpMethodOption
  },
  {
    prop: "protocol",
    label: "接口协议",
    width: 200,
    enum: protocolOption
  },
  {
    prop: "service",
    label: "API服务名",
    width: 200,
    showOverflowTooltip: false
  },
  {
    prop: "method",
    label: "API接口名",
    width: 250,
    showOverflowTooltip: false
  },
  { prop: "createUser", label: "创建人", minWidth: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", minWidth: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 如果表格需要初始化请求参数，直接定义传给 ProTable (之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({});

// dataCallback 是对于返回的表格数据做处理，如果你后台返回的数据不是 list && total && pageNum && pageSize 这些字段，可以在这里进行处理成这些字段
// 或者直接去 hooks/useTable.ts 文件中把字段改为你后端对应的就行
const dataCallback = (data: any) => {
  console.log("dataCallback", data);
  return { list: data.data, total: data.pageTotal, pageNum: data.pageNum, pageSize: data.pageSize };
};

// 如果你想在请求之前对当前请求参数做一些操作，可以自定义如下函数：params 为当前所有的请求参数（包括分页），最后返回请求列表接口
// 默认不做操作就直接在 ProTable 组件上绑定	:requestApi="getTableList"
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  console.log("getTableList", newParams);
  const search = {};
  if (treeFilterValues.treeNodeId.length != 0) {
    search.tree_node_id = {
      value: treeFilterValues.treeNodeId,
      condition: "in"
    };
  }
  newParams.search = JSON.stringify(search);
  return getResourcesApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除接口资源表
const deleteBakerAuthResource = async (row: Resource.ResourceModel) => {
  await useHandleData(deleteResourceApi, row.id, `删除【${row.name}】接口资源表`);
  proTable.value?.getTableList();
};

// 接口状态改变
const handleStatusChange = (row: Resource.ResourceModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateResourceApi(newParams);
};

// 标准输出改变
const handleIsNeedStdResChange = (row: Resource.ResourceModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateResourceApi(newParams);
};

// 是否需要登录改变
const handleIsNeedLoginChange = (row: Resource.ResourceModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateResourceApi(newParams);
};

// 是否需要权限改变
const handleIsNeedPermissionChange = (row: Resource.ResourceModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateResourceApi(newParams);
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof ResourceDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Resource.ResourceModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createResourceApi : title === "编辑" ? updateResourceApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
