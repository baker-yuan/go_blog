<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}页面`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="页面名" prop="pageName">
        <el-input v-model="drawerProps.row!.pageName" placeholder="请填写页面名" clearable></el-input>
      </el-form-item>
      <el-form-item label="页面标签" prop="pageLabel">
        <el-input v-model="drawerProps.row!.pageLabel" placeholder="请填写页面标签" clearable></el-input>
      </el-form-item>
      <el-form-item label="页面封面" prop="pageCover">
        <SingleUpload v-model:value="drawerProps.row!.pageCover" mode="article">
          <template #tip> 上传图片最大为3M </template>
        </SingleUpload>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="PageDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Page } from "@/api/interface";
import SingleUpload from "@/components/Upload/SingleUpload.vue";

const rules = reactive({
  pageName: [{ required: true, message: "请填写页面名" }],
  pageLabel: [{ required: true, message: "请填写页面标签" }],
  pageCover: [{ required: true, message: "请填写页面封面" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Page.Model>; // 数据
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

// 接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;
};

// 提交数据（新增/编辑）
const ruleFormRef = ref<FormInstance>();
const handleSubmit = () => {
  ruleFormRef.value!.validate(async valid => {
    if (!valid) return;
    try {
      await drawerProps.value.api!(drawerProps.value.row);
      ElMessage.success({ message: `${drawerProps.value.title}页面成功！` });
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
