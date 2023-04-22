import React from "react";

export const BoxLayout: React.FC<React.PropsWithChildren> = ({ children }) => {
  return (
    <div className="flex flex-col bg-box-color w-full p-two gap-one">
      {children}
    </div>
  );
};
