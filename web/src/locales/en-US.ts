import en from 'element-plus/dist/locale/en.mjs';

const messages = {
    common: {
        select: 'Select',
        search: 'Search',
        delete: 'Delete',
        homepage: 'Home',
        add: 'Insert',
        modify: 'Modify',
        refresh: 'Refresh',
        confirm: 'Confirm',
        cancel: 'Cancel',
    },
    database: {
        button: {
            format: 'Format',
            run: 'Run',
            new_query: 'New Query',
        },
        lable: {
            total: 'Total',
            elapsed_time: 'Elapsed Time',
        },
    },
    connect: {
        title: 'Connections',
    },
    setting: {
        title: 'Setting',
        menu: {
            basic: 'Basic Settings',
            other: 'Other'
        },
        basic: {
            color: 'Setting theme color',
            restore: 'Restore default',
            select_language: 'Select Language',
        }
    }
}

export default {
    ...en,
    ...messages,
}