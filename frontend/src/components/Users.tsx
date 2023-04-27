import React from "react";
import { useQuery } from "react-query";
import { getUsers } from "../api/queries";
import { NavigateFunction } from "react-router-dom";
import { ToggleFollowButton } from "./ToggleFollowButton";

type props = {
  navigate: NavigateFunction;
};
export const UsersComponent: React.FC<props> = ({ navigate }) => {
  const { data, isLoading } = useQuery("users", getUsers);

  return (
    <div
      data-testid={"users-component-test-id"}
      className={"flex flex-col sm:w-2/5"}
    >
      <div className="p-one border-b border-white w-full">Who to follow?</div>
      {isLoading ? (
        <p>is loading</p>
      ) : (
        data?.users.map((user) => (
          <div className={`flex flex-row gap-two p-one w-full`}>
            <div className="flex flex-col w-3/4">
              <img
                className=" w-8 aspect-square"
                src={
                  user.avatar ||
                  "https://ionicframework.com/docs/img/demos/avatar.svg"
                }
                alt=""
              />
              <div
                className="w-full text-smm cursor-pointer hover:text-blue"
                onClick={() => {
                  navigate(`/users/detail?id=${user.id}`);
                }}
              >
                @{user?.username}
              </div>
              <div className="w-full text-smm">{user?.bio}</div>
            </div>
            <ToggleFollowButton
              isFollowing={user.is_following}
              userId={user.id}
            />
          </div>
        ))
      )}
    </div>
  );
};
