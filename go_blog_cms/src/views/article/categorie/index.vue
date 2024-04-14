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
        <el-button type="primary" link :icon="Delete" @click="deleteCategory(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <CategoryDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="CategoryManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { Category } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import CategoryDrawer from "@/views/article/categorie/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { createCategoryApi, updateCategoryApi, deleteCategoryApi, getCategorysApi } from "@/api/modules/category.ts";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Category.Model>[]>([
  {
    prop: "id",
    label: "分类ID",
    width: 100,
    search: { el: "input" }
  },
  {
    prop: "name",
    label: "名称",
    search: { el: "input" },
    width: 100,
    showOverflowTooltip: false
  },
  {
    prop: "description",
    label: "描述",
    search: { el: "input" },
    minWidth: 100,
    showOverflowTooltip: false
  },
  {
    prop: "sort",
    label: "排序",
    width: 100
  },
  { prop: "createUser", label: "创建人", width: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", width: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", width: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", width: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
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
  return getCategorysApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除文章分类
const deleteCategory = async (row: Category.Model) => {
  await useHandleData(deleteCategoryApi, row.id, `删除【${row.name}】文章分类`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof CategoryDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Category.Model> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createCategoryApi : title === "编辑" ? updateCategoryApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
