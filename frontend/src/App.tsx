import React from "react";
import { QueryClient, QueryClientProvider, useQuery } from "react-query";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import { Main } from "./components/Main";
import { LoginPage } from "./components/LoginPage";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  return (
    <div className="bg-black text-white">
      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<LoginPage />} />
          </Routes>
          <Main />
        </BrowserRouter>
      </QueryClientProvider>
    </div>
  );
}

export default App;
