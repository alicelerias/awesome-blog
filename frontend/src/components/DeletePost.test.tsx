import { render, screen } from "@testing-library/react";
import { TestsContext } from "./testComponents/Context";

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
});
