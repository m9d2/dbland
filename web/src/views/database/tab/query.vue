<template>
  <div class="query-main">
    <div class="query-tool">
      <el-select v-model="configValueRef" size="small" @change="changeConfig" default-first-option>
        <el-option style="font-size: var(--font-size);" v-for="item in configs" :key="item.id" :label="item.name"
          :value="item.id" />
      </el-select>
      <el-select v-model="dbValueRef" size="small">
        <el-option style="font-size: var(--font-size);" v-for="item in databases" :key="item" :label="item"
          :value="item" @change="changeConfig" default-first-option />
      </el-select>
      <el-button @click="formatSql" size="small">{{ $t('database.button.format') }}</el-button>
      <el-button @click="querySql" size="small">{{ $t('database.button.run') }}</el-button>
    </div>
    <div class="query-console">
      <Console class="code" ref="consoleRef" :sql="sqlStr" />
    </div>
    <div class="table-tool" v-if="showQueryStatus">
      <div style="margin: 0 8px;" class="clearfix">
        <div class="fl">
          <el-tooltip effect="dark" :content="$t('common.add')">
            <el-link :underline="false" @click="insertRow">
              <el-icon>
                <Plus />
              </el-icon>
            </el-link>
          </el-tooltip>
        </div>

        <div class="fl">
          <el-tooltip effect="dark" :content="$t('common.delete')">
            <el-link :underline="false" @click="deleteRow">
              <el-icon style="font-weight: 700;">
                <Minus />
              </el-icon>
            </el-link>
          </el-tooltip>
        </div>


        <div class="fl">
          <el-tooltip effect="dark" :content="$t('common.modify')">
            <el-link :underline="false" @click="modifyRow">
              <el-icon style="font-weight: 700;">
                <Edit />
              </el-icon></el-link>
          </el-tooltip>
        </div>

        <div class="fl">
          <el-tooltip effect="dark" :content="$t('common.refresh')">
            <el-link :underline="false" @click="refresh">
              <el-icon>
                <Refresh />
              </el-icon>
            </el-link>
          </el-tooltip>
        </div>

        <div class="fr">
          <el-dropdown>
            <el-button size="small" style="margin-top: 3px">
              Import<el-icon class="el-icon--right"><arrow-down /></el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>csv</el-dropdown-item>
                <el-dropdown-item>insert sql</el-dropdown-item>
                <el-dropdown-item>excel</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>

      </div>
    </div>
    <Table class="query-table" :columns="tableColumns" :data="tableData" :loading="loading"
      @row-current-change="handlerCurrentChange" @row-contextmenu="tableDbClick">
    </Table>

    <Form v-model="showModify" :row="currentRow" :actionType="actionType" :columns="tableColumns" @cancel="handlerCancel" @confirm="handlerConfirm"></Form>

    <div class="query-status">
      <div class="status-left">
        <span v-show="showQueryStatus" style="margin-left: 8px">{{ $t('database.label.total') }}: {{ total }}</span>
      </div>
      <div class="status-center">
        <label>{{ sqlStr }}</label>
      </div>
      <div class="status-right">
        <span v-show="showQueryStatus">{{ $t('database.label.elapsed_time') }}: {{ elapsedTime }}s</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import Console from "@/components/console/index.vue";
import Table from "@/components/table/index.vue";
import Form from "@/components/form/index.vue";
import { format } from "sql-formatter";
import { getConfigs } from "@/api/config";
import { getDatabases, query, execute } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import type { QueryReq } from "@/api/connector/type";
import type { AxiosPromise } from "axios";
import { ElNotification } from "element-plus";
import ContextMenu from "@imengyu/vue3-context-menu";
import { ArrowDown } from '@element-plus/icons-vue'
import {
  Plus,
  Edit,
  Minus,
  Refresh,
} from '@element-plus/icons-vue'
import { ActionTypeEnum } from '@/common/enums/'
import { CreateDatabase } from '@/service'
import type {Database} from "@/service/database";

const configs = ref();
const configValueRef = ref();
const databases = ref();
let currentConfig
const dbValueRef = ref<string>();
const consoleRef = ref();
const tableData = ref();
const tableColumns = ref();
const loading = ref(false);
let sqlStr: string;
const elapsedTime = ref<number>();
const showQueryStatus = ref(false);
const total = ref();
let currentRow: any
let actionType: any
const showModify = ref(false)

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

function insertRow() {
  showModify.value = true
  actionType = ActionTypeEnum.INSERT
}

async function deleteRow() {

  try {
    const database:Database = CreateDatabase(currentConfig, dbValueRef, sqlStr)
    if (!currentRow) {
      ElNotification({
        message: 'please choose row',
        type: 'error'
      })
    }
    const sql = database.createDeleteSql(tableColumns.value, currentRow)
    const {data} = await execute({sql: sql, cid: configValueRef.value})
    ElNotification({
      message: 'Affected rows: ' + data,
      type: "success",
    });
    await querySql()
  } catch (error) {
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
  currentRow = row
}

function handlerCancel() {
  showModify.value = false
}

async function handlerConfirm(formData: any) {
  let sql
  const database:Database = CreateDatabase(currentConfig, dbValueRef, sqlStr)
  if (actionType == ActionTypeEnum.INSERT) {
    sql = database.createInsertSql(tableColumns.value, formData)
  }
  if (actionType == ActionTypeEnum.MODIFY) {
    if (!currentRow) {
      ElNotification({
        message: 'No rows selected',
        type: 'error'
      })
    }
    sql = database.createUpdateSql(tableColumns.value, formData, currentRow)
  }
  try {
    const { data } = await execute({cid: configValueRef.value, sql: sql})
    await querySql()
    ElNotification({
      message: 'Affected rows: ' + data,
      type: "success",
    });
  } catch (error) {
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
    };
    const { data } = await query(params);
    tableColumns.value = data.columns;
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
    loading.value = false;
  }
}

function tableDbClick(row: any, column: any, event: any) {
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
            const database:Database = CreateDatabase(currentConfig, dbValueRef, sqlStr)
            const sql = database.createDeleteSql(tableColumns.value, currentRow)
            const {data} = await execute({sql: sql, cid: configValueRef.value})
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
    height: 260px;
  }
}

.table-tool {
  height: 30px;
  min-height: 30px;
  border: 1px solid var(--db-c-border);
  border-bottom: none;
  line-height: 30px;

  .el-link {
    margin-right: 8px;
    font-weight: 700;
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
  //color: var(--db-c-text-1);
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
</style>
