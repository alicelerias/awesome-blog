import { fireEvent, render, screen } from "@testing-library/react";
import { User } from "../types";
import * as queries from "../api/queries";
import * as mutations from "../api/mutations";
import { TestsContext } from "./testComonents/Context";
import { Profile } from "./Profile";

describe("tests for user detail componet", () => {
  const user: User = {
    id: "64",
    username: "lukeskywalker",
    email: "email@email.com",
    bio: "hsajash",
    avatar: "hasashh",
    is_following: false,
  };

  const fakeGet = jest
    .spyOn(queries, "getCurrentUser")
    .mockImplementation(() => {
      return Promise.resolve(user);
    });

  const renderComponent = () => {
    const handleSubmit = jest.fn();
    const register = jest.fn();
    const reset = jest.fn();
    const setValue = jest.fn();
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <Profile
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

    const componentId = screen.getByTestId("user-detail-component-test-id");
    expect(componentId).toBeInTheDocument();

    const imgId = screen.getByTestId("user-detail-img-test-id");
    expect(imgId).toBeInTheDocument();

    const inputId = screen.getByTestId("input-user-test-id");
    expect(inputId).toBeInTheDocument();

    const buttonId = screen.getByTestId("input-button-component-test-id");
    expect(buttonId).toBeInTheDocument();

    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test get post detail", async () => {
    renderComponent();

    expect(fakeGet).toHaveBeenCalledTimes(1);
  });

  it("test update post", async () => {
    renderComponent();
    const componentId = screen.getByTestId("user-detail-component-test-id");
    fireEvent.submit(componentId);
  });
});
