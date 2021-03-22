import {Media} from "./Media";

export type Collection = {
    CollectionID: string;
    Title: string;
    Visibility: number;
    IsDefault: boolean;
    Collections: Collection[] | null;
    Medias: Media[] | null;
}