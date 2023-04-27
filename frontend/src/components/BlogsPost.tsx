import { useQuery } from "react-query";
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
  const { isLoading, data } = useQuery("getBlogsPost", () => getBlogsPost(id));
  return <BoxPosts isLoading={isLoading} data={data} navigate={navigate} />;
};
