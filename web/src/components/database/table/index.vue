<template>
  <el-table
      v-loading="loading"
      :data="data"
      stripe size="small"
      border
      empty-text=" "
      :height="height"
      highlight-current-row
      show-overflow-tooltip>
    <template #empty>
      <el-empty :image-size="200"></el-empty>
    </template>
    <el-table-column
        v-for="column in columns"
        :key="column"
        :prop="column"
        :label="column"
        width="200"
        min-width="90"
        sortable="true"
        resizable />
  </el-table>
</template>
<script setup lang="ts">
import type {Column, Data} from './type';
defineProps({
  columns: {
    type: Array as () => Column[]
  },
  data: {
    type: Array as () => Data[]
  },
  loading: {
    type: Boolean,
    required: true,
    default: false,
  },
  height: {
    type: String,
    required: false,
  }
})

defineOptions({
  name: 'Table'
})
</script>
<style lang="scss">
.el-empty__description {
  display: none;
}
.el-scrollbar__view {
  height: 100%;
}
</style>