import { PropsWithChildren, useState } from "react";
import { AiOutlineHeart } from "react-icons/ai";
import { useMutation, useQuery } from "react-query";
import { favorite } from "../api/mutations";
import { getFavoritesCount } from "../api/queries";

type props = {
  postId: string;
};

export const Heart: React.FC<PropsWithChildren<props>> = ({ postId }) => {
  const [id, setId] = useState("");

  const { mutate } = useMutation(() => favorite(id));

  return (
    <>
      <AiOutlineHeart
        className="h-6 w-6 cursor-pointer"
        onClick={() => {
          setId(postId);
          setTimeout(() => {
            mutate();
          }, 1000);
        }}
      />
    </>
  );
};
