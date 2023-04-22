import { useQuery } from "react-query";
import { healthCheck } from "../api/queries";

export const HealthCheck: React.FC<{}> = () => {
  const { data, isLoading } = useQuery("healthCheck", healthCheck);

  return <div>{isLoading ? <p>is loading</p> : data?.status}</div>;
};
