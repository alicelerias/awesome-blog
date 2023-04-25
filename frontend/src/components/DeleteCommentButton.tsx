import React, { PropsWithChildren } from "react";
import { BsFillTrash3Fill } from "react-icons/bs";
import { useMutation, useQuery } from "react-query";
import { deleteComment } from "../api/mutations";
import { getComments, getPost } from "../api/queries";
import { useNavigate } from "react-router-dom";

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
    <button onClick={() => mutate()}>
      <BsFillTrash3Fill />
    </button>
  );
};
