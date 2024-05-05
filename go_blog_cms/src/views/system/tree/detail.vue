<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}树形结构表`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="业务类型" prop="bizType">
        <el-select
          :disabled="drawerProps.row!.id"
          @change="handleBizTypeChange"
          v-model="drawerProps.row!.bizType"
          placeholder="请选择接口协议"
        >
          <el-option v-for="item in bizTypeOptions" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="父节点" prop="parentID">
        <el-cascader
          style="width: 100%"
          v-model="selectedValue"
          :options="treeNodeOptions"
          :props="{ value: 'id', label: 'name', checkStrictly: true }"
          @change="handleTreeNodeIdChange"
          filterable
          clearable
        ></el-cascader>
      </el-form-item>
      <el-form-item label="节点名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写节点名称" clearable></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="TreeNodeDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { TreeNode } from "@/api/interface";
import { bizTypeOptions } from "@/views/system/tree/index";
import { getTreeByBizType } from "@/api/modules/tree_node";

const rules = reactive({
  name: [{ required: true, message: "请填写节点名称" }],
  bizType: [{ required: true, message: "请填写业务类型" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<TreeNode.Model>; // 数据
  api?: (params: any) => Promise<any>; // 提交api
  getTableList?: () => void; // 刷新列表
}

// 是否显示弹框
const drawerVisible = ref(false);
// 参数
const drawerProps = ref<DrawerProps>({
  isView: false,
  title: "",
  row: {}
});

// 父节点下拉框
const treeNodeOptions = reactive([]);
// 存储选中的值
const selectedValue = ref([]);
// 当treeNodeId级联选择器的值发生变化时调用的方法
const handleTreeNodeIdChange = (value: string[]) => {
  if (value.length > 0) {
    drawerProps.value.row.parentId = value[value.length - 1];
  } else {
    drawerProps.value.row.parentId = 0;
  }
};

// 业务类型改变
const handleBizTypeChange = () => {
  // 加载树节点
  getTreeByBizType(drawerProps.value.row.bizType).then(value => {
    // 清空现有的数组内容
    treeNodeOptions.splice(0, treeNodeOptions.length);
    // 添加新的数据到数组
    treeNodeOptions.push(...value.data);
  });
};

// 接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;

  selectedValue.value = params.row.parentId;
  // 加载树节点
  handleBizTypeChange();
};

// 提交数据（新增/编辑）
const ruleFormRef = ref<FormInstance>();
const handleSubmit = () => {
  ruleFormRef.value!.validate(async valid => {
    if (!valid) return;
    try {
      await drawerProps.value.api!(drawerProps.value.row);
      ElMessage.success({ message: `${drawerProps.value.title}树形结构表成功！` });
      drawerProps.value.getTableList!();
      drawerVisible.value = false;
    } catch (error) {
      console.log(error);
    }
  });
};

defineExpose({
  acceptParams
});
</script>
