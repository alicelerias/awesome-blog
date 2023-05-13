import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";
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

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      navigate={navigate}
      children={<MenuPosts />}
      hasNextPage={hasNextPage}
      fetchNextPage={fetchNextPage}
    />
  );
};
