<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}文章标签`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="标签名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写标签名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="标签描述" prop="description">
        <el-input v-model="drawerProps.row!.description" type="textarea" placeholder="请填写标签描述" clearable></el-input>
      </el-form-item>
      <el-form-item label="标签颜色" prop="color">
        <el-color-picker v-model="drawerProps.row!.color" :predefine="predefineColors" clearable></el-color-picker>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="TagDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Tag } from "@/api/interface";

const rules = reactive({
  name: [
    { required: true, message: "请填写标签名称" },
    { max: 50, message: "标签名称最多50个字符" }
  ],
  description: [
    { required: true, message: "请填写标签描述" },
    { max: 100, message: "标签描述最多100个字符" }
  ],
  color: [{ required: true, message: "请填写标签颜色" }]
});

const predefineColors = [
  "#ff4500", // 橙红色
  "#ff6347", // 番茄色
  "#ff8c00", // 深橙色
  "#ffd700", // 金色
  "#90ee90", // 亮绿色
  "#00fa9a", // 适中的春天的绿色
  "#00ffff", // 青色
  "#1e90ff", // 道奇蓝
  "#0000ff", // 蓝色
  "#8a2be2", // 蓝紫罗兰色
  "#ff00ff", // 洋红色
  "#ff1493", // 深粉色
  "#c71585", // 中紫罗兰红色
  "#a52a2a", // 褐色
  "#ffffff", // 白色
  "#c0c0c0", // 银色
  "#808080", // 灰色
  "#000000" // 黑色
];

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Tag.Model>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}文章标签成功！` });
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
