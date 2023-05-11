import { fireEvent, render, screen } from "@testing-library/react";
import { DeleteCommentButton } from "./DeleteCommentButton";
import { TestsContext } from "./testComonents/Context";
import * as mutations from "../api/mutations";

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

  // it("test call function", () => {
  //   renderComponent()
  //   const fakeDelete = jest.spyOn(mutations, "deleteComment").mockImplementation()

  //   const componentId = screen.getByTestId("delete-button")
  //   fireEvent.click(componentId)
  //   expect(fakeDelete).toHaveBeenCalledTimes(1)
  // })
});
