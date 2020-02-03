// package: blog
// file: blog/blog.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class PostList extends jspb.Message {
  clearPostsList(): void;
  getPostsList(): Array<Post>;
  setPostsList(value: Array<Post>): void;
  addPosts(value?: Post, index?: number): Post;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostList.AsObject;
  static toObject(includeInstance: boolean, msg: PostList): PostList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PostList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostList;
  static deserializeBinaryFromReader(message: PostList, reader: jspb.BinaryReader): PostList;
}

export namespace PostList {
  export type AsObject = {
    postsList: Array<Post.AsObject>,
  }
}

export class Post extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Post.AsObject;
  static toObject(includeInstance: boolean, msg: Post): Post.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Post, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Post;
  static deserializeBinaryFromReader(message: Post, reader: jspb.BinaryReader): Post;
}

export namespace Post {
  export type AsObject = {
    id: string,
    title: string,
    body: string,
  }
}

export class PostQuery extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostQuery.AsObject;
  static toObject(includeInstance: boolean, msg: PostQuery): PostQuery.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PostQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostQuery;
  static deserializeBinaryFromReader(message: PostQuery, reader: jspb.BinaryReader): PostQuery;
}

export namespace PostQuery {
  export type AsObject = {
    id: string,
  }
}

