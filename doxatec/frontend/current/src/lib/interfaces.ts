export interface User {
  id: string;
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


export interface Device {
  id: string;
  owner: string;
  name: string;
  temp_sup: number;
  temp_mid: number;
  temp_sub: number;
  created: Date;
  modified: Date;
}

