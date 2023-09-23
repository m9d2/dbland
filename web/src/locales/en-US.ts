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
        save: 'Save',
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
            color: 'Global Theme Colors',
            restore: 'Default',
            select_language: 'Language',
            font_size: 'Font size',
        }
    }
}

export default {
    ...en,
    ...messages,
}