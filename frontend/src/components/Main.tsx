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
import { useForm } from "react-hook-form";

export const Main: React.FC = () => {
  const navigate = useNavigate();

  const {
    handleSubmit,
    register,
    reset,
    formState: { errors },
  } = useForm();

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
                <FeedComponent navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/favorites"
            element={
              <Layout title="favorites">
                <FavoritesPosts navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/users/detail"
            element={
              <Layout title="user detail">
                <UserDetail navigate={navigate} />
              </Layout>
            }
          />

          <Route
            path="/posts/you"
            element={
              <Layout title="your posts">
                <PostsByUserComponent navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/posts"
            element={
              <Layout title="all posts">
                <AllPostsComponent navigate={navigate} />
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
                <PostDetailBox
                  navigate={navigate}
                  handleSubmit={handleSubmit}
                  register={register}
                  errors={errors}
                />
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
                <CreatePost
                  handleSubmit={handleSubmit}
                  register={register}
                  reset={reset}
                  errors={errors}
                  navigate={navigate}
                />
                <UsersComponent />
              </Layout>
            }
          />
        </Routes>
      </div>
    </CurrentUserContext.Provider>
  );
};
