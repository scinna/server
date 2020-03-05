export default class File {

    parent: File|null;
    filename: string;
    isFolder: boolean;
    content: File[];

    constructor(filename: string, isFolder: boolean, parent: File|null = null) {
        this.parent = parent;
        this.filename = filename;
        this.isFolder = isFolder;
        this.content = [];
    }

    GetFullPath(): string {
        if (this.parent !== null) {
            return this.parent?.GetFullPath() + '/' + this.filename;
        }    

        return '';
    }

    IsRoot(): boolean {
        return this.parent === null || this.parent === undefined;
    }

}