<template>
  <el-dialog>
    <el-form :model="formData" v-for="row in list">
      <el-form-item :label="row.key" :label-width="formLabelWidth">
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

const formLabelWidth = '140px'
const props = defineProps({
  row: Object,
})
const formData = reactive({})
const emit = defineEmits(['confirm', 'cancel'])

interface Row {
  key: string,
  value: any,
}

const list = computed(() => {
  const keys: string[] = Object.keys(props.row)
  const data: Row[] = []
  for (const key of keys) {
    data.push({ key: key, value: props.row[key] })
  }
  console.log(data)
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
.el-input {
  width: 300px;
}
</style>