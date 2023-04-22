import { useCallback, useState } from "react";

export const useAlert = () => {
  const [show, setShow] = useState(false);
  const showAlert = useCallback(() => {
    setShow(true);
    setTimeout(() => {
      setShow(false);
    }, 5000);
  }, []);

  const alert: React.FC<{ message: string; type: string }> = ({
    message,
    type,
  }) => {
    const color = type === "error" ? "bg-red-800" : "bg-green-300";

    return (
      <>
        {show && (
          <div className={`w-full p-2 m-2 ${color}`}>
            <p className="w-full text-white text-center font-semibold text-sm">
              {message}
            </p>
          </div>
        )}
      </>
    );
  };

  return [showAlert, alert] as [() => void, typeof alert];
};
