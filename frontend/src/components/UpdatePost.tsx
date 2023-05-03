import { useMutation, useQuery } from "react-query";
import { useSearchParams, NavigateFunction } from "react-router-dom";
import { getPost } from "../api/queries";
import {
  FieldErrors,
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormReset,
  UseFormSetValue,
} from "react-hook-form";
import { updatePost } from "../api/mutations";
import { Post } from "../types";
import { InputForm } from "./InputForm";
import { InputButton } from "./InputButton";
import { BoxLayout } from "./BoxLayout";
import { DeletePost } from "./DeletePost";

type props = {
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  reset: UseFormReset<FieldValues>;
  errors?: FieldErrors<FieldValues>;
  setValue: UseFormSetValue<FieldValues>;
  navigate: NavigateFunction;
};

export const UpdatePost: React.FC<props> = ({
  handleSubmit,
  register,
  reset,
  errors,
  setValue,
  navigate,
}) => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");

  const { data } = useQuery("getPost", () => getPost(id), {
    onSuccess: (data) => {
      setValue("title", data?.title);
      setValue("content", data?.content);
    },
  });

  const { mutate } = useMutation((data: Post) => updatePost(id, data), {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/");
      }, 2000);
    },
  });

  const onSubmit = (data: FieldValues) => {
    mutate(data as Post);
    setTimeout(() => {
      reset();
    }, 2000);
  };
  return (
    <BoxLayout>
      <div className="flex justify-end">
        {" "}
        <DeletePost id={id} navigate={navigate} />
      </div>
      <form
        data-testid={"update-post-component-test-id"}
        onSubmit={handleSubmit(onSubmit)}
      >
        <span>
          <InputForm
            className="bg-transparent text-3xl p-one"
            placeholder="insert your title"
            data-testid={"input-title-test-id"}
            controller={register("title", {
              required: false,
            })}
            error={errors?.title}
          />
        </span>

        <span>
          <InputForm
            placeholder="insert your content"
            data-testid={"input-content-test-id"}
            controller={register("content", {
              required: false,
            })}
          />
        </span>

        <div className="p-one">@{data?.author.username}</div>
        <div>{data?.author.bio}</div>
        <div className="m-one">
          <InputButton data-testid={"post-detail-button-test-id"} name="save" />
        </div>
      </form>
    </BoxLayout>
  );
};
