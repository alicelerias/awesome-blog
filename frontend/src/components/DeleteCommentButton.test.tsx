import { render, screen } from "@testing-library/react";
import { DeleteCommentButton } from "./DeleteCommentButton";
import { TestsContext } from "./testComponents/Context";

describe("test for delete component button", () => {
  const renderComponent = () => {
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <DeleteCommentButton navigate={navigate} commentId="1" />
      </TestsContext>
    );
  };

  it("test render component", () => {
    renderComponent();
    const componentId = screen.getByTestId("delete-comment-button-test-id");
    const buttonId = screen.getByTestId("delete-button");
    expect(componentId).toBeInTheDocument();
    expect(buttonId).toBeInTheDocument();
  });
});
