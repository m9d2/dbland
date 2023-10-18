import request from "@/plugins/axios";
import type { QueryReq, Table, Query } from './type'
import type { AxiosPromise } from "axios";

export enum API {
    DATABASE_URL = '/v1/databases',
    TABLE_URL = '/v1/tables',
    QUERY_URL = '/v1/query',
    PING_URL = '/v1/ping',
    EXECUTE_URL = '/v1/execute',
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

export const execute = (data: any) => {
    return request.post(API.EXECUTE_URL, data)
}