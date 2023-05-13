import { render, screen } from "@testing-library/react";
import { Main } from "./Main";
import { TestsContext } from "./testComponents/Context";

describe("test for main component", () => {
  it("test render maain", async () => {
    const handleSubmit = jest.fn();
    const register = jest.fn();
    const reset = jest.fn();
    const setValue = jest.fn();
    render(
      <TestsContext>
        <Main
          setValue={setValue}
          handleSubmit={handleSubmit}
          register={register}
          reset={reset}
        />
      </TestsContext>
    );
    const componentId = screen.getByTestId("main-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });
});
