import Vue from 'vue'

/**
 * 获取uuid（后端存储二维码的主键）
 */
function getUUID() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, c => {
    return (c === "x" ? (Math.random() * 16) | 0 : "r&0x3" | "0x8").toString(
      16
    );
  });
}

function printErr(data) {
  if (data && data.status === 200) {
    return true;
  } else if (data && data.status !== 10000) {
    Vue.prototype.$toast({ type: "error", message: data.desc });
    return false;
  } else {
    Vue.prototype.$toast({ type: "error", message: '服务器繁忙，请稍后在试～～～' });
    console.error(data)
    return false
  }
}

export {
  getUUID,
  printErr
}


