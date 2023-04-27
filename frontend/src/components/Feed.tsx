import { useQuery } from "react-query";
import { getFeed } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  navigate: NavigateFunction;
};

export const FeedComponent: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { isLoading, data } = useQuery("getPosts", getFeed);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      navigate={navigate}
      children={<MenuPosts />}
    />
  );
};
