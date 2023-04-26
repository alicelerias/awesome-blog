import { useQuery } from "react-query";
import { useSearchParams } from "react-router-dom";
import { getBlogsPost } from "../api/queries";
import { BoxPosts } from "./BoxPosts";
import { PropsWithChildren } from "react";

type props = {
  id: string | null;
};
export const BlogsPost: React.FC<PropsWithChildren<props>> = ({ id }) => {
  const { isLoading, data } = useQuery("getBlogsPost", () => getBlogsPost(id));
  return <BoxPosts isLoading={isLoading} data={data} url={"posts"} />;
};
