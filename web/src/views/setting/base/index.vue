<template>
    <div class="base-box">
        <div class="title-box">
            <span class="title" style="margin-right: 20px;">设置主题颜色</span>
            <el-color-picker v-model="themeColor" @change="setColor" />
            <el-link style="margin-left: 8px;" @click="restore">恢复默认</el-link>
        </div>
        <div class="title-box">
            <span class="title" style="margin-right: 20px;">选择语言</span>
            <el-select v-model="defaultLanguage" class="m-2" placeholder="language" size="small">
                <el-option key="chinese" label="中文" value="chinese" />
                <el-option key="chinese" label="English" value="english" />
            </el-select>
        </div>
        <el-divider />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const themeColor = ref('#1973ED')
const defaultLanguage = ref('chinese')

function setColor() {
    const root = document.documentElement
    root.style.setProperty('--el-color-primary', themeColor.value);
    root.style.setProperty('--vt-c-primary', themeColor.value);

    // save
    localStorage.setItem('color-primary', themeColor.value)
}

function restore() {
    const color = '#1973ED'
    const root = document.documentElement
    root.style.setProperty('--el-color-primary', color);
    root.style.setProperty('--vt-c-primary', color);

    themeColor.value = color

    // save
    localStorage.setItem('color-primary', color)
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