<template>
    <div class="d-list">
        <div class="d-list-row" :class="{ active: item.active }" @click="nodeClick(index, item)"
            v-for="(item, index) in list">
            <div :value="index" class="d-row-content clearfix" @mouseover="handleMouseEnter(index)"
                @mouseleave="handleMouseLeave(index)">
                <span class="d-row-name fl">{{ item.name }}</span>
                <div class="d-row-meun fr">
                    <slot :index="index">
                    </slot>
                </div>
            </div>

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

const emuis = defineEmits(["node-click", 'node-mouse-enter', 'node-mouse-leave']);

function nodeClick(index: number, row: Item) {
    if (props.list) {
        for (let i = 0; i < props.list.length; i++) {
            props.list[i].active = i === index;
        }
    }
    emuis('node-click', index, row)
}

function handleMouseEnter(index: number) {
    emuis('node-mouse-enter', index)
}

function handleMouseLeave(index: number) {
    emuis('node-mouse-leave', index)
}
</script>

<style lang="scss" scoped>
.d-list {
    .d-list-row {

        box-sizing: border-box;
        height: 30px;
        line-height: 30px;
        margin-top: 2px;
        padding: 0 8px;
        cursor: pointer;
        border-radius: 5px;

        .d-row-content {

            .d-row-name {
                font-weight: 500;
            }

            .d-row-meun {
                width: 30px;
            }
        }


    }

    .d-list-row:hover {
        background-color: var(--tree-node-hover-bg-color);
    }

    .active {
        color: var(--color-text-hover);
        background-color: var(--tree-node-hover-bg-color);
    }
}
</style>