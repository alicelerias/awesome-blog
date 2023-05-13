import { render, screen } from "@testing-library/react";
import { Posts, User } from "../types";
import * as queries from "../api/queries";
import { TestsContext } from "./testComponents/Context";
import { BlogsPost } from "./BlogsPost";
import { QueryFunctionContext } from "react-query";

describe("test for posts component", () => {
  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "may the force be with you",
    avatar: "xwing.jpg",
    is_following: false,
  };
  const posts: Posts = {
    content: [
      {
        id: "1",
        title: "star wars",
        content: "may the force be with you",
        img: "lightsaber.jpg",
        author: user,
        author_id: "64",
        comments_count: 0,
        favorites_count: 0,
        is_favorite: false,
        created_at: "4-5",
      },
    ],
    next_link: "",
  };

  const renderComponent = () => {
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <BlogsPost id={user.id ? user.id : "0"} navigate={navigate} />
      </TestsContext>
    );
  };

  const fakeGet = jest
    .spyOn(queries, "getBlogsPost")
    .mockImplementation((id) => {
      return async ({
        pageParam = `/posts/blog/${id}`,
      }: QueryFunctionContext) => {
        return Promise.resolve(posts);
      };
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
