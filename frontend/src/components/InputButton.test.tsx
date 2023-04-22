import { fireEvent, render, screen } from "@testing-library/react";
import { InputButton } from "./InputButton";
import { TestsContext } from "./testComonents/Context";

describe("test for input form component", () => {
  const renderComponent = () => {
    return render(
      <TestsContext>
        <InputButton name="testName" />
      </TestsContext>
    );
  };

  it("test render input button", async () => {
    renderComponent();
    const componentId = screen.getByTestId("input-button-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test name", async () => {
    renderComponent();
    expect(screen.getByText("testName")).toBeInTheDocument();
  });
});
