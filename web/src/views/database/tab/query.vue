<template>
  <div class="query-main">
    <div class="query-tool">
      <el-select v-model="configValueRef" size="small" @change="changeConfig" default-first-option>
        <el-option v-for="item in configs" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
      <el-select v-model="dbValueRef" size="small">
        <el-option v-for="item in databases" :key="item" :label="item" :value="item" @change="changeConfig"
          default-first-option />
      </el-select>
      <el-button @click="formatSql" size="small" :icon="MagicStick">{{ $t('database.button.format') }}</el-button>
      <el-button @click="querySql" size="small" :icon="CaretRight">{{ $t('database.button.run') }}</el-button>
    </div>
    <div class="query-console" :style="{ height: contentHeight + 'px' }">
      <Console class="code" ref="consoleRef" :sql="sqlStr" />
    </div>

    <div class="query-content">
      <div class="resize-handle" @mousedown="startResize"></div>
      <div class="table-tool" v-if="showQueryStatus">
        <div class="clearfix">
          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.add')">
              <el-link :underline="false" @click="insertRow" :icon="Plus" />
            </el-tooltip>
          </div>

          <div class="menu fl">
            <el-tooltip effect="dark" :content="$t('common.delete')">
              <el-link :underline="false" @click="deleteRow" :disabled="!currentRow" :icon="Minus" />
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

          <div class="menu fr">
            <el-dropdown>
              <el-button size="small">
                Import<el-icon class="el-icon--right"><arrow-down /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>CSV</el-dropdown-item>
                  <el-dropdown-item>Insert SQL</el-dropdown-item>
                  <el-dropdown-item>Excel</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>

        </div>
      </div>
      <Table class="query-table" :columns="tableColumns" :data="tableData" :loading="loading" @sort-change="sortChange"
        @row-current-change="handlerCurrentChange" @row-contextmenu="handlerContextmenu">
      </Table>

      <Form v-model="showModify" :row="currentRow" :actionType="actionType" :columns="tableColumns"
        @cancel="handlerCancel" @confirm="handlerConfirm"></Form>

      <div class="query-status unselectable">
        <div class="status-left">
          <span v-show="showQueryStatus">{{ $t('database.label.elapsed_time') }}: {{ elapsedTime }}s</span>
        </div>
        <div class="status-center">
          
        </div>
        <div class="status-right" v-show="showQueryStatus">
          <span v-show="showQueryStatus" style="margin-left: 8px">{{ $t('database.label.total') }}: {{ total }}</span>
          <span v-show="showQueryStatus" style="margin-left: 8px">{{ $t('database.label.total_page') }}: {{ totalPage }}</span>
          <el-link style="margin-left: 8px;" :icon="ArrowLeftBold" :underline="false" @click="arrowLeft"/>
          <el-input v-model="page" size="small" style="width: 36px; margin-left: 8px;" @change="changePage"></el-input>
          <el-link style="margin-left: 8px;" :icon="ArrowRightBold" :underline="false" @click="arrowRight"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import Console from "@/components/console/index.vue";
import Table from "@/components/table/index.vue";
import Form from "@/components/form/index.vue";
import { format } from "sql-formatter";
import { getConfigs } from "@/api/config";
import { getDatabases, query, execute } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import type { QueryReq, Sort } from "@/api/connector/type";
import type { AxiosPromise } from "axios";
import { ElNotification } from "element-plus";
import ContextMenu from "@imengyu/vue3-context-menu";
import { ArrowDown, ArrowLeftBold, ArrowRightBold } from '@element-plus/icons-vue'
import {
  Plus,
  Edit,
  Minus,
  Refresh,
  CaretRight,
  MagicStick,
} from '@element-plus/icons-vue'
import { ActionTypeEnum } from '@/common/enums/'
import { CreateDatabase } from '@/service'
import type { Database } from "@/service/database";

const configs = ref();
const configValueRef = ref();
const databases = ref();
let currentConfig: Record<string, any>
const dbValueRef = ref<string>("");
const consoleRef = ref();
const tableData = ref();
const tableColumns = ref();
const loading = ref(false);
let sqlStr: string;
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
const props = defineProps({
  config: {
    type: Object,
  },
  database: {
    type: Object,
  },
  sql: {
    type: String,
  },
});
const resizeData = reactive({
  isResizing: false,
  startY: 0,
  startHeight: 0,
});
let sort: Sort

function insertRow() {
  showModify.value = true
  actionType = ActionTypeEnum.INSERT
}

