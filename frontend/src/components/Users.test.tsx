import { render, screen } from "@testing-library/react";
import { Users } from "../types";
import * as queries from "./../api/queries";
import { TestsContext } from "./testComonents/Context";
import { UsersComponent } from "./Users";

describe("test for users component", () => {
  const users: Users = {
    users: [
      {
        id: "64",
        username: "username",
        email: "email",
        bio: "bio",
        avatar: "avatar",
        is_following: false,
      },
    ],
  };

  const fakeGet = jest.spyOn(queries, "getUsers").mockImplementation(() => {
    return Promise.resolve(users);
  });

  const renderComponent = () => {
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <UsersComponent navigate={navigate} />
      </TestsContext>
    );
  };

  it("test render posts", async () => {
    renderComponent();
    const componentId = screen.getByTestId("users-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get users", async () => {
    renderComponent();
    expect(fakeGet).toHaveBeenCalledTimes(1);
  });
});
