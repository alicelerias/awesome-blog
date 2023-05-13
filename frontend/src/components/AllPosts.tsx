import { useInfiniteQuery } from "react-query";
import { getAllPosts } from "../api/queries";

import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  navigate: NavigateFunction;
};

export const AllPostsComponent: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { isLoading, data, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "getAllPosts",
    getAllPosts,
    {
      getNextPageParam: (data) => data.next_link,
    }
  );

  return (
    <BoxPosts
      isLoading={isLoading}
      fetchNextPage={fetchNextPage}
      hasNextPage={hasNextPage}
      data={data}
      navigate={navigate}
      children={<MenuPosts />}
    />
  );
};
