<template>
  <div class="setting">
    <div class="setting-left">
      <el-menu style="height: 100%; font-size: 26px;"
        default-active="1"
        class="el-menu-vertical-demo"
      >
        <el-menu-item index="1" @click="menu = 'base'">
          <el-icon><setting /></el-icon>
          <span>{{ $t('setting.menu.basic') }}</span>
        </el-menu-item>
      </el-menu>
    </div>
    <div class="setting-content">
      <BaseComponent v-if="menu === 'base'"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import {
  Setting,
} from '@element-plus/icons-vue';
import BaseComponent from './base/index.vue'

let menu = ref('base')

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
}
.el-menu-item {
  font-size: 14px;
}

</style>