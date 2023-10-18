class Base {

    databaseName!: string;
    tableName!: string;
    sqlStr: string

    constructor(database: string, sql: string) {
        const regex:RegExp = /FROM\s+`?([^.\s`]+)`?\.`?([^.\s`]+)`?/i;
        const matches:RegExpExecArray = regex.exec(sql);
        if (matches && matches.length === 3) {
            let databaseName:string = matches[1];
            let tableName:string = matches[2];
            if (databaseName == null && database == null) {
                this.databaseName = ''
            } else if (databaseName == null) {
                this.databaseName = database
            } else {
                this.databaseName = databaseName
            }
            if (tableName == null) {
                this.tableName = ''
            } else {
                this.tableName = tableName
            }
        }
        this.sqlStr = sql
    }

    parseSqlQuery(database: string, sqlQuery: string): { databaseName: string, tableName: string } | null {
        const regex:RegExp = /FROM\s+`?([^.\s`]+)`?\.`?([^.\s`]+)`?/i;
        const matches:RegExpExecArray = regex.exec(sqlQuery);
        if (matches && matches.length === 3) {
            let databaseName:string = matches[1];
            let tableName:string = matches[2];
            if(databaseName == null && database == null) {
                databaseName = ''
            }
            if (databaseName == null) {
                databaseName = database
            }
            if (tableName == null) {
                tableName = ''
            }
            return {databaseName, tableName}
        }
        return null;
    }
}

export {Base}