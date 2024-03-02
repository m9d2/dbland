<template>
  <div class="query-main">
    <div class="query-content">
      <div class="condition">
        <span>WHERE</span>
        <input v-model="whereSql" class="where" @keyup.enter="handleEnterKey" />
        <div class="enter">
          <span class="iconfont icon-enter"></span>
        </div>
      </div>
      <div class="table-tool" v-if="showQueryStatus">
        <div class="clearfix">
          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.add')">
              <el-link :underline="false" @click="insertRow" :icon="Plus" />
            </el-tooltip>
          </div>

          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.delete')">
              <el-link :underline="false" @click="deleteRecord" :disabled="!currentRow" :icon="Minus" />
            </el-tooltip>
          </div>

          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.modify')">
              <el-link :underline="false" @click="modifyRow" :disabled="!currentRow" :icon="Edit" />
            </el-tooltip>
          </div>

          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.refresh')">
              <el-link :underline="false" @click="refresh" :icon="Refresh" />
            </el-tooltip>
          </div>

<!--          <div class="menu fr">-->
<!--            <el-dropdown>-->
<!--              <el-button size="small">-->
<!--                Import<el-icon class="el-icon&#45;&#45;right"><arrow-down /></el-icon>-->
<!--              </el-button>-->
<!--              <template #dropdown>-->
<!--                <el-dropdown-menu>-->
<!--                  <el-dropdown-item>CSV</el-dropdown-item>-->
<!--                  <el-dropdown-item>Insert SQL</el-dropdown-item>-->
<!--                  <el-dropdown-item>Excel</el-dropdown-item>-->
<!--                </el-dropdown-menu>-->
<!--              </template>-->
<!--            </el-dropdown>-->
<!--          </div>-->

        </div>
      </div>
      <Table class="query-table" header-row="header-row" header-cell="header-cell" :sortable="true"
        :columns="tableColumns" :data="tableData" :loading="loading" @sort-change="sortChange" :column-width="columnWidth"
        @row-current-change="handlerCurrentChange" @row-contextmenu="handlerContextmenu">
      </Table>

      <Form v-model="showModify" :row="currentRow" :actionType="actionType" :columns="tableColumns"
        @cancel="handlerCancel" @confirm="handlerConfirm"></Form>

      <div class="query-status unselectable">
        <div class="status-left">
          <span v-show="showQueryStatus">{{ $t('database.label.elapsed_time') }}: {{ elapsedTime }}s</span>
        </div>
        <div class="status-center">
          <span v-show="showQueryStatus">{{ sqlStr }}</span>
        </div>
        <div class="status-right" v-show="showQueryStatus">
          <span v-show="showQueryStatus" style="margin-left: 8px">{{ $t('database.label.total') }}: {{ total }}</span>
          <el-link style="margin-left: 8px;" :icon="ArrowLeftBold" :underline="false" @click="arrowLeft" />
          <el-input v-model="page" size="small" style="width: 36px; margin-left: 8px;" @change="changePage"></el-input>
          <el-link style="margin-left: 8px;" :icon="ArrowRightBold" :underline="false" @click="arrowRight" />
        </div>
      </div>
    </div>

    <el-dialog v-model="showDeleteConfirmDialog" width="300" top="35vh" :show-close="false" :close-on-click-modal="false">
      <span style="color: var(--db-c-text);">{{ $t('message.delete_record') }}</span>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="deleteRecord" style="margin-right: 8px;color: var(--db-c-text);">
            {{ $t('common.confirm') }}
          </el-button>
          <el-button type="primary" @click="showDeleteConfirmDialog = false">
            {{ $t('common.cancel') }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import Table from "@/components/table/index.vue";
import Form from "@/components/form/index.vue";
import { query, execute } from "@/api/connector";
import type { QueryReq } from "@/api/connector/type";
import { ElNotification } from "element-plus";
import ContextMenu from "@imengyu/vue3-context-menu";
import { ArrowDown, ArrowLeftBold, ArrowRightBold } from '@element-plus/icons-vue'
import {
  Plus,
  Edit,
  Minus,
  Refresh,
} from '@element-plus/icons-vue'
import { ActionTypeEnum, Direction } from '@/common/enums/'
import { CreateDatabase } from '@/service'
import type { Database } from "@/service/database";
import type { Sort } from "@/service/model/sort";

const tableData = ref();
const tableColumns = ref();
const loading = ref(false);
const elapsedTime = ref<number>();
const showQueryStatus = ref(false);
const total = ref();
const currentRow = ref();
let actionType: any
const showModify = ref(false)
const contentHeight = ref(260);
const page = ref(1)
const pageSize = ref(100)
const totalPage = ref()
const columnWidth = ref(160)
const showDeleteConfirmDialog = ref(false)
const whereSql = ref('')
const sqlStr = ref()
let table: string = '';
let database: string = '';
let databaseType: string = '';
let cid: number = 0;
const propsData = defineProps({
  node: {
    type: Object,
  }
});
const resizeData = reactive({
  isResizing: false,
  startY: 0,
  startHeight: 0,
});
let sort: Sort

const insertRow = () => {
  showModify.value = true
  actionType = ActionTypeEnum.INSERT
}

function modifyRow() {
  showModify.value = true
  actionType = ActionTypeEnum.MODIFY
}

async function refresh() {
  await queryPage()
}

const handlerCurrentChange = (row: any) => {
  currentRow.value = row
}

const handlerCancel = () => {
  showModify.value = false
}

const handleEnterKey = () => {
  queryPage()
}

async function handlerConfirm(formData: any) {
  const db: Database = CreateDatabase(databaseType, database)
  if (actionType == ActionTypeEnum.INSERT) {
    sqlStr.value = db.createInsertSql(table, tableColumns.value, formData.value)
  }
  if (actionType == ActionTypeEnum.MODIFY) {
    if (!currentRow.value) {
      ElNotification({
        message: 'No rows selected',
        type: 'error'
      })
    }
    sqlStr.value = db.createUpdateSql(table, tableColumns.value, formData.value, currentRow.value)
  }
  try {
    const { data } = await execute({ cid: cid, sql: sqlStr.value })
    await queryPage()
    ElNotification({
      message: 'Affected rows: ' + data,
      type: "success",
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
  showModify.value = false
}

onMounted(() => {
  const height = localStorage.getItem('height')
  if (height) {
    contentHeight.value = parseInt(height)
  }
  if (propsData.node) {
    databaseType = propsData.node.parent.parent.data.type;
    database = propsData.node.parent.data.name
    table = propsData.node.data.name
    cid = propsData.node.parent.parent.data.id
    queryPage()
  }
});

async function queryPage() {
  loading.value = true;
  let where = ''
  if (whereSql.value) {
    where = 'WHERE ' + whereSql.value
  }
  const db: Database = CreateDatabase(databaseType, database)
  sqlStr.value = db.createQueryPageSql(table, page.value, pageSize.value, where, sort)
  try {
    const params: QueryReq = {
      cid: cid,
      database: database,
      sql: sqlStr.value,
      table: table,
    };
    const { data } = await query(params);

    tableColumns.value = data.columns;

    let maxWidth = 0;
    for (const column of tableColumns.value) {
      const fieldLength = String(column.field).length;
      if (fieldLength > maxWidth) {
        maxWidth = fieldLength;
      }
    }

    tableData.value = data.rows;
    total.value = data.total;
    elapsedTime.value = data.elapsed_time;
    showQueryStatus.value = true;
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  } finally {
    setTimeout(() => {
      loading.value = false;
    }, 200);
  }
}

function handlerContextmenu(row: any, column: any, event: any) {
  event.preventDefault();
  ContextMenu.showContextMenu({
    theme: 'flat',
    x: event.x,
    y: event.y,
    items: [
      {
        label: "Copy as",
        children: [
          {
            label: "Insert Statement",
            onClick: () => {
              console.log("Insert Statement");
            },
          },
          {
            label: "Update Statement",
            onClick: () => {
              console.log("Update Statement");
            },
          },
        ],
      },
      {
        label: "Delete Record",
        onClick: async () => {
          showDeleteConfirmDialog.value = true
        },
      },
      {
        label: "Modify Record",
        onClick: async () => {
          modifyRow()
        },
      },
    ],
  });
}

function deleteRecord() {
  try {
    const db: Database = CreateDatabase(databaseType, database)
    const sql = db.createDeleteSql(table, tableColumns.value, currentRow.value)

    execute({ sql: sql, cid: cid }).then((res: any) => {
      ElNotification({
        message: 'Affected rows: ' + res.data,
        type: "success",
      });
      queryPage()
    })
    queryPage()
  } catch (error: any) {
    ElNotification({
      title: "error",
      message: error.message,
      type: "error",
    });
  } finally {
    showDeleteConfirmDialog.value = false
  }
}

function arrowLeft() {
  if (page.value > 1) {
    page.value = Number(page.value) - 1
    queryPage()
  }
}

function arrowRight() {
  page.value = Number(page.value) + 1
  queryPage()
}

function sortChange({ column, prop, order }: any) {
  if (!prop || !order) {
    return;
  }
  if (order === 'descending') {
    sort.direction = Direction.DESC
  } else {
    sort.direction = Direction.ASC
  }
  sort.field = column
  queryPage();
}

function changePage() {
  queryPage()
}
</script>

<style lang="scss" scoped>
.query-main {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.el-button+.el-button {
  margin: 0;
}

.query-tool>* {
  margin-right: 8px;
}

.query-console {
  margin: 8px 0;
  width: 100%;

  .code {
    height: 100%;
  }
}

.table-tool {
  height: 30px;
  min-height: 30px;
  border: 1px solid var(--db-c-border);
  border-bottom: none;
  line-height: 30px;
  //background-color: var(--db-c-bg-nav);

  .menu {
    margin-left: 8px;
    font-size: 12px;
    font-weight: bold;

    .el-link {
      font-size: 14px;
    }
  }

  .menu:last-child {
    margin-right: 8px;
    margin-top: 3px;
  }
}

.query-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;

  .condition {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin: 4px 8px;

    span {
      margin-right: 8px;
      color: var(--db-c-primary);
    }

    .el-input {
      border: none;
    }

    .where {
      outline: none;
      width: 100%;
      border: none;
    }

    .where:focus {
      outline: none;
    }

    .enter {
      width: 60px;
      display: flex;
      justify-content: center;
      align-items: center;

      .iconfont {
        color: var(--db-c-primary);
        font-size: 16px;
      }
    }
  }


}

.query-table {
  flex-grow: 1;
}

.query-status {
  display: flex;
  flex-direction: row;
  height: 36px;
  min-height: 36px;
  line-height: 36px;
  border: 1px solid var(--db-c-border);
  border-top: none;
  background-color: var(--db-c-bg-nav);

  .status-left {
    display: flex;
    align-items: center;
    margin-left: 8px;
  }

  .status-center {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-grow: 1;
  }

  .status-right {
    display: flex;
    align-items: center;
    margin-right: 8px;

    .el-icon {
      color: var(--db-c-text-1);
    }
  }
}

.iconfont {
  color: var(--db-c-text-1);
}

.el-dropdown {
  height: 30px;
}

.mx-context-menu-item .label {
  font-size: 12px;
}

.mx-context-menu {
  font-weight: var(--font-family);
}

.el-icon {
  font-size: 13px;
}

.resize-handle {
  width: 100%;
  height: 6px;
}

.resize-handle:hover {
  cursor: row-resize;
  background-color: var(--db-c-border);

}

.el-select {
  font-size: var(--font-size);
}

:deep(.header-cell) {
  .cell {
    overflow: hidden;
    white-space: nowrap;
  }
}

:deep(.header-row) {
  background: red;
}

:deep(.th.el-table__cell) {
  background: #faf7fa;
}
</style>
