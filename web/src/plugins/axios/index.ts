import axios from 'axios';
const baseUrl = import.meta.env.MODE == "release" ? window.location.origin : import.meta.env.VITE_SERVER
const request = axios.create({
    baseURL: baseUrl,
    timeout: 10000,
    responseType: "json",
    headers: {
        'Content-Type': 'application/json',
    }
});

request.interceptors.request.use(
    config => {
      config.headers['Authorization'] = 'Bearer token';
      return config;
    },
    error => {
      return Promise.reject(error);
    }
);

request.interceptors.response.use(
    response => {
        const { code, message } = response.data;
        if (code === 200) {
            return response.data;
        } else if (code == 400) {
            throw new Error(message)
        }
    },
    error => {
        return Promise.reject(error);
    } 
);

export default request;
