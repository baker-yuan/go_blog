import { OssUploadRecord } from "@/api/interface";

export const formatOssType = (ossType: number) => {
  if (ossType === OssUploadRecord.OssType.OssType_ALI_YUN) {
    return "阿里云";
  }
  return "未知";
};

export const formatOssStatus = (ossStatus: number) => {
  if (ossStatus === OssUploadRecord.OssStatus.OssStatus_UNPROCESSED) {
    return "未处理";
  } else if (ossStatus == OssUploadRecord.OssStatus.OssStatus_SUCCESS) {
    return "上传成功";
  } else if (ossStatus == OssUploadRecord.OssStatus.OssStatus_FAILURE) {
    return "上传失败";
  }
  return "";
};

export const getOssStatusType = (ossStatus: number) => {
  if (ossStatus === OssUploadRecord.OssStatus.OssStatus_UNPROCESSED) {
    return "info"; // 未处理状态显示为蓝色
  } else if (ossStatus == OssUploadRecord.OssStatus.OssStatus_SUCCESS) {
    return "success"; // 上传成功状态显示为绿色
  } else if (ossStatus == OssUploadRecord.OssStatus.OssStatus_FAILURE) {
    return "danger"; // 上传失败状态显示为红色
  }
  return "";
};

// 上传状态
export const ossStatusOptions = [
  {
    label: "未处理",
    value: OssUploadRecord.OssStatus.OssStatus_UNPROCESSED
  },
  {
    label: "上传成功",
    value: OssUploadRecord.OssStatus.OssStatus_SUCCESS
  },
  {
    label: "上传失败",
    value: OssUploadRecord.OssStatus.OssStatus_FAILURE
  }
];
