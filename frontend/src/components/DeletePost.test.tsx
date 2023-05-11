import { fireEvent, render, screen } from "@testing-library/react";
import { DeleteCommentButton } from "./DeleteCommentButton";
import { TestsContext } from "./testComonents/Context";
import * as mutations from "../api/mutations";
import { DeletePost } from "./DeletePost";

describe("test for delete component button", () => {
  const renderComponent = () => {
    const navigate = jest.fn();
    return render(
      <TestsContext>
        <DeletePost navigate={navigate} id="1" />
      </TestsContext>
    );
  };

  it("test render component", () => {
    renderComponent();
    const componentId = screen.getByTestId("delete-component-button-test-id");
    expect(componentId).toBeInTheDocument();
  });

  // it("test call function", () => {
  //   renderComponent()
  //   const fakeDelete = jest.spyOn(mutations, "deletePost")

  //   const componentId = screen.getByTestId("delete-component-button-test-id")
  //   fireEvent.click(componentId)
  //   expect(fakeDelete).toHaveBeenCalledTimes(1)
  // })
});
