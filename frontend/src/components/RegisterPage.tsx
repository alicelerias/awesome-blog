import React, { useState } from "react";
import { useMutation } from "react-query";
import { registerUser } from "../api/mutations";
import { useNavigate } from "react-router-dom";
import { BoxLayout } from "./BoxLayout";
import { InputForm } from "./InputForm";
import { InputButton } from "./InputButton";
import {
  FieldErrors,
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormReset,
} from "react-hook-form";
import { User } from "../types";
import { Alert } from "./Alert";
import { error } from "console";

type props = {
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
  errors: FieldErrors<FieldValues> | undefined;
  reset: UseFormReset<FieldValues>;
};

export const Register: React.FC<props> = ({
  handleSubmit,
  register,
  errors,
  reset,
}) => {
  const navigate = useNavigate();
  const [alert, showAlert] = useState(false);
  const { mutate } = useMutation("register", registerUser, {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/login");
      }, 2000);
    },
    onError: () => {
      showAlert(true);
      setTimeout(() => {
        showAlert(false);
      }, 5000);
    },
  });

  const onSubmit = (data: FieldValues) => {
    mutate(data as User);
    setTimeout(() => {
      reset();
    }, 2000);
  };

  return (
    <div className="flex flex-col justify-center w-screen h-screen bg-black">
      <div className="flex flex-col h-screen sm:w-auto sm:shadow-lg p-8 rounded bg-box-color sm:h-5/6 sm:mx-auto sm:space-y-4 sm:p-4 ">
        <span className="flex justify-end text-title1">Register</span>
        {alert && <Alert type="error" message="Invalid fields" />}
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="flex flex-row gap-four">
            <img
              className="w-36 aspect-square"
              src={"https://ionicframework.com/docs/img/demos/avatar.svg"}
              alt=""
            />
            <div className="flex flex-col gap-two">
              <span>
                <InputForm
                  controller={register("username", {
                    required: true,
                  })}
                  type="text"
                  placeholder="Username"
                  error={errors?.username}
                />
              </span>

              <span>
                <InputForm
                  controller={register("email", {
                    required: true,
                  })}
                  type="text"
                  placeholder="Email"
                  error={errors?.email}
                />
              </span>

              <span>
                <input
                  className="bg-transparent text-sm italic p-one"
                  placeholder="Password"
                  {...register("password", {
                    required: true,
                    ...errors?.password,
                  })}
                  type="password"
                />
              </span>
            </div>
          </div>

          <div className="flex flex-col gap-two mt-two">
            <span>
              <InputForm
                controller={register("bio", {
                  required: true,
                })}
                type="text"
                placeholder="Bio"
                error={errors?.bio}
              />
            </span>
            <InputButton name="Register" />
          </div>
        </form>
      </div>
    </div>
  );
};
