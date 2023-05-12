import { Application, Router } from "https://deno.land/x/oak/mod.ts";
import { ClientApi } from "./client.ts";

const clientApi = new ClientApi({
    hostname: "localhost",
    port: 3306,
    username: "root",
    password: "password",
    db: "my_database",
});

const app = new Application();
const router = new Router();

router.get("/clients", async (ctx) => {
    ctx.response.body = await clientApi.getAll();
});

router.get("/clients/:id", async (ctx) => {
    const client = await clientApi.getById(ctx.params.id);
    if (client) {
        ctx.response.body = client;
    } else {
        ctx.response.status = 404;
    }
});

router.post("/clients", async (ctx) => {
    const client = await clientApi.create(ctx.request.body);
    ctx.response.status = 201;
    ctx.response.body = client;
});

router.put("/clients/:id", async (ctx) => {
    const client = await client
    Api.update(ctx.params.id, ctx.request.body);
    if (client) {
        ctx.response.body = client;
    } else {
        ctx.response.status = 404;
    }
});

router.delete("/clients/:id", async (ctx) => {
    const client = await clientApi.delete(ctx.params.id);
    if (client) {
        ctx.response.body = client;
    } else {
        ctx.response.status = 404;
    }
});

app.use(router.routes());

await app.listen({ port: 8000 });