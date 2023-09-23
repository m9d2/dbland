<template>
  <div class="connect-box">
    <div class="connect-left">
      <span style="text-align: center;font-size: 20px; font-weight: bold; margin: 20px 0">{{ $t('connect.title') }}</span>
      <List :list="configs" style="margin: 0 20px;" @node-click="clickConfig" @node-mouse-enter="handleMouseEnter"
        @node-mouse-leave="handleMouseLeave">
        <template #default="{ index }">
          <el-dropdown style="vertical-align: center;" size="small" trigger="click">
            <el-link v-show="index === activeIndex" :underline="false" :icon="ArrowDown" />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :icon="Delete" @click="deleteConfig(index)">{{ $t('common.delete') }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </List>
    </div>
    <div class="connect-content">
      <div class="connect-main">
        <ul v-if="itemDisplay">
          <li>
            <Item name="MySQL" icon="&#xec6d;" @click="save(DbTypeEnum.MySQL)"></Item>
          </li>
          <li>
            <Item name="ORACLE" icon="&#xec48;" @click="save(DbTypeEnum.ORACLE)"></Item>
          </li>
          <li>
            <Item name="PostgreSQL" icon="&#xe8b7;" @click="save(DbTypeEnum.PostgreSQL)"></Item>
          </li>
          <li>
            <Item name="SQLite" icon="&#xe65a;" @click="save(DbTypeEnum.SQLite)"></Item>
          </li>
          <li>
            <Item name="MariaDB" icon="&#xec6d;" @click="save(DbTypeEnum.MariaDB)"></Item>
          </li>
        </ul>
        <div class="connect-form" v-if="formDisplay">
          <Form :db-type="dbType" :type="formTypeRef" @cancel="cancel" @success="success" ref="formRef"></Form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Item from './item/index.vue'
import Form from './form/index.vue'
import { onMounted, ref, nextTick } from "vue";
import type { DBConfig } from "@/api/config/type";
import { getConfigs, deleteById } from "@/api/config";
import { ElNotification } from "element-plus";
import { Delete, ArrowDown } from '@element-plus/icons-vue'
import { findKeyByValue } from '@/common/utils'
import type { AxiosPromise } from 'axios';
import { DbTypeEnum } from '@/common/enums'
import List from '@/components/layout/list/index.vue'

const itemDisplay = ref(true)
const formDisplay = ref(false)
const dbType = ref()
const configs = ref()
const formRef = ref()
const formTypeRef = ref()
const activeIndex = ref()

onMounted(() => {
  loadConfigs();
})

function loadConfigs() {
  try {
    let childNodes: any = []
    const response: AxiosPromise<DBConfig[]> = getConfigs();
    response.then((res: any) => {
      if (res.data) {
        res.data.forEach((data: any) => {
          data.active = false
          childNodes.push(data)
        })
        configs.value = childNodes
      }
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
}

async function deleteConfig(index) {
  try {
    const id = configs.value[index].id
    console.log(id)
    await deleteById(id)
    loadConfigs()
  } catch (error) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
}

function save(name: string) {
  itemDisplay.value = false
  formDisplay.value = true
  dbType.value = name
  formTypeRef.value = 'save'
}

function cancel() {
  itemDisplay.value = true
  formDisplay.value = false
}

function success() {
  loadConfigs()
  itemDisplay.value = true
  formDisplay.value = false
}

function handleMouseEnter(index: number) {
  activeIndex.value = index;
}

function handleMouseLeave(index: number) {
  activeIndex.value = null;
}

async function drop(id: number) {
  try {
    await deleteConfig(id)
    loadConfigs()
    ElNotification({
      message: 'Delete Successful',
      type: 'success'
    })
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
}

async function clickConfig(index: number) {
  for (let i = 0; i < configs.value.length; i++) {
    configs.value[i].active = i == index;
  }
  const config = configs.value[index]
  itemDisplay.value = false
  formDisplay.value = true
  dbType.value = findKeyByValue(DbTypeEnum, config.type)?.toLowerCase()
  await nextTick();
  formTypeRef.value = 'edit'
  formRef.value.setConfig(config)
}
</script>
<style scoped lang="scss">
.connect-left {
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--db-c-border);
  color: var(--db-c-text-menu);
  background-color: var(--db-c-bg-nav);
  position: relative;
  display: flex;
  flex-direction: column;

  .el-dropdown {
    
    line-height: 30px;
    vertical-align: middle;
    .el-link {
      font-size: var(--font-size);
    }
  }
}

.show {
  display: none;
}

.hide {
  display: block;
}

.connect-box {
  display: flex;
  height: 100%;
}

.connect-content {
  flex-grow: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.connect-main {
  width: 600px;

  ul {
    display: flex;
    flex-wrap: wrap;

    li {
      padding: 20px;
      width: 200px;

      .el-button {
        height: 100%;
        width: 100%;
      }
    }
  }
}
</style>