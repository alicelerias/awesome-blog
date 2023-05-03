import React from "react";

const Skeleton: React.FC<{}> = () => {
  return (
    <div className="animate-pulse">
      <div className="flex gap-one pb-one">
        <div className="w-auto h-64">
          <img
            className=" w-20 aspect-square"
            src={"https://ionicframework.com/docs/img/demos/avatar.svg"}
            alt=""
          />
        </div>
        <div className="p-two bg-box-color w-full"></div>
      </div>

      <div className="flex gap-one pb-one">
        <div className="w-auto h-64">
          <img
            className=" w-20 aspect-square"
            src={"https://ionicframework.com/docs/img/demos/avatar.svg"}
            alt=""
          />
        </div>
        <div className="p-two bg-box-color w-full"></div>
      </div>
    </div>
  );
};

export default Skeleton;
