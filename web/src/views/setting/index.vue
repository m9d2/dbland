<template>
  <div class="setting-box">
    <div class="setting-left">
      <span style="text-align: center;font-size: 20px; font-weight: bold; margin: 20px 0">设置</span>
      <List :list="settings" style="margin: 0 20px;" @node-click="nodeClick"></List>
    </div>
    <div class="setting-content">
      <div class="setting-main">
        <RouterView></RouterView>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import List from '@/components/layout/list/index.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const settings = ref(
  [
    {
      name: '基础设置',
      active: true,
      path: '/setting'
    },
    {
      name: '关于我们',
      path: '/about'
    }
  ]
)

function nodeClick(index: number, row: any) {
  router.replace(row.path)
}
</script>

<style scoped>
.setting-left {
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--color-border);
  color: var(--color-text) !important;
  background-color: var(--color-background-deep);
  position: relative;
  display: flex;
  flex-direction: column;
}

.setting-box {
  display: flex;
  height: 100%;
}

.setting-content {
  flex-grow: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.setting-main {
  width: 80%;
  height: 100%;
}
</style>