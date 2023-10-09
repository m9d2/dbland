<template>
  <div class="chart-box">
    <div class="chart-left">
      <span style="text-align: center;font-size: 20px; font-weight: bold; margin: 20px 0">{{ $t('chart.title') }}</span>
      <List class="chart-list" style="margin: 0 8px;" :list="charts" @node-click="handleNodeClick"
        @node-mouse-enter="handleMouseEnter" @node-mouse-leave="handleMouseLeave">
        <template #default="{ index }">
          <el-dropdown style="vertical-align: center;" size="small" trigger="click">
            <el-link v-show="index === activeIndex" :underline="false" :icon="ArrowDown" />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :icon="Edit" @click="deleteChart(index)">{{ $t('common.edit') }}</el-dropdown-item>
                <el-dropdown-item :icon="Delete" @click="deleteChart(index)">{{ $t('common.delete') }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </List>
      <div class="new-chart">
        <el-button type="primary" style="width: 100%;" @click="createChart">{{ $t('chart.button.new_chart') }}</el-button>
      </div>
    </div>
    <div class="chart-content">
      <div class="chart-main">
        <div class="chart-container" v-for="item in chartsContent">
          <div :id="item.name" style="width: 800px; height: 400px;">

          </div>
        </div>
      </div>
      <Form v-model="formVisiable" @cancel="handlerCancel" @confirm="handlerConfirm"></Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts';
import List from '@/components/layout/list/index.vue';
import { getCharts } from '@/api/chart/index';
import { ElNotification } from "element-plus";
import type { Chart } from "@/api/chart/type";
import type { AxiosPromise } from "axios";
import { Delete, ArrowDown, Edit } from '@element-plus/icons-vue'
import Form from './form/index.vue';

const charts = ref()
const chartsContent = [
  {
    id: 1,
    name: 'test1',
  },
  {
    id: 2,
    name: 'test2',
  },
  {
    id: 3,
    name: 'test3',
  }
]
const activeIndex = ref()
const formVisiable = ref(false)

onMounted(() => {
  loadCharts();
})

function handleNodeClick(index: any, node: any) {
  for (const item of chartsContent) {
    var chart = echarts.init(document.getElementById(item.name));
    const option = {
      title: {
        text: node.name
      },
      tooltip: {},
      xAxis: {
        data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
      },
      yAxis: {},
      series: [
        {
          name: '销量',
          type: 'bar',
          data: [5, 20, 36, 10, 10, 20]
        }
      ]
    };
    chart.setOption(option);
  }
}

function loadCharts() {
  try {

    let list: any = []
    const response: AxiosPromise<Chart[]> = getCharts();

    response.then((res: any) => {
      if (res.data) {
        res.data.forEach((data: any) => {
          data.name = data.title
          list.push(data)
        })
        charts.value = list
      }
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: 'error'
    })
  }
}

function deleteChart(index: number) {

}

function handleMouseEnter(index: number) {
  activeIndex.value = index;
}

function handleMouseLeave(index: number) {
  activeIndex.value = null;
}

function createChart() {
  formVisiable.value = true
}


function handlerCancel() {
  formVisiable.value = false
}

function handlerConfirm() {
  formVisiable.value = false
}
</script>

<style scoped lang="scss">
.chart-left {
  height: 100%;
  width: 220px;
  min-width: 220px;
  border-right: 1px solid var(--db-c-border);
  color: var(--db-c-text-menu);
  background-color: var(--db-c-bg-nav);
  position: relative;
  display: flex;
  flex-direction: column;

  .chart-list {
    overflow: auto;
    flex-grow: 1;
  }

  .new-chart {
    width: 200px;
    height: 48px;
    min-height: 48px;
    margin: 20px auto 0;
    justify-content: center;
    box-sizing: border-box;
  }
}

.chart-box {
  display: flex;
  height: 100%;
}

.chart-content {
  flex-grow: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: auto;
}

.chart-main {
  width: 80%;
  height: 100%;
}

.chart-container {
  height: 400px;
  width: 800px;
  margin: 20px auto;
}

.el-dropdown {
  line-height: 30px;
  vertical-align: middle;

  .el-link {
    font-size: var(--font-size);
  }
}
</style>