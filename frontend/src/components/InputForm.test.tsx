import { fireEvent, render, screen } from "@testing-library/react";
import { InputForm } from "./InputForm";
import { TestsContext } from "./testComonents/Context";

describe("test for input form component", () => {
  const renderComponent = () => {
    const setValue = jest.fn();
    return render(
      <TestsContext>
        <InputForm
          controller={""}
          type="text"
          onChange={(e) => {
            setValue(e.target.value);
          }}
        />
      </TestsContext>
    );
  };

  it("test render input form", async () => {
    renderComponent();
    const componentId = screen.getByTestId("input-form-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test input", async () => {
    renderComponent();
    const componentId = screen.getByTestId("input-form-component-test-id");
    fireEvent.change(componentId, {
      target: {
        value: "test",
      },
    });
    expect(componentId).toHaveValue("test");
  });
});
