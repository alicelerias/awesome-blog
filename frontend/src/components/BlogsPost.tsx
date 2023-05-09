import { useInfiniteQuery } from "react-query";
import { getBlogsPost } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  id: string | null;
  navigate: NavigateFunction;
};
export const BlogsPost: React.FC<PropsWithChildren<props>> = ({
  id,
  navigate,
}) => {
  const { isLoading, data, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "getBlogsPost",
    getBlogsPost(id),
    {
      getNextPageParam: (data) => data.next_link,
    }
  );
  return (
    <BoxPosts
      key={`blogs-post-${id}`}
      isLoading={isLoading}
      data={data}
      navigate={navigate}
      fetchNextPage={fetchNextPage}
      hasNextPage={hasNextPage}
    />
  );
};
