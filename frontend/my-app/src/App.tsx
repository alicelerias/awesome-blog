import React from 'react';
import { QueryClient, QueryClientProvider} from "react-query"
import './App.css';
import { HealthCheck } from './components/healthCheck';

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
    
       <HealthCheck />
   </QueryClientProvider>
  );
}

export default App;
