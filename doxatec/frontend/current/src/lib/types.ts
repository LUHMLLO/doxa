export type User = {
  id: string;
  jwt: string;
  username: string;
  password: string;
  avatar: string;
  name: string;
  email: string;
  phone: string;
  role: string;
  created: Date;
  modified: Date;
}

export type Device = {
  id: string;
  owner: string;
  name: string;
  pin: string;
  temp_sup: number;
  temp_mid: number;
  temp_sub: number;
  created: Date;
  modified: Date;
}

