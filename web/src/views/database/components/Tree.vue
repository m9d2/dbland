<template>
  <div class="search">
    <el-input v-model="filterText" :placeholder="$t('common.search')" size="small" clearable>
      <template #prefix>
        <el-icon>
          <Search />
        </el-icon>
      </template>
    </el-input>
    <el-link :icon="Refresh" @click="refresh" :underline="false" />
    <el-link :icon="Plus" @click="handleClickPlus" :underline="false" />
  </div>
  <el-tree ref="treeRef" :props="props" :load="loadNode" empty-text="" lazy @check-change="handleCheckChange" :indent="12"
    :highlight-current="true" :filter-node-method="filterNode" @node-click="handleNodeClick">
    <template v-slot="{ node }">
      <div class="tree-node unselectable" @dblclick="handleDoubleClick(node)" @mouseover="handleMouseEnter(node)"
        @mouseleave="handleMouseLeave(node)">
        <div class="node-content unselectable">
          <Icon :node="node"/>
          <span>{{ node.data.name }}</span>
        </div>
        <div class="node-menu" v-show="false">
          <slot name="node-menu" :node="node">
          </slot>
        </div>
      </div>
    </template>
  </el-tree>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { TreeLevelEnum } from "@/common/enums";
import { getConfigs } from "@/api/config";
import { getDatabases, getTables } from "@/api/connector";
import type { DBConfig } from "@/api/config/type";
import type { AxiosPromise } from "axios";
import { ElNotification, ElTree } from "element-plus";
import Icon from "./Icon.vue"
const treeRef = ref<InstanceType<typeof ElTree>>()
import {
  Refresh,
  Search,
  Plus,
} from "@element-plus/icons-vue";

const filterText = ref("");
const emit = defineEmits(['node-click', 'dbclick', 'click-plus'])

const props = {
  label: "name",
  children: "children",
  isLeaf: "leaf",
};

interface Tree {
  name: string
  leaf?: boolean
}

const loadNode = async (node: any, resolve: (data: Tree[]) => void) => {
  let nodes: any;
  if (node.level === TreeLevelEnum.ROOT) {
    nodes = await loadConfigs();
    nodes = nodes.map((childNode: any) => {
      return childNode;
    })
    return resolve(nodes);
  }
  if (node.level === TreeLevelEnum.CONFIG) {
    nodes = await loadDatabases(node.data.id);
    nodes = nodes.map((childNode: any) => {
      return childNode;
    })
    return resolve(nodes);
  }
  if (node.level === TreeLevelEnum.DATABASE) {
    nodes = await loadTables(node.data.cid, node.data.name);
    nodes = nodes.map((childNode: any) => {
      childNode.leaf = true
      return childNode;
    });
    return resolve(nodes);
  }
  return resolve([])
};

const handleDoubleClick = (node: any) => {
  emit('dbclick', node);
}

const handleClickPlus = () => {
  emit('click-plus')
}

const handleCheckChange = (
  data: Tree,
  checked: boolean,
  indeterminate: boolean
) => {
  console.log(data, checked, indeterminate)
}

const handleNodeClick = (node: any) => {
  emit('node-click', node)
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

const loadTables = async (cid: number, datbase: string) => {
  const { data } = await getTables({
    cid: cid,
    database: datbase,
  });
  return data
    ? data.map((table) => ({
      name: table.name,
      cid: cid
    }))
    : []
}

const handleMouseEnter = (node: any) => {
  // if (node.level == TreeLevelEnum.TABLE) {
  //   node.data.isCurrent = true;
  // }
};

const handleMouseLeave = (node: any) => {
  // if (node.level == TreeLevelEnum.TABLE) {
  //   node.data.isCurrent = false;
  // }
};

const filterNode = (value: string, data: any) => {
  if (!value) return true;
  return data.name.includes(value);
};

const refresh = async () => {
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})
</script>
<style lang="scss" scoped>
.search {
  border-radius: 5px;
  padding: 12px;
  display: flex;
  align-items: center;

  .el-icon {
    margin: 0 5px 0 0;
  }

  .el-link {
    font-size: 16px;
    vertical-align: bottom;
    margin-left: 8px;
  }
}

.resize-handle {
  position: absolute;
  top: 0;
  right: 0;
  height: 100%;
  z-index: 999;
  width: 6px;
  border-right: 1px solid var(--db-c-border);
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

.node-content {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
@/stores/global