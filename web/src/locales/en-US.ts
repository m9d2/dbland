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
        connect: 'Connect',
    },
    message: {
        delete_record: 'Delete this record?',
    },
    database: {
        button: {
            format: 'Format',
            run: 'Run',
            new_query: 'New Query',
            query: 'Query',
            refresh: 'Refresh',
        },
        label: {
            total: 'Total',
            elapsed_time: 'Elapsed Time',
            total_page: 'Total Page',
        },
    },
    connect: {
        title: 'Connections',
    },
    setting: {
        title: 'Setting',
        menu: {
            basic: 'Basic Settings',
            about: 'About'
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