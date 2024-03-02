<template>
  <div class="main">
    <Menu></Menu>
    <div class="content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useDark, useToggle } from '@vueuse/core'
import Menu from './layout/Menu.vue';

const isDark = useDark({
  storageKey: "theme",
  valueDark: "dark",
  valueLight: "light",
});

onMounted(() => {
  const root = document.documentElement
  // theme color
  let color = localStorage.getItem('color-primary')
  if (color) {
    root.style.setProperty('--db-c-primary', color);
    const currentColor = root.style.getPropertyValue('--db-c-primary').trim();
    const rgbaArray = currentColor.match(/\d+/g);
    if (rgbaArray) {
      rgbaArray[3] = '0.8';
      const newColor = "rgba(" + rgbaArray.join(", ") + ")";
      root.style.setProperty('--el-color-primary-light-3', newColor);
      root.style.setProperty('--el-color-primary-light-5', newColor);
      root.style.setProperty('--el-color-primary-dark-2', newColor);
      root.style.setProperty('--el-color-primary-light-7', color);
    }
  }

  // font size
  let fontSize = localStorage.getItem('font-size')
  if (fontSize) {
    root.style.setProperty('--font-size', fontSize);
  }

  let theme = localStorage.getItem('theme')
  if (theme) {
    useToggle(isDark)
  }
})
</script>

<style lang="scss" scoped>
.main {
  height: 100vh;
  display: flex;
  overflow: hidden;
  flex-direction: row;

  .content {
    flex-grow: 1;
    overflow: hidden;
  }
}
</style>