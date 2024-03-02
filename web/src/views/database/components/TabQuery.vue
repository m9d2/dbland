<template>
  <div class="query-main">
    <div class="query-tool">
      <el-select v-model="config" size="small" @change="changeConfig" default-first-option>
        <el-option v-for="item in configs" :key="item.id" :label="item.name" :value="item.id"/>
      </el-select>
      <el-select v-model="database" size="small" @change="changeDatabase">
        <el-option v-for="item in databases" :key="item" :label="item" :value="item"
                   default-first-option/>
      </el-select>

<!--      <el-link :underline="false" @click="formatSql">-->
<!--        <span class="iconfont icon-yijianmeihua"></span>-->
<!--      </el-link>-->

      <el-link :underline="false" @click="querySql">
        <span class="iconfont icon-yunhang" style="color: var(--db-c-primary);"></span>
      </el-link>

<!--      <el-link :underline="false" @click="formatSql">-->
<!--        <span class="iconfont icon-24gl-stop" style="color: red;"></span>-->
<!--      </el-link>-->
    </div>
    <div id="query-console" class="query-console" :style="{ height: contentHeight + 'px' }">
      <Console ref="console" class="code" @change="handleCodeChange" :tables="tables" :schemas="schemas"/>
    </div>

    <Resize :resize-y="true" ref-id="query-console"></Resize>
    <div class="query-content">
      <Table class="query-table" :columns="tableColumns" :data="tableData" :loading="loading"
             header-row-class-name="custom-header-cell" header-cell-class-name="custom-header-row"
             :column-width="columnWidth" @row-current-change="handlerCurrentChange" :sortable="false">
      </Table>

      <div class="query-status unselectable">
        <div class="status-left">
          <span v-show="showQueryStatus">{{ $t('database.label.elapsed_time') }}: {{ elapsedTime }}s</span>
        </div>
        <div class="status-center">

        </div>
        <div class="status-right" v-show="showQueryStatus">
          <span v-show="showQueryStatus" style="margin-left: 8px">{{ $t('database.label.total') }}: {{ total }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, onBeforeMount} from "vue";
import Console from "@/components/console/index.vue";
import Table from "@/components/table/index.vue";
import {getConfigs} from "@/api/config";
import {getDatabases, query, getTables} from "@/api/connector";
import type {DBConfig} from "@/api/config/type";
import type {QueryReq} from "@/api/connector/type";
import type {AxiosPromise} from "axios";
import {ElNotification} from "element-plus";
import Resize from '@/components/resize/index.vue';

const configs = ref();
const config = ref();
const databases = ref();
const database = ref()
const tableData = ref();
const tableColumns = ref();
const loading = ref(false);
let sqlStr: string;
const elapsedTime = ref<number>();
const showQueryStatus = ref(false);
const total = ref();
const currentRow = ref();
let actionType: any
const contentHeight = ref(300);
const totalPage = ref()
const columnWidth = ref(160)
const tables = ref([])
const schemas = ref([])

const handlerCurrentChange = (row: any) => {
  currentRow.value = row
}

onMounted(() => {
  sqlStr = "";
  loadConfigs();
});

onBeforeMount(() => {
  const height = localStorage.getItem('height')
  if (height) {
    contentHeight.value = parseInt(height)
  }
})

const changeConfig = async (id: any) => {
  const params = {
    cid: id,
  };
  const {data} = await getDatabases(params);
  databases.value = data
  schemas.value = data
  database.value = data[0]
}

const changeDatabase = async (e: any) => {
  const req: QueryReq = {
    cid: config.value,
    database: e,
  }
  const response: AxiosPromise<Table[]> = getTables(req);
  response.then((res: any) => {
    if (res.data) {
      for (const item of res.data) {
        item.key = item.name
        tables.value.push(item.name)
      }
    }
  });
}

const handleCodeChange = (e) => {
  sqlStr = e
}

const loadConfigs = () => {
  try {
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    response.then((res: any) => {
      if (res.data) {
        configs.value = res.data;
        if (configs.value.length > 0) {
          config.value = configs.value[0].id;
          changeConfig(configs.value[0].id)
        }
      }
    });
  } catch (error) {
    //
  }
}

const formatSql = () => {
  const cfg = {
    language: "sql",
    indent: "    ",
    uppercase: true,
    linesBetweenQueries: 2,
  };
  const sql = format(sqlStr, cfg);
}

const querySql = async () => {
  loading.value = true;
  try {
    const params: QueryReq = {
      cid: config.value,
      database: database.value,
      sql: sqlStr
    };
    const {data} = await query(params);
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
    totalPage.value = data.total_page;
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
</script>

<style lang="scss" scoped>
.query-main {
  height: 100%;
  display: flex;
  flex-direction: column;

  .query-tool {
    border-bottom: 1px solid var(--db-c-border);
    padding: 8px;

    .el-link {
      font-weight: bold;
    }
  }

  .query-tool > * {
    margin-right: 8px;
  }


  .el-button + .el-button {
    margin: 0;
  }

  .query-console {
    //margin: 8px 0;
    width: 100%;

    .code {
      height: 100%;
    }
  }

  .query-content {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;

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
      background-color: var(--db-c-bg-nav);
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
  }
}

:deep(.custom-header-cell) {
  .cell {
    overflow: hidden;
    white-space: nowrap;
  }
}

.custom-header-row {
  //background: red;
}

:deep(.el-table tr) {
  //background-color: red;
}
</style>
  