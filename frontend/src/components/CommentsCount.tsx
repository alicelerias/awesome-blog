import React, { PropsWithChildren, useEffect, useState } from "react";
import { getComments } from "../api/queries";

type props = {
  id: string | null;
};

export const CommentsCount: React.FC<PropsWithChildren<props>> = ({ id }) => {
  const [count, setCount] = useState(0);
  useEffect(() => {
    getComments(id).then((response) => {
      setCount(response.comments.length);
    });
  }, [id]);

  return <span className="text-sm">{count}</span>;
};
