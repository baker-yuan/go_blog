<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}友链表`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="友链名称" prop="linkName">
        <el-input v-model="drawerProps.row!.linkName" placeholder="请填写友链名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="友链头像" prop="linkAvatar">
        <UploadImg v-model:image-url="drawerProps.row!.linkAvatar" mode="article">
          <template #tip> 上传图片最大为3M </template>
        </UploadImg>
      </el-form-item>
      <el-form-item label="友链地址" prop="linkAddress">
        <el-input v-model="drawerProps.row!.linkAddress" placeholder="请填写友链地址" clearable></el-input>
      </el-form-item>
      <el-form-item label="友链介绍" prop="linkIntro">
        <el-input v-model="drawerProps.row!.linkIntro" type="textarea" placeholder="请填写友链介绍" clearable></el-input>
      </el-form-item>
      <el-form-item label="友链状态" prop="status">
        <el-radio-group v-model="drawerProps.row!.status">
          <el-radio :label="FriendLink.FriendLinkStatus.STATUS_PUBLISHED">已发布</el-radio>
          <el-radio :label="FriendLink.FriendLinkStatus.STATUS_OFFLINE">以下线</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="友链排序" prop="sort">
        <el-input-number v-model="drawerProps.row!.sort" placeholder="请填写友链排序" clearable></el-input-number>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="FriendLinkDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { FriendLink } from "@/api/interface";
import UploadImg from "@/components/Upload/Img.vue";

const rules = reactive({});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<FriendLink.Model>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}友链表成功！` });
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
