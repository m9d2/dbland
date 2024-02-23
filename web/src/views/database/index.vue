<template>
  <div class="content-box">
    <div class="content-main">
      <div class="content-left" :style="{ width: contentWidth + 'px' }">
        <div class="search">
          <el-input v-model="filterText" :placeholder="$t('common.search')" size="small" clearable>
            <template #prefix>
              <el-icon class="el-input__icon">
                <search />
              </el-icon>
            </template>
          </el-input>
          <el-link :icon="Refresh" @click="refresh" :underline="false"/>
          <el-link :icon="Plus" @click="newTab('')" :underline="false"/>
        </div>
        <el-tree ref="treeRef" :props="props" :load="loadNode" empty-text="" lazy @check-change="handleCheckChange"
          :indent="12" :highlight-current="true" :filter-node-method="filterNode" @node-click="clickNode">
          <template v-slot="{ node, data }">
            <div class="tree-node" @dblclick="doubleClickNode(data, node)" @mouseover="handleMouseEnter(node, data)"
              @mouseleave="handleMouseLeave(node, data)">
              <div class="node-content unselectable">
                <el-icon v-show="node.level == TreeLevelEnum.CONFIG">
                  <span v-for="item in icons" v-show="data.type == item.dbType" class="iconfont"
                    :class="item.icon"></span>
                </el-icon>
                <el-icon v-show="node.level == TreeLevelEnum.DATABASE" class="iconfont">
                  <Folder />
                </el-icon>
                <el-icon v-show="node.level == TreeLevelEnum.TABLE" class="iconfont">
                  <Document />
                </el-icon>
                <span>{{ data.name }}</span>
              </div>

              <div class="node-menu">
                <el-dropdown size="small" trigger="click">
                  <el-link v-show="data.isCurrent" :underline="false" :icon="MoreFilled" />
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item :icon="Search" @click="newTab(node.name)">{{
                        $t('database.button.query')
                      }}
                      </el-dropdown-item>
                      <el-dropdown-item :icon="Edit">{{ $t('common.edit') }}</el-dropdown-item>
                      <el-dropdown-item :icon="Delete">{{ $t('common.delete') }}</el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </div>
            </div>
          </template>
        </el-tree>
        <div class="resize-handle" @mousedown="startResize"></div>
      </div>
      <div class="content-right">
        <div class="info" v-if="shortcutsVisible">
          <span class="logo-text">DBLAND</span>
        </div>
        <Tab :isResultVisible="isResultVisible" ref="childRef" @messageUpdated="handlerMessageUpdate" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref, watch } from "vue";
import { ElNotification, ElTree } from "element-plus";
import { Coin, Document, Folder, Refresh, Edit, Delete, Search, MoreFilled, Connection, Plus } from '@element-plus/icons-vue';
import { getConfigs } from "@/api/config";
import { getDatabases, getTables } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import Tab from "@/views/database/tab/index.vue";
import type { AxiosPromise } from "axios";
import { DbTypeEnum, TreeLevelEnum } from "@/common/enums";
import type Node from 'element-plus/es/components/tree/src/model/node';

const filterText = ref("");
const isResultVisible = ref();
const shortcutsVisible = ref(true);
const childRef = ref();
const resizing = ref(false);
const contentWidth = ref(220)
const resizeData = reactive({
  isResizing: false,
  startX: 0,
  startWidth: 0,
});
let currentConfig: Object
let currentDatabase: Object
const props = {
  label: 'name',
  children: 'children',
  isLeaf: 'leaf',
}
const icons = [
  { dbType: DbTypeEnum.MySQL, icon: 'icon-mysql' },
  { dbType: DbTypeEnum.SQLite, icon: 'icon-sqlite' },
  { dbType: DbTypeEnum.ORACLE, icon: 'icon-oracle' },
  { dbType: DbTypeEnum.PostgreSQL, icon: 'icon-PostgreSQL' },
  { dbType: DbTypeEnum.MariaDB, icon: 'icon-mysql' },
]

interface Tree {
  name: string
}

const handleCheckChange = (
  data: Tree,
  checked: boolean,
  indeterminate: boolean
) => {
  console.log(data, checked, indeterminate)
}

