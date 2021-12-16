import { Application } from "https://deno.land/x/oak/mod";

const app = new Application();

app.use((ctx) => {
    console.log("receive request!!")
    ctx.response.body = "Xin chao!";
});

await app.listen({ port: 9500 });