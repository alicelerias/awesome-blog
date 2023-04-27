import React from "react";
import {
  FieldErrors,
  FieldValues,
  SubmitHandler,
  UseFormHandleSubmit,
  UseFormRegister,
} from "react-hook-form";
import { InputButton } from "./InputButton";
import { InputForm } from "./InputForm";

type props = {
  onSubmit: SubmitHandler<FieldValues>;
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  errors: FieldErrors<FieldValues> | undefined;
};
export const CreateComment: React.FC<React.PropsWithChildren<props>> = ({
  onSubmit,
  handleSubmit,
  register,
  errors,
}) => {
  return (
    <div>
      <form className="flex flex-col gap-one" onSubmit={handleSubmit(onSubmit)}>
        <InputForm
          placeholder="insert your comment"
          controller={register("content", {
            required: true,
          })}
          error={errors?.content}
        />
        <InputButton name="Comment" />
      </form>
    </div>
  );
};
