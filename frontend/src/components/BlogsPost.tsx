import { useQuery } from "react-query";
import { useSearchParams } from "react-router-dom";
import { getBlogsPost } from "../api/queries";
import { BoxPosts } from "./BoxPosts";

export const BlogsPost: React.FC<{}> = () => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");
  const { isLoading, data } = useQuery("getBlogsPost", () => getBlogsPost(id));
  return <BoxPosts isLoading={isLoading} data={data} url={"posts"} />;
};
