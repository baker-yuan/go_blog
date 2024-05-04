<template>
  <div class="editor-flex-container">
    <div ref="jsonEditorContainer" class="json-editor-container"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref, watch } from "vue";
import JSONEditor, { JSONEditorOptions } from "jsoneditor";
import "jsoneditor/dist/jsoneditor.min.css";

// 定义 props 的类型
interface JsonEditorProps {
  modelValue: any;
}

// 定义 emits
const emits = defineEmits(["update:modelValue"]);

// 接收外部传入的 JSON 数据
const props = defineProps<JsonEditorProps>();

// 创建一个 DOM 引用
const jsonEditorContainer = ref<HTMLElement | null>(null);

// 创建 JSONEditor 实例的变量
let jsonEditor: JSONEditor | null = null;

onMounted(() => {
  if (jsonEditorContainer.value) {
    // 定义 JSONEditor 的配置选项
    const options: JSONEditorOptions = {
      mode: "code", // 设置为纯代码编辑模式
      mainMenuBar: false, // 隐藏主菜单栏
      onChange: () => {
        if (jsonEditor) {
          // 当编辑器内容变化时，发出事件更新外部变量
          emits("update:modelValue", JSON.parse(jsonEditor.getText()));
        }
      }
    };
    // 实例化 JSONEditor
    jsonEditor = new JSONEditor(jsonEditorContainer.value, options);
    // 设置编辑器的初始内容
    jsonEditor.setText(JSON.stringify(props.modelValue, null, 2));
  }
});

// 监听外部数据的变化，并更新编辑器内容
watch(
  () => props.modelValue,
  newValue => {
    if (jsonEditor) {
      jsonEditor.updateText(JSON.stringify(newValue, null, 2));
    }
  },
  { deep: true }
);

onBeforeUnmount(() => {
  // 销毁 JSONEditor 实例
  if (jsonEditor) {
    jsonEditor.destroy();
    jsonEditor = null;
  }
});
</script>

<style>
/* 设置外层容器使用 Flexbox 布局 */
.editor-flex-container {
  display: flex;
  flex-direction: column;
  height: 85vh; /* 或者使用其他高度，如 100% 如果父元素有明确的高度 */
}

/* 设置 JSONEditor 容器的 flex 属性为 1，使其填满剩余空间 */
.json-editor-container {
  flex: 1;
  /* 如果需要，可以设置最小高度 */
  min-height: 200px; /* 例如，设置一个最小高度 */
}
</style>
