import {ColumnType} from "@/common/enums";
import {Base} from "./base";
import type {Database} from "./database";

class Sqlite extends Base implements Database{

    createInsertSql(column: any, row): string | Error {
        const values:any[] = [];
        const columns:any[] = [];
        for (const item: any of column) {
            columns.push(`\"${item.column_name}\"`)
            const val = row.value[item.column_name]
            if (val == null) {
                values.push('NULL')
            } else if (item.column_type == ColumnType.NUMBER) {
                values.push(val)
            } else if (item.column_type == ColumnType.TEXT) {
                values.push(`\"${val}\"`)
            }
        }
        const columnsString:string = columns.join(', ')
        const valuesString:string = values.join(', ');
        const result  = super.parseSqlQuery(this.databaseName, this.sqlStr);
        if (result == null) {
            return new Error("database or table cant be null")
        }
        return `INSERT INTO \"${result.databaseName}\".\"${result.tableName}\" (${columnsString}) VALUES (${valuesString})`;
    }

    createUpdateSql(column: any, data: any, originalData): string | Error {
        const conditions: any[] = [];
        const content: any[] = [];
        for (const item: any of column) {
            const val = data.value[item.column_name]
            const originalVal = originalData[item.column_name]
            if (val == null) {
                content.push(`\"${item.column_name}\" = NULL`)
                conditions.push(`\"${item.column_name}\" is NULL`)
            } else if (item.column_type == ColumnType.NUMBER) {
                content.push(`\"${item.column_name}\" = ${val}`)
                conditions.push(`\"${item.column_name}\" = ${originalVal}`)
            } else if (item.column_type == ColumnType.TEXT) {
                content.push(`\"${item.column_name}\" = \'${val}\'`)
                conditions.push(`\"${item.column_name}\" = \'${originalVal}\'`)
            }
        }
        const valuesString: string = content.join(', ')
        const conditionsString: string = conditions.join(' AND ')
        const result  = super.parseSqlQuery(this.databaseName, this.sqlStr);
        if (result == null) {
            return new Error("database or table cant be null")
        }
        return `UPDATE \"${result.databaseName}\".\"${result.tableName}\" SET ${valuesString} WHERE ${conditionsString}`;
    }

    createDeleteSql(column: any, row: any): string | Error {
        const conditions:any[] = [];
        for (const item: any of column) {
            const val = row[item.column_name]
            if (val == null) {
                conditions.push(`\"${item.column_name}\" is NULL`)
            } else if (item.column_type == ColumnType.NUMBER) {
                conditions.push(`\"${item.column_name}\" = ${val}`)
            } else if (item.column_type == ColumnType.TEXT) {
                conditions.push(`\"${item.column_name}\" = \'${val}\'`)
            }
        }
        const conditionsString:string = conditions.join(' AND ');
        const result  = super.parseSqlQuery(this.database, this.sqlStr);
        if (result == null) {
            return new Error("database or table cant be null")
        }
        return `DELETE FROM \"${result.databaseName}\".\"${result.tableName}\" WHERE ${conditionsString}`;
    }
}

export {Sqlite}