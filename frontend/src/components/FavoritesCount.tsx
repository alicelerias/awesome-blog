import React, { PropsWithChildren } from "react";
import { useQuery } from "react-query";
import { getFavoritesCount } from "../api/queries";

type props = {
  id: string | null;
};

export const FavoritesCount: React.FC<PropsWithChildren<props>> = ({ id }) => {
  const { data } = useQuery("getFavoritesCount", () => getFavoritesCount(id));

  const count = data?.feed.length;

  return <span className="text-sm">{count}</span>;
};
