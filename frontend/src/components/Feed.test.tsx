import { render, screen } from "@testing-library/react";
import { Posts, User } from "../types";
import * as queries from "../api/queries";
import { FeedComponent } from "./Feed";
import { TestsContext } from "./testComponents/Context";

describe("test for feed component", () => {
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
        <FeedComponent navigate={navigate} />
      </TestsContext>
    );
  };

  const fakeGet = jest.spyOn(queries, "getFeed").mockImplementation(() => {
    return Promise.resolve(posts);
  });

  it("test render feed", async () => {
    renderComponent();

    const componentId = screen.getByTestId("posts-component-test-id");
    expect(componentId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get feed", async () => {
    renderComponent();
    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
