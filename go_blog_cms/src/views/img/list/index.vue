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
        <el-button type="primary" :icon="CirclePlus" @click="openUploadImgDrawer">上传</el-button>
      </template>
      <!-- OSS厂商 -->
      <template #osType="scope">
        <el-tag>{{ formatOssType(scope.row.osType) }}</el-tag>
      </template>
      <!-- 状态 -->
      <template #status="scope">
        <el-tag :type="getOssStatusType(scope.row.status)">{{ formatOssStatus(scope.row.status) }}</el-tag>
      </template>
      <!-- 图片地址 -->
      <template #url="scope">
        <el-image style="width: 80px; height: 80px" :src="scope.row.url" :preview-src-list="[scope.row.url]"></el-image>
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
        <el-button type="primary" link @click="copyLink(scope.row)">复制</el-button>
      </template>
    </ProTable>
    <OssUploadImgDrawer ref="uploadImgDrawerRef" />
  </div>
</template>

<script setup lang="ts" name="OssUploadRecordManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { OssUploadRecord } from "@/api/interface";
import { ref, reactive } from "vue";
import { CirclePlus } from "@element-plus/icons-vue";
import { getOssUploadRecordsApi } from "@/api/modules/oss_upload_record.ts";
import { formatTimestamp } from "@/utils/time";
import { formatOssStatus, formatOssType, getOssStatusType } from "@/views/img/list/index";
import OssUploadImgDrawer from "@/views/img/list/upload_img.vue";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<OssUploadRecord.Model>[]>([
  {
    prop: "id",
    label: "主键ID",
    width: 100,
    search: { el: "input" }
  },
  {
    prop: "uid",
    label: "用户ID",
    width: 100,
    search: { el: "input" }
  },
  {
    prop: "mode",
    label: "模块",
    width: 100,
    search: { el: "input" },
    showOverflowTooltip: false
  },
  {
    prop: "url",
    label: "图片地址",
    width: 250
  },
  {
    prop: "osType",
    label: "OSS厂商",
    width: 100,
    search: { el: "input" }
  },
  {
    prop: "status",
    label: "状态",
    width: 100,
    search: { el: "input" }
  },
  { prop: "createUser", label: "创建人", minWidth: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", minWidth: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", width: 100 }
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
  return getOssUploadRecordsApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

const copyLink = async (row: OssUploadRecord.Model) => {
  console.log("copyLink row: " + JSON.stringify(row));
};

// 打开 drawer(新增、查看、编辑)
const uploadImgDrawerRef = ref<InstanceType<typeof OssUploadImgDrawer> | null>(null);
const openUploadImgDrawer = (row: Partial<OssUploadRecord.Model> = {}) => {
  const params = {
    row: { ...row },
    getTableList: proTable.value?.getTableList
  };
  uploadImgDrawerRef.value?.acceptParams(params);
};
</script>
