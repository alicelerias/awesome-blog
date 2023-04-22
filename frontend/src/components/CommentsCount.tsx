import React, { PropsWithChildren } from "react";
import { useQuery } from "react-query";
import { getComments } from "../api/queries";

type props = {
  id: string | null;
};

export const CommentsCount: React.FC<PropsWithChildren<props>> = ({ id }) => {
  const { data } = useQuery("getComments", () => getComments(id));

  const count = data?.comments.length;

  return <span className="text-sm">{count}</span>;
};
