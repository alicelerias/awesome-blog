import axios from "axios";
import configs from "../configs/configs";

const HTTP_FORBIDDEN_STATUS = 403;

axios.interceptors.request.use((request) => {
  request.withCredentials = true;
  return request;
});

axios.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    if (error.response?.status === HTTP_FORBIDDEN_STATUS) {
      window.location.assign(configs.HOST + "/login");
    }
    return Promise.reject(error);
  }
);

export default axios;
