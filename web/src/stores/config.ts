import {defineStore} from 'pinia'
import {Names} from './type'

export const ConfigStore = defineStore(Names.CONFIG, {
    state: () => {
        return {}
    },
})
