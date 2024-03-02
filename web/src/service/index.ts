import {DbTypeEnum} from "@/common/enums";
import {Mysql} from './mysql';
import {Sqlite} from './sqlite';
import type { Database } from "./database";

export function CreateDatabase(type: string, database: string): Database {
    let instance: Database | undefined;
    if (type === DbTypeEnum.SQLite) {
        instance = new Sqlite(database);
    }
    if (type === DbTypeEnum.MySQL) {
        instance = new Mysql(database);
    }
    return instance as Database;
}