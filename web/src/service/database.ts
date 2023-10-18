interface Database {
    createInsertSql(column: any, row: any): string | Error
    createUpdateSql(column: any, data: any, originalData: any): string | Error
    createDeleteSql(column: any, row: any): string | Error
}

export type {Database}