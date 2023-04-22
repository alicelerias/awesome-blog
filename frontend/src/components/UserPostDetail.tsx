import { useQuery } from "react-query";
import { useSearchParams, useNavigate } from "react-router-dom";
import { getPost } from "../api/queries";
import { useState } from "react";
import { BoxLayout } from "./BoxLayout";
import { PostDetailBox } from "./PostDetailBox";

export const UserPostDetail: React.FC<{}> = () => {
  const [searchParam] = useSearchParams();
  const id = searchParam.get("id");

  const [post, setUpdatePost] = useState("");
  const navigate = useNavigate();

  const { data } = useQuery("getPost", () => getPost(id));

  return (
    <BoxLayout>
      <PostDetailBox>
        <button
          className="flex justify-end"
          onClick={() => {
            data ? setUpdatePost(data.id) : navigate("/");

            navigate(`/posts/update?id=${data?.id}`);
          }}
        >
          {" "}
          ...{" "}
        </button>
      </PostDetailBox>
    </BoxLayout>
  );
};
