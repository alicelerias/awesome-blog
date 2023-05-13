import { createContext } from "react";
import { User } from "../types";

export const CurrentUserContext = createContext({} as User | undefined);
