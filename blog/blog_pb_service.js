// package: blog
// file: blog/blog.proto

var blog_blog_pb = require("../blog/blog_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Blog = (function () {
  function Blog() {}
  Blog.serviceName = "blog.Blog";
  return Blog;
}());

Blog.ListAllPosts = {
  methodName: "ListAllPosts",
  service: Blog,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: blog_blog_pb.PostList
};

Blog.GetPost = {
  methodName: "GetPost",
  service: Blog,
  requestStream: false,
  responseStream: false,
  requestType: blog_blog_pb.PostQuery,
  responseType: blog_blog_pb.Post
};

Blog.CreatePost = {
  methodName: "CreatePost",
  service: Blog,
  requestStream: false,
  responseStream: false,
  requestType: blog_blog_pb.Post,
  responseType: blog_blog_pb.Post
};

Blog.DeletePost = {
  methodName: "DeletePost",
  service: Blog,
  requestStream: false,
  responseStream: false,
  requestType: blog_blog_pb.PostQuery,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Blog = Blog;

function BlogClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BlogClient.prototype.listAllPosts = function listAllPosts(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Blog.ListAllPosts, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

BlogClient.prototype.getPost = function getPost(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Blog.GetPost, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

BlogClient.prototype.createPost = function createPost(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Blog.CreatePost, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

BlogClient.prototype.deletePost = function deletePost(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Blog.DeletePost, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

exports.BlogClient = BlogClient;

