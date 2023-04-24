import React, { useState, useEffect, PropsWithChildren } from "react";

import { favorite, unfavorite } from "../api/mutations";
import { AiFillHeart } from "react-icons/ai";
import { FavoritesCount } from "./FavoritesCount";

type props = {
  postId: string | null;
  isFavorite: boolean | undefined;
};

export const ToggleFavoriteButton: React.FC<PropsWithChildren<props>> = ({
  postId,
  isFavorite,
}) => {
  const [isFavoriteConst, setIsFavoriteConst] = useState(false);

  useEffect(() => {
    isFavorite ? setIsFavoriteConst(true) : setIsFavoriteConst(false);
  }, [postId, isFavorite]);

  const handleClick = () => {
    if (isFavoriteConst) {
      unfavorite(postId).then(() => {
        setIsFavoriteConst(false);
      });
    } else {
      favorite(postId).then(() => {
        setIsFavoriteConst(true);
      });
    }
  };

  return (
    <>
      <FavoritesCount id={postId} isFavorite={isFavoriteConst} />
      <AiFillHeart
        className={
          isFavoriteConst
            ? "text-red-700 h-6 w-6 transition duration-150 ease-in cursor-pointer"
            : "h-6 w-6 transition duration-150 ease-in cursor-pointer"
        }
        onClick={handleClick}
      />
    </>
  );
};
