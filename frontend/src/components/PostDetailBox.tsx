import { useMutation, useQuery } from "react-query";
import { NavigateFunction, useSearchParams } from "react-router-dom";
import { getComments, getPost } from "../api/queries";
import React, { useContext } from "react";
import { BoxLayout } from "./BoxLayout";
import { CommentsComponent } from "./Comments";
import { CreateComment } from "./CreateComment";
import { createComment } from "../api/mutations";
import { Comment } from "../types";
import {
  FieldErrors,
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
} from "react-hook-form";
import { AiOutlineComment } from "react-icons/ai";
import { ToggleFavoriteButton } from "./ToggleFavoriteButton";
import { UpdateButton } from "./UpdateButton";
import { CurrentUserContext } from "../context/CurrentUserContext";

type props = {
  navigate: NavigateFunction;
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  errors: FieldErrors<FieldValues> | undefined;
};

export const PostDetailBox: React.FC<React.PropsWithChildren<props>> = ({
  navigate,
  handleSubmit,
  register,
  errors,
}) => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");
  const { data } = useQuery("getPost", () => getPost(id));
  const { refetch } = useQuery("getComments", () => getComments(id));

  const currentUserContext = useContext(CurrentUserContext);

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
      {currentUserContext?.id === data?.author_id ? (
        <UpdateButton id={id} />
      ) : (
        ""
      )}
      <div className="flex flex-col gap-one p-two  border-b border-b-white">
        <span className="bg-transparent text-3xl text-blue ">
          {data?.title}
        </span>

        <span className="bg-transparent">
          <img src={data?.img} alt="" />
        </span>

        <span className="bg-transparent text-sm italic">"{data?.content}"</span>

        <span
          className="text-blue text-sm cursor-pointer"
          onClick={() => {
            navigate(`/users/detail?id=${data?.author_id}`);
          }}
        >
          {data?.author.username}
        </span>
        <span className="bg-transparent text-sm">
          Created at: {data?.created_at.replace(/-/g, "/").replace(/T/g, " ")}
        </span>
        <div className="flex flex-row justify-end gap-one">
          <span>{data?.comments_count}</span>
          <AiOutlineComment className="h-6 w-6" />

          <ToggleFavoriteButton
            postId={id}
            isFavorite={data?.is_favorite}
            favoritesCount={data?.favorites_count}
          />
        </div>
      </div>

      <CreateComment
        onSubmit={onSubmit}
        handleSubmit={handleSubmit}
        register={register}
        errors={errors}
      />
      <CommentsComponent
        id={id}
        currentUser={currentUserContext}
        navigate={navigate}
      />
    </BoxLayout>
  );
};
