import { useInfiniteQuery } from "react-query";
import { getFavorites } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  navigate: NavigateFunction;
};
export const FavoritesPosts: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { isLoading, data, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "getFavorites",
    () => getFavorites(),
    {
      getNextPageParam: (data) => data.next_link,
    }
  );

  return (
    <div data-testid="favorites-component-test-id">
      <BoxPosts
        isLoading={isLoading}
        fetchNextPage={fetchNextPage}
        hasNextPage={hasNextPage}
        data={data}
        navigate={navigate}
        children={
          <div className="bg-box-color p-one flex justify-center text-2xl">
            Your favorites posts
          </div>
        }
      />
    </div>
  );
};
