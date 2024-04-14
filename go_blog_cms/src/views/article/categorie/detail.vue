<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}文章分类`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写分类名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="分类描述" prop="description">
        <el-input v-model="drawerProps.row!.description" type="textarea" placeholder="请填写分类描述" clearable></el-input>
      </el-form-item>
      <el-form-item label="分类排序" prop="sort">
        <el-input-number v-model="drawerProps.row!.sort" placeholder="请填写分类排序" clearable></el-input-number>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="CategoryDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Category } from "@/api/interface";

const rules = reactive({
  name: [
    { required: true, message: "请填写分类名称" },
    { max: 50, message: "分类名称最多50个字符" }
  ],
  description: [
    { required: true, message: "请填写分类描述" },
    { max: 200, message: "分类描述最多200个字符" }
  ],
  sort: [{ required: true, message: "请填写分类排序" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Category.Model>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}文章分类成功！` });
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
