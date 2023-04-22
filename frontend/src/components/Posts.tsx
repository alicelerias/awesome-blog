import { useQuery } from "react-query";
import { getPosts } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";

export const PostsComponent: React.FC<{}> = () => {
  const { isLoading, data } = useQuery("getPosts", getPosts);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      url={"posts"}
      children={<MenuPosts />}
    />
  );
};
