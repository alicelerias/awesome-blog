import React from "react";
import { Routes, Route, useNavigate } from "react-router-dom";
import { UsersComponent } from "./Users";
import { Profile } from "./Profile";
import { FeedComponent } from "./Feed";
import { CreatePost } from "./CreatePost";
import { Layout } from "./Layout";
import { UpdatePost } from "./UpdatePost";
import { AllPostsComponent } from "./AllPosts";
import { PostsByUserComponent } from "./PostsByUser";
import { UserDetail } from "./UserDetail";
import { FavoritesPosts } from "./FavoritesPosts";
import { useQuery } from "react-query";
import { getCurrentUser } from "../api/queries";
import { CurrentUserContext } from "../context/CurrentUserContext";
import { PostDetailBox } from "./PostDetailBox";

export const Main: React.FC = () => {
  const navigate = useNavigate();
  const { data } = useQuery("getCurrentUser", getCurrentUser);
  return (
    <CurrentUserContext.Provider value={data}>
      <div
        data-testid={"main-component-test-id"}
        className="flex flex-row sm:mx-three"
      >
        <Routes>
          <Route
            path="/"
            element={
              <Layout title="posts">
                <FeedComponent />
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
            path="/posts"
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
                <PostDetailBox navigate={navigate} />
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
                <CreatePost navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />
        </Routes>
      </div>
    </CurrentUserContext.Provider>
  );
};
