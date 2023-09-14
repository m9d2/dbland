<template>
  <div :id="id"></div>
</template>

<script lang="ts" setup>
import * as monaco from 'monaco-editor/esm/vs/editor/edcore.main';
import 'monaco-editor/esm/vs/basic-languages/sql/sql.contribution';
import {language as sqlLanguage} from 'monaco-editor/esm/vs/basic-languages/sql/sql';
import {onMounted, ref, nextTick} from 'vue'
import { generateRandomString } from '@/common/utils'
let editor:any
const id = ref()

const props = defineProps({
  sql: ''
})

onMounted(() => {
  id.value = generateRandomString(8)
  nextTick(() => {
    initEditor()
  })

})

const initEditor = () => {
  editor = monaco.editor.create(document.getElementById(id.value), {
    value: props.sql,
    theme: 'vs', // 官方自带三种主题vs, hc-black, or vs-dark
    minimap: { // 关闭小地图
      enabled: false,
    },
    foldingStrategy: 'indentation', // 代码可分小段折叠
    overviewRulerBorder: false, // 不要滚动条的边框
    scrollbar: { // 滚动条设置
      verticalScrollbarSize: 0, // 竖滚动条
      horizontalScrollbarSize: 6, // 横滚动条
    },
    cursorStyle: 'line', // 光标样式
    fontSize: 12, // 字体大小
    tabSize: 2, // tab缩进长度
    autoIndent: true, // 自动布局
    lineNumbers: 'on',
    autoIndex: true,
    language: 'sql',
    tabCompletion: 'on',
    cursorSmoothCaretAnimation: true,
    formatOnPaste: true,
    mouseWheelZoom: true,
    folding: true,
    automaticLayout: true, // 自动布局
    autoClosingBrackets: 'always',
    autoClosingOvertype: 'always',
    autoClosingQuotes: 'always',
  })

  monaco.languages.registerCompletionItemProvider('sql', {
    provideCompletionItems() {
      let suggestions: any = [];
      sqlLanguage.keywords.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Keyword,
          insertText: item
        });
      })
      sqlLanguage.operators.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Operator,
          insertText: item
        });
      })
      sqlLanguage.builtinFunctions.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Function,
          insertText: item
        });
      })
      sqlLanguage.builtinVariables.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Variable,
          insertText: item
        });
      })
      return {
        suggestions: suggestions
      };
    },
  });
}

defineExpose({
  getValue: () => {return editor.getValue()},
  setValue: (content: string) => editor.setValue(content),
  dispose: () => editor.dispose(),
})
</script>