import {defineStore} from 'pinia'
import {Names} from './type'

export const useDatabaseStore = defineStore(Names.DATABASE, {
    state: () => ({
       currentDatabase: {},
       currentConfig: {},
    }),

    actions: {
        getCurrentDatabase() {
            return this.currentDatabase
        },
        setCurrentDatabase(database: any) {
            this.currentDatabase = database
        },
        getCurrentConfig() {
            return this.currentConfig
        },
        setCurrentConfig(config: any) {
            this.currentConfig = config
        }
    }
})
