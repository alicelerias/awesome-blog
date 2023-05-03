import { useQuery } from "react-query";
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
  const { isLoading, data } = useQuery("getFavorites", () => getFavorites());

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data?.content || []}
      navigate={navigate}
      children={
        <div className="bg-box-color p-one flex justify-center text-2xl">
          Your favorites posts
        </div>
      }
    />
  );
};
