export enum TreeLevelEnum {
    ROOT = 0,
    CONFIG = 1,
    DATABASE = 2,
    TABLE = 3
}

export enum DbTypeEnum {
    MySQL = 'mysql',
    SQLite = 'sqlite',
    ORACLE = 'oracle',
    PostgreSQL = 'postgresql',
    MariaDB = 'mariadb'
}

export enum ActionTypeEnum {
    MODIFY = "modify",
    INSERT = "insert",
}

export enum ColumnType {
    NUMBER = "number",
    TEXT = "text"
}

export enum ColumnKey {
    PRIMARY = "primary",
}

export enum Direction {
    ASC = "asc",
    DESC = "desc"
}