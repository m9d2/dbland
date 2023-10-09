<template>
    <el-dialog title="New Chart" class="d-dialog" style="width: 50%; overflow: auto; border-radius: 5px;" :close-on-click-modal="false">
        <el-form :model="form" label-width="120px">
            <el-form-item label="Activity name">
                <el-input v-model="form.name" />
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="cancel">Cancel</el-button>
                <el-button type="primary" @click="confirm">
                    Save
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';

const form = reactive({
    name: '',
    region: '',
    date1: '',
    date2: '',
    delivery: false,
    type: [],
    resource: '',
    desc: '',
})

const emit = defineEmits(['confirm', 'cancel'])

function confirm() {
    console.log('confirm!')
  emit('confirm', form)
}

function cancel() {
  emit('cancel')
}
</script>

<style scoped>
</style>