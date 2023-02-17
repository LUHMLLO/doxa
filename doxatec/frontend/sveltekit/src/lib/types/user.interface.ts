import type Customer from "./customer.interface";

interface User {
  ID: string;
  Avatar: string;
  Username: string;
  Password: string;
  Customer: Customer
}

export default User;
