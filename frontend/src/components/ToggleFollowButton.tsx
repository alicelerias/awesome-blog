import React, { useState, useEffect, PropsWithChildren } from "react";

import { Follow, unfollow } from "../api/mutations";

type props = {
  userId: string | null;
  isFollowing: boolean | undefined;
};

export const ToggleFollowButton: React.FC<PropsWithChildren<props>> = ({
  userId,
  isFollowing,
}) => {
  const [isFollowingConst, setIsFollowingConst] = useState(false);

  useEffect(() => {
    isFollowing ? setIsFollowingConst(true) : setIsFollowingConst(false);
  }, [userId, isFollowing]);

  const handleClick = () => {
    if (isFollowingConst) {
      unfollow(userId).then(() => {
        setIsFollowingConst(false);
      });
    } else {
      Follow(userId).then(() => {
        setIsFollowingConst(true);
      });
    }
  };

  return (
    <button
      className="bg-blue p-1 w-1/5 text-smm h-6 truncate transition duration-150 ease-in"
      onClick={handleClick}
    >
      {isFollowingConst ? "Unfollow" : "Follow"}
    </button>
  );
};
