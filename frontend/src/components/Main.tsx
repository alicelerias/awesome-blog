import { Routes, Route, Link } from "react-router-dom"
import { HealthCheck } from "./HealthCheck"

export const Main: React.FC = () => {
    return (
    <>
    <Link to="/healthcheck">
        HealthCheck
       </Link>
       <Routes>
        <Route path="/healthcheck" element={ <HealthCheck />} />

       </Routes>
       </>

     
    )
}