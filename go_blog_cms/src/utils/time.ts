import moment from "moment"; // pnpm install moment --save

/**
 * 格式化时间戳为日期字符串
 * @param timestamp 时间戳（秒）
 * @returns 格式化的日期字符串，格式为 "YYYY-MM-DD HH:mm:ss"，如果时间戳为0，则返回空字符串
 */
export const formatTimestamp = (timestamp: number) => {
  // 如果时间戳为0，返回空字符串
  if (timestamp === 0) {
    return "N/A";
  }
  // 使用 moment 库将时间戳（秒）转换为日期字符串
  return moment(timestamp * 1000).format("YYYY-MM-DD HH:mm:ss");
};
