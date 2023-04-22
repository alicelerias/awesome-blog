import { render, screen } from "@testing-library/react";
import { Main } from "./Main";
import { TestsContext } from "./testComonents/Context";

describe("test for main component", () => {
  it("test render maain", async () => {
    render(
      <TestsContext>
        <Main />
      </TestsContext>
    );
    const componentId = screen.getByTestId("main-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });
});
