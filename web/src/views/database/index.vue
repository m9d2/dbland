<template>
  <div
    style="padding: 0 8px;background-color: var(--color-background-deep); border-bottom: 1px solid var(--color-border); height: 30px; line-height: 30px;">
    <el-link :underline="false" v-show="firstData" @click="loadConfigs">首页</el-link>
    <el-icon v-show="firstData" >
      <ArrowRight />
    </el-icon>
    <el-link :underline="false" @click="loadDatabases">{{ firstData }}</el-link>
    <el-icon v-show="seconedData">
      <ArrowRight />
    </el-icon>
    <el-link :underline="false">{{ seconedData }}</el-link>
  </div>

  <div class="main-box">
    <div class="tree-box">
      <div class="search" style="display: flex; border-radius: 5px; padding: 8px;">
        <el-input v-model="filterText" style="height: 32px" :placeholder="placeholder" size="small" />
      </div>

      <List class="tree" :list="listData" style="margin: 0 20px;" @node-click="clickNode"
        :row-name-style="{fontWeight: 300, height: '26px', lineHeight: '26px', marginTop: '0'}"
        @node-mouse-enter="handleMouseEnter" @node-mouse-leave="handleMouseLeave">
      </List>

      <div class="new-query">
        <el-button type="primary" style="width: 100%" @click="newTab(null)">{{ $t('database.button.new_query')
        }}</el-button>
      </div>
    </div>
    <div class="main">
      <div class="shortcuts" v-if="shortcutsVisible">
        <span class="logo-text">DBLAND</span>
      </div>
      <Tab :isResultVisible="isResultVisible" ref="childRef" @messageUpdated="handlerMessageUpdate" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import { ElNotification } from "element-plus";
import { ElTree } from "element-plus";
import { ArrowRight } from '@element-plus/icons-vue'
import { getConfigs } from "@/api/config";
import { getDatabases, getTables } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import Tab from "@/views/database/tab/index.vue";
import i18n from '@/plugins/i18n'
import type { AxiosPromise } from "axios";
import List from '@/components/layout/list/index.vue'
import { TreeLevelEnum } from "@/common/enums";

const treeRef = ref<InstanceType<typeof ElTree>>();
// data
const defaultProps = {
  children: "children",
  label: "name",
  isLeaf: "leaf",
};

enum Level {
  ROOT = 0,
  CONFIG = 1,
  DATABASE = 2,
  TABLE = 3,
}

interface Tree {
  [key: string]: any;
}

const filterText = ref("");
const isResultVisible = ref();
const shortcutsVisible = ref(true);
const childRef = ref();
const currentConfig = ref();
const currentDatabase = ref();
const placeholder = i18n.global.t('common.search')
const listData = ref()
const firstData = ref()
const seconedData = ref()

const handleNodeDblClick = (node: any) => {
  if (node.level === 3) {
    newTab(node.data.name);
  }
};

function newTab(table: string) {
  isResultVisible.value = true;
  shortcutsVisible.value = false;
  childRef.value?.newTab(currentConfig.value, currentDatabase.value, table);
}

function nodeClick(node: any) {
  if (node.level == Level.CONFIG) {
    currentConfig.value = node;
  }
  if (node.level == Level.DATABASE) {
    currentDatabase.value = node;
  }
}
const handlerMessageUpdate = (data: any) => {
  isResultVisible.value = data;
  shortcutsVisible.value = !data;
};

function handleMouseEnter(index: number) {
}

function handleMouseLeave(index: number) {
}

function clickNode(index: number, row: any) {
  const level = row.level
  if (level == TreeLevelEnum.CONFIG) {
    loadDatabases(row.id)
    firstData.value = row.name
  }
  if (level == TreeLevelEnum.DATABASE) {
    loadTables(row.cid, row.name)
    seconedData.value = row.name
  }
}

onMounted(() => {
  loadConfigs();
})

function loadConfigs() {
  try {
    
    let nodes: any = []
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    response.then((res: any) => {
      if (res.data) {
        res.data.forEach((data: any) => {
          data.active = false
          data.level = TreeLevelEnum.CONFIG
          nodes.push(data)
        })
        listData.value = nodes
      }
    });
  } catch (error: any) {
    ElNotification({
      title: 'error',
      message: error.message,
      type: 'error'
    })
  } finally {
    firstData.value = ''
    seconedData.value = ''
  }
}

async function loadDatabases(cid: number) {
  let nodes: any = []
  const { data } = await getDatabases({ cid: cid });
  for (let i = 0; i < data.length; i++) {
    let child = {
      name: data[i],
      level: Level.DATABASE,
      cid: cid,
    };
    nodes.push(child)
  }
  listData.value = nodes
}

async function loadTables(cid: number, db: string) {
  const { data } = await getTables({
    cid: cid,
    db: db,
  });
  const nodes = data
    ? data.map((table) => ({
      name: table.name,
      level: Level.TABLE,
      cid: cid,
      leaf: true,
    }))
    : []
  listData.value = nodes
}

watch(filterText, (val) => {
  treeRef.value!.filter(val);
});

const filterNode = (value: string, data: Tree) => {
  if (!value) return true;
  return data.label.includes(value);
};
</script>

<style lang="scss" scoped>
.main-box {
  height: 100%;
  display: flex;
  flex-direction: row;
}

.tree-box {
  resize: both;
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--color-border);
  color: var(--color-text) !important;
  background-color: var(--color-background-deep);
  position: relative;
  display: flex;
  flex-direction: column;


  .tree {
    height: 100%;
    overflow: auto;
    flex-grow: 1;
  }

  .new-query {
    width: 200px;
    height: 48px;
    margin: 20px auto 0;
    justify-content: center;
    box-sizing: border-box;
  }
}

.el-tree {
  padding: 0 8px 8px;
  background-color: var(--color-background-deep);
  font-size: var(--font-size) !important;
  --el-tree-node-hover-bg-color: var(--tree-node-hover-bg-color);
}

.shortcuts {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;

  .logo-text {
    font-size: 26px;
    color: var(--color-text);
    opacity: 0.15;
    font-weight: 900;
  }
}

.main {
  flex-grow: 1;
  height: 100%;
  overflow: hidden;
  margin: 0 8px;
}
</style>
