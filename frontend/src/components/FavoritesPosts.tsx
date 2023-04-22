import { useQuery } from "react-query";
import { getFavorites } from "../api/queries";
import { BoxPosts } from "./BoxPosts";

export const FavoritesPosts: React.FC<{}> = () => {
  const { isLoading, data } = useQuery("getFavorites", () => getFavorites());

  return (
    <BoxPosts
      isLoading={isLoading}
      data={data}
      url={"posts"}
      children={
        <div className="bg-box-color p-one flex justify-center text-2xl">
          Your favorites posts
        </div>
      }
    />
  );
};
