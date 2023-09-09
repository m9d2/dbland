import axios from 'axios';

const request = axios.create({
    baseURL: 'http://127.0.0.1:8080',
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