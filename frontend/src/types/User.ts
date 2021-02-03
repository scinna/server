import {Collection} from "@/types/Collection";

export type User = {
    UserID: string;
    Name: string;
    Email: string;
    IsAdmin: boolean;
    Collections?: Collection[];
}