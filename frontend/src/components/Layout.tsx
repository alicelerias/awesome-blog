import React from "react";
import { Link, NavigateFunction } from "react-router-dom";
import { Menu } from "./Menu";
import { UsersComponent } from "./Users";

type props = {
  title: string;
  navigate: NavigateFunction;
  usersComponent: boolean;
};
export const Layout: React.FC<React.PropsWithChildren<props>> = ({
  children,
  navigate,
  usersComponent,
}) => {
  return (
    <div
      data-testid="layout-component-test-id"
      className="flex flex-col justify-between gap w-full h-auto p-one gap-one"
    >
      <div className="w-full flex flex-col sm:flex-row justify-between sm:fixed bg-black top-0">
        <div className="flex flex-row text-3xl sm:py-4 space-x-6 justify-center">
          <Link to="/"> !AWESOME</Link>
        </div>

        <div className="sm:fixed right-24 top-2">
          <Menu navigate={navigate} />
        </div>
      </div>

      <div className="flex flex-col sm:flex-row gap-one sm:gap-four sm:mt-four">
        {children}

        {usersComponent && <UsersComponent navigate={navigate} />}
      </div>
    </div>
  );
};
