import axios from "axios";

const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_GATEWAY_URL,
  withCredentials: true,
});

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },

  async (error) => {
    if (error.response && error.response.status === 401) {
      try {
        await axiosInstance.get("/auth/refresh");

        return axiosInstance(error.config);
        // same request retried after refreshing token
        // same as: return axios.get('api/protected/resource', {withCredentials:true});
        // as error.config has get method, url and withCredentials:true
      } catch (refreshError) {
        window.location.href = import.meta.env.VITE_LOGIN_APP_URL;
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

export default axiosInstance;
