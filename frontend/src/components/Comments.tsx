import { useQuery } from "react-query";
import { useSearchParams } from "react-router-dom";
import { getComments } from "../api/queries";
import { BoxLayout } from "./BoxLayout";
import { CurrentUserContext } from "../context/CurrentUserContext";
import { useContext } from "react";
import { DeleteCommentButton } from "./DeleteCommentButton";

export const Comments: React.FC<{}> = () => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");
  const currentUser = useContext(CurrentUserContext);

  const { data, isLoading } = useQuery("getComments", () => getComments(id));
  return (
    <BoxLayout>
      <h1 className="sm:text3xl">Comments</h1>
      <div>
        {isLoading
          ? "is Loading"
          : data?.comments.map((comment) => (
              <div className="flex flex-col p-two gap-two opacity-3 w-full shadow-md">
                <div className="flex justify-end w-full">
                  {currentUser?.id === comment.author.id ? (
                    <DeleteCommentButton commentId={comment.id} postId={id} />
                  ) : (
                    ""
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
            ))}
      </div>
    </BoxLayout>
  );
};
