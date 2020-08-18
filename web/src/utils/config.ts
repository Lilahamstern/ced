export const API_URL: string =
  process.env.NODE_ENV === "development"
    ? "http://host.docker.internal:5000/"
    : "";
