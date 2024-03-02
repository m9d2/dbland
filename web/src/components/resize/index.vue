<template>
    <div id="resize-handle" :class="propsData.resizeX ? 'resize-handle-x' : 'resize-handle-y'" @mousedown="startResize"></div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";

const contentWidth = ref(0)
const contentHeight = ref(0);
const resizing = ref(false);
const resizeData = reactive({
    isResizing: false,
    startX: 0,
    startWidth: 0,
    startY: 0,
    startHeight: 0,
});
const propsData = defineProps({
    resizeX: {
        type: Boolean,
        default: false
    },
    resizeY: {
        type: Boolean,
        default: false
    },
    refId: {
        type: String,
        default: '',
        required: true
    }
})

onMounted(() => {
    var element = document.getElementById(propsData.refId);
    contentWidth.value = Number(element?.clientWidth);
    contentHeight.value = Number(element?.clientHeight);
})

const startResize = (event: MouseEvent) => {
    resizeData.isResizing = true;
    resizeData.startX = event.clientX;
    resizeData.startY = event.clientY;
    resizeData.startWidth = contentWidth.value;
    resizeData.startHeight = contentHeight.value;

    document.addEventListener('mousemove', resize);
    document.addEventListener('mouseup', stopResize);
}

const resize = (event: MouseEvent) => {
    if (resizeData.isResizing) {
        const element = document.getElementById(propsData.refId);
        if (element && propsData.resizeX) {
            let deltaX = event.clientX - resizeData.startX;
            contentWidth.value = resizeData.startWidth + deltaX;
            element.style.width = contentWidth.value + 'px';
            localStorage.setItem('width', contentWidth.value.toString())
        }
        if (element && propsData.resizeY) {
            let deltaY = event.clientY - resizeData.startY  ;
            contentHeight.value = resizeData.startHeight + deltaY;
            element.style.height = contentHeight.value + 'px';
            localStorage.setItem('height', contentHeight.value.toString())
        }
    }
}

const stopResize = () => {
    resizing.value = false;
    document.removeEventListener('mousemove', resize);
    document.removeEventListener('mouseup', stopResize);
}
</script>

<style scoped>
.resize-handle-x {
    width: 4px;
    height: 100%;
}

.resize-handle-y {
    height: 4px;
}

.resize-handle-x:hover {
    background-color: var(--db-c-border);
    cursor: col-resize;
}

.resize-handle-y:hover {
    background-color: var(--db-c-border);
    cursor: row-resize;
}
</style>