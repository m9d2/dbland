<template>
  <div class="base-box">
    <el-form :model="form" label-width="auto" label-position="left">
      <el-form-item :label="$t('setting.basic.color')">
        <el-color-picker v-model="form.themeColor" />
        <el-link style="margin-left: 8px;" @click="restore">{{ $t('setting.basic.restore') }}</el-link>
      </el-form-item>
      <el-form-item :label="$t('setting.basic.select_language')">
        <el-select v-model="form.language" class="m-2" placeholder="language" size="small">
          <el-option key="zh" label="中文" value="zh" />
          <el-option key="en" label="English" value="en" />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('setting.basic.font_size')">
        <el-input v-model="form.fontSize" size="small" style="width: 181px;" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">{{ $t('common.save') }}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'

const defaultColor = '#1973ED'
const root = document.documentElement
const form = reactive({
  themeColor: defaultColor,
  language: 'en',
  fontSize: '12px',
})

onMounted(() => {
  const lang = localStorage.getItem('language')
  if (lang) {
    form.language = lang
  }
  const color = localStorage.getItem('color-primary')
  if (color) {
    form.themeColor = color
  }
  const size = localStorage.getItem('font-size')
  if (size) {
    form.fontSize = size
  }
})

function restore() {
  root.style.setProperty('--db-c-primary', defaultColor);
  form.themeColor = defaultColor
  localStorage.setItem('color-primary', defaultColor)
}


function onSubmit() {
  // color
  root.style.setProperty('--db-c-primary', form.themeColor);
  localStorage.setItem('color-primary', form.themeColor)

  // language
  localStorage.setItem('language', form.language)

  // font size
  if (form.fontSize) {
    localStorage.setItem('font-size', form.fontSize)
  }

}

</script>

<style scoped>
.base-box {
  padding: 20px
}
</style>