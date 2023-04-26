import React from "react";
import { FieldError, FieldErrorsImpl, Merge } from "react-hook-form";

type props = React.DetailedHTMLProps<
  React.InputHTMLAttributes<HTMLInputElement>,
  HTMLInputElement
> & {
  error?: FieldError | Merge<FieldError, FieldErrorsImpl<any>>;
  controller: any;
};

export const InputForm: React.FC<props> = ({ error, controller, ...rest }) => {
  return (
    <div className="flex flex-col gap-one">
      <textarea
        data-testid={"input-form-component-test-id"}
        className="bg-transparent text-sm w-full h-three overflow-visible p-one italic"
        {...controller}
        {...rest}
      />
      {error && (
        <div className="flex justify-center bg-red-800 p-1 text-smm mb-2">
          {(error?.message as string) || "Invalid field"}
        </div>
      )}
    </div>
  );
};
