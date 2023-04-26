import { useQuery } from "react-query";
import { getAllPosts } from "../api/queries";

import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";

export const AllPostsComponent: React.FC<{}> = () => {
  const { isLoading, data } = useQuery("getAllPosts", getAllPosts);

  return (
    <BoxPosts isLoading={isLoading} data={data} children={<MenuPosts />} />
  );
};
