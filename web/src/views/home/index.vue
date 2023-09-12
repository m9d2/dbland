<template>
  <div class="wrapper">
    <div class="left">
      <div class="logo">
        <a href="#" style="color: var(--color-text); font-size: 14px; font-weight: bold">
          <el-image style="width: 40px; height: 40px" src="/src/assets/logo.svg" />
        </a>
      </div>
      <div class="nav">
        <ul>
          <router-link v-for="(item, index) in menuItems" :key="index" :to="item.route" @click="setActiveIndex(index)"
            replace>
            <li :class="{ active: activeIndex === index }">
              <el-icon>
                <component :is="item.icon" :class="{ active: activeIndex === index }" />
              </el-icon>
            </li>
          </router-link>
        </ul>
      </div>
    </div>
    <div class="content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Coin,
  Edit,
  Monitor,
  Setting,
} from '@element-plus/icons-vue'
import { onMounted } from 'vue';

const menuItems = [
  { route: '/', icon: Coin },
  { route: '/connect', icon: Edit },
  // { route: '/chart', icon: Monitor },
  { route: '/setting', icon: Setting }
]
let activeIndex: number

onMounted(() => {
  let color = localStorage.getItem('color-primary')
  if (!color) {
    const rootStyles = getComputedStyle(document.documentElement);
    color = rootStyles.getPropertyValue('--vt-c-primary');
  }
  const root = document.documentElement
  root.style.setProperty('--el-color-primary', color);
  root.style.setProperty('--vt-c-primary', color);
})

const setActiveIndex = (index: number) => {
  activeIndex = index;
}
</script>

<style lang="scss" scoped>
.wrapper {
  height: 100%;
  display: flex;
  overflow: hidden;
  flex-direction: row;

  .left {
    width: 68px;
    min-width: 68px;
    background-color: var(--color-background-deep);
    border: 1px solid var(--color-border);
    border-top: none;
    height: 100vh;

    .logo {
      text-align: center;
      height: 40px;
      width: 100%;
      margin-top: 20px;
    }

    .nav {
      margin-top: 20px;
    }
  }

  .content {
    flex-grow: 1;
    height: 100vh;
    overflow: hidden;
  }
}

ul li:not(:first-child) {
  margin: 8px 0;
  background-color: red;
}

li {
  border-radius: 5px;
  background-color: var(--color-background-deep);
  height: 40px;
  margin: 8px;
  display: flex;
  align-items: center;
  justify-content: center;

  .el-icon {
    font-size: 28px;
    color: var(--color-text);
  }

  .el-icon:hover {
    color: var(--color-text-hover);
  }
}

li:hover {
  background-color: var(--color-background-hover);
}

.active {
  background-color: var(--color-background-hover);
  color: var(--color-text-hover);
}

.iconfont {
  font-family: 'iconfont', serif !important;
  font-size: 30px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>