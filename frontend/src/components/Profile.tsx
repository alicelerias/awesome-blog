import { useMutation, useQuery } from "react-query";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { getCurrentUser } from "../api/queries";
import {
  FieldErrors,
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormReset,
  UseFormSetValue,
  useForm,
} from "react-hook-form";
import { InputForm } from "./InputForm";
import { User } from "../types";
import { updateCurrentUser } from "../api/mutations";
import { BoxLayout } from "./BoxLayout";
import { InputButton } from "./InputButton";
import { useContext } from "react";
import { CurrentUserContext } from "../context/CurrentUserContext";

type props = {
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  reset: UseFormReset<FieldValues>;
  navigate: NavigateFunction;
  setValue: UseFormSetValue<FieldValues>;
};

export const Profile: React.FC<props> = ({
  navigate,
  handleSubmit,
  register,
  setValue,
}) => {
  const currentUserContext = useContext(CurrentUserContext);

  setValue("bio", currentUserContext?.bio);

  const { mutate } = useMutation((data: User) => updateCurrentUser(data), {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/");
      }, 2000);
    },
  });

  const onSubmit = async (data: FieldValues) => {
    await mutate(data as User);
  };
  return (
    <BoxLayout>
      <form
        data-testid={"user-detail-component-test-id"}
        onSubmit={handleSubmit(onSubmit)}
      >
        <p className="flex justify-end text-title1 p-one">Profile</p>

        <div className="flex flex-row gap-two my-two">
          <div className="w-1/5">
            <img
              data-testid={"user-detail-img-test-id"}
              className={"w-three aspect-square"}
              src={
                currentUserContext?.avatar ||
                "https://ionicframework.com/docs/img/demos/avatar.svg"
              }
              alt="avatar"
            />
          </div>
          <div className="flex flex-col gap-two w-4/5">
            <span>{currentUserContext?.username}</span>

            <InputForm
              data-testid={"input-user-test-id"}
              placeholder="insert your bio"
              controller={register("bio", {
                required: false,
              })}
            />
          </div>
        </div>

        <InputButton name="Save" data-testid={"user-detail-button-test-id"} />
      </form>
    </BoxLayout>
  );
};
