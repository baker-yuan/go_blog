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
      <!-- 内容 -->
      <template #content="scope">
        <div @click="viewContentDialog(scope.row)" style="cursor: pointer">{{ scope.row.content }}</div>
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
        <el-row :gutter="10">
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openEditJsonDialog(scope.row)"
            >
              编辑内容
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openEditJsonSchemaDialog(scope.row)"
              >编辑约束
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openDrawer('编辑', scope.row)"
              >编辑其他
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="Delete"
              @click="deleteDynamicConfig(scope.row)"
              >删除
            </el-button>
          </el-col>
        </el-row>
      </template>
    </ProTable>
    <DynamicConfigDrawer ref="drawerRef" />

    <!-- 编辑内容 -->
    <el-drawer v-model="showEditJsonContent" :destroy-on-close="true" size="600px" title="编辑内容">
      <VueForm
        @submit="handlerContentFormSubmit"
        @cancel="handlerContentFormCancel"
        v-model="contentFormData"
        :schema="contentSchemaData"
      ></VueForm>
    </el-drawer>
    <!-- 编辑约束 -->
    <el-drawer
      class-name="custom-drawer"
      :style="{ padding: '0px' }"
      v-model="showEditJsonSchema"
      :destroy-on-close="true"
      size="700px"
      title="编辑 JSON Schema"
    >
      <div class="dialog-content">
        <json-editor v-model="contentSchemaData"></json-editor>
      </div>
      <template #footer>
        <el-button @click="showEditJsonSchema = false">取消</el-button>
        <el-button type="primary" @click="saveJsonSchema">确认</el-button>
      </template>
    </el-drawer>
    <!-- 编辑json schema -->
    <el-drawer
      class-name="custom-drawer"
      :style="{ padding: '0px' }"
      v-model="showContent"
      :destroy-on-close="true"
      size="700px"
      title="查看内容"
    >
      <div class="dialog-content">
        <json-editor v-model="contentFormData"></json-editor>
      </div>
    </el-drawer>
  </div>
</template>

<script setup lang="ts" name="DynamicConfigManage">
import VueForm from "@lljj/vue3-form-element";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { DynamicConfig, ResultCode } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import DynamicConfigDrawer from "@/views/system/config/detail.vue";
import JsonEditor from "@/components/JsonEditor/index.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import {
  createDynamicConfigApi,
  updateDynamicConfigApi,
  deleteDynamicConfigApi,
  getDynamicConfigsApi
} from "@/api/modules/dynamic_config";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<DynamicConfig.Model>[]>([
  {
    prop: "bizType",
    label: "业务类型",
    showOverflowTooltip: false,
    search: { el: "input" },
    width: 200
  },
  {
    prop: "bizDesc",
    label: "业务描述",
    search: { el: "input" },
    width: 200
  },
  {
    prop: "content",
    label: "内容",
    search: { el: "input" },
    width: 200
  },
  {
    prop: "jsonSchema",
    label: "json内容约束",
    search: { el: "input" },
    width: 200
  },
  { prop: "createUser", label: "创建人", minWidth: 120, showOverflowTooltip: false },
  { prop: "updateUser", label: "更新人", minWidth: 120, showOverflowTooltip: false },
  { prop: "createTime", label: "创建时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "updateTime", label: "修改时间", minWidth: 180, showOverflowTooltip: false },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

const rowData = ref<DynamicConfig.Model | null>(null);

// 控制谈框展示
const showEditJsonSchema = ref(false); // 编辑约束
const showEditJsonContent = ref(false); // 编辑内容
const showContent = ref(false); // 展示内容

const contentFormData = ref<any>({}); // 编辑内容
const contentSchemaData = ref<any>({}); // 编辑约束

// 编辑内容-点击操作栏
const openEditJsonDialog = (row: DynamicConfig.Model) => {
  rowData.value = row;
  if (!row) {
    return;
  }
  showEditJsonContent.value = true;
  fillData(row);
};

// 编辑约束-点击操作栏
const openEditJsonSchemaDialog = (row: DynamicConfig.Model) => {
  rowData.value = row;
  showEditJsonSchema.value = true;
  fillData(row);
};

// 编辑约束-点击操作栏
const viewContentDialog = (row: DynamicConfig.Model) => {
  rowData.value = row;
  showContent.value = true;
  fillData(row);
};

const fillData = (row: DynamicConfig.Model) => {
  rowData.value = row;
  if (row.jsonSchema) {
    contentSchemaData.value = JSON.parse(row.jsonSchema);
  }
  if (row.content) {
    contentFormData.value = JSON.parse(row.content);
  }
};

// 编辑内容-提交
const handlerContentFormSubmit = () => {
  let newParams: DynamicConfig.Model = JSON.parse(JSON.stringify(rowData.value));
  newParams.content = JSON.stringify(contentFormData.value);
  updateDynamicConfigApi(newParams).then(value => {
    if (value.code == ResultCode.SUCCESS) {
      showEditJsonContent.value = false;
      // 重新记载列表数据
      proTable.value?.getTableList();
    }
  });
};

// 编辑内容-取消
const handlerContentFormCancel = () => {
  showEditJsonContent.value = false;
};

// 编辑约束-确定
const saveJsonSchema = () => {
  if (rowData.value == null) {
    return;
  }
  let newParams: DynamicConfig.Model = JSON.parse(JSON.stringify(rowData.value));
  newParams.jsonSchema = JSON.stringify(contentSchemaData.value);
  updateDynamicConfigApi(newParams).then(value => {
    if (value.code == ResultCode.SUCCESS) {
      showEditJsonSchema.value = false;
      // 重新记载列表数据
      proTable.value?.getTableList();
    }
  });
};

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
  return getDynamicConfigsApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除JSON动态配置
const deleteDynamicConfig = async (row: DynamicConfig.Model) => {
  await useHandleData(deleteDynamicConfigApi, row.id, `删除【${row.bizType}】JSON动态配置`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof DynamicConfigDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<DynamicConfig.Model> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createDynamicConfigApi : title === "编辑" ? updateDynamicConfigApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>

<style scoped>
.custom-drawer ::v-deep .el-drawer__body {
  padding: 0 !important;
}
</style>
