import request from "@/plugins/axios";

enum API {
    README_URL = '/v1/readme',
}

export const getReadme = (data: any) => {
    return request.get(API.README_URL, data)
}