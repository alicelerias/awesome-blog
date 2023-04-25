import { useQuery } from "react-query";
import { getFeed } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";

export const FeedComponent: React.FC<{}> = () => {
  const { isLoading, data } = useQuery("getPosts", getFeed);

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      url={"posts"}
      children={<MenuPosts />}
    />
  );
};
