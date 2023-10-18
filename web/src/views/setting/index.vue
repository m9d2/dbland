<template>
  <div class="setting">
    <div class="setting-left">
      <List :list="settings" style="background-color: var(db-c-bg);" :rowStyle="rowStyle">
        <template #begin="{ node }">
          <el-icon style="margin: 8px">
            <span v-show="node.name == 'Basic Settings'">
              <Setting />
            </span>
            <span v-show="node.name == 'About Us'">
              <ChatDotSquare />
            </span>
          </el-icon>
        </template>
      </List>
    </div>
    <div class="setting-content">
      <BaseSetting></BaseSetting>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import List from '@/components/layout/list/index.vue';
import { useRouter } from 'vue-router';
import {
  Setting,
  ChatDotSquare,
} from '@element-plus/icons-vue';
import i18n from '@/plugins/i18n';
import BaseSetting from '@/views/setting/base/index.vue'

const router = useRouter()
const settings = ref<any>([]);

onMounted(() => {
  settings.value = [
    {
      name: i18n.global.t('setting.menu.basic')
    },
    {
      name: i18n.global.t('setting.menu.other')
    }
  ];
});
const rowStyle = {
  height: '46px',

}
</script>

<style scoped>
.setting {
  display: flex;
  height: 480px;
}

.setting-left {
  height: 100%;
  width: 220px;
  color: var(--db-c-text-menu);
  position: relative;
  display: flex;
  flex-direction: column;
}

.setting-content {
  flex-grow: 1;
  overflow: auto;
}</style>