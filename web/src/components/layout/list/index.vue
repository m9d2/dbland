<template>
    <div class="d-list">
        <div class="d-list-row" :style="rowStyle" :class="{ active: item.active }"
            v-for="(item, index) in list">
            <div :value="index" class="d-row-content" @mouseover="handleMouseEnter(index)" @contextmenu.native="handleContextMenu(index, item)"
                @dblclick.prevent="handleDoubleClick(index, item)" @mouseleave="handleMouseLeave(index)">
                <div class="d-row-begin">
                    <slot name="begin" :node="item">
                    </slot>
                </div>
                <span class="d-row-name" :style="rowNameStyle" @click="nodeClick(index, item)">
                    {{ item.name }}
                </span>
                <div class="d-row-menu">
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
    rowNameStyle: Object,
    rowStyle: Object,
})

const emits = defineEmits(["node-click", 'node-mouse-enter', 'node-mouse-leave', 'node-db-click', 'contextmenu']);

function nodeClick(index: number, row: Item) {
    if (props.list) {
        for (let i = 0; i < props.list.length; i++) {
            props.list[i].active = i === index;
        }
    }
  emits('node-click', index, row)
}

function handleMouseEnter(index: number) {
  emits('node-mouse-enter', index)
}

function handleDoubleClick(index: number, row: any) {
  emits('node-db-click', index, row)
}

function handleMouseLeave(index: number) {
  emits('node-mouse-leave', index)
}

function handleContextMenu(index: number, row: any) {
    emits('contextmenu', index, row, event)
}
</script>

<style lang="scss" scoped>
.d-list {
    overflow: auto;
    .d-list-row {
        box-sizing: border-box;
        height: 30px;
        line-height: 30px;
        margin-top: 2px;
        padding: 0 8px;
        cursor: pointer;
        border-radius: 5px;

        .d-row-content {
            display: flex;
            .d-row-name {
                font-weight: 500;
                flex-grow: 1;
            }
        }
        .d-row-name {
            user-select:none;
            overflow: hidden;
            text-overflow: ellipsis; 
        }
        .d-row-name:hover {
          color: var(--db-c-text-hover);
        }
    }

    .d-list-row:hover {
        color: var(--db-c-text-hover);
        background-color: var(--db-c-bg-hover);
    }

    .active {
        color: var(--db-c-text-hover);
        background-color: var(--db-c-bg-hover);
    }
}
</style>