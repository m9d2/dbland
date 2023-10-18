import en from 'element-plus/dist/locale/en.mjs';

const messages = {
    common: {
        select: 'Select',
        search: 'Search',
        delete: 'Delete',
        homepage: 'Home',
        add: 'Insert',
        modify: 'Modify',
        edit: 'Edit',
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
            query: 'Query',
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
            basic: 'Basic Settings',
            other: 'About Us'
        },
        basic: {
            color: 'Global Theme Colors',
            restore: 'Default',
            select_language: 'Language',
            font_size: 'Font Size',
            theme: 'Theme',
            dark: 'Dark',
            light: 'Light',
            auto: 'Auto',
        }
    },
    chart: {
        title: 'Charts',
        button: {
            new_chart: 'New Chart',
        }
    }
}

export default {
    ...en,
    ...messages,
}