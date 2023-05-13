import { PropsWithChildren } from "react";
import { NavigateFunction } from "react-router-dom";

type props = {
  id: string | null;
  navigate: NavigateFunction;
};
export const UpdateButton: React.FC<PropsWithChildren<props>> = ({
  id,
  navigate,
}) => {
  return (
    <button
      className="flex justify-end"
      onClick={() => {
        navigate(`/posts/update?id=${id}`);
      }}
    >
      {" "}
      Edit{" "}
    </button>
  );
};
