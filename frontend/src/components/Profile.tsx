import { useMutation, useQuery } from "react-query";
import { useSearchParams, useNavigate } from "react-router-dom";
import { getCurrentUser } from "../api/queries";
import { FieldValues, useForm } from "react-hook-form";
import { InputForm } from "./InputForm";
import { User } from "../types";
import { updateCurrentUser } from "../api/mutations";
import { BoxLayout } from "./BoxLayout";
import { InputButton } from "./InputButton";

export const Profile: React.FC<{}> = () => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");
  const navigate = useNavigate();

  const {
    formState: { errors },
    handleSubmit,
    register,
    setValue,
  } = useForm();

  const { data } = useQuery("getCurrentUser", () => getCurrentUser(), {
    onSuccess: (data) => {
      setValue("bio", data.bio);
    },
  });

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
                data?.avatar ||
                "https://ionicframework.com/docs/img/demos/avatar.svg"
              }
              alt="avatar"
            />
            {/* <input
              type="file"
              placeholder="New avatar"
              {...register("avatar", {
                required: false,
              })} 
            /> */}
          </div>
          <div className="flex flex-col gap-two w-4/5">
            <span>{data?.username}</span>

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
