import { NavigateFunction } from "react-router-dom";
import { AiOutlineComment } from "react-icons/ai";
import { Post, Posts } from "../types";

import { ToggleFavoriteButton } from "./ToggleFavoriteButton";
import React from "react";
import { Observer } from "./IntersectionObserver";
import {
  FetchNextPageOptions,
  InfiniteData,
  InfiniteQueryObserverResult,
} from "react-query";

type props = {
  isLoading: boolean;
  data?: InfiniteData<Posts>;
  navigate: NavigateFunction;
  hasNextPage?: boolean;
  fetchNextPage: (
    options?: FetchNextPageOptions | undefined
  ) => Promise<InfiniteQueryObserverResult<Posts, unknown>>;
};

export const BoxPosts: React.FC<React.PropsWithChildren<props>> = ({
  isLoading,
  data,
  children,
  navigate,
  fetchNextPage,
  hasNextPage,
}) => {
  const posts =
    data?.pages.reduce(
      (previous, current) => [...previous, ...current.content],
      [] as Post[]
    ) || [];

  return (
    <div className="flex flex-col gap-one sm:w-3/5">
      {children}
      <div
        data-testid={"posts-component-test-id"}
        className={`m-0 p-0 w-full sm:w-full`}
      >
        {isLoading
          ? "is Loading"
          : posts.map((post) => (
              <div key={post.id} className={`flex gap-one pb-one`}>
                <div className="w-auto h-auto">
                  <img
                    className=" w-20 aspect-square"
                    src={
                      post.author.avatar ||
                      "https://ionicframework.com/docs/img/demos/avatar.svg"
                    }
                    alt=""
                  />
                </div>
                <div className="p-two bg-box-color w-full">
                  <img className="w-20" src={post?.img} alt="" />
                  <h1
                    onClick={() => {
                      navigate(`/posts/detail?id=${post.id}`);
                    }}
                    className="text-xl text-bold cursor-pointer"
                  >
                    {post?.title}
                  </h1>
                  <p className="text-sm italic">"{post?.content}"</p>
                  <div className="flex flex-row gap-two justify-end p-one">
                    <span> {post.comments_count} </span>
                    <AiOutlineComment
                      onClick={() => {
                        navigate(`/$posts/detail?id=${post.id}`);
                      }}
                      className="h-6 w-6 text-blue cursor-pointer"
                    />
                    <ToggleFavoriteButton
                      postId={post.id}
                      isFavorite={post.is_favorite}
                      favoritesCount={post.favorites_count}
                    />
                  </div>
                </div>
                <div></div>
              </div>
            ))}
      </div>
      <Observer fetchNextPage={fetchNextPage} />
    </div>
  );
};
