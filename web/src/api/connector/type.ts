export interface QueryReq {
    cid?: number,
    sql?: string,
    db?: string,
    table?: string,
    page?: number,
    size?: number, 
    sort?: Sort,
}

export interface Table {
    name: string,
    rows: string,
    collation: string,
    create_time: string,
    update_time: string
}

export interface Query {
[x: string]: any;
    columns: string,
	tableName: string,
	rows: any,
	total: number,   
    elapsed_time: number, 
    total_page: number, 
}

export type Sort = {
    prop: string,
    order: string
  }