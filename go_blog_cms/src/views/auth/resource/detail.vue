<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}接口资源表`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="接口名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写接口名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="接口描述" prop="describe">
        <el-input v-model="drawerProps.row!.describe" type="textarea" placeholder="请填写接口描述" clearable></el-input>
      </el-form-item>
      <el-form-item label="接口URL" prop="url">
        <el-input v-model="drawerProps.row!.url" placeholder="请填写接口URL" clearable></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-radio-group v-model="drawerProps.row!.status">
          <el-radio :label="Resource.ResourceStatus.PUBLISHED">已发布</el-radio>
          <el-radio :label="Resource.ResourceStatus.OFFLINE">以下线</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="请求方式" prop="httpMethod">
        <el-radio-group v-model="drawerProps.row!.httpMethod" size="small" clearable>
          <el-radio v-for="(item, index) in httpMethodOption" :key="index" :label="item.value">
            {{ item.label }}
          </el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="登录校验" prop="isNeedLogin">
        <el-radio-group v-model="drawerProps.row!.isNeedLogin">
          <el-radio :label="Resource.NeedLogin.LOGIN_REQUIRED">是</el-radio>
          <el-radio :label="Resource.NeedLogin.LOGIN_NOT_REQUIRED">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="权限校验" prop="isNeedPermission">
        <el-radio-group v-model="drawerProps.row!.isNeedPermission">
          <el-radio :label="Resource.NeedPermission.PERMISSION_REQUIRED">是</el-radio>
          <el-radio :label="Resource.NeedPermission.PERMISSION_NOT_REQUIRED">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="API服务名" prop="service">
        <el-input v-model="drawerProps.row!.service" placeholder="请填写API服务名" clearable></el-input>
      </el-form-item>
      <el-form-item label="API接口" prop="method">
        <el-input v-model="drawerProps.row!.method" placeholder="请填写API接口" clearable></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="ResourceDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Resource } from "@/api/interface";
import { httpMethodOption } from "@/views/auth/resource/index";
const rules = reactive({
  name: [{ required: true, message: "请填写接口名称" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Resource.ResourceModel>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}接口资源表成功！` });
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