async function deleteRow() {
  try {
    const database: Database = CreateDatabase(currentConfig, dbValueRef.value, sqlStr)
    if (!currentRow.value) {
      ElNotification({
        message: 'please choose row',
        type: 'error'
      })
    }
    const sql = database.createDeleteSql(tableColumns.value, currentRow.value)
    const { data } = await execute({ sql: sql, cid: configValueRef.value })
    ElNotification({
      message: 'Affected rows: ' + data,
      type: "success",
    });
    await querySql()
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  }
}

function modifyRow() {
  showModify.value = true
  actionType = ActionTypeEnum.MODIFY
}

async function refresh() {
  await querySql()
}

function handlerCurrentChange(row: any) {
  currentRow.value = row
}

function handlerCancel() {
  showModify.value = false
}

async function handlerConfirm(formData: any) {
  let sql
  const database: Database = CreateDatabase(currentConfig, dbValueRef.value, sqlStr)
  if (actionType == ActionTypeEnum.INSERT) {
    sql = database.createInsertSql(tableColumns.value, formData)
  }
  if (actionType == ActionTypeEnum.MODIFY) {
    if (!currentRow.value) {
      ElNotification({
        message: 'No rows selected',
        type: 'error'
      })
    }
    sql = database.createUpdateSql(tableColumns.value, formData, currentRow.value)
  }
  try {
    const { data } = await execute({ cid: configValueRef.value, sql: sql })
    await querySql()
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
  sqlStr = "";
  if (props.config) {
    configValueRef.value = props.config.id;
    currentConfig = props.config
  }
  if (props.database) {
    dbValueRef.value = props.database.name;
  }
  if (props.sql) {
    sqlStr = props.sql;
  }
  const height = localStorage.getItem('height')
  if (height) {
    contentHeight.value = parseInt(height)
  }
  loadConfigs();
});

async function changeConfig(id: any) {
  const params = {
    cid: id,
  };
  const { data } = await getDatabases(params);
  databases.value = data
  dbValueRef.value = data[0]
}

function loadConfigs() {
  try {
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    response.then((res: any) => {
      if (res.data) {
        if (!configValueRef.value) {
          configValueRef.value = res.data[0].id;
          currentConfig = res.data[0]
        }
        changeConfig(configValueRef.value);
        configs.value = res.data;
      }
    });
  } catch (error) {
    //
  }
}

function formatSql() {
  const cfg = {
    language: "sql",
    indent: "    ",
    uppercase: true,
    linesBetweenQueries: 2,
  };
  const sql = format(consoleRef.value.getValue(), cfg);
  consoleRef.value.setValue(sql);
}

async function querySql() {
  loading.value = true;
  
  sqlStr = consoleRef.value.getValue();
  try {
    const params: QueryReq = {
      cid: configValueRef.value,
      db: dbValueRef.value,
      sql: sqlStr,
      page: Number(page.value),
      size: Number(pageSize.value),
      sort: sort,
    };
    const { data } = await query(params);
    tableColumns.value = data.columns;
    tableData.value = data.rows;
    total.value = data.total;
    totalPage.value = data.total_page;
    elapsedTime.value = data.elapsed_time;
    showQueryStatus.value = true;
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  } finally {
    loading.value = false;
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
          try {
            const database: Database = CreateDatabase(currentConfig, dbValueRef.value, sqlStr)
            const sql = database.createDeleteSql(tableColumns.value, currentRow.value)
            const { data } = await execute({ sql: sql, cid: configValueRef.value })
            await querySql()
            ElNotification({
              message: 'Affected rows: ' + data,
              type: "success",
            });
          } catch (error: any) {
            ElNotification({
              title: "error",
              message: error.message,
              type: "error",
            });
          }
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

function startResize(event: MouseEvent) {
  resizeData.isResizing = true;
  resizeData.startY = event.clientY;
  resizeData.startHeight = contentHeight.value;

  document.addEventListener('mousemove', resize);
  document.addEventListener('mouseup', stopResize);
};

function resize(event: MouseEvent) {
  if (resizeData.isResizing) {
    const deltaY = event.clientY - resizeData.startY;
    contentHeight.value = resizeData.startHeight + deltaY;
    localStorage.setItem('height', String(contentHeight.value))
  }
};

function stopResize() {
  resizeData.isResizing = false;
  document.removeEventListener('mousemove', resize);
  document.removeEventListener('mouseup', stopResize);
}

function arrowLeft() {
  if (page.value > 1) {
    page.value = Number(page.value) - 1
    querySql()
  }
}

function arrowRight() {
  page.value = Number(page.value) + 1
  querySql()
}

function sortChange({ column, prop, order }: any) {
  if (!prop || !order) {
    return;
  }
  if (order === 'descending') {
    order = 'desc';
  } else {
    order = 'asc';
  }
  sort = {prop, order}
  querySql();
}

function changePage() {
  querySql()
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
</style>
