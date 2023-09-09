<template>
  <el-button plain>
    <template #default>
      <span class="iconfont">{{icon}}</span>
      <span style="margin-left: 8px">{{name}}</span>
    </template>
  </el-button>
</template>
<script setup lang="ts">
defineProps({
  name: '',
  icon: '',
})
</script>
<style scoped lang="scss">

</style>