const loadNode = async (node: Node, resolve: (data: Tree[]) => void) => {
  if (node.level === 0) {
    const nodes: any = await loadConfigs();
    return resolve(nodes);
  }
  if (node.level === 1) {
    const nodes: any = await loadDatabases(node.data.id)
    return resolve(nodes);
  }
  if (node.level === 2) {
    const nodes: any = await loadTables(node.data.cid, node.data.name)
    const formattedNodes = nodes.map((childNode: Node) => ({
      ...childNode,
      leaf: true,
    }));
    return resolve(formattedNodes);
  }
  return resolve([])
};


onMounted(() => {
  let width = localStorage.getItem('width')
  if (width) {
    contentWidth.value = parseInt(width)
  }
})

const newTab = (table: string) => {
  isResultVisible.value = true;
  shortcutsVisible.value = false;
  childRef.value?.newTab(currentConfig, currentDatabase, table);
}

const refresh = async () => {
  loadConfigs()
}

const handlerMessageUpdate = (data: any) => {
  isResultVisible.value = data;
  shortcutsVisible.value = !data;
};

const handleMouseEnter = (node: any, data: any) => {
  if (node.level == TreeLevelEnum.TABLE) {
    data.isCurrent = true
  }
}

const handleMouseLeave = (node: any, data: any) => {
  if (node.level == TreeLevelEnum.TABLE) {
    data.isCurrent = false
  }
}

const clickNode = (data: any, node: any) => {
  const level = node.level
  if (!level || level == TreeLevelEnum.CONFIG) {
    currentConfig = node.data
  }
  if (level == TreeLevelEnum.DATABASE) {
    currentDatabase = node.data
  }
}

const doubleClickNode = (data: any, node: any) => {
  if (node.level === 3) {
    newTab(node.data.name);
  }
};

const startResize = (event: MouseEvent) => {
  resizeData.isResizing = true;
  resizeData.startX = event.clientX;
  resizeData.startWidth = contentWidth.value;

  document.addEventListener('mousemove', resize);
  document.addEventListener('mouseup', stopResize);
}

const resize = (event: MouseEvent) => {
  if (resizeData.isResizing) {
    const deltaX = event.clientX - resizeData.startX;
    contentWidth.value = resizeData.startWidth + deltaX;
    localStorage.setItem('width', contentWidth.value.toString())
  }
}

const stopResize = () => {
  resizing.value = false;
  document.removeEventListener('mousemove', resize);
  document.removeEventListener('mouseup', stopResize);
}

const loadConfigs = async () => {
  try {
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    const res = await response;
    if (res.data) {
      return res.data;
    }
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: 'error',
    });
    return [];
  }
}

const loadDatabases = async (cid: number) => {
  const { data } = await getDatabases({ cid: cid });
  return data
    ? data.map((name) => ({
      name: name,
      cid: cid,
    }))
    : []
}

const loadTables = async (cid: number, db: string) => {
  const { data } = await getTables({
    cid: cid,
    db: db,
  });
  return data
    ? data.map((table) => ({
      name: table.name,
      cid: cid,
      children: [],
      isLeaf: true,
    }))
    : []
}

const treeRef = ref<InstanceType<typeof ElTree>>()

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const filterNode = (value: string, data: any) => {
  if (!value) return true;
  return data.name.includes(value);
};
</script>

<style lang="scss" scoped>
.content-box {
  display: flex;
  flex-direction: column;
  height: 100%;

  .el-link {
    font-size: 16px;
    vertical-align: bottom;
    margin-left: 8px;
  }
}

.content-main {
  display: flex;
  flex-direction: row;
  flex-grow: 1;
  overflow: hidden;

  .content-left {
    flex: 0 0 auto;
    height: 100%;
    color: var(--db-c-text-menu);
    background-color: var(--db-c-bg-nav);
    display: flex;
    flex-direction: column;
    position: relative;

    .resize-handle {
      position: absolute;
      top: 0;
      right: 0;
      height: 100%;
      z-index: 999;
      width: 6px;
      border-right: 1px solid var(--db-c-border);
    }

    .search {
      border-radius: 5px;
      padding: 12px;
      display: flex;
      align-items: center;
    }

    .el-icon {
      margin: 0 5px 0 0;
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

.info {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;

  .logo-text {
    font-size: 26px;
    color: var(--db-c-text);
    opacity: 0.15;
  }

  .version {
    display: block;
    font-size: 16px;
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

.node-content {
  display: flex;
  justify-content: center;
  align-items: center;
}

.el-tree {
  padding: 0 5px;
  color: var(--db-c-text-menu);
  background: var(--db-c-bg-nav);
  overflow: auto;
}

.tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8px;
}
</style>
