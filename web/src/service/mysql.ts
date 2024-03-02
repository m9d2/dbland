import {ColumnKey} from "@/common/enums";
import {Base} from "./base";
import type {Database} from "./database";
import type { Sort } from "./model/sort";

class Mysql extends Base implements Database {

    createQueryPageSql(table: string, page: number, size: number, where: string, sort: Sort): string | undefined {
        let orderStr = '';
        if (sort) {
            orderStr = 'ODERY BY ' + sort.field + ' ' + sort.direction;
        }
        return `SELECT * FROM \`${this.database}\`.\`${table}\` ${where} ${orderStr} LIMIT ${(page-1) * size},${(size)}`;
    }

    createQuerySql(table: string): string | undefined {
        return `SELECT * FROM \`${this.database}\`.\`${table}\``;
    }

    createInsertSql(table: string, columns: Column[], row: any): string | undefined {
        let columnsStr = ''
        let valuesStr = ''
        for (let column of columns) {
            let value = row[column.field];
            columnsStr += `\`${column.field}\`, `
            valuesStr += `${this.stringifyValue(value)}, `
        }
        return `INSERT INTO \`${this.database}\`.\`${table}\` (${this.removeLastComma(columnsStr)}) VALUES (${this.removeLastComma(valuesStr)})`;
    }

    createUpdateSql(table: string, columns: Column[], row: any, originalRow: any): string | undefined {
        let valuesStr = ''
        for (let column of columns) {
            let value = row[column.field];
            valuesStr += `\`${column.field}\` = ${this.stringifyValue(value)}, `
        }
        return `UPDATE \`${this.database}\`.\`${table}\` SET ${this.removeLastComma(valuesStr)} WHERE ${this.removeLastAnd(this.prepareWhereSql(columns, originalRow))}`;
    }

    createDeleteSql(table: string, columns: Column[], row: any): string | undefined {
        return `DELETE FROM "${this.database}"."${table}" WHERE ${this.prepareWhereSql(columns, row)}`;
    }

    prepareWhereSql(columns: Column[], row: any) {
        let sql: string = '';
        for (let column of columns) {
            if (column.key === ColumnKey.PRIMARY) {
                return `\`${column.field}\` = ${row[column.field]}`
            }
        }

        for (let column of columns) {
            let value = row[column.field];
            if (value === null) {
                sql += `\`${column.field}\` IS ${this.stringifyValue(value)} AND `
            } else {
                sql += `\`${column.field}\` = ${this.stringifyValue(value)} AND `
            }
            
        }
        return this.removeLastAnd(sql);
    }
}

export {Mysql}