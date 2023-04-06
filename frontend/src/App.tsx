import React from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter } from "react-router-dom";
import "./App.css";
import { HealthCheck } from "./components/HealthCheck";
import { Main } from "./components/Main";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  return (
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        learn react
        <Main />
      </QueryClientProvider>
    </BrowserRouter>
  );
}

export default App;
