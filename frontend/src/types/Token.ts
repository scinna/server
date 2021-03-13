export type Token = {
  Token: string;
  LoginIP: string;
  CreatedAt: Date;
  LastSeen: Date|null;
  RevokedAt: Date|null;
};
