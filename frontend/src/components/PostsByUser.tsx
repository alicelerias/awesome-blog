import { useInfiniteQuery } from "react-query";
import { getPostsByUser } from "../api/queries";

import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  navigate: NavigateFunction;
};
export const PostsByUserComponent: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { isLoading, data, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "getPostsByUser",
    getPostsByUser,
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
      fetchNextPage={fetchNextPage}
      hasNextPage={hasNextPage}
    />
  );
};
