import { InputButton } from "./InputButton";
import { InputForm } from "./InputForm";
import {
  FieldValues,
  UseFormHandleSubmit,
  UseFormRegister,
} from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useMutation } from "react-query";
import { login } from "../api/mutations";
import { Credential } from "../types";
import { useAlert } from "./Alert";

type props = {
  handleSubmit: UseFormHandleSubmit<FieldValues>;
  register: UseFormRegister<FieldValues>;
};

export const LoginPage: React.FC<props> = ({ handleSubmit, register }) => {
  const navigate = useNavigate();
  const [showAlert, Alert] = useAlert();

  const { mutate } = useMutation(login, {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/");
      }, 2000);
    },
    onError: () => {
      showAlert();
    },
  });

  const onSubmit = (data: FieldValues) => {
    mutate(data as Credential);
  };
  return (
    <>
      <div className="flex flex-col justify-center w-screen h-screen">
        <div className="flex flex-col justify-between w-full h-screen bg-box-color sm:shadow-lg p-8 rounded-md sm:px-4  sm:w-auto sm:h-5/6 sm:aspect-[1/2] sm:mx-auto sm:space-y-4 sm:p-4 ">
          <div className="text-center self-center text-title1 sm:font-semibold">
            !AWESOME
          </div>
          <div className="font-medium self-center text-xl sm:text-2xl">
            <Alert message="Wrong username or password!" type="error" />
            <form onSubmit={handleSubmit(onSubmit)}>
              <span>
                <InputForm
                  className="bg-transparent text-3xl p-one"
                  placeholder="username"
                  controller={register("username", {
                    required: false,
                  })}
                  type="text"
                />
              </span>

              <span>
                <input
                  className="bg-transparent text-3xl p-one"
                  placeholder="password"
                  {...register("password", {
                    required: false,
                  })}
                  type="password"
                />
              </span>
              <div className="flex justify-center m-two">
                <InputButton name="Login" />
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  );
};
