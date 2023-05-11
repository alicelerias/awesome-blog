import { render, screen } from "@testing-library/react";
import { Observer } from "./IntersectionObserver";

describe("Observer", () => {
  it("should render without errors", () => {
    const fetchNextPage = jest.fn();
    render(<Observer fetchNextPage={fetchNextPage} />);
    expect(
      screen.getByTestId("intersection-observer-test-id")
    ).toBeInTheDocument();
  });
});
