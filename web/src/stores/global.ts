import { defineStore } from 'pinia'

export const useGlobalStore = defineStore({
    id: 'global',
    state: () => ({
        help: true,
    }),
    actions: {
        toggleHelp() {
            this.help = !this.help
        },
        setHelp(value: boolean) {
            this.help = value
        },
        getHelp() {
            return this.help
        }
    }
})
