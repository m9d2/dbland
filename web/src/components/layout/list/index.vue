<template>
    <div class="d-list">
        <div class="d-list-row" :class="{ active: item.active }" @click="nodeClick(index, item)" v-for="(item, index) in list">
            <span class="d-list-name">{{ item.name }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { Item } from './type';

const props = defineProps({
    list: {
        type: Array as () => Item[],
    },
})

const emuis = defineEmits(["node-click"]);

function nodeClick(index: number, row: Item) {
    if (props.list) {
        for (let i = 0; i < props.list.length; i++) {
            props.list[i].active = i === index;
        }
    }
    emuis('node-click', index, row)
}
</script>

<style lang="scss" scoped>
.d-list {
    .d-list-row {
        display: flex;
        box-sizing: border-box;
        height: 30px;
        line-height: 30px;
        padding-left: 10px;
        margin-top: 2px;
        cursor: pointer;
        border-radius: 5px;

        .d-list-name {
            flex-grow: 1;
        }
    }

    .active {
        color: var(--color-text-hover);
        background-color: var(--tree-node-hover-bg-color);
    }
}
</style>