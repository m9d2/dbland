export interface QueryReq {
    cid?: number,
    sql?: string,
    db?: string,
    table?: string,
    page?: number,
    size?: number, 
}

export interface Table {
    name: string,
    rows: string,
    collation: string,
    create_time: string,
    update_time: string
}

export interface Query {
    columns: string,
	tableName: string,
	rows: any,
	total: number,    
}