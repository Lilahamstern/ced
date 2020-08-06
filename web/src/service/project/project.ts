import API from "../index";
import { IResponse } from "../types";

export async function FetchProjects(): Promise<IResponse> {
  return new Promise((resolve, reject) => {
    API.request({ url: "space/projects", method: "GET" })
      .then((res) => {
        resolve(res.data);
      })
      .catch((err) => {
        if (err.response) {
          resolve(err.response);
        }
      });
  });
}
