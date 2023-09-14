<template>
  <el-dialog style="width: 50%;">
    <el-form :model="formData" v-for="row in list" :label-width="labelWIdth" label-position="right"
    >
      <el-form-item :label="row.key" >
        <el-input v-model="row.value" />
      </el-form-item>

    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="cancel">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="confirm">
          {{ $t('common.confirm') }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'

const labelWIdth = ref()
const props = defineProps({
  row: Object,
})
const formData = reactive({})
const emit = defineEmits(['confirm', 'cancel'])
const formLabelAlign = reactive({
  name: '',
  region: '',
  type: '',
})
interface Row {
  key: string,
  value: any,
}

function calculateStringWidth(text: string): number {
  const canvas = document.createElement('canvas');
  const context = canvas.getContext('2d');
  if (context === null) {
    return 100;
  }
  context.font = '12px Arial';
  const metrics = context.measureText(text);
  return metrics.width;
}

const list = computed(() => {
  const keys: string[] = Object.keys(props.row)
  const data: Row[] = []
  let width = 0
  for (const key of keys) {
    data.push({ key: key, value: props.row[key] })
    const stringWidth = calculateStringWidth(key)
    if (stringWidth > width) {
      width = stringWidth
    }
  }
  labelWIdth.value = width + 30
  return data
})

function confirm() {
  console.log(formData)
  emit('confirm')
}

function cancel() {
  emit('cancel')
}

const form = {

}

</script>
<style scoped>
.el-form {
  padding: 0 48px;
}

.el-form-item{
  font-size: var(--font-size);
}

</style>