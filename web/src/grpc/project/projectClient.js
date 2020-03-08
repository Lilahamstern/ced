import { getProjectsParams } from '../gen/project_pb';
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

  static getProjects() {
    if (this._client == null) {
      this._connection();
    }

    let req = new getProjectsParams();

    return new Promise((resolve, reject) => {
      this._client.getProjects(req, {}, (err, res) => {
        if (err != null) {
          reject(err);
        }
        console.log(res.toObject().projectsList);
        resolve(res.toObject().projectsList);
      });
    });
  }
}
