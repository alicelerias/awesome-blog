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
import {
  FieldErrors,
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormReset,
  UseFormSetValue,
} from "react-hook-form";

type props = {
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  reset: UseFormReset<FieldValues>;
  errors?: FieldErrors<FieldValues>;
  setValue: UseFormSetValue<FieldValues>;
};
export const Main: React.FC<props> = ({
  handleSubmit,
  register,
  reset,
  setValue,
  errors,
}) => {
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
              <Layout navigate={navigate} title="posts">
                <FeedComponent navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/favorites"
            element={
              <Layout navigate={navigate} title="favorites">
                <FavoritesPosts navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/users/detail"
            element={
              <Layout navigate={navigate} title="user detail">
                <UserDetail navigate={navigate} />
              </Layout>
            }
          />

          <Route
            path="/posts/you"
            element={
              <Layout navigate={navigate} title="your posts">
                <PostsByUserComponent navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/posts"
            element={
              <Layout navigate={navigate} title="all posts">
                <AllPostsComponent navigate={navigate} />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/profile"
            element={
              <Layout title="user detail" navigate={navigate}>
                <Profile
                  setValue={setValue}
                  navigate={navigate}
                  handleSubmit={handleSubmit}
                  register={register}
                  reset={reset}
                />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/posts/detail"
            element={
              <Layout navigate={navigate} title="post detail">
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
              <Layout navigate={navigate} title="update post">
                <UpdatePost
                  handleSubmit={handleSubmit}
                  register={register}
                  reset={reset}
                  errors={errors}
                  setValue={setValue}
                  navigate={navigate}
                />
                <UsersComponent />
              </Layout>
            }
          />

          <Route
            path="/posts/new"
            element={
              <Layout navigate={navigate} title="update post">
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
