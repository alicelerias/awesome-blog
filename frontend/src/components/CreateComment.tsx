import React from "react";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import { QueryFunction, useMutation, UseQueryResult } from "react-query";
import { useNavigate, useSearchParams } from "react-router-dom";
import { createComment } from "../api/mutations";
import { Comment } from "../types";
import { InputButton } from "./InputButton";
import { InputForm } from "./InputForm";

type props = {
  onSubmit: SubmitHandler<FieldValues>;
};
export const CreateComment: React.FC<React.PropsWithChildren<props>> = ({
  onSubmit,
}) => {
  const {
    formState: { errors },
    handleSubmit,
    register,
  } = useForm();

  return (
    <div>
      <form className="flex flex-col gap-one" onSubmit={handleSubmit(onSubmit)}>
        <InputForm
          placeholder="insert your comment"
          controller={register("content", {
            required: true,
          })}
          error={errors.content}
        />
        <InputButton name="Comment" />
      </form>
    </div>
  );
};
