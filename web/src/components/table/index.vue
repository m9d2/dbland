<template>
  <el-table v-loading="loading" :data="data" size="small" :border="true" :height="height"
            @row-contextmenu="handlerContextmenu"
            :header-cell-class-name="propsData.headerCellClassName"
            :header-row-class-name="propsData.headerRowClassName"
            @current-change="handlerCurrentChange"
            :header-cell-style="{background:'#f6f8fa'}"
            :highlight-current-row="true"
            element-loading-background="rgba(0, 0, 0, 0)"
            element-loading-svg-view-box="0, 0, 20, 20"
            :element-loading-svg="svg"
            @sort-change="sortChange"
            :show-overflow-tooltip="true">
    <template #empty>
      <div>

      </div>
    </template>
    <el-table-column v-for="column in columns" :column-key="column.field" :prop="column.field" :label="column.field"
                     :width="columnWidth"
                     :sortable="sortable" resizable>
      <template #default="{ row, column }">
        <span :class="{ 'null-value': row[column.property] == null, }">
          {{ row[column.property] != null ? row[column.property] : '(NULL)' }}
        </span>
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup lang="ts">
const propsData = defineProps({
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
  },
  sortable: {
    type: Boolean,
    required: false,
    default: false,
  },
  headerRowClassName: {
    type: String,
    default: 'custom-header-row',
  },
  headerCellClassName: {
    type: String,
    default: 'custom-header-cell',
  }
});

defineOptions({
  name: "Table",
});
const svg = '<svg t="1709349712442" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="24881" width="20" height="20"><path d="M307.146667 821.333333a39.253333 39.253333 0 1 0 53.333333 14.346667 39.253333 39.253333 0 0 0-53.333333-14.346667zM496 874.666667a34.88 34.88 0 1 0 47.68 12.746666A34.826667 34.826667 0 0 0 496 874.666667zM685.12 828.586667a30.346667 30.346667 0 1 0 41.386667 11.093333 30.346667 30.346667 0 0 0-41.386667-11.093333zM824.906667 697.333333a24.106667 24.106667 0 1 0 32.906666 8.853334 24.16 24.16 0 0 0-32.906666-8.853334zM879.573333 518.666667a14.826667 14.826667 0 1 0 20.32 5.333333 14.88 14.88 0 0 0-20.32-5.333333zM168.693333 680.746667a43.2 43.2 0 1 0 59.04 16 43.253333 43.253333 0 0 0-59.04-16zM161.28 294.613333a58.026667 58.026667 0 1 0 79.306667 21.333334 58.08 58.08 0 0 0-79.306667-21.333334zM182.773333 506.986667a49.12 49.12 0 1 0-17.973333 67.093333 49.173333 49.173333 0 0 0 17.973333-67.093333zM476.586667 94.08a74.026667 74.026667 0 1 0 101.333333 27.093333 74.026667 74.026667 0 0 0-101.333333-27.093333zM294.24 151.573333a65.386667 65.386667 0 1 0 89.333333 23.946667 65.386667 65.386667 0 0 0-89.333333-23.946667zM742.72 282.186667a85.333333 85.333333 0 1 0-116.693333-31.253334 85.333333 85.333333 0 0 0 116.693333 31.253334zM920.053333 296.853333a96 96 0 1 0-35.146666 131.146667 96 96 0 0 0 35.146666-131.146667z" p-id="24882" fill="#bfbfbf"></path></svg>'
const emits = defineEmits(["row-contextmenu", "row-current-change", "sort-change"]);

function handlerContextmenu(row: any, column: any, event: any) {
  emits("row-contextmenu", row, column, event);
}

function handlerCurrentChange(currentRow: any) {
  emits("row-current-change", currentRow);
}

function sortChange({column, prop, order}: any) {
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

:deep(.el-loading-spinner .circular) {
  height: 16px;
  width: 16px;

}

.null-value {
  color: var(--db-c-text-light-2);
}
</style>
