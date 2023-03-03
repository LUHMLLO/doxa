interface Device {
  id: string;
  owner: string;
  name: string;
  temp_sup: number;
  temp_mid: number;
  temp_sub: number;
  created: Date;
  modified: Date;
}

export default Device;
