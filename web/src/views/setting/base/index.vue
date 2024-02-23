<template>
  <div class="base-box">
    <div class="setting-base" v-if="active == 'base'">
      <div class="setting-item">
        <div class="base-title">
          <el-link :underline="false" target="_blank" href="https://github.com/m9d2/dbland/releases">
            <img alt="GitHub Release" src="https://img.shields.io/github/v/release/m9d2/dbland?labelColor=rgba(25%2C115%2C237%2C1)&color=%23fff">
          </el-link>
          <el-link style="margin-left: 8px;" :underline="false" target="_blank" href="https://github.com/m9d2/dbland/releases">
            <img alt="GitHub License" src="https://img.shields.io/github/license/m9d2/dbland">
          </el-link>
          <el-link style="margin-left: 8px;" :underline="false" target="_blank" href="https://github.com/m9d2/dbland">
            <img alt="GitHub Org's stars" src="https://img.shields.io/github/stars/m9d2%2Fdbland">
          </el-link>
        </div>
      </div>
      <el-divider />
      <div class="setting-item">
        <div class="base-title">
          <span>{{ $t('setting.basic.color') }}</span>
        </div>
        <div class="base-content">
          <el-color-picker v-model="themeColor" show-alpha @change="changeColor" />
          <el-link style="margin-left: 8px;" @click="restore">{{ $t('setting.basic.restore') }}</el-link>
        </div>
      </div>
      <div class="setting-item">
        <div class="base-title">
          <span>{{ $t('setting.basic.select_language') }}</span>
        </div>
        <div class="base-content">
          <el-select v-model="language" @change="changeLanguage" size="default">
            <el-option key="en" label="English" value="en" />
            <el-option key="zh" label="中文" value="zh" />
          </el-select>
        </div>
      </div>
      <div class="setting-item">
        <div class="base-title">
          <span>{{ $t('setting.basic.theme') }}</span>
        </div>
        <div class="base-content">
          <el-select v-model="theme" @change="changeTheme" size="default">
            <el-option key="auto" :label="$t('setting.basic.auto')" value="auto" />
            <el-option key="light" :label="$t('setting.basic.light')" value="light" />
            <el-option key="dark" :label="$t('setting.basic.dark')" value="dark" />
          </el-select>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useDark, useToggle } from '@vueuse/core'

let localTheme = localStorage.getItem('theme')
let defaultColor = 'rgba(25,115,237,1)'
if (localTheme) {
  if (localTheme == 'dark') {
    defaultColor = 'rgba(64, 158, 255, 1)'
  }
}

const root = document.documentElement
const themeColor = ref(defaultColor)
const language = ref('en')
const fontSize = ref('12px')
const theme = ref('auto')
const active = ref('base')
const isDark = useDark({
  storageKey: "theme",
  valueDark: "dark",
  valueLight: "light",
});

onMounted(() => {
  const lang = localStorage.getItem('language')
  if (lang) {
    language.value = lang
  }
  const color = localStorage.getItem('color-primary')
  if (color) {
    themeColor.value = color
  }
  const size = localStorage.getItem('font-size')
  if (size) {
    fontSize.value = size
  }
  const themeValue = localStorage.getItem('theme')
  if (themeValue) {
    theme.value = themeValue
  }
})



function restore() {
  root.style.setProperty('--db-c-primary', defaultColor);
  themeColor.value = defaultColor
  localStorage.setItem('color-primary', defaultColor)
}

function changeColor(value: any) {
  root.style.setProperty('--db-c-primary', value);
  localStorage.setItem('color-primary', value)
}


function changeLanguage(value: any) {
  localStorage.setItem('language', value)
  window.location.reload()
}

function changeTheme(value: any) {
  localStorage.setItem('theme', value)
  if (value == 'dark') {
    isDark.value = true
    useToggle(isDark);
  } else {
    isDark.value = false
  }
}

function changeFontSize(value: any) {
  localStorage.setItem('font-size', value)
}

</script>

<style scoped>
.base-box {
  padding: 0 20px;
}

.preview {
  width: 134px;
  height: 48px;
}

.ui-dark {
  background-color: #2E2E2E;
}

.ui-light {
  background-color: #DFDFDF;
}

.setting-item {
  margin-bottom: 16px;

  .el-radio-group {
    height: 32px;
  }

  .el-input {
    height: 32px;
  }
}

.base-title {
  margin-bottom: 8px;
  font-size: var(--font-size);
  color: var(--db-c-text);
  font-weight: bold;
}
</style>