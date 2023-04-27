import { fireEvent, render, screen } from "@testing-library/react";
import { Post, User } from "../types";
import * as queries from "../api/queries";
import { TestsContext } from "./testComonents/Context";
import { UpdatePost } from "./UpdatePost";

describe("tests for post detail componet", () => {
  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "hsajash",
    avatar: "hasashh",
    is_following: false,
  };
  const post: Post = {
    id: "8",
    title: "star wars",
    content: ".sa.",
    img: "hahash",
    author: user,
    author_id: "jsjaj",
    is_favorite: false,
    favorites_count: 0,
    comments_count: 0,
    created_at: "jsasj",
  };

  const fakeGet = jest.spyOn(queries, "getPost").mockImplementation(() => {
    return Promise.resolve(post);
  });

  const renderComponent = () => {
    const handleSubmit = jest.fn();
    const register = jest.fn();
    const reset = jest.fn();
    const setValue = jest.fn();
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <UpdatePost
          navigate={navigate}
          setValue={setValue}
          handleSubmit={handleSubmit}
          register={register}
          reset={reset}
        />
      </TestsContext>
    );
  };

  it("test render post detail", async () => {
    renderComponent();

    const componentId = screen.getByTestId("update-post-component-test-id");
    expect(componentId).toBeInTheDocument();

    const imgId = screen.getByTestId("input-title-test-id");
    expect(imgId).toBeInTheDocument();

    const inputId = screen.getByTestId("input-content-test-id");
    expect(inputId).toBeInTheDocument();

    const buttonId = screen.getByTestId("input-button-component-test-id");
    expect(buttonId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get post detail", async () => {
    renderComponent();

    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
