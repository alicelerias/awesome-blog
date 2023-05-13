export const Alert: React.FC<{ message: string; type: string }> = ({
  message,
  type,
}) => {
  const color = type === "error" ? "bg-red-800" : "bg-green-300";

  return (
    <div data-testid="alert-component-id" className={`w-full p-2 m-2 ${color}`}>
      <p className="w-full text-white text-center font-semibold text-sm">
        {message}
      </p>
    </div>
  );
};
