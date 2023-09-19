import { createI18n, type LocaleMessages, type VueMessageType } from 'vue-i18n'

import zh from '@/locales/zh-CN';
import en from '@/locales/en-US';

const i18n = createI18n({
    legacy: false, 
    locale: localStorage.getItem('language') || 'en',
    globalInjection: true,
    messages: {
        en,
        zh,
    },
})

export default i18n;