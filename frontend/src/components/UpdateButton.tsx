import { PropsWithChildren } from "react";
import { useNavigate } from "react-router-dom";

type props = {
  id: string | null;
};
export const UpdateButton: React.FC<PropsWithChildren<props>> = ({ id }) => {
  const navigate = useNavigate();

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
