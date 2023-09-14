import zhCn from 'element-plus/dist/locale/zh-cn.mjs';

const messages = {
    common: {
        select: '选择',
        search: '查询',
        delete: '删除',
        homepage: '首页',
        add: '添加',
        modify: '编辑',
        refresh: '刷新',
        confirm: '确认',
        cancel: '取消',
    },
    database: {
        button: {
            format: '格式化',
            run: '执行',
            new_query: '新建查询',
        },
        lable: {
            total: '总数',
            elapsed_time: '执行时间',
        },
    },
    connect: {
        title: '数据库连接',
    },
    setting: {
        title: '设置',
        menu: {
            basic: '基础设置',
            other: '其他'
        },
        basic: {
            color: '设置主题颜色',
            restore: '恢复默认',
            select_language: '选择语言',
        }
    }
}

export default {
    ...zhCn,
    ...messages,
}