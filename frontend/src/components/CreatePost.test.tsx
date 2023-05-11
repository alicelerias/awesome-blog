import React from "react";
import { render, fireEvent, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { CreatePost } from "./CreatePost";
import * as mutations from "../api/mutations";

jest.mock("react-query");

describe("CreatePost component", () => {
  const navigate = jest.fn();
  const handleSubmit = jest.fn();
  const register = jest.fn();
  const errors = {
    title: {
      type: "required",
      message: "Title is required",
    },
  };
  const reset = jest.fn();

  const renderComponent = () => {
    render(
      <CreatePost
        navigate={navigate}
        handleSubmit={handleSubmit}
        register={register}
        errors={errors}
        reset={reset}
      />
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders the form", () => {
    renderComponent();

    const componentId = screen.getByTestId("create-post-test-id");

    const buttonId = screen.getByTestId("input-button-component-test-id");

    expect(componentId).toBeInTheDocument();
    expect(buttonId).toBeInTheDocument();
  });

  it("submits the form when the save button is clicked", async () => {
    renderComponent();
    // const fakePost = jest.spyOn(mutations, "createPost").mockImplementation()

    const titleInput = screen.getByPlaceholderText("Insert title");
    const contentInput = screen.getByPlaceholderText("Insert content");
    const saveButton = screen.getByText("save");

    userEvent.type(titleInput, "Test Title");
    userEvent.type(contentInput, "Test Content");

    fireEvent.click(saveButton);
  });
});
