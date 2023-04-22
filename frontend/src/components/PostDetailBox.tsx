import { useMutation, useQuery } from "react-query";
import { useSearchParams } from "react-router-dom";
import { getComments, getPost } from "../api/queries";
import React from "react";
import { BoxLayout } from "./BoxLayout";
import { Comments } from "./Comments";
import { CreateComment } from "./CreateComment";
import { createComment } from "../api/mutations";
import { Comment } from "../types";
import { FieldValues, useForm } from "react-hook-form";
import { Heart } from "./Heart";
import { FavoritesCount } from "./FavoritesCount";
import { CommentsCount } from "./CommentsCount";
import { AiOutlineComment } from "react-icons/ai";

type props = {};

export const PostDetailBox: React.FC<React.PropsWithChildren<props>> = ({
  children,
}) => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");
  const { data } = useQuery("getPost", () => getPost(id));
  const { refetch } = useQuery("getComments", () => getComments(id));
  const { reset } = useForm();
  const { mutate } = useMutation(
    (comment: Comment) => createComment(id, comment),
    {
      onSuccess: () => {
        setTimeout(() => {
          refetch();
        }, 2000);
      },
    }
  );
  const onSubmit = (data: FieldValues) => {
    mutate(data as Comment);
  };

  return (
    <BoxLayout>
      {children}

      <div className="flex flex-col gap-one p-two  border-b border-b-white">
        <span className="bg-transparent text-3xl text-blue ">
          {data?.title}
        </span>

        <span className="bg-transparent">
          <img src={data?.img} alt="" />
        </span>

        <span className="bg-transparent text-sm italic">"{data?.content}"</span>

        <span className="text-blue text-sm">{data?.author.username}</span>
        <span className="bg-transparent text-sm">{data?.created_at}</span>
        <div className="flex flex-row justify-end gap-one">
          <CommentsCount id={id} />
          <AiOutlineComment className="h-6 w-6" />
          <FavoritesCount id={id} />
          <Heart postId={data ? data.id : ""} />
        </div>
      </div>

      <CreateComment onSubmit={onSubmit} />
      <Comments />
    </BoxLayout>
  );
};
