<template>
  <el-table v-loading="loading" :data="data" size="small" :border="true" empty-text="" :height="height"
    @row-contextmenu="handlerContextmenu" 
    :header-cell-class-name="headerCellClassName"
    :element-loading-svg="svg"
    :header-row-class-name="headerRowClassName"
    @current-change="handlerCurrentChange"
    :highlight-current-row="true"
    @sort-change="sortChange"
    :show-overflow-tooltip="true">
    <template #empty>
      <el-empty :image-size="200"></el-empty>
    </template>
    <el-table-column v-for="column in columns" :column-key="column.field" :prop="column.field" :label="column.field" :width="columnWidth"
      sortable="true" resizable>
      <template #default="{ row, column }">
        <span :class="{ 'null-value': row[column.property] == null, }">
          {{ row[column.property] != null ? row[column.property] : '(NULL)' }}
        </span>
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup lang="ts">
import {ref} from 'vue'

defineProps({
  columns: {
    type: Array as () => any[],
  },
  data: {
    type: Array as () => any[],
  },
  loading: {
    type: Boolean,
    required: true,
    default: false,
  },
  height: {
    type: String,
    required: false,
  },
  columnWidth: {
    type: Number,
    required: false,
  }
});

defineOptions({
  name: "Table",
});
const svg = '<svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ea893728=""><path fill="currentColor" d="M512 64a32 32 0 0 1 32 32v192a32 32 0 0 1-64 0V96a32 32 0 0 1 32-32zm0 640a32 32 0 0 1 32 32v192a32 32 0 1 1-64 0V736a32 32 0 0 1 32-32zm448-192a32 32 0 0 1-32 32H736a32 32 0 1 1 0-64h192a32 32 0 0 1 32 32zm-640 0a32 32 0 0 1-32 32H96a32 32 0 0 1 0-64h192a32 32 0 0 1 32 32zM195.2 195.2a32 32 0 0 1 45.248 0L376.32 331.008a32 32 0 0 1-45.248 45.248L195.2 240.448a32 32 0 0 1 0-45.248zm452.544 452.544a32 32 0 0 1 45.248 0L828.8 783.552a32 32 0 0 1-45.248 45.248L647.744 692.992a32 32 0 0 1 0-45.248zM828.8 195.264a32 32 0 0 1 0 45.184L692.992 376.32a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0zm-452.544 452.48a32 32 0 0 1 0 45.248L240.448 828.8a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0z"></path></svg>'
const emits = defineEmits(["row-contextmenu", "row-current-change", "sort-change"]);

function handlerContextmenu(row: any, column: any, event: any) {
  emits("row-contextmenu", row, column, event);
}

function headerCellClassName(row: any) {
  return 'header-cell'
}

function headerRowClassName(row: any) {
  return 'header-row'
}
function handlerCurrentChange(currentRow: any) {
  emits("row-current-change", currentRow);
}
function sortChange({ column, prop, order }: any) {
  emits("sort-change", {column, prop, order});
}
</script>
<style lang="scss">
.el-empty__description {
  display: none;
}

.el-scrollbar__view {
  height: 100%;
}

.null-value {
  color: var(--db-c-text-light-2);
}

</style>
