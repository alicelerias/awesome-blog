import React, { useEffect } from "react";
import { FetchNextPageOptions, InfiniteQueryObserverResult } from "react-query";
import { Posts } from "../types";

type props = {
  fetchNextPage: (
    options?: FetchNextPageOptions | undefined
  ) => Promise<InfiniteQueryObserverResult<Posts, unknown>>;
};

export const Observer: React.FC<props> = ({ fetchNextPage }) => {
  useEffect(() => {
    if (!window.IntersectionObserver) return;

    const intersectionObserver = new IntersectionObserver((entries) => {
      if (entries.some((entry) => entry.isIntersecting)) {
        fetchNextPage();
      }
    });
    intersectionObserver.observe(document.querySelector("#observer")!);
    return () => intersectionObserver.disconnect();
  }, []);

  return <li id="observer" className="text-black"></li>;
};
