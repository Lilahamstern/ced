import Axios, { AxiosInstance } from "axios";

const instance: AxiosInstance = Axios.create({
  baseURL: "http://localhost",
  headers: { "Content-Type": "application/json" },
  timeout: 1000,
});

export default instance;
