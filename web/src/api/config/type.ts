export interface DBConfig {
    name: string,
    type: string,
    host: string,
    port: string,
    username: string,
    password: string,

    // additional
    level: number,
    leaf: boolean,
}