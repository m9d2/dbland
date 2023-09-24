<template>
  <div class="base-box">
    <el-form :model="form" label-width="auto" label-position="left">
      <el-form-item :label="$t('setting.basic.color')">
        <el-color-picker v-model="form.themeColor" show-alpha/>
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

      <el-form-item :label="$t('setting.basic.theme')">
        <el-radio-group v-model="form.theme" class="ml-4">
          <el-radio label="auto" size="large">{{ $t('setting.basic.auto') }}
          </el-radio>
          <el-radio label="light" size="large">{{ $t('setting.basic.light') }}
          </el-radio>
          <el-radio label="dark" size="large">{{ $t('setting.basic.dark') }}
          </el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">{{ $t('common.save') }}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { useDark, useToggle } from '@vueuse/core'

const defaultColor = 'rgba(25,115,237,1)'
const root = document.documentElement
const form = reactive({
  themeColor: defaultColor,
  language: 'en',
  fontSize: '12px',
  theme: 'light'
})


const isDark = useDark({
  storageKey: "theme",
  valueDark: "dark",
  valueLight: "light",
});

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
  const theme = localStorage.getItem('theme')
  if (theme) {
    form.theme = theme
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

  if (form.theme) {
    localStorage.setItem('theme', form.theme)
    if (form.theme == 'dark') {
      useToggle(isDark);
    } else {
      isDark.value = false
    }
  }

}

</script>

<style scoped>
.base-box {
  padding: 20px
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
</style>