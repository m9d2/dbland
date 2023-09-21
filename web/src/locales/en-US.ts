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
        export: 'Export',
        import: 'Import',
    },
    database: {
        button: {
            format: 'Format',
            run: 'Run',
            new_query: 'New Query',
        },
        label: {
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
            basic: 'Basic',
            other: 'About'
        },
        basic: {
            color: 'Color',
            restore: 'Default',
            select_language: 'Language',
        }
    }
}

export default {
    ...en,
    ...messages,
}