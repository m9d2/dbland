<template>
  <div class="content-box">
    <div class="content-main">
      <div class="content-left">
        <div id="content-tree" :style="{ width: contentWidth + 'px' }">
          <Tree @dbclick="doubleClickNode" @click-plus="handleClickPlus" @node-click="handleNodeClick">
            <template #node-menu="{ node }">
              <el-dropdown size="small" trigger="click">
                <el-link v-show="node" :underline="false" :icon="MoreFilled" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item :icon="Search" @click="newTab(node, 'query')">{{
                      $t("database.button.query") }}</el-dropdown-item>
                    <el-dropdown-item :icon="Edit">{{
                      $t("common.edit")
                    }}</el-dropdown-item>
                    <el-dropdown-item :icon="Delete">{{
                      $t("common.delete")
                    }}</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
          </Tree>
        </div>
        <Resize :resize-x="true" ref-id="content-tree" class="resize" :size="contentWidth"></Resize>
      </div>

      <div class="content-right">
        <div class="tab-content">
          <el-tabs v-model="activeTab" class="tabs-content" closable @tab-remove="removeTab">
            <el-tab-pane v-for="item in tabsData" :key="item.name" :label="item.title" :name="item.name" lazy>
              <template #label>
                <span class="custom-tabs-label unselectable">
                  <span class="iconfont icon-biaoge blue" style="padding-right: 8px;"></span>
                  <span>{{ item.title }}</span>
                </span>
              </template>
              <template #default>
                <TabTable :node="currentNode" v-if="item.type == TabType.TABLE" />
                <TabQuery :node="currentNode" v-if="item.type == TabType.QUERY" />
              </template>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, onBeforeMount} from "vue";
import TabTable from './components/TabTable.vue';
import TabQuery from './components/TabQuery.vue';
import Tree from './components/Tree.vue';
import Resize from '@/components/resize/index.vue';
import {
  Edit,
  Delete,
  Search,
  MoreFilled,
} from "@element-plus/icons-vue";

const isResultVisible = ref();
const shortcutsVisible = ref(true);
const resizing = ref(false);
const contentWidth = ref(220)
const currentNode = ref()
const activeTab = ref()
const tabsData = ref<any[]>([])
let tabIndex = 0
const resizeData = reactive({
  isResizing: false,
  startX: 0,
  startWidth: 0,
});

enum TabType {
  TABLE = 'table',
  QUERY = 'query',
}

const handleClickPlus = () => {
  newTab(currentNode.value, 'query')
}

onBeforeMount(() => {
  let width = localStorage.getItem('width')
  if (width) {
    contentWidth.value = parseInt(width)
  }
})

// 新建tab
const newTab = (node: any, type: string) => {
  let newTabName;
  if (type == TabType.QUERY) {
    newTabName = `${++tabIndex}`;
    newTabName = `new query ${tabIndex}`;
    tabsData.value.push({
      title: `new query ${tabIndex}`,
      name: newTabName,
      type: TabType.QUERY,
    });
  }

  if (type == TabType.TABLE) {
    let tabExists = false;
    newTabName = node.data.name + "@" + node.parent.data.name;
    tabExists = false;
    for (const tab in tabsData.value) {
      if (tabsData.value[tab].name == newTabName) {
        tabExists = true;
        return;
      }
    }
    if (!tabExists) {
      tabsData.value.push({
        title: newTabName,
        name: newTabName,
        type: TabType.TABLE,
      });
    }
  }
  activeTab.value = newTabName;
};

const handleNodeClick = (node: any) => {
  // currentNode.value = node
}

const doubleClickNode = (node: any) => {
  currentNode.value = node
  if (node.level === 3) {
    newTab(node, TabType.TABLE);
  }
};

const refresh = async () => {
}

function removeTab(targetName: string) {
  const tabs = tabsData.value
  let activeName = activeTab.value
  if (activeName === targetName) {
    tabs.forEach((tab: any, index: any) => {
      if (tab.name === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1]
        if (nextTab) {
          activeName = nextTab.name
        }
      }
    })
  }
  activeTab.value = activeName
  tabsData.value = tabs.filter((tab: any) => tab.name !== targetName)
  if (tabsData.value.length == 0) {
    tabIndex = 0
    shortcutsVisible.value = true
    isResultVisible.value = false
  }
}

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
</script>

<style lang="scss" scoped>
.content-box {
  display: flex;
  flex-direction: column;
  height: 100%;
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
    flex-direction: row;
    border-right: 1px solid var(--db-c-border);
    // margin-right: 8px;

    .content-tree {
      height: 100%;
      display: flex;
      flex-direction: column;
    }

    .resize {
      // background-color: var(--db-c-bg-nav);
    }
  }



  .content-right {
    flex-grow: 1;
    height: 100%;
    overflow: hidden;
    background-color: var(--db-c-bg);
  }
}

.el-dropdown {
  line-height: 30px;
  vertical-align: middle;

  .el-link {
    font-size: var(--font-size);
  }
}

.tab-content {
  height: 100%;
  border-left: none;
}

.el-tabs {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.el-tab-pane {
  height: 100%;
}

.el-tabs__header {
  margin: 0 0 8px;
}

:deep(.el-tabs__nav-scroll) {
  background-color: var(--db-c-bg-nav);
  border-left: 1px solid var(--db-c-border);
  border-right: 1px solid var(--db-c-border);
  border-bottom: none;
}

:deep(.el-tabs__item) {
  padding: 0 8px;
}

.el-tabs__item:last-child {
  border-right: none
}

.custom-tabs-label {
  font-size: 12px;
  color: var(--db-c-text);
  vertical-align: middle;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: 8px;
}

:deep(.el-tabs__active-bar) {
  height: 1px;
}

:deep(.el-tabs__header) {
  margin: 0;
}

.no-right-border {
  border-right: none;
}
.right-border {
  border-right: 1px solid var(--db-c-border);
}
</style>
