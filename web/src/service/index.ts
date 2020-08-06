import Axios, { AxiosInstance } from "axios";
import { API_URL } from "../utils/config";

const instance: AxiosInstance = Axios.create({
  baseURL: API_URL,
  headers: {
    common: {
      "Cache-Control": "no-cache, no-store, must-revalidate",
      Pragma: "no-cache",
      "Content-Type": "application/json",
      Accept: "application/json",
    },
  },
  withCredentials: false,
});

export default instance;
