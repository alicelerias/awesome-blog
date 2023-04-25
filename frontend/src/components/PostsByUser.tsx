import { useQuery } from "react-query";
import { getPostsByUser } from "../api/queries";

import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";

export const PostsByUserComponent: React.FC<{}> = () => {
  const { isLoading, data } = useQuery("getPostsByUser", getPostsByUser);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      url={"posts"}
      children={<MenuPosts />}
    />
  );
};
