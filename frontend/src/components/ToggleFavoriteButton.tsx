import React, { useState, useEffect, PropsWithChildren } from "react";

import { favorite, unfavorite } from "../api/mutations";
import { AiFillHeart } from "react-icons/ai";

type props = {
  postId: string | null;
  isFavorite: boolean | undefined;
  favoritesCount: number | undefined;
};

export const ToggleFavoriteButton: React.FC<PropsWithChildren<props>> = ({
  postId,
  isFavorite,
  favoritesCount,
}) => {
  const [isFavoriteConst, setIsFavoriteConst] = useState(false);

  const [count, setCount] = useState(favoritesCount);

  useEffect(() => {
    isFavorite ? setIsFavoriteConst(true) : setIsFavoriteConst(false);
  }, [postId, isFavorite]);

  const handleClick = () => {
    if (isFavoriteConst) {
      unfavorite(postId).then(() => {
        setIsFavoriteConst(false);
        setCount(count ? count - 1 : 0);
      });
    } else {
      favorite(postId).then(() => {
        setIsFavoriteConst(true);
        setCount(count ? count + 1 : 1);
      });
    }
  };

  return (
    <div className="flex flex-row gap-one" data-testid="follow-button-testid">
      <span className="text-sm">{count}</span>
      <AiFillHeart
        className={
          isFavoriteConst
            ? "text-red-700 h-6 w-6 transition duration-150 ease-in cursor-pointer"
            : "h-6 w-6 transition duration-150 ease-in cursor-pointer"
        }
        onClick={handleClick}
      />
    </div>
  );
};
