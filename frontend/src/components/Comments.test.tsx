import { render, screen } from "@testing-library/react";
import { CommentsComponent } from "./Comments";
import { Comments, User } from "../types";
import * as queries from "./../api/queries";
import { QueryFunctionContext } from "react-query";
import { TestsContext } from "./testComponents/Context";

describe("Tests for commements component", () => {
  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "may the force be with you",
    avatar: "xwing.jpg",
    is_following: false,
  };

  const comments: Comments = {
    content: [
      {
        id: "1",
        postId: "1",
        authorId: "64",
        author: user,
        content: "content comment 1",
        created_at: "4-5",
      },
      {
        id: "2",
        postId: "1",
        authorId: "64",
        author: user,
        content: "content comment 2",
        created_at: "4-5",
      },
    ],
    next_link: "/comments?page=2",
  };

  const navigate = jest.fn();

  beforeEach(() => {
    jest.resetAllMocks();
  });

  const renderComponent = () =>
    render(
      <TestsContext>
        <CommentsComponent id="1" currentUser={user} navigate={navigate} />
      </TestsContext>
    );

  const fakeGet = jest
    .spyOn(queries, "getComments")
    .mockImplementation((id) => {
      return async ({ pageParam = `/comment/${id}` }: QueryFunctionContext) => {
        return Promise.resolve(comments);
      };
    });

  it("render component", () => {
    renderComponent();
    const componentId = screen.getByTestId("comments-component-test-id");
    expect(componentId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get comments", async () => {
    renderComponent();
    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
