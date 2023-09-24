<template>
  <div class="setting-box">
    <div class="setting-left">
      <span style="text-align: center;font-size: 20px; font-weight: bold; margin: 20px 0">{{ $t('setting.title') }}</span>
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
import { ref, onMounted } from 'vue';
import List from '@/components/layout/list/index.vue'
import { useRouter } from 'vue-router'
import i18n from '@/plugins/i18n'

const router = useRouter()
const settings = ref<any>([]);

onMounted(() => {
  settings.value = [
    {
      name: i18n.global.t('setting.menu.basic'),
      path: '/setting'
    },
    {
      name: i18n.global.t('setting.menu.other'),
      path: '/about'
    }
  ];
});

function nodeClick(index: number, row: any) {
  router.replace(row.path)
}
</script>

<style scoped>
.setting-left {
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--db-c-border);
  color: var(--db-c-text-menu);
  background-color: var(--db-c-bg-nav);
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
  overflow: auto;
}

.setting-main {
  width: 80%;
  height: 100%;
}
</style>