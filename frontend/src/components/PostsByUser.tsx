import { useQuery } from "react-query";
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
  const { isLoading, data } = useQuery("getPostsByUser", getPostsByUser);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data?.content || []}
      navigate={navigate}
      children={<MenuPosts />}
    />
  );
};
