import request from "@/plugins/axios";
import type { Chart } from "./type"
import type { AxiosPromise } from "axios";

enum API {
    CHART_URL = '/v1/chart',
}

export const getCharts = (): AxiosPromise<Chart[]> => {
    return request.get<Chart[]>(API.CHART_URL);
};

export const createChart = (data: Chart) => {
    return request.post(API.CHART_URL, data);
};

export const updateChart = (data: Chart) => {
    return request.put(API.CHART_URL, data);
};

export const deleteById = (id: number) => {
    return request.delete(API.CHART_URL + '/' + id);
};