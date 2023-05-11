import { render, screen } from "@testing-library/react";
import { CreateComment } from "./CreateComment";

describe("tests for create comment component", () => {
  const renderComponent = () => {
    const mockFunc = jest.fn();

    return render(
      <CreateComment
        onSubmit={mockFunc}
        handleSubmit={mockFunc}
        register={mockFunc}
      />
    );
  };

  it("test render component", async () => {
    renderComponent();

    const formId = screen.getByTestId("comment-form-test-id");
    expect(formId).toBeInTheDocument();

    const inputId = screen.getByTestId("input-form-component-test-id");
    expect(inputId).toBeInTheDocument();

    const buttonId = screen.getByTestId("input-button-component-test-id");
    expect(buttonId).toBeInTheDocument();
  });
});
