import { fireEvent, render, screen } from "@testing-library/react";
import { Posts, User } from "../types";
import * as queries from "./../api/queries";
import { PostsComponent } from "./Posts";
import { TestsContext } from "./testComonents/Context";

describe("test for posts component", () => {
  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "hsajash",
    avatar: "hasashh",
    is_following: false,
  };
  const posts: Posts = {
    feed: [
      {
        id: "asa",
        title: "star wars",
        content: ".sa.",
        img: "hahash",
        author: user,
        author_id: "jsjaj",
        is_favorite: false,
        created_at: "hasahs",
      },
    ],
  };

  const renderComponent = () => {
    return render(
      <TestsContext>
        <PostsComponent />
      </TestsContext>
    );
  };

  const fakeGet = jest.spyOn(queries, "getPosts").mockImplementation(() => {
    return Promise.resolve(posts);
  });

  it("test render posts", async () => {
    renderComponent();

    const componentId = screen.getByTestId("posts-component-test-id");
    expect(componentId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get posts", async () => {
    renderComponent();
    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
