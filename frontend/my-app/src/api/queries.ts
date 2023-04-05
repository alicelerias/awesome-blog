import configs from "../configs/configs"
import { HealthCheck } from "../types/HealthCheck"
import axios from "./axios"


export const healthCheck = async (): Promise<HealthCheck> => {
    const url = new URL(configs.API_URL + "/healthcheck")

    const { data } = await axios.get<HealthCheck>(url.toString())

    return data
}