import request from "@/axios";
import type { QueryReq, Table } from './type'
import type { AxiosPromise } from "axios";

enum API {
    DATABASE_URL = '/v1/databases',
    TABLE_URL = '/v1/tables',
    QUERY_URL = '/v1/query',
    PING_URL = '/v1/ping',
}

export const getDatabases = (data: QueryReq): AxiosPromise<string[]> => {
    return request.post<string[]>(API.DATABASE_URL, data);
};

export const getTables = (data: QueryReq): AxiosPromise<Table[]> => {
    return request.post<Table[]>(API.TABLE_URL, data);
};

export const query = (data: QueryReq): AxiosPromise<Query> => {
    return request.post<Query>(API.QUERY_URL, data);
};

export const ping = (data: any) => {
    return request.post(API.PING_URL, data);
};