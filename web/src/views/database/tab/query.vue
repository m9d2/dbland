<template>
  <div class="query-main">
    <div class="query-tool">
      <el-select v-model="configValueRef" size="small" @change="changeConfig" default-first-option>
        <el-option
            v-for="item in configs"
            :key="item.name"
            :label="item.name"
            :value="item.id"
        />
      </el-select>
      <el-select v-model="dbValueRef" size="small">
        <el-option
            v-for="item in databases"
            :key="item"
            :label="item"
            :value="item"
            @change="changeConfig"
            default-first-option
        />
      </el-select>
      <el-button @click="formatSql" size="small">Format</el-button>
      <el-button @click="querySql" size="small">Run</el-button>
    </div>
    <div class="query-console">
      <Console class="code" ref="consoleRef" :sql="sqlStr"/>
    </div>
    <div class="table-tool"></div>
    <Table class="query-table" :columns="tableColumns" :data="tableData" :loading="loading"></Table>
    <div class="query-status">
      <div class="status-left">
        <span v-show="showQueryStatus" style="margin-left: 8px">Total: {{total}}</span>
      </div>
      <div class="status-center">
        <label>{{ sqlStr }}</label>
      </div>
      <div class="status-right">
        <span v-show="showQueryStatus">Elapsed Time: {{elapsedTime}}s</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref } from 'vue'
import Console from '@/components/database/console/index.vue'
import Table from "@/components/database/table/index.vue"
import {format} from "sql-formatter"
import {getConfigs} from "@/api/config";
import {getDatabases, query} from "@/api/connector";
import type {DBConfig} from "@/api/config/type"
import type {QueryReq} from "@/api/connector/type";
import type { AxiosPromise } from "axios";
import {ElNotification} from "element-plus";

const configs = ref()
const configValueRef = ref()
const databases = ref()
const dbValueRef = ref<string>()
const consoleRef = ref()
const tableData = ref()
const tableColumns = ref()
const loading = ref(false)
let sqlStr:string
const elapsedTime = ref<string>()
const showQueryStatus = ref(false)
const total = ref()

const props = defineProps({
  config: {
    type: Object,
  },
  database: {
    type: Object,
  },
  sql: {
    type: String,
  }
})

onMounted(() => {
  sqlStr = ''
  if (props.config) {
    configValueRef.value = props.config.id
  }
  if (props.database) {
    dbValueRef.value = props.database.name
  }
  if (props.sql) {
    sqlStr = props.sql
  }
  loadConfigs();
})

async function changeConfig(id:any) {
  databases.value = []
  const data = {
    cid: id
  }
  const response: AxiosPromise<string[]> = getDatabases(data);
  response.then((res:any) => {
    databases.value = res.data
  });
}

function loadConfigs() {
  try {
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    response.then((res:any) => {
      if (res.data) {
        if (!configValueRef.value) {
          configValueRef.value = res.data[0].id
        }
        changeConfig(configValueRef.value)
        configs.value = res.data;
      }
    });
  } catch (error) {
    //
  }
}

function formatSql() {
  const cfg = {
    language: 'sql',
    indent: '    ',
    uppercase: true,
    linesBetweenQueries: 2,
  }
}

async function querySql() {
  loading.value = true
  try {
    const params: QueryReq = {
      cid: 2,
      db: dbValueRef.value,
      sql: sqlStr,
    }
    const { data } = await query(params)
    tableColumns.value = data.columns
    tableData.value = data.rows
    total.value = data.total
    elapsedTime.value = data.elapsed_time
    showQueryStatus.value = true
  } catch (error: any) {
    ElNotification({
      title: 'error',
      message: error.message,
      type: 'error'
    })
  } finally {
    loading.value = false
  }
}

</script>

<style lang="scss" scoped>
.query-main {
  height: 100%;
  display: flex;
  flex-direction: column;
}
.el-button + .el-button {
  margin: 0;
}
.query-tool > * {
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
  border: 1px solid var(--color-border);
  border-bottom: none;
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
  border: 1px solid var(--color-border);
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
      color: var(--color-text);
    }
  }
}
</style>
