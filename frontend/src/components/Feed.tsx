import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";
import { Post } from "../types";
import { useInfiniteQuery } from "react-query";
import { getFeed } from "../api/queries";

type props = {
  navigate: NavigateFunction;
};

export const FeedComponent: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { data, isLoading, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "feed",
    getFeed,
    {
      getNextPageParam: (data) => data.next_link,
    }
  );

  const posts =
    data?.pages.reduce(
      (previous, current) => [...previous, ...current.content],
      [] as Post[]
    ) || [];
  return (
    <BoxPosts
      isLoading={isLoading}
      data={posts}
      navigate={navigate}
      children={<MenuPosts />}
      button={<button onClick={() => fetchNextPage()}>More</button>}
    />
  );
};
