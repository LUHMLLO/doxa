interface Device {
  ID: string;
  Owner: string;
  Name: string;
  TempSup: number;
  TempMid: number;
  TempSub: number;
  Created: Date;
  Modified: Date;
}

export default Device;
