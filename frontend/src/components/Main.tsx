import React from "react";
import { Routes, Route } from "react-router-dom";
import { UsersComponent } from "./Users";
import { Profile } from "./Profile";
import { PostsComponent } from "./Posts";
import { PostDetail } from "./PostDetail";
import { CreatePost } from "./CreatePost";
import { Layout } from "./Layout";
import { UpdatePost } from "./UpdatePost";
import { LoginPage } from "./LoginPage";
import { AllPostsComponent } from "./AllPosts";
import { PostsByUserComponent } from "./PostsByUser";
import { UserPostDetail } from "./UserPostDetail";
import { UserDetail } from "./UserDetail";
import { FavoritesPosts } from "./FavoritesPosts";

export const Main: React.FC = () => {
  return (
    <div
      data-testid={"main-component-test-id"}
      className="flex flex-row sm:mx-three"
    >
      <Routes>
        <Route path="/login" element={<LoginPage />} />

        <Route
          path="/posts"
          element={
            <Layout title="posts">
              <PostsComponent />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/favorites"
          element={
            <Layout title="favorites">
              <FavoritesPosts />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/users/detail"
          element={
            <Layout title="user detail">
              <UserDetail />
            </Layout>
          }
        />

        <Route
          path="/posts/you"
          element={
            <Layout title="your posts">
              <PostsByUserComponent />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/"
          element={
            <Layout title="all posts">
              <AllPostsComponent />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/profile"
          element={
            <Layout title="user detail">
              <Profile />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/posts/detail"
          element={
            <Layout title="post detail">
              <PostDetail />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/posts/you/detail"
          element={
            <Layout title="post detail">
              <UserPostDetail />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/posts/update"
          element={
            <Layout title="update post">
              <UpdatePost />
              <UsersComponent />
            </Layout>
          }
        />

        <Route
          path="/posts/new"
          element={
            <Layout title="update post">
              <CreatePost />
              <UsersComponent />
            </Layout>
          }
        />
      </Routes>
    </div>
  );
};
