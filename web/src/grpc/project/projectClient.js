import {
  getProjectsParams,
  getProjectsByProjectIdParams
} from '../gen/project_pb';
import { ProjectClient } from '../gen/project_grpc_web_pb';

export default class ProjectGRPC {
  static _connection() {
    if (this._client == null) {
      try {
        this._client = new ProjectClient('https://localhost:32770', null, null);
      } catch (error) {
        console.log(error);
      }
    }
  }

  static getProjects(search = '') {
    this.check();
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

  static getProjectInfoByProjectId(id) {
    this.check();

    let req = new getProjectsByProjectIdParams();
    req.setProjectId(id);

    console.log('shit');
    return new Promise((resolve, reject) => {
      this._client.getProjectByProjectId(req, {}, (err, res) => {
        if (err != null) {
          reject(err);
          return;
        }
        console.log(res.toObject());
        resolve(res.toObject());
      });
    });
  }

  static check() {
    if (this._client == null) {
      this._connection();
    }
  }
}
