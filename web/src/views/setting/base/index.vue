<template>
    <div class="base-box">
        <div class="title-box">
            <span class="title" style="margin-right: 20px;">{{ $t('setting.basic.color') }}</span>
            <el-color-picker v-model="themeColor" @change="setColor" />
            <el-link style="margin-left: 8px;" @click="restore">{{ $t('setting.basic.restore') }}</el-link>
        </div>
        <div class="title-box">
            <span class="title" style="margin-right: 20px;">{{ $t('setting.basic.select_language') }}</span>
            <el-select v-model="language" class="m-2" placeholder="language" size="small" @change="changeLanguage">
                <el-option key="zh" label="中文" value="zh" />
                <el-option key="en" label="English" value="en" />
            </el-select>
        </div>
        <el-divider />
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
const { locale } = useI18n()

const themeColor = ref('#1973ED')
const language = ref('en')

function setColor() {
    const root = document.documentElement
    root.style.setProperty('--color-primary', themeColor.value);
    // save
    localStorage.setItem('color-primary', themeColor.value)
}

onMounted(() => {
    const lang = localStorage.getItem('language')
    if(lang) {
        language.value = lang
    }
    const color = localStorage.getItem('color-primary')
    if (color) {
        themeColor.value = color
    }
})

function restore() {
    const color = '#1973ED'
    const root = document.documentElement
    root.style.setProperty('--color-primary', color);
    themeColor.value = color
    localStorage.setItem('color-primary', color)
}

function changeLanguage() {
    locale.value = language.value;
    localStorage.setItem('language', language.value)
}

</script>

<style scoped>
.base-box {
    padding: 20px
}

.title-box {
    margin-bottom: 20px;
}

.title {
    font-size: 14px;
    font-weight: 700;
}
</style>