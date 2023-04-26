import React, { PropsWithChildren } from "react";
import { BsFillTrash3Fill } from "react-icons/bs";
import { useMutation, useQuery } from "react-query";
import { deleteComment } from "../api/mutations";
import { getComments, getPost } from "../api/queries";
import { useNavigate } from "react-router-dom";
import { Sidebar } from "./Sidebar";

type props = {
  commentId: string;
  postId: string | null;
};

export const DeleteCommentButton: React.FC<PropsWithChildren<props>> = ({
  commentId,
  postId,
}) => {
  const navigate = useNavigate();
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
