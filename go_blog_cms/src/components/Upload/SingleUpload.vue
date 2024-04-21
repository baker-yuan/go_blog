<template>
  <div>
    <el-upload
      action="#"
      list-type="picture"
      :multiple="false"
      :show-file-list="showFileList"
      :file-list="fileList"
      :on-remove="handleRemove"
      :on-success="handleUploadSuccess"
      :on-preview="handlePreview"
      :http-request="httpRequest"
      :before-upload="beforeUpload"
    >
      <el-tooltip content="只能上传jpg/png文件，且不超过10MB" placement="right">
        <el-button :icon="CirclePlus" type="primary" size="small">点击上传</el-button>
      </el-tooltip>
    </el-upload>
    <el-dialog v-model="dialogVisible">
      <img width="100%" :src="fileList[0].url" alt="" />
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, defineProps, defineEmits } from "vue";
import * as imageConversion from "image-conversion";
import axios from "axios";
import { CirclePlus } from "@element-plus/icons-vue";
import { getUploadSignApi } from "@/api/modules/oss.ts";
import SparkMD5 from "spark-md5";

const props = defineProps({
  value: String,
  mode: {
    type: String,
    required: true
  }
});

const emit = defineEmits(["update:value"]);

// 图片
const imageUrl = ref(props.value);
// 弹窗是否可见
const dialogVisible = ref(false);

// 图片名称
const imageName = computed(() => {
  return imageUrl.value ? imageUrl.value.substr(imageUrl.value.lastIndexOf("/") + 1) : null;
});

// 是否显示已上传文件列表
const showFileList = computed(() => {
  return imageUrl.value !== null && imageUrl.value !== "" && imageUrl.value !== undefined;
});

// 上传的文件列表, 例如: [{name: 'food.jpg', url: 'https://xxx.cdn.com/xxx.jpg'}]
const fileList = computed(() => {
  return imageUrl.value ? [{ name: imageName.value, url: imageUrl.value }] : [];
});

watch(imageUrl, newValue => {
  emit("update:value", newValue);
});

const handleRemove = () => {
  console.log(">>>> handleRemove");
  imageUrl.value = "";
};

const handlePreview = () => {
  dialogVisible.value = true;
};

// 读取文件并计算MD5值的函数
const getFileMd5 = file => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    // eslint-disable-next-line
    reader.onload = function(e) {
      const data = e.target.result;
      const md5 = SparkMD5.ArrayBuffer.hash(data);
      resolve(md5);
    };
    // eslint-disable-next-line
    reader.onerror = function(error) {
      reject(error);
    };
    reader.readAsArrayBuffer(file);
  });
};

const httpRequest = async param => {
  console.log(">>>> httpRequest");
  let file = param.file;
  // 获取文件md5值
  const fileMd5 = await getFileMd5(file);
  // 获取签名
  const ossSignData = await getUploadSignApi({
    mode: props.mode,
    fileName: file.name,
    contentMd5: fileMd5
  });
  // 上传文件
  const formData = new FormData();
  Object.entries(ossSignData.data.formData).forEach(([key, value]) => {
    formData.append(key, value);
  });
  formData.append("file", file); // 文件本身
  try {
    const response = await axios.post(ossSignData.data.uploadUrl, formData);
    console.log("httpRequest response", response);
    param.onSuccess(ossSignData.data.imageUrl, file);
  } catch (error) {
    param.onError(error);
  }
};

const beforeUpload = file => {
  return new Promise(resolve => {
    let sizeCheck = file.size / 1024 < 10240; // Assuming UPLOAD_SIZE is 10MB
    if (sizeCheck) {
      resolve(file);
    } else {
      // 压缩到200KB,这里的200就是要压缩的大小,可自定义
      imageConversion.compressAccurately(file, 200).then(res => {
        resolve(res);
      });
    }
  });
};

const handleUploadSuccess = (url, file) => {
  console.log(">>>>>> handleUploadSuccess", url, file);
  // 这里需要根据实际返回的数据结构来获取URL
  imageUrl.value = url;
};
</script>

<style></style>
