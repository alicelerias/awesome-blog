import React from "react";
import { Link } from "react-router-dom";
import { Menu } from "./Menu";

type props = {
  title: string;
};
export const Layout: React.FC<React.PropsWithChildren<props>> = ({
  children,
}) => {
  return (
    <div className="flex flex-col justify-between gap-one w-full h-auto p-one sm:p-one sm:mx-auto ">
      <div className="w-full flex flex-col sm:flex-row justify-between sm:fixed bg-black top-0">
        <div className="flex flex-row text-3xl sm:py-4 space-x-6 justify-center">
          <Link to="/"> !AWESOME</Link>
        </div>

        <div className="sm:fixed sm:right-24 sm:top-2">
          <Menu />
        </div>
      </div>

      <div className="flex flex-col sm:flex-row gap-one sm:gap-four sm:mt-four">
        {children}
      </div>
    </div>
  );
};
