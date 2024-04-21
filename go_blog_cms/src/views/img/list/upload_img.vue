<template>
  <el-drawer v-model="drawerVisible" direction="rtl" :destroy-on-close="true" size="800px" :title="`文件上传`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="新增图片" prop="patch">
        <SingleUpload v-model:value="drawerProps.row!.patch" mode="article">
          <template #tip> 上传图片最大为3M </template>
        </SingleUpload>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button v-show="!drawerProps.isView" @click="handleSubmit">关闭</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="OssUploadImgDrawer">
import { ref } from "vue";
import { FormInstance } from "element-plus";
import { OssUploadRecord } from "@/api/interface";
import SingleUpload from "@/components/Upload/SingleUpload.vue";

interface DrawerProps {
  row: Partial<OssUploadRecord.Model>; // 数据
  getTableList?: () => void; // 刷新列表
}

// 是否显示弹框
const drawerVisible = ref(false);
// 参数
const drawerProps = ref<DrawerProps>({
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
