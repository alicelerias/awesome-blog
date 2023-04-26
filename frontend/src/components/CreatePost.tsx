import { useMutation } from "react-query";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { FieldValues, useForm } from "react-hook-form";
import { createPost } from "../api/mutations";
import { Post } from "../types";
import { InputForm } from "./InputForm";
import { BoxLayout } from "./BoxLayout";
import { InputButton } from "./InputButton";
import { PropsWithChildren } from "react";

type props = {
  navigate: NavigateFunction;
};

export const CreatePost: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { mutate } = useMutation(createPost, {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/");
      }, 2000);
    },
  });

  const {
    formState: { errors },
    handleSubmit,
    register,
    reset,
  } = useForm();

  const onSubmit = (data: FieldValues) => {
    mutate(data as Post);
    reset(data);
  };

  return (
    <BoxLayout>
      <div className="flex flex-col gap-two">
        <span className="flex justify-end text-title1">New Post</span>
        <form onSubmit={handleSubmit(onSubmit)}>
          <span>
            <InputForm
              controller={register("title", {
                required: true,
              })}
              type="text"
              placeholder="Insert title"
              error={errors.title}
            />
          </span>

          <span>
            <InputForm
              controller={register("content", {
                required: true,
              })}
              type="text"
              placeholder="Insert content"
              error={errors.content}
            />
          </span>
          <InputButton name="save" />
        </form>
      </div>
    </BoxLayout>
  );
};
