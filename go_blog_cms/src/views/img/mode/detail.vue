<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}OSS模块配置`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="名称" prop="mode">
        <el-input v-model="drawerProps.row!.mode" placeholder="请填写名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="路径" prop="path">
        <el-input v-model="drawerProps.row!.path" placeholder="请填写路径" clearable></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="desc">
        <el-input v-model="drawerProps.row!.desc" type="textarea" placeholder="请填写描述" clearable></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="OssModeCfgDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { OssModeCfg } from "@/api/interface";

const rules = reactive({
  mode: [
    { required: true, message: "请填写模块" },
    { max: 100, message: "模块最多50个字符" }
  ],
  path: [
    { required: true, message: "请填写路径" },
    { max: 100, message: "路径最多50个字符" }
  ],
  desc: [
    { required: true, message: "请填写描述称" },
    { max: 255, message: "描述最多50个字符" }
  ]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<OssModeCfg.Model>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}OSS模块配置成功！` });
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
