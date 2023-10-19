<template>
  <div class="content-box">

    <div class="header"
      style="padding: 0 8px;background-color: var(--db-c-bg-nav); border-bottom: 1px solid var(--db-c-border); height: 30px; line-height: 30px;">
      <el-breadcrumb :separator-icon="ArrowRight"
        style="overflow: hidden;height: 30px; line-height: 30px; font-size: var(--font-size); color: var(--db-c-text)">
        <el-breadcrumb-item v-if="firstData" :to="{ path: '/' }" @click="loadConfigs">{{ $t('common.homepage')
        }}</el-breadcrumb-item>
        <el-breadcrumb-item v-show="firstData" @click="back" :to="{ path: '/' }">{{ firstData }}</el-breadcrumb-item>
        <el-breadcrumb-item>

          <el-text truncated>
            {{ secondData }}
          </el-text>
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="content-main">
      <div class="content-left" :style="{ width: contentWidth + 'px' }">
        <div class="search">
          <el-input v-model="filterText" style="height: 32px" :placeholder="placeholder" size="small" />
          <el-dropdown>
            <el-link style="font-size: 16px; margin-left: 8px;" :icon="Menu" :underline="false">
            </el-link>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="newTab('')">{{ $t('database.button.new_query') }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>

        <List class="menu-list" style="margin: 8px 8px;" :list="listData" @node-click="clickNode"
          @node-mouse-enter="handleMouseEnter" @node-mouse-leave="handleMouseLeave" @node-db-click="handleNodeDblClick">

          <template #begin="{ node }">
            <el-icon v-show="node.level == TreeLevelEnum.CONFIG" class="iconfont" style="margin: 5px">
              <span v-show="node.type == DbTypeEnum.MySQL" class="iconfont">&#xec6d;</span>
              <span v-show="node.type == DbTypeEnum.SQLite" class="iconfont">&#xe65a;</span>
              <span v-show="node.type == DbTypeEnum.ORACLE" class="iconfont">&#xec48;</span>
              <span v-show="node.type == DbTypeEnum.PostgreSQL" class="iconfont">&#xe8b7;</span>
              <span v-show="node.type == DbTypeEnum.MariaDB" class="iconfont">&#xec6d;</span>
            </el-icon>
            <el-icon v-show="node.level == TreeLevelEnum.DATABASE" class="iconfont" style="margin: 5px">
              <Coin />
            </el-icon>
            <el-icon v-show="node.level == TreeLevelEnum.TABLE" class="iconfont" style="margin: 5px">
              <Document />
            </el-icon>
          </template>

          <template #default="{ index, node }">
            <el-dropdown style="vertical-align: center;" size="small" trigger="click">
              <el-link v-show="index === activeIndex && node.level == Level.TABLE" :underline="false"
                :icon="MoreFilled" />
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item :icon="Search" @click="newTab(node.name)">{{ $t('database.button.query')
                  }}</el-dropdown-item>
                  <el-dropdown-item :icon="Edit">{{ $t('common.edit') }}</el-dropdown-item>
                  <el-dropdown-item :icon="Delete">{{ $t('common.delete') }}</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </template>
        </List>
        <div class="resize-handle" @mousedown="startResize"></div>
      </div>
      <div class="content-right">
        <div class="shortcuts" v-if="shortcutsVisible">
          <span class="logo-text">DBLAND</span>
        </div>
        <Tab :isResultVisible="isResultVisible" ref="childRef" @messageUpdated="handlerMessageUpdate" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, reactive } from "vue";
import { ElNotification } from "element-plus";
import { ArrowRight, Coin, Document, Search, Menu } from '@element-plus/icons-vue'
import { getConfigs } from "@/api/config";
import { getDatabases, getTables } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import Tab from "@/views/database/tab/index.vue";
import i18n from '@/plugins/i18n'
import type { AxiosPromise } from "axios";
import List from '@/components/layout/list/index.vue'
import { TreeLevelEnum } from "@/common/enums";
import { DbTypeEnum } from '@/common/enums';
import { Delete, Edit, MoreFilled } from '@element-plus/icons-vue'

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

