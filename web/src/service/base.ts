class Base {

    database!: string;

    constructor(database: string) {
        this.database = database
    }

    removeLastAnd(sql: string): string {
        return sql.replace(/AND\s*$/, '');
    }

    removeLastComma(sql: string): string {
        return sql.replace(/,\s*$/, '');
    }

    stringifyValue(value: any): string {
        if (value === null || value === undefined) {
            return 'NULL'
        }
        if (typeof value === 'string') {
            return `'${value}'`;
        } else if (typeof value === 'number') {
            return `${value}`;
        } else if (typeof value === 'boolean') {
            return value ? 'true' : 'false';
        } else {
            return ''; // 其他类型暂不处理
        }
    }
}

export {Base}