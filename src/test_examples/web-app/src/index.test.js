const { default: fetch } = require("node-fetch")


test("saves one object", async () => {
    const response = await fetch("http://localhost:3000", {method: 'POST'});

    expect(await response.text()).toEqual("item added");
});

test("returns amount of saved objects", async () => {
    await fetch("http://localhost:3000", {method: 'POST'});
    await fetch("http://localhost:3000", {method: 'POST'});
    const response = await fetch("http://localhost:3000");

    expect(await response.text()).toEqual("Got 2");
});
