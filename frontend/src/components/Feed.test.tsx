import { render, screen } from "@testing-library/react";
import { Posts, User } from "../types";
import * as queries from "../api/queries";
import { FeedComponent } from "./Feed";
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
        comments_count: 0,
        favorites_count: 0,
        is_favorite: false,
        created_at: "SS",
      },
    ],
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
