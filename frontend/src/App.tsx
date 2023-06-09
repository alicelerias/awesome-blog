import React from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Main } from "./components/Main";
import { LoginPage } from "./components/LoginPage";
import { useForm } from "react-hook-form";
import { Register } from "./components/RegisterPage";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  const {
    handleSubmit,
    register,
    reset,
    setValue,
    formState: { errors },
  } = useForm();

  return (
    <div className="bg-black text-white">
      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <Routes>
            <Route
              path="/register"
              element={
                <Register
                  handleSubmit={handleSubmit}
                  register={register}
                  errors={errors}
                  reset={reset}
                />
              }
            />
            <Route
              path="/login"
              element={
                <LoginPage handleSubmit={handleSubmit} register={register} />
              }
            />
            <Route
              path="*"
              element={
                <Main
                  setValue={setValue}
                  handleSubmit={handleSubmit}
                  register={register}
                  reset={reset}
                  errors={errors}
                />
              }
            />
          </Routes>
        </BrowserRouter>
      </QueryClientProvider>
    </div>
  );
}

export default App;
