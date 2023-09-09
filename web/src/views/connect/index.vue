<template>
  <div class="connect-box">
    <div class="connect-left">
        <span style="text-align: center;font-size: 20px; font-weight: bold; margin-top: 20px">Connections</span>
        <div class="configs">
          <div class="config" :class="{active: item.active}" @click="clickConfig(index)" v-for="(item, index) in configs">
            <span class="config-name">{{item.name}}</span>
            <span class="config-menu">
              <el-popconfirm title="Are you sure delete it?" @confirm="drop(item.id)">
                <template #reference>
                   <el-link :underline="false" :icon="Delete" />
                </template>
              </el-popconfirm>
              </span>
          </div>
        </div>
    </div>
    <div class="connect-content">
      <div class="connect-main">
        <ul v-if="itemDisplay">
          <li>
            <Item name="MySQL" icon="&#xec6d;" @click="save('MySQL')"></Item>
          </li>
          <li>
            <Item name="ORACLE" icon="&#xec48;" @click="save('ORACLE')"></Item>
          </li>
          <li>
            <Item name="PostgreSQL" icon="&#xe8b7;" @click="save('PostgreSQL')"></Item>
          </li>
          <li>
            <Item name="SQLite" icon="&#xe65a;" @click="save('SQLite')"></Item>
          </li>
          <li>
            <Item name="MariaDB" icon="&#xec6d;" @click="save('MariaDB')"></Item>
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
import {onMounted, ref, nextTick} from "vue";
import {DBConfig} from "@/api/config/type";
import {getConfigs, deleteConfig} from "@/api/config";
import {ElNotification} from "element-plus";
import { Delete } from '@element-plus/icons-vue'

const itemDisplay = ref(true)
const formDisplay = ref(false)
const dbType = ref()
const configs = ref()
const formRef = ref()
const formTypeRef = ref()

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
  } catch (error) {
    ElNotification({
      title: 'error',
      message: error.message,
      type: 'error'
    })
  }
}

function save(name) {
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

async function drop(id: number) {
  try {
    await deleteConfig(id)
    await loadConfigs()
    ElNotification({
      message: 'Delete Successful',
      type: 'success'
    })
  } catch (error) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
}

async function clickConfig(index: number) {
  for (let i = 0; i < configs.value.length; i++) {
    configs.value[i].active = i == index;
    if (i == index) {

    }
  }
  const config = configs.value[index]
  itemDisplay.value = false
  formDisplay.value = true
  dbType.value = 'MySQL'
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
  border-right: 1px solid var(--color-border);
  color: var(--color-text) !important;
  background-color: var(--color-background-deep);
  position: relative;
  display: flex;
  flex-direction: column;
  .configs {
    padding: 20px 20px;
    .config {
      display: flex;
      box-sizing: border-box;
      height: 30px;
      line-height: 30px;
      padding-left: 10px;
      margin-top: 2px;
      .config-name{
        flex-grow: 1;
      }
      .config-menu {
        width: 30px;
        line-height: 30px;
        text-align: center;
        display: none;
      }
    }

    .config:hover {
      background-color: var(--tree-node-hover-bg-color);
    }

    .config:hover .config-menu {
      display: block;
    }
  }
  .active {
    color: var(--color-text-hover);
    background-color: var(--tree-node-hover-bg-color);
  }
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