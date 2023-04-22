type props = React.DetailedHTMLProps<
  React.InputHTMLAttributes<HTMLInputElement>,
  HTMLInputElement
> & {
  name: string;
};

export const InputButton: React.FC<props> = ({ name }) => {
  return (
    <button
      className={` bg-blue p-1 w-1/5 text-sm`}
      data-testid={"input-button-component-test-id"}
      type="submit"
    >
      {name}
    </button>
  );
};
