// deno-lint-ignore-file no-explicit-any
import { Client } from "https://deno.land/x/mysql@v2.11.0/mod.ts";

interface ClientData {
  id: string;
  name: string;
  email: string;
  phone: string;
  created?: Date;
  modified?: Date | null;
}

export class ClientApi {
  private client: Client;

  constructor(config: any) {
    this.client = new Client();
    this.client.connect(config);
  }

  async getAll(): Promise<ClientData[]> {
    const result = await this.client.query("SELECT * FROM clients;");
    return result.rows;
  }

  async getById(id: string): Promise<ClientData | null> {
    const result = await this.client.query(
      "SELECT * FROM clients WHERE id = ?;",
      [id],
    );
    if (result.rows.length > 0) {
      return result.rows[0];
    } else {
      return null;
    }
  }

  async create(clientData: ClientData): Promise<ClientData | null> {
    const result = await this.client.query(
      "INSERT INTO clients (id, name, email, phone, created, modified) VALUES (?, ?, ?, ?, ?, ?);",
      [
        clientData.id,
        clientData.name,
        clientData.email,
        clientData.phone,
        clientData.created || new Date(),
        clientData.modified || null,
      ],
    );
    if (result.affectedRows > 0) {
      return { id: result.lastInsertId.toString(), ...clientData };
    } else {
      return null;
    }
  }

  async update(id: string, clientData: ClientData): Promise<ClientData | null> {
    const result = await this.client.query(
      "UPDATE clients SET name = ?, email = ?, phone = ?, modified = ? WHERE id = ?;",
      [
        clientData.name,
        clientData.email,
        clientData.phone,
        clientData.modified || new Date(),
        id,
      ],
    );
    if (result.affectedRows > 0) {
      return { id, ...clientData };
    } else {
      return null;
    }
  }

  async delete(id: string): Promise<ClientData | null> {
    const result = await this.client.query(
      "DELETE FROM clients WHERE id = ?;",
      [id],
    );
    if (result.affectedRows > 0) {
      return { id };
    } else {
      return null;
    }
  }
}
