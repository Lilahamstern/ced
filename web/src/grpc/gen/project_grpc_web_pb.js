/* eslint-disable */
/**
 * @fileoverview gRPC-Web generated client stub for project
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!

const grpc = {};
grpc.web = require('grpc-web');

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
const proto = {};
proto.project = require('./project_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.project.ProjectClient = function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.project.ProjectPromiseClient = function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.project.addProjectParams,
 *   !proto.project.addProjectResponse>}
 */
const methodDescriptor_Project_CreateProject = new grpc.web.MethodDescriptor(
  '/project.Project/CreateProject',
  grpc.web.MethodType.UNARY,
  proto.project.addProjectParams,
  proto.project.addProjectResponse,
  /**
   * @param {!proto.project.addProjectParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.project.addProjectResponse.deserializeBinary
);

/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.project.addProjectParams,
 *   !proto.project.addProjectResponse>}
 */
const methodInfo_Project_CreateProject = new grpc.web.AbstractClientBase.MethodInfo(
  proto.project.addProjectResponse,
  /**
   * @param {!proto.project.addProjectParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.project.addProjectResponse.deserializeBinary
);

/**
 * @param {!proto.project.addProjectParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.project.addProjectResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.project.addProjectResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.project.ProjectClient.prototype.createProject = function(
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + '/project.Project/CreateProject',
    request,
    metadata || {},
    methodDescriptor_Project_CreateProject,
    callback
  );
};

/**
 * @param {!proto.project.addProjectParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.project.addProjectResponse>}
 *     A native promise that resolves to the response
 */
proto.project.ProjectPromiseClient.prototype.createProject = function(
  request,
  metadata
) {
  return this.client_.unaryCall(
    this.hostname_ + '/project.Project/CreateProject',
    request,
    metadata || {},
    methodDescriptor_Project_CreateProject
  );
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.project.getProjectsParams,
 *   !proto.project.projectResponse>}
 */
const methodDescriptor_Project_GetProjects = new grpc.web.MethodDescriptor(
  '/project.Project/GetProjects',
  grpc.web.MethodType.UNARY,
  proto.project.getProjectsParams,
  proto.project.projectResponse,
  /**
   * @param {!proto.project.getProjectsParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.project.projectResponse.deserializeBinary
);

/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.project.getProjectsParams,
 *   !proto.project.projectResponse>}
 */
const methodInfo_Project_GetProjects = new grpc.web.AbstractClientBase.MethodInfo(
  proto.project.projectResponse,
  /**
   * @param {!proto.project.getProjectsParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.project.projectResponse.deserializeBinary
);

/**
 * @param {!proto.project.getProjectsParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.project.projectResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.project.projectResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.project.ProjectClient.prototype.getProjects = function(
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + '/project.Project/GetProjects',
    request,
    metadata || {},
    methodDescriptor_Project_GetProjects,
    callback
  );
};

/**
 * @param {!proto.project.getProjectsParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.project.projectResponse>}
 *     A native promise that resolves to the response
 */
proto.project.ProjectPromiseClient.prototype.getProjects = function(
  request,
  metadata
) {
  return this.client_.unaryCall(
    this.hostname_ + '/project.Project/GetProjects',
    request,
    metadata || {},
    methodDescriptor_Project_GetProjects
  );
};

module.exports = proto.project;
