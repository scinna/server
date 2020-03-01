export default class File {

    filename: string;
    isFolder: boolean;

    constructor(filename: string, isFolder: boolean) {
        this.filename = filename;
        this.isFolder = isFolder;
    }

}