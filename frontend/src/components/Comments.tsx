import { useQuery } from "react-query";
import { getComments } from "../api/queries";
import { BoxLayout } from "./BoxLayout";
import { PropsWithChildren } from "react";
import { DeleteCommentButton } from "./DeleteCommentButton";
import { User } from "../types";
import { NavigateFunction } from "react-router-dom";
import Skeleton from "./Skeleton";

type props = {
  id: string | null;
  currentUser: User | undefined;
  navigate: NavigateFunction;
};
export const CommentsComponent: React.FC<PropsWithChildren<props>> = ({
  id,
  currentUser,
  navigate,
}) => {
  const { data, isLoading } = useQuery("getComments", () => getComments(id));
  return (
    <BoxLayout>
      <h1 className="sm:text3xl">Comments</h1>
      <div>
        {isLoading ? (
          <Skeleton />
        ) : (
          data?.comments.map((comment) => (
            <div
              key={comment.id}
              className="flex flex-col p-two gap-two opacity-3 w-full shadow-md"
            >
              <div className="flex justify-end w-full">
                {currentUser?.id === comment.author.id && (
                  <DeleteCommentButton
                    commentId={comment.id}
                    navigate={navigate}
                  />
                )}
              </div>

              <div className="flex flex-row gap-two">
                <div className="flex flex-col">
                  <img
                    className="w-10 aspect-square"
                    src={
                      comment?.author.avatar ||
                      "https://ionicframework.com/docs/img/demos/avatar.svg"
                    }
                    alt=""
                  />
                  <p className="text-sm">@{comment.author.username}</p>
                </div>
                <p className="text-sm italic p-one">{comment.content}</p>
              </div>
            </div>
          ))
        )}
      </div>
    </BoxLayout>
  );
};
