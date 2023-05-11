import { render, screen } from "@testing-library/react";
import { Layout } from "./Layout";
import { TestsContext } from "./testComonents/Context";

describe("test for Layout component", () => {
  it("test render layout with users component", () => {
    const navigate = jest.fn();
    render(
      <TestsContext>
        <Layout
          title="test component"
          navigate={navigate}
          usersComponent={true}
        />
      </TestsContext>
    );

    const componentId = screen.getByTestId("layout-component-test-id");
    const menuId = screen.getByTestId("menu-component-test-id");
    const usersId = screen.getByTestId("users-component-test-id");
    expect(componentId).toBeInTheDocument();
    expect(menuId).toBeInTheDocument();
    expect(usersId).toBeInTheDocument();
    expect(screen.queryByTestId("bla")).not.toBeInTheDocument();
  });

  it("test render layout without users component", () => {
    const navigate = jest.fn();
    render(
      <TestsContext>
        <Layout
          title="test component"
          navigate={navigate}
          usersComponent={false}
        />
      </TestsContext>
    );

    const usersId = screen.queryByText("users-component-test-id");
    expect(usersId).not.toBeInTheDocument();
  });
});
