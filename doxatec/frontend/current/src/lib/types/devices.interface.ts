import type Customer from "$lib/types/customer.interface";

interface Device {
  ID: string;
  Name: string;
  Customer: Customer;
  TempSup: number;
  TempMid: number;
  TempSub: number;
  Date: string;
}

export default Device;
