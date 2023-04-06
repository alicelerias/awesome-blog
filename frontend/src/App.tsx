import React from 'react';
import { QueryClient, QueryClientProvider} from "react-query"
import './App.css'
import { HealthCheck } from './components/HealthCheck'
import { Main } from './components/Main';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
})

function App() {
  
  return (
   <QueryClientProvider client={queryClient}>
       <Main />
   </QueryClientProvider>
  );
}

export default App;
