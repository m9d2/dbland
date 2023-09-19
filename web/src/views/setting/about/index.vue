<template>
    <div v-html="markdownContent" style="width: 100%;margin: 20px 0;">
    </div>
</template>

<script setup lang="ts">
import MarkdownIt from 'markdown-it';
import { onMounted, ref } from 'vue';
import { getReadme } from '@/api/other'

const markdownContent = ref()

onMounted(() => {
    fetchMarkdownFile();
})

async function fetchMarkdownFile() {
      const { data } = await getReadme()
      const md = new MarkdownIt({
        html: true,
      })
      markdownContent.value = md.render(data)
}
</script>

<style scoped></style>