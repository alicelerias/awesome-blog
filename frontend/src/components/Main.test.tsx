import { render, screen } from "@testing-library/react";
import { Main } from "./Main";
import { TestsContext } from "./testComonents/Context";

describe("test for main component", () => {
  it("test render maain", async () => {
    const handleSubmit = jest.fn();
    const register = jest.fn();
    const reset = jest.fn();
    render(
      <TestsContext>
        <Main handleSubmit={handleSubmit} register={register} reset={reset} />
      </TestsContext>
    );
    const componentId = screen.getByTestId("main-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });
});
