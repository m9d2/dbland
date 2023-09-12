<template>
  <el-form :model="form" label-width="120px">
    <el-form-item>
      <span style="margin: 0 auto; height: 30px; font-size: 28px">{{
        getName
      }}</span>
    </el-form-item>
    <el-input v-show="false" v-model="form.id" />
    <el-form-item label="Name">
      <el-input v-model="form.name" />
    </el-form-item>
    <el-form-item label="Host" v-show="showHost">
      <el-col :span="11">
        <el-input v-model="form.host" />
      </el-col>
      <el-col :span="2">
        <span style="display: block; text-align: center">Port</span>
      </el-col>
      <el-col :span="11" class="text-center">
        <el-input v-model="form.port" />
      </el-col>
    </el-form-item>
    <el-form-item label="Path" v-show="showPath">
      <el-input v-model="form.db_file" type="file" />
    </el-form-item>
    <el-form-item label="User" v-show="showUser">
      <el-input v-model="form.username" />
    </el-form-item>
    <el-form-item label="Password" v-show="showPassword">
      <el-input v-model="form.password" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button @click="test">Test</el-button>
      <el-button type="primary" v-show="type == 'save'" @click="save"
        >Save</el-button
      >
      <el-button type="primary" v-show="type == 'edit'" @click="edit"
        >Edit</el-button
      >
      <el-button @click="cancel">Cancel</el-button>
    </el-form-item>
  </el-form>
</template>
<script setup lang="ts">
import { onMounted, reactive, computed } from "vue";
import { createConfig, updateConfig } from "@/api/config";
import { ping } from "@/api/connector";
import { ElNotification } from "element-plus";
import { DbTypeEnum } from "@/common/enums";
import {findKeyByValue} from '@/common/utils'

const props = defineProps({
  dbType: {
    type: String,
    default: ''
  },
  type: {
    type: String,
    default: ''
  }
});
const emits = defineEmits(["cancel", "success"]);
const form = reactive({
  id: null,
  name: "",
  host: "",
  port: "",
  username: "",
  password: "",
  db_file: null,
});
defineExpose({
  setConfig,
});

function cancel() {
  emits("cancel");
}

function success() {
  emits("success");
}

onMounted(() => {
  if (props.dbType.toLowerCase() == "mysql") {
    form.host = "127.0.0.1";
    form.port = "3306";
    form.username = "root";
  }
  if (props.dbType.toLowerCase() == "oracle") {
    form.host = "127.0.0.1";
    form.port = "1521";
  }
});

const getName = computed(() => {
  return findKeyByValue(DbTypeEnum, props.dbType)
})

const showHost = computed(() => {
  switch (props.dbType) {
    case DbTypeEnum.MySQL:
      return true;
    case DbTypeEnum.MariaDB:
      return true;
    case DbTypeEnum.ORACLE:
      return true;
    case DbTypeEnum.PostgreSQL:
      return true;
    case DbTypeEnum.SQLite:
      return false;
    default:
      return false;
  }
});

const showUser = computed(() => {
  const key = DbTypeEnum.MySQL
  switch (props.dbType) {
    case DbTypeEnum.MySQL:
      return true;
    case DbTypeEnum.MariaDB:
      return true;
    case DbTypeEnum.ORACLE:
      return true;
    case DbTypeEnum.PostgreSQL:
      return true;
    case DbTypeEnum.SQLite:
      return true;
    default:
      return false;
  }
});

const showPassword = computed(() => {
  switch (props.dbType) {
    case DbTypeEnum.MySQL:
      return true;
    case DbTypeEnum.MariaDB:
      return true;
    case DbTypeEnum.ORACLE:
      return true;
    case DbTypeEnum.PostgreSQL:
      return true;
    case DbTypeEnum.SQLite:
      return true;
    default:
      return false;
  }
});

const showPath = computed(() => {
  return props.dbType == DbTypeEnum.SQLite;
});

function getConfig() {
  return {
    id: form.id,
    type: props.dbType,
    name: form.name,
    host: form.host,
    port: form.port,
    username: form.username,
    password: form.password,
  };
}

function setConfig(config: any) {
  form.id = config.id;
  form.host = config.host;
  form.port = config.port;
  form.username = config.username;
  form.name = config.name;
  form.password = config.password;
}

const save = async () => {
  try {
    const config: any = getConfig();
    await createConfig(config);
    success();
    ElNotification({
      message: "Save Successful",
      type: "success",
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  }
};

const edit = async () => {
  try {
    const config: any = getConfig();
    await updateConfig(config);
    success();
    ElNotification({
      message: "Save Successful",
      type: "success",
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  }
};

const test = async () => {
  try {
    await ping(getConfig());
    ElNotification({
      message: "Connection Successful",
      type: "success",
    });
  } catch (error: any) {
    ElNotification({
      message: error.message,
      type: "error",
    });
  }
};
</script>
<style scoped lang="scss"></style>
