import { render, screen, fireEvent } from "@testing-library/react";
import { NavigateFunction } from "react-router-dom";
import { InfiniteData } from "react-query";
import { Posts, User } from "../types";
import { BoxPosts } from "./BoxPosts";
import { TestsContext } from "./testComponents/Context";

describe("BoxPosts", () => {
  const navigate = jest.fn();

  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "may the force be with you",
    avatar: "xwing.jpg",
    is_following: false,
  };
  const posts: InfiniteData<Posts> = {
    pages: [
      {
        content: [
          {
            id: "1",
            title: "Post 1",
            content: "This is the content of post 1",
            img: "https://picsum.photos/id/1/200/300",
            author_id: "64",
            author: user,
            comments_count: 2,
            favorites_count: 3,
            is_favorite: false,
            created_at: "4-5",
          },
          {
            id: "2",
            title: "Post 2",
            content: "This is the content of post 2",
            img: "https://picsum.photos/id/1/200/300",
            author_id: "64",
            author: user,
            comments_count: 1,
            favorites_count: 5,
            is_favorite: false,
            created_at: "4-5",
          },
        ],
        next_link: "/api/posts?page=2",
      },
      {
        content: [
          {
            id: "3",
            title: "Post 3",
            content: "This is the content of post 3",
            img: "https://picsum.photos/id/1/200/300",
            author_id: "64",
            author: user,
            comments_count: 1,
            favorites_count: 5,
            is_favorite: false,
            created_at: "4-5",
          },
        ],
        next_link: "",
      },
    ],
    pageParams: [],
  };

  const fetchNextPage = jest.fn();

  const props = {
    isLoading: false,
    posts,
    navigate: navigate as NavigateFunction,
    hasNextPage: true,
    fetchNextPage,
  };

  beforeEach(() => {
    jest.resetAllMocks();
  });

  const renderComponent = () => {
    render(
      <TestsContext>
        <BoxPosts
          isLoading={props.isLoading}
          data={posts}
          fetchNextPage={props.fetchNextPage}
          navigate={props.navigate}
        />
      </TestsContext>
    );
  };

  it("loading posts", () => {
    render(
      <TestsContext>
        <BoxPosts
          isLoading={true}
          data={posts}
          fetchNextPage={props.fetchNextPage}
          navigate={props.navigate}
        />
      </TestsContext>
    );
    expect(screen.getByTestId("skeleton-test-id")).toBeInTheDocument();
  });

  it("renders the posts", () => {
    renderComponent();
    expect(screen.getByText("Post 1")).toBeInTheDocument();
    expect(screen.getByText("Post 2")).toBeInTheDocument();
    expect(screen.getByText("Post 3")).toBeInTheDocument();
    expect(screen.queryByTestId("skeleton-test-id")).not.toBeInTheDocument();
  });

  it("navigates to post detail when title is clicked", () => {
    renderComponent();

    fireEvent.click(screen.getByText("Post 1"));
    expect(navigate).toHaveBeenCalledWith("/posts/detail?id=1");
  });

  it("navigates to post detail when comment icon is clicked", () => {
    renderComponent();

    fireEvent.click(screen.getAllByTestId("comment-icon-test-id")[1]);
    expect(navigate).toHaveBeenCalledWith("/posts/detail?id=2");
  });

  it("should fetch next page when intersection observer is triggered", () => {
    renderComponent();

    const intersectionObserverElement = screen.getByTestId(
      "intersection-observer-test-id"
    );
    expect(intersectionObserverElement).toBeInTheDocument();
  });
});