const placeholder = i18n.global.t('common.search')
const firstData = ref()
const secondData = ref()
let currentCid: number
let currentConfig: Object
let currentDatabase: Object
const activeIndex = ref()
const resizing = ref(false);
const listData = ref()
const contentWidth = ref(220)
const resizeData = reactive({
  isResizing: false,
  startX: 0,
  startWidth: 0,
});

onMounted(() => {
  let width = localStorage.getItem('width')
  if (width) {
    contentWidth.value = parseInt(width)
  }
})

const handleNodeDblClick = (index: number, node: any) => {
  if (node.level === 3) {
    newTab(node.name);
  }
};

function newTab(table: string) {
  isResultVisible.value = true;
  shortcutsVisible.value = false;
  childRef.value?.newTab(currentConfig, currentDatabase, table);
}

function back() {
  secondData.value = ''
  loadDatabases(currentCid)
}

const handlerMessageUpdate = (data: any) => {
  isResultVisible.value = data;
  shortcutsVisible.value = !data;
};

function handleMouseEnter(index: number) {
  activeIndex.value = index;
}

function handleMouseLeave(index: number) {
  activeIndex.value = null;
}

function clickNode(index: number, row: any) {
  const level = row.level
  if (level == TreeLevelEnum.CONFIG) {

    loadDatabases(row.id)
    firstData.value = row.name
    currentCid = row.id
    currentConfig = row
  }
  if (level == TreeLevelEnum.DATABASE) {
    loadTables(row.cid, row.name)
    secondData.value = row.name
    currentDatabase = row
  }
}

function startResize(event: MouseEvent) {
  resizeData.isResizing = true;
  resizeData.startX = event.clientX;
  resizeData.startWidth = contentWidth.value;
  
  document.addEventListener('mousemove', resize);
  document.addEventListener('mouseup', stopResize);
};

function resize(event: MouseEvent) {
  if (resizeData.isResizing) {
    const deltaX = event.clientX - resizeData.startX;
    contentWidth.value = resizeData.startWidth + deltaX;
    // console.log(contentWidth.value)
    localStorage.setItem('width', contentWidth.value)
  }
};

function stopResize() {
  resizing.value = false;
  document.removeEventListener('mousemove', resize);
  document.removeEventListener('mouseup', stopResize);
};

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
      message: error.message,
      type: 'error'
    })
  } finally {
    firstData.value = ''
    secondData.value = ''
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

const filterNode = (value: string, data: Tree) => {
  if (!value) return true;
  console.log(value)
  return data.label.includes(value);
};
</script>

<style lang="scss" scoped>
.content-box {
  display: flex;
  flex-direction: column;
  height: 100%;

  .header {
    color: var(--db-c-text-1);

    .el-link {
      font-size: var(--font-size);
      vertical-align: bottom;
    }
  }
}

.content-main {
  display: flex;
  flex-direction: row;
  flex-grow: 1;
  overflow: hidden;

  .content-left {
    height: 100%;
    color: var(--db-c-text-menu);
    background-color: var(--db-c-bg-nav);
    position: relative;
    display: flex;
    flex-direction: column;
    position: relative;

    .resize-handle {
    position: absolute;
    top: 0;
    right: 0;
    width: 1px;
    height: 100%;
    z-index: 999;
    width: 6px;
    border-right: 1px solid var(--db-c-border);
  }

    .search {
      border-radius: 5px;
      padding: 8px;
      display: flex;
      align-items: center;
    }
  }

    .content-right {
      flex-grow: 1;
      height: 100%;
      overflow: hidden;
      margin: 0 8px;
      background-color: var(--db-c-bg);
    }

  }

  .resize-handle:hover {
    background-color: var(--db-c-border);
    cursor: col-resize;
  }
 

  .menu-list {
    flex-grow: 1;
  }

  .new-query {
    height: 48px;
    padding: 8px;
    justify-content: center;
    box-sizing: border-box;
  }

.shortcuts {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;

  .logo-text {
    font-size: 26px;
    color: var(--db-c-text);
    opacity: 0.15;
  }
}


.el-dropdown {
  line-height: 30px;
  vertical-align: middle;

  .el-link {
    font-size: var(--font-size);
  }
}
</style>
