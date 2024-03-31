<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}菜单资源表`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="菜单名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写菜单名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="图标" prop="icon">
        <SelectIcon v-model:icon-value="drawerProps.row!.icon" />
      </el-form-item>
      <el-form-item label="组件路径" prop="component">
        <el-input v-model="drawerProps.row!.component" placeholder="请填写组件路径" clearable></el-input>
      </el-form-item>
      <el-form-item label="访问路径" prop="path">
        <el-input v-model="drawerProps.row!.path" placeholder="请填写访问路径" clearable></el-input>
      </el-form-item>
      <el-form-item label="显示排序" prop="sortValue">
        <el-input-number v-model="drawerProps.row!.sortValue" placeholder="请填写显示排序" clearable></el-input-number>
      </el-form-item>
      <el-form-item label="显示状态" prop="isHidden">
        <el-radio-group v-model="drawerProps.row!.isHidden">
          <el-radio :label="true">显示</el-radio>
          <el-radio :label="false">不显示</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="MenuDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Menu } from "@/api/interface";
import SelectIcon from "@/components/SelectIcon/index.vue";

const rules = reactive({});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Menu.Model>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}菜单资源表成功！` });
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
