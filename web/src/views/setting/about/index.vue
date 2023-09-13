<template>
    <div v-html="renderedMarkdown" style="margin-top: 20px;">
    </div>
</template>

<script setup lang="ts">
import MarkdownIt from 'markdown-it';
import { onMounted, ref } from 'vue';
import { readFileContent } from '@/common/utils'
import { ElNotification } from "element-plus";
import axios from 'axios';

const markdownContent = ref()
const renderedMarkdown = ref()
onMounted(() => {
    fetchMarkdownFile();
})

function fetchMarkdownFile() {
    markdownContent.value = "# hello"
    renderMarkdown();
}

function renderMarkdown() {
    const md = new MarkdownIt();
    renderedMarkdown.value = md.render(markdownContent.value);
}
</script>

<style scoped></style>