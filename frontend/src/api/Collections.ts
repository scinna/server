import axios from "axios";

export const FetchCollection = (username: string, collection: string) => axios.get('/api/browse/' + username + '/' + collection);
