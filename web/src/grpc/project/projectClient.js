import { getProjectsParams } from '../gen/project_pb';
import { ProjectClient } from '../gen/project_grpc_web_pb';

export default class ProjectGRPC {
  static _connection() {
    if (this._client == null) {
      try {
        this._client = new ProjectClient('https://localhost:32768', null, null);
      } catch (error) {
        console.log(error);
      }
    }
  }

  static getProjects() {
    if (this._client == null) {
      this._connection();
    }

    let req = new getProjectsParams();

    return new Promise((resolve, reject) => {
      this._client.getProjects(req, {}, (err, res) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(res.toObject().projectsList);
      });
    });
  }

  static searchProjects(search) {
    if (this._client == null) {
      this._connection();
    }

    let req = new getProjectsParams();
    req.setSearch(search);

    return new Promise((resolve, reject) => {
      this._client.getProjects(req, {}, (err, res) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(res.toObject().projectsList);
      });
    });
  }
}
