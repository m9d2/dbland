<template>
  <div class="main">
    <div class="left">
      <div class="logo">
        <a href="#" style="color: var(--db-c-text); font-size: 14px; font-weight: bold">
          <el-image style="width: 40px; height: 40px" :src="logo" />
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
import logo from '@/assets/img/logo.svg'

const menuItems = [
  { route: '/', icon: Coin },
  { route: '/connect', icon: Edit },
  { route: '/chart', icon: Monitor },
  { route: '/setting', icon: Setting }
]
let activeIndex: number

onMounted(() => {
  const root = document.documentElement
  // theme color
  let color = localStorage.getItem('color-primary')
  if (color) {
    root.style.setProperty('--db-c-primary', color);
  }

  // font size
  let fontSize = localStorage.getItem('font-size')
  if (fontSize) {
    root.style.setProperty('--font-size', fontSize);
  }

})

const setActiveIndex = (index: number) => {
  activeIndex = index;
}
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  display: flex;
  overflow: hidden;
  flex-direction: row;

  .left {
    width: 68px;
    min-width: 68px;
    background-color: var(--db-c-bg-nav);
    border: 1px solid var(--db-c-border);
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
  background-color: var(--db-c-bg-nav);
  height: 40px;
  margin: 8px;
  display: flex;
  align-items: center;
  justify-content: center;

  .el-icon {
    font-size: 28px;
    color: var(--db-c-text-nav);
  }

  .el-icon:hover {
    color: var(--db-c-text-hover);
  }
}

li:hover {
  background-color: var(--db-c-bg-hover);
}

.active {
  background-color: var(--db-c-bg-hover);
  color: var(--db-c-text-hover);
}

.iconfont {
  font-family: 'iconfont', serif !important;
  font-size: 30px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>