import { useQuery } from "react-query"
import { healthCheck } from "../api/queries"

export const HealthCheck: React.FC = () => {
    const { data, isLoading } = useQuery("healthCheck", healthCheck)

    return (
        <div>
            {isLoading ? (
                "is loading"
            ) : (
                 data?.status 
            )}
        </div>
    )
}