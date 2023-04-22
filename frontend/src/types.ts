export type User = {
  id: string;
  username: string;
  bio: string;
  avatar: string;
  email: string;
  is_following: boolean;
};

export type Users = ResultList<"users", User>;

export type Post = {
  id: string;
  title: string;
  content: string;
  img: string;
  author: User;
  author_id: string;
  created_at: string;
};

export type Credential = {
  username: string;
  password: string;
};

export type ResultList<K extends string, T> = {
  [P in K]: T[];
};

export type PostCreate = Pick<Post, "title" | "content" | "img">;

export type PostUpdate = Partial<PostCreate>;

export type Posts = ResultList<"feed", Post>;

export type HealthCheck = {
  status: string;
};

export type Favorite = {
  postId: string;
  userId: string;
};

export type Favorites = ResultList<"feed", Favorite>;

export type Comment = {
  postId: string;
  authorId: string;
  author: User;
  content: string;
  created_at: string;
};

export type Comments = ResultList<"comments", Comment>;