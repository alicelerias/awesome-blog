import { render, screen } from "@testing-library/react";
import { Alert } from "./Alert";

describe("test alaert component", () => {
  test("alert component renders correctly", () => {
    render(<Alert message="test message" type="error" />);
    const alertComponent = screen.getByTestId("alert-component-id");
    expect(alertComponent).toBeInTheDocument();
    expect(alertComponent).toHaveClass("bg-red-800");
    expect(alertComponent).toHaveTextContent("test message");
  });
});
