import { BoxPosts } from "./BoxPosts";
import { MenuPosts } from "./MenuPosts";
import { PropsWithChildren, useEffect } from "react";
import { NavigateFunction } from "react-router-dom";
import { Post } from "../types";
import { useInfiniteQuery } from "react-query";
import { getFeed } from "../api/queries";

type props = {
  navigate: NavigateFunction;
};

export const FeedComponent: React.FC<PropsWithChildren<props>> = ({
  navigate,
}) => {
  const { data, isLoading, fetchNextPage, hasNextPage } = useInfiniteQuery(
    "feed",
    getFeed,
    {
      getNextPageParam: (data) => data.next_link,
    }
  );

  const posts =
    data?.pages.reduce(
      (previous, current) => [...previous, ...current.content],
      [] as Post[]
    ) || [];

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

  return (
    <BoxPosts
      isLoading={isLoading}
      data={posts}
      navigate={navigate}
      children={<MenuPosts />}
      button={<li id="observer" className="text-black"></li>}
    />
  );
};
