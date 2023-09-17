import request from "@/plugins/axios";
import type { DBConfig } from "./type"
import type { AxiosPromise } from "axios";

enum API {
    CONFIG_URL = '/v1/config',
}

export const getConfigs = (): AxiosPromise<DBConfig[]> => {
    return request.get<DBConfig[]>(API.CONFIG_URL);
};

export const createConfig = (data: DBConfig) => {
    return request.post(API.CONFIG_URL, data);
};

export const updateConfig = (data: DBConfig) => {
    return request.put(API.CONFIG_URL, data);
};

export const deleteById = (id: number) => {
    return request.delete(API.CONFIG_URL + '/' + id);
};

