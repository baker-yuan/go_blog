<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="600px" :title="`${drawerProps.title}接口资源表`">
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
      <el-form-item label="目录" prop="treeNodeId">
        <el-cascader
          style="width: 100%"
          v-model="selectedValue"
          :options="treeNodeOptions"
          :props="{ value: 'id', label: 'name' }"
          @change="handleTreeNodeIdChange"
        ></el-cascader>
      </el-form-item>
      <el-form-item label="标输输出" prop="isNeedStdRes">
        <el-radio-group v-model="drawerProps.row!.isNeedStdRes">
          <el-radio :label="Resource.NeedStdRes.RES_REQUIRED">是</el-radio>
          <el-radio :label="Resource.NeedStdRes.RES_NOT_REQUIRED">否</el-radio>
        </el-radio-group>
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
      <el-form-item label="接口协议" prop="protocol">
        <el-select v-model="drawerProps.row!.protocol" placeholder="请选择接口协议">
          <el-option v-for="item in protocolOption" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
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
import { httpMethodOption, protocolOption } from "@/views/auth/resource/index";
import { getResourceTree } from "@/api/modules/tree_node";
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

// 树节点下拉框
const treeNodeOptions = reactive([]);

// 接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;
  selectedValue.value = params.row.treeNodeId;
  // 加载树节点
  getResourceTree().then(value => {
    // 清空现有的数组内容
    treeNodeOptions.splice(0, treeNodeOptions.length);
    // 添加新的数据到数组
    treeNodeOptions.push(...value.data);
  });
};

// 存储选中的值
const selectedValue = ref([]);

// 当treeNodeId级联选择器的值发生变化时调用的方法
const handleTreeNodeIdChange = (value: string[]) => {
  if (value.length > 0) {
    drawerProps.value.row.treeNodeId = value[value.length - 1];
  } else {
    drawerProps.value.row.treeNodeId = 0;
  }
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
