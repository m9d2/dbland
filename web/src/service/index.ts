import {DbTypeEnum} from "@/common/enums";
import {Mysql} from './mysql';
import {Sqlite} from './sqlite';

export function CreateDatabase(config: any, database: string, sql: string): Database {
    let instance: Database;
    if (config.type == DbTypeEnum.SQLite) {
        instance = new Sqlite(database, sql)
    }
    if (config.type == DbTypeEnum.MySQL) {
        instance = new Mysql(database, sql)
    }
    return instance
}