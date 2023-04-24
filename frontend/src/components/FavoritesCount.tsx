import React, { PropsWithChildren, useEffect, useState } from "react";
import { useQuery } from "react-query";
import { getFavoritesCount } from "../api/queries";

type props = {
  id: string | null;
  isFavorite: boolean | undefined;
};

export const FavoritesCount: React.FC<PropsWithChildren<props>> = ({
  id,
  isFavorite,
}) => {
  const [count, setCount] = useState(0);

  useEffect(() => {
    getFavoritesCount(id).then((response) => {
      setCount(response.feed.length);
    });
  }, [id, isFavorite]);

  return <span className="text-sm">{count}</span>;
};
