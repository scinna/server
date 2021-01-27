import axios from "axios";
import {AuthRequest} from "@/types/Requests";

export const FetchUserInfos = () => axios.get('/api/account');
export const Authenticate = (infos: AuthRequest) => axios.post('/api/auth', infos)