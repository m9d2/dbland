<template>
  <el-dialog class="d-dialog" style="width: 50%; overflow: auto; border-radius: 5px;">
    <el-form :model="formData" :label-width="labelWidth" label-position="right">
      <el-form-item v-for="row in list" :label="row.key" :key="row.key">
        <el-input v-model="formData[row.key]" :type="row.type" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="cancel">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="confirm">{{ $t('common.confirm') }}</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { ActionTypeEnum } from '@/common/enums/'


const labelWidth = ref()
const props = defineProps({
  columns: {
    type: Object,
    default: null
  },
  row: Object,
  actionType: String,
})
const formData = ref({})
const emit = defineEmits(['confirm', 'cancel'])
interface Row {
  key: string,
  value: any,
  type: string,
}

function calculateStringWidth(text: string): number {
  const canvas = document.createElement('canvas');
  const context = canvas.getContext('2d');
  if (!context) {
    return 100;
  }
  context.font = '12px Arial';
  const metrics = context.measureText(text);
  return metrics.width;
}

const list = computed(() => {
  formData.value = {}
  if (props.actionType == ActionTypeEnum.MODIFY) {
    Object.assign(formData.value, props.row)

  }
  const data: Row[] = []
  let width = 0
  for (const key of props.columns) {
    let value
    if (props.row) {
      value = props.row[key.column_name]
    }
    data.push({ key: key.column_name, value: value, type: key.column_type })
    const stringWidth = calculateStringWidth(key.column_name)
    if (stringWidth > width) {
      width = stringWidth
    }
  }
  labelWidth.value = width + 30
  return data
})

function confirm() {
  emit('confirm', formData)
}

function cancel() {
  emit('cancel')
}

</script>
<style scoped>
.d-dialog {
  width: 50%; 
  height: 60%;
  overflow: scroll;
}
.el-form {
  padding: 0 48px;
}

.el-form-item {
  font-size: var(--font-size);
}
</style>