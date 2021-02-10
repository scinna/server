import axios from "axios";

export const Browse = (username: string, collection: string) => axios.get('/api/browse/' + username + '/' + collection);
