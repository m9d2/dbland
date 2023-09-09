<template>
  <ul>
    <router-link v-for="item in data" :to="item.path">
      <li>
          <span class="iconfont">
            {{ item.icon }}
          </span>
      </li>
    </router-link>
  </ul>
</template>

<script setup lang="ts">
defineProps({
  data: {
    type: Array,
    required: true
  },
})
</script>

<style lang="scss" scoped>
ul li:not(:first-child) {
  margin: 8px 0;
  background-color: red;
}

li {
  border-radius: 5px;
  background-color: var(--color-background-deep);
  text-align: center;
  height: 40px;
  line-height: 40px;
  margin: 8px;
}

li:hover {
  background-color: var(--color-background-hover)
}

span {
  color: var(--color-text)
}

span:hover {
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