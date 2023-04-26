import React, { PropsWithChildren } from "react";
import { useMutation } from "react-query";
import { deleteComment } from "../api/mutations";
import { NavigateFunction } from "react-router-dom";
import { Sidebar } from "./Sidebar";

type props = {
  commentId: string;
  navigate: NavigateFunction;
};

export const DeleteCommentButton: React.FC<PropsWithChildren<props>> = ({
  commentId,
  navigate,
}) => {
  const { mutate } = useMutation(
    "deleteComment",
    () => deleteComment(commentId),
    {
      onSuccess: () => {
        setTimeout(() => {
          navigate("/");
        }, 2000);
      },
    }
  );

  return (
    <Sidebar
      name={"..."}
      children={
        <p
          onClick={() => mutate()}
          className="text-black text-sm block py-2 px-4 hover:text-blue cursor-pointer"
        >
          Delete
        </p>
      }
    />
  );
};
