<template>
  <el-tabs v-model="activeTab" class="tabs-content" closable @tab-click="tabClick" @tab-remove="removeTab">
    <el-tab-pane v-for="item in tabsData" :key="item.name" :label="item.title" :name="item.name" lazy>
      <template #label>
        <span class="custom-tabs-label">
          <span class="iconfont icon-biaoge blue" style="padding-right: 8px;"></span>
          <span>{{ item.title }}</span>
        </span>
      </template>
      <Query :config="configRef" :database="databaseRef" :sql="sqlRef"> </Query>
    </el-tab-pane>
  </el-tabs>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { createQuerySql } from '@/common/utils'
import Query from './query.vue'

// tabs data array
const tabsData = ref([])

const props = defineProps()
const isResultVisible = ref(props.modelValue)
const emit = defineEmits()

let tabIndex = 0
const activeTab = ref()
const configRef = ref({})
const databaseRef = ref({})
const sqlRef = ref('')

// change tab
function tabClick(node: any) {
}

// remove tab
function removeTab(targetName: string) {
  const tabs = tabsData.value
  let activeName = activeTab.value
  if (activeName === targetName) {
    tabs.forEach((tab, index) => {
      if (tab.name === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1]
        if (nextTab) {
          activeName = nextTab.name
        }
      }
    })
  }
  activeTab.value = activeName
  tabsData.value = tabs.filter((tab) => tab.name !== targetName)
  if (tabsData.value.length == 0) {
    isResultVisible.value = false
    tabIndex = 0
    emit('messageUpdated', isResultVisible.value)
  }
}

// new tab
function newTab(config: any, database: any, table: any) {
  if (config) {
    configRef.value = config
  }
  if (database) {
    databaseRef.value = database
  }

  isResultVisible.value = true
  if (table) {
    sqlRef.value = createQuerySql(database.name, table)
  }
  const newTabName = `${++tabIndex}`
  tabsData.value.push({
    title: `New Query ${tabIndex}`,
    name: newTabName,
  })
  activeTab.value = newTabName
}

// expose function
defineExpose({ newTab })
</script>

<style lang="scss">
.el-tabs {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.el-tabs__content {
  height: 100%;
  flex-grow: 1;
}

.el-tab-pane {
  height: 100%;
}

.el-tabs__header {
  margin: 0 0 8px;
}

.el-tabs__nav-scroll {
  background-color: var(--db-c-bg-nav);
  border-left: 1px solid var(--db-c-border);
  border-right: 1px solid var(--db-c-border);
  padding: 0 8px;
}

.el-tabs__item {
  border-right: 1px solid var(--db-c-border);
}

.el-tabs__item:last-child {
  border-right: none
}

.custom-tabs-label {
  font-size: 12px;
  vertical-align: middle;
}</style>
