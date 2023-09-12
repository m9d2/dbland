import { ref, watch } from '@vue/composition-api';

export function useLocalStorage<T>(key: string, initialValue: T) {
  const storedValue = localStorage.getItem(key);
  const value = ref<T>(storedValue ? JSON.parse(storedValue) : initialValue);

  watch(value, (newValue) => {
    localStorage.setItem(key, JSON.stringify(newValue));
  });

  return value;
}

export function createQuerySql(database: string, table: string): string {
    return `SELECT * FROM \`${database}\`.\`${table}\``
}

export function generateRandomString(length: number): string {
    const characters = 'abcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        result += characters.charAt(randomIndex);
    }
    return result;
}

export function findKeyByValue<T extends string>(enumObject: Record<string, T>, value: T): keyof typeof enumObject | undefined {
    return Object.keys(enumObject).find((key) => enumObject[key] === value);
}

export function createDeleteSql(database: string, sqlStr: string, row: any): string | Error {
    const keys = Reflect.ownKeys(row);
    const conditions = [];
    let table
    for (const key of keys) {
        if (typeof row[key] === 'string') {
            conditions.push(`${key}='${row[key]}'`)
        } else{
            conditions.push(`${key}=${row[key]}`)
        }
    }

    const conditionsString = conditions.join(' AND ');
    const result  = parseSqlQuery(sqlStr);
    if (result != null) {
        if(result.databaseName == null && database == null) {
            throw new Error('Database name cannot be null')
        }
        if (result.tableName == null) {
            throw new Error('Table name cannot be null')
        }
        if (result.databaseName != null) {
            database = result.databaseName
        }
        table = result.tableName

    }
    return `DELETE FROM \`${database}\`.\`${table}\` WHERE ${conditionsString}`;
}

export function createInsertSql(database: string, sqlStr: string, row: any, column: any): string | Error {
    const keys = Reflect.ownKeys(row);
    const values = [];
    const columns = [];
    let table
    for (const key of keys) {
        if (typeof row[key] === 'string') {
            values.push(`${row[key]}'`)
        } else{
            values.push(`${row[key]}'`)
        }
    }

    const valuesString = values.join(', ');
    const columnsString = values.join(', ')
    const result  = parseSqlQuery(sqlStr);
    if (result != null) {
        if(result.databaseName == null && database == null) {
            database = ''
        }
        if (result.tableName == null) {
            table = ''
        }
        if (result.databaseName != null) {
            database = result.databaseName
        }
        table = result.tableName

    }
    return `INSERT INTO \`${database}\`.\`${table}\` (${columnsString}) VALUES (${valuesString})`;
}

function parseSqlQuery(sqlQuery: string): { databaseName: string, tableName: string } | null {
    const regex = /FROM\s+`?([^.\s`]+)`?\.`?([^.\s`]+)`?/i;
    const matches = regex.exec(sqlQuery);
    if (matches && matches.length === 3) {
        const databaseName = matches[1];
        const tableName = matches[2];
        return { databaseName, tableName };
    }
    return null;
}
