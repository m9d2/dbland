<template>
    <el-icon v-show="propsData.node?.level == TreeLevelEnum.CONFIG">
        <span v-for="item in icons" v-show="propsData.node?.data.type == item.dbType" class="iconfont"
            :class="item.icon"></span>
    </el-icon>
    <el-icon v-show="propsData.node?.level == TreeLevelEnum.DATABASE" class="iconfont">
        <Folder />
    </el-icon>
    <el-icon v-show="propsData.node?.level == TreeLevelEnum.TABLE" class="iconfont">
        <span class="iconfont icon-biaoge blue"></span>
    </el-icon>
</template>

<script setup lang="ts">
import { DbTypeEnum, TreeLevelEnum } from "@/common/enums";
import { defineProps } from "vue";
import {
    Folder,
} from "@element-plus/icons-vue";
const propsData = defineProps({
    node: Object,
})

const icons = [
    { dbType: DbTypeEnum.MySQL, icon: "icon-mysql" },
    { dbType: DbTypeEnum.SQLite, icon: "icon-sqlite" },
    { dbType: DbTypeEnum.ORACLE, icon: "icon-oracle" },
    { dbType: DbTypeEnum.PostgreSQL, icon: "icon-PostgreSQL" },
    { dbType: DbTypeEnum.MariaDB, icon: "icon-mysql" },
]
</script>

<style lang="scss" scoped>
.el-icon {
    margin-right: 8px;
}
</style>