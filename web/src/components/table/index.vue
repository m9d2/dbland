<template>
  <el-table v-loading="loading" :data="data" size="small" :border="true" empty-text="" :height="height"
    @row-contextmenu="contextmenu" 
    :header-cell-class-name="headerCellClassName" 
    :header-row-class-name="headerRowClassName"
    @current-change="handlerCurrentChange"
    :highlight-current-row="true"
    :show-overflow-tooltip="true">
    <template #empty>
      <el-empty :image-size="200"></el-empty>
    </template>
    <el-table-column v-for="column in columns" :column-key="column.column_name" :prop="column.column_name" :label="column.column_name" width="200" min-width="90"
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
});

defineOptions({
  name: "Table",
});

const emits = defineEmits(["row-contextmenu", "row-current-change"]);

function contextmenu(row: any, column: any, event: any) {
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
</script>
<style lang="scss">
.el-empty__description {
  display: none;
}

.el-scrollbar__view {
  height: 100%;
}

.null-value {
  color: var(--color-text-light);
}

.header-cell {
  background-color: var(--color-background-deep)!important;
  font-weight: 700;
}
.header-row {
  color: var(--color-text);
}

.current-row {
  background-color: green!important;
}
</style>
