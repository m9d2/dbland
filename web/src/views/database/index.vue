<template>
  <div class="main-box">
    <div class="tree-box">
      <div class="search" style="display: flex; border-radius: 5px; padding: 8px">
        <el-input v-model="filterText" style="height: 32px" placeholder="查询"/>
      </div>
      <div class="tree">
        <el-tree
            ref="treeRef"
            :data="data"
            :props="defaultProps"
            lazy
            :load="loadData"
            empty-text=""
            :filter-node-method="filterNode"
            :node-key="data.id"
            @node-click="nodeClick"
        >
          <template #default="{ node }">
          <span @dblclick="handleNodeDblClick(node)">
            <span
                v-if="node.level === Level.DATABASE"
                class="iconfont"
                style="margin-right: 5px"
            >
              &#xe685;
            </span>
            <span
                v-if="node.level === Level.TABLE"
                class="iconfont .icon-biaoge"
                style="margin-right: 5px"
            >
              &#xe615;
            </span>
            <span>{{ node.label }}</span>
          </span>
          </template>
        </el-tree>
      </div>
      <div class="new-query">
        <el-button type="primary" style="width: 100%" @click="newTab(null)">New Query</el-button>
      </div>
    </div>
    <div class="main">
      <div class="shortcuts" v-if="shortcutsVisible">
        <span class="logo-text">DBLAND</span>
      </div>
      <Tab
          :isResultVisible="isResultVisible"
          ref="childRef"
          @messageUpdated="handlerMessageUpdate"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref, reactive, provide, watch} from 'vue'
import {ElNotification} from 'element-plus'
import {ElTree} from 'element-plus'
import {getConfigs} from '@/api/config'
import {getDatabases, getTables} from '@/api/connector'
import type {DBConfig} from '@/api/config/type'
import Tab from '@/views/database/tab/index.vue'

const treeRef = ref<InstanceType<typeof ElTree>>()
// data
const defaultProps = {
  children: 'children',
  label: 'name',
  isLeaf: 'leaf'
}

enum Level {
  ROOT = 0,
  CONFIG = 1,
  DATABASE = 2,
  TABLE = 3
}

interface Tree {
  [key: string]: any
}

let data = reactive([])
const filterText = ref('')
const isResultVisible = ref()
const shortcutsVisible = ref(true)
const childRef = ref()
const currentConfig = ref()
const currentDatabase = ref()

const handleNodeDblClick = (node: any) => {
  if (node.level === 3) {
    newTab(node.data.name)
  }
}

function newTab(table: string) {
  isResultVisible.value = true
  shortcutsVisible.value = false
  childRef.value?.newTab(currentConfig.value, currentDatabase.value, table)
}

function nodeClick(node:any) {
  if (node.level == Level.CONFIG) {
    currentConfig.value = node
  }
  if (node.level == Level.DATABASE) {
    currentDatabase.value = node
  }

}
const handlerMessageUpdate = (data: any) => {
  isResultVisible.value = data
  shortcutsVisible.value = !data
}

const loadData = (node: any, resolve: any) => {
  let childNodes: any = []
  setTimeout(async () => {
    switch (node.level) {
      case Level.ROOT:
        try {
          const {data} = await getConfigs()
          data.forEach((config: DBConfig) => {
            config.level = Level.CONFIG
            config.leaf = false
            childNodes.push(config)
          })
        } catch (error: any) {
          ElNotification({
            title: 'error',
            message: error.message,
            type: 'error'
          })
        } finally {
          resolve(childNodes)
        }

        break
      case Level.CONFIG:
        try {
          const {data} = await getDatabases({cid: node.data.id})
          for (let i = 0; i < data.length; i++) {
            let child = {name: data[i], level: Level.DATABASE, cid: node.data.id}
            childNodes.push(child)
          }
        } catch (error: any) {
          ElNotification({
            title: 'error',
            message: error.message,
            type: 'error'
          })
        } finally {
          resolve(childNodes)
        }

        break
      case Level.DATABASE:
        try {
          const {data} = await getTables({cid: node.data.cid, db: node.data.name})
          childNodes = data ? data.map(table => ({
            name: table.name,
            level: Level.TABLE,
            cid: node.id,
            leaf: true
          })) : []

        } catch (error: any) {
          ElNotification({
            title: 'error',
            message: error.message,
            type: 'error'
          })
        } finally {
          resolve(childNodes)
        }
        break
      case Level.TABLE:
        // 处理 Level.TABLE 的情况
        break
    }
  }, 100)

  node.loading = false
  node.expanded = true
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const filterNode = (value: string, data: Tree) => {
  if (!value) return true
  return data.label.includes(value)
}
</script>

<style lang="scss" scoped>
.main-box {
  height: 100%;
  display: flex;
  flex-direction: row;
  border-top: 1px solid var(--color-border);
}

.tree-box {
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--color-border);
  color: var(--color-text) !important;
  background-color: var(--color-background-deep);
  position: relative;
  display: flex;
  flex-direction: column;
  .search {
    height: 60px;
  }
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
  font-size: var(--font-size)!important;
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
