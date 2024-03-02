import type { Sort } from "./model/sort"

interface Database {

    createQuerySql(table: string): string | undefined
    createQueryPageSql(table: string, page: number, size: number, where: string, sort: Sort): string | undefined
    createInsertSql(table: string, columns: Column[], row: any): string | undefined
    createUpdateSql(table: string, columns: Column[], data: any, originalData: any): string | undefined
    createDeleteSql(table: string, columns: Column[], row: any): string | undefined
}

export type {Database}