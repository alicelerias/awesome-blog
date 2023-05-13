import { useMutation } from "react-query";
import { NavigateFunction } from "react-router-dom";
import { deletePost } from "../api/mutations";
import { PropsWithChildren } from "react";

type props = {
  id: string | null;
  navigate: NavigateFunction;
};

export const DeletePost: React.FC<PropsWithChildren<props>> = ({
  id,
  navigate,
}) => {
  const { mutate } = useMutation(() => deletePost(id), {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/");
      }, 2000);
    },
  });
  const onClick = () => {
    mutate();
  };
  return (
    <button
      data-testid="delete-component-button-test-id"
      className={` bg-yellow p-1 w-1/5 text-sm`}
      onClick={onClick}
    >
      Delete Post
    </button>
  );
};
