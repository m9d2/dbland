import zhCn from 'element-plus/dist/locale/zh-cn.mjs';

const messages = {
    common: {
        select: '选择',
        search: '查询',
        delete: '删除',
        homepage: '首页',
        add: '添加',
        modify: '编辑',
        edit: '编辑',
        refresh: '刷新',
        confirm: '确认',
        cancel: '取消',
        export: '导入',
        import: '导出',
        save: '保存',
        connect: '连接',
    },
    database: {
        button: {
            format: '格式化',
            run: '执行',
            new_query: '新建查询',
            query: '查询',
            refresh: '刷新',
        },
        label: {
            total: '总数',
            elapsed_time: '执行时间',
            total_page: '总页数',
        },
    },
    connect: {
        title: '数据库连接',
    },
    setting: {
        title: '设置',
        menu: {
            basic: '基础设置',
            about: '关于'
        },
        basic: {
            color: '全局主题颜色',
            restore: '恢复默认',
            select_language: '选择语言',
            font_size: '字体大小',
            theme: '主题',
            dark: '暗黑模式',
            light: '明亮模式',
            auto: '自动',
        }
    },
    chart: {
        title: '图表',
        button: {
            new_chart: '新建图表',
        }
    }
}

export default {
    ...zhCn,
    ...messages,
}