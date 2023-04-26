import React, { PropsWithChildren, ReactNode, useState } from "react";

type props = {
  name: ReactNode | string | undefined;
};

export const Sidebar: React.FC<PropsWithChildren<props>> = ({
  children,
  name,
}) => {
  const [isOpen, setOpen] = useState(false);

  const handleDropDown = () => {
    setOpen(!isOpen);
  };

  return (
    <div className="w-6">
      <button onClick={handleDropDown}>{name}</button>

      <div
        id="dropdown"
        className={`z-10 w-32 bg-white rounded divide-y divide-gray-100 shadow ${
          isOpen ? "block" : "hidden"
        }`}
      >
        <ul className=" z-10 w-32 bg-white rounded divide-y divide-gray-100 shadow ">
          <li>{children}</li>
        </ul>
      </div>
    </div>
  );
};
