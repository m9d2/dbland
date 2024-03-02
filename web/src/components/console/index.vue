<template>
  <div :id="id"></div>
</template>

<script lang="ts" setup>
import * as monaco from "monaco-editor";
import 'monaco-editor/esm/vs/basic-languages/sql/sql.contribution';
import {language} from 'monaco-editor/esm/vs/basic-languages/sql/sql';
import {onMounted, ref, nextTick, onUnmounted} from 'vue'
import {generateRandomString} from '@/common/utils'

let monacoEditor: any
const id = ref()
const emits = defineEmits(['change'])
const props = defineProps({
  tables: {
    type: Array,
    default: () => []
  },
  schemas: {
    type: Array,
    default: () => []
  }
})

onMounted(() => {
  id.value = generateRandomString(8)
  nextTick(() => {
    initEditor()
  })
})

onUnmounted(() => {
  monacoEditor.dispose()
})

const initEditor = () => {
  let localTheme = localStorage.getItem('theme')
  let theme = 'light'
  if (localTheme) {
    if (localTheme == 'dark') {
      theme = 'vs-dark'
    }
  }

  monaco.editor.defineTheme('light', {
    base: 'vs',
    inherit: true,
    rules: [],
    colors: {
      'editor.background': '#ffffff',
      'editor.lineHighlightBackground': '#f6f8fa'
    }
  })

  monacoEditor = monaco.editor.create(document.getElementById(id.value), {
    value: '',
    theme: theme, // 官方自带三种主题vs, hc-black, or vs-dark
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
    wordWrap: 'on',
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
    contextmenu: false, //右键菜单
  })

  monacoEditor.onDidChangeModelContent((e) => {
    emits('change', monacoEditor.getValue());
  })

  monaco.languages.registerCompletionItemProvider('sql', {
    provideCompletionItems() {
      let suggestions: any = [];
      language.keywords.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Keyword,
          insertText: item
        });
      })

      language.operators.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Operator,
          insertText: item
        });
      })

      props.tables?.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.File,
          insertText: item,
          documentation: 'table'
        });
      })

      props.schemas?.forEach((item: any) => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Folder,
          insertText: item
        });
      })
      return {
        suggestions: suggestions
      };
    },
  });
}
</script>