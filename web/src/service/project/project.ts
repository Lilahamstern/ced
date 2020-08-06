import API from "../index";

export function FetchProjects() {
  API.request({ url: "/space/projects", method: "GET" })
    .then((res) => {
      console.log(res);
      console.log(res.data);
    })
    .catch((error) => {
      console.log(error);
      console.log(error);
    });
}
