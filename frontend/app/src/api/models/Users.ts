export interface ILoginUser {
    ID: number,
    CreatedAt: Date,
    Email: string,
    Username: string
}

export interface ILoginResponse {
    CurrentUser: ILoginUser,
    Token: string,
}