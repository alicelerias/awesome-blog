import { useQuery } from "react-query";
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
  const { isLoading, data } = useQuery("getAllPosts", getAllPosts);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      navigate={navigate}
      children={<MenuPosts />}
    />
  );
};
