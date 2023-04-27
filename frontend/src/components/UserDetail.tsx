import { useQuery } from "react-query";
import { NavigateFunction, useSearchParams } from "react-router-dom";
import { getUser } from "../api/queries";
import { BlogsPost } from "./BlogsPost";
import { BoxLayout } from "./BoxLayout";
import { ToggleFollowButton } from "./ToggleFollowButton";
import { PropsWithChildren } from "react";

type props = {
  navigate: NavigateFunction;
};

export const UserDetail: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");

  const { isLoading, data } = useQuery("getUser", () => getUser(id));

  return (
    <BoxLayout>
      <div className="flex flex-col sm:flex-row justify-center gap-two">
        {isLoading ? (
          "is Loading"
        ) : (
          <>
            <div className="sm:flex sm:flex-col sm:justify-center sm:gap-two sm:w-2/6 w-0"></div>
            <div className="flex flex-col  justify-center gap-two w-full mb-two sm:mb-0 sm:w-2/6 sm:fixed sm:left-40">
              <div className="flex justify-center sm:justify-start">
                <img
                  className="w-3/5 aspect-square"
                  src={
                    data?.avatar ||
                    "https://ionicframework.com/docs/img/demos/avatar.svg"
                  }
                  alt=""
                />
              </div>

              <ToggleFollowButton
                isFollowing={data?.is_following}
                userId={id ? id : ""}
              />
              <span className="flex justify-center sm:justify-start text-title1 text-blue">
                {data?.username}{" "}
              </span>
              <span className="text-sm italic sm:w-4/6">"{data?.bio}"</span>
            </div>
          </>
        )}

        <BlogsPost id={id} navigate={navigate} />
      </div>
    </BoxLayout>
  );
};
