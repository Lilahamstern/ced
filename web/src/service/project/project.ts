import API from "../index";

export function FetchProjects() {
  console.log(process.env.API_URL);
  console.log(process.env["API_URL"]);
  API.get("/space/projects")
    .then((res) => {
      console.log(res.data);
    })
    .catch((error) => {
      console.log(error);
    });
}
