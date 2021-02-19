// package: clusters
// file: clusters.proto

var clusters_pb = require("./clusters_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Clusters = (function () {
  function Clusters() {}
  Clusters.serviceName = "clusters.Clusters";
  return Clusters;
}());

Clusters.GetClusters = {
  methodName: "GetClusters",
  service: Clusters,
  requestStream: false,
  responseStream: false,
  requestType: clusters_pb.GetClustersRequest,
  responseType: clusters_pb.GetClustersResponse
};

Clusters.GetNamespaces = {
  methodName: "GetNamespaces",
  service: Clusters,
  requestStream: false,
  responseStream: false,
  requestType: clusters_pb.GetNamespacesRequest,
  responseType: clusters_pb.GetNamespacesResponse
};

Clusters.GetResources = {
  methodName: "GetResources",
  service: Clusters,
  requestStream: false,
  responseStream: false,
  requestType: clusters_pb.GetResourcesRequest,
  responseType: clusters_pb.GetResourcesResponse
};

exports.Clusters = Clusters;

function ClustersClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ClustersClient.prototype.getClusters = function getClusters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Clusters.GetClusters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ClustersClient.prototype.getNamespaces = function getNamespaces(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Clusters.GetNamespaces, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ClustersClient.prototype.getResources = function getResources(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Clusters.GetResources, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ClustersClient = ClustersClient;
