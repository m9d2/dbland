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