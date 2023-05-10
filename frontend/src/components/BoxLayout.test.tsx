import React from "react";
import { render, screen } from "@testing-library/react";

import { BoxLayout } from "./BoxLayout";

describe("Box Layout", () => {
  const childElement = <div>Child element</div>;

  const renderComponent = () => {
    return render(<BoxLayout>{childElement}</BoxLayout>);
  };

  it("render component", () => {
    renderComponent();

    const getComponent = screen.getByTestId("box-layout-test-id");
    expect(getComponent).toBeInTheDocument();
  });

  it("should render children", () => {
    renderComponent();

    const getChild = screen.getByText("Child element");
    expect(getChild).toBeInTheDocument();
  });
});
