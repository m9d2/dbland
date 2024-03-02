<template>
    <div class="left">
        <div class="top">
            <div class="logo">
                <a href="#" style="color: var(--db-c-text); font-size: 14px; font-weight: bold">
                    <el-image style="width: 30px; height: 30px" :src="logo" />
                </a>
            </div>
            <div class="nav">
                <ul>
                    <router-link v-for="(item, index) in menuItems" :key="index" :to="item.route"
                        @click="setActiveIndex(index, item)" replace>
                        <li :class="{ active: activeIndex === index }">
                            <el-icon>
                                <component :is="item.icon" />
                            </el-icon>
                        </li>
                    </router-link>
                </ul>
            </div>
        </div>
        <div class="setting centered">
            <el-link @click="setting" :underline="false">
                <el-icon>
                    <Setting />
                </el-icon>
            </el-link>
        </div>
    </div>
    <el-dialog v-model="settingVisible" :close-on-click-modal="false" :title="$t('setting.title')" width="60%" align-center>
        <SettingView></SettingView>
    </el-dialog>
</template>

<script setup lang="ts">
import {
    Coin,
    Edit,
    Setting,
    Monitor,
} from '@element-plus/icons-vue'
import { ref } from 'vue';
import logo from '@/assets/img/logo.png';
import { useRoute } from 'vue-router';
import SettingView from '@/views/setting/index.vue';

const menuItems = [
    { route: '/', icon: Coin },
    { route: '/connect', icon: Edit },
    { route: '/chart', icon: Monitor },
]
const settingVisible = ref(false)
const route = useRoute();
let activeIndex: number = findMenuItemIndex(route.path)

function findMenuItemIndex(route: string): number {
    return menuItems.findIndex(item => item.route === route);
}

const setActiveIndex = (index: number, item: any) => {
    activeIndex = index;
}

const setting = () => {
    settingVisible.value = true
}
</script>

<style lang="scss" scoped>
.left {
    display: flex;
    flex-direction: column;
    width: 68px;
    min-width: 68px;
    background-color: var(--db-c-bg-nav);
    border: 1px solid var(--db-c-border);
    border-top: none;
    box-sizing: border-box;

    .top {
        flex-grow: 1;

        .logo {
            text-align: center;
            height: 40px;
            width: 100%;
            margin-top: 20px;
        }

        .nav {
            margin-top: 8px;
        }
    }

    .setting {
        margin-bottom: 20px;
    }
}

ul li:not(:first-child) {
    margin: 8px 0;
}

li {
    background-color: var(--db-c-bg-nav);
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    border: 2px solid var(--db-c-bg-nav);
}

li:hover {
    background-color: var(--db-c-bg-hover);
}

.el-icon {
    font-size: 28px;
    color: var(--db-c-text-nav);
}

.el-icon:hover {
    color: var(--db-c-text-hover);
}

.active {
    background-color: var(--db-c-bg-hover);
    color: var(--db-c-text-hover);
    border-left: 2px solid var(--db-c-text-hover);
}

.iconfont {
    font-family: 'iconfont', serif !important;
    font-size: 30px;
    font-style: normal;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}
</style>