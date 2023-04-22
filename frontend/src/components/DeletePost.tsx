import { useMutation } from "react-query";
import { useNavigate, useSearchParams } from "react-router-dom";
import { deletePost } from "../api/mutations";
import { Post } from "../types";

export const DeletePost: React.FC<{}> = () => {
  const [searchParam] = useSearchParams();
  const navigate = useNavigate();
  const id = searchParam.get("id");
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
    <button className={` bg-yellow p-1 w-1/5 text-sm`} onClick={onClick}>
      Delete Post
    </button>
  );
};
