import { render, screen } from "@testing-library/react";
import { Posts, User } from "../types";
import * as queries from "../api/queries";
import { TestsContext } from "./testComponents/Context";
import { FavoritesPosts } from "./FavoritesPosts";

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
        <FavoritesPosts navigate={navigate} />
      </TestsContext>
    );
  };

  const fakeGet = jest.spyOn(queries, "getFavorites").mockImplementation(() => {
    return Promise.resolve(posts);
  });

  it("test render favorites", async () => {
    renderComponent();

    const componentId = screen.getByTestId("favorites-component-test-id");
    expect(componentId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get posts", async () => {
    renderComponent();
    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
