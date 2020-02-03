// package: blog
// file: blog/blog.proto

import * as blog_blog_pb from "../blog/blog_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type BlogListAllPosts = {
  readonly methodName: string;
  readonly service: typeof Blog;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof blog_blog_pb.PostList;
};

type BlogGetPost = {
  readonly methodName: string;
  readonly service: typeof Blog;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_blog_pb.PostQuery;
  readonly responseType: typeof blog_blog_pb.Post;
};

type BlogCreatePost = {
  readonly methodName: string;
  readonly service: typeof Blog;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_blog_pb.Post;
  readonly responseType: typeof blog_blog_pb.Post;
};

type BlogDeletePost = {
  readonly methodName: string;
  readonly service: typeof Blog;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_blog_pb.PostQuery;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Blog {
  static readonly serviceName: string;
  static readonly ListAllPosts: BlogListAllPosts;
  static readonly GetPost: BlogGetPost;
  static readonly CreatePost: BlogCreatePost;
  static readonly DeletePost: BlogDeletePost;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class BlogClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listAllPosts(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<blog_blog_pb.PostList>;
  getPost(
    requestMessage: blog_blog_pb.PostQuery,
    metadata?: grpc.Metadata,
  ): Promise<blog_blog_pb.Post>;
  createPost(
    requestMessage: blog_blog_pb.Post,
    metadata?: grpc.Metadata,
  ): Promise<blog_blog_pb.Post>;
  deletePost(
    requestMessage: blog_blog_pb.PostQuery,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
}

