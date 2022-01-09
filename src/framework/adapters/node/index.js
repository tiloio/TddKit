const test = async (name, fn) => {
    try {
        await fn();
        console.info(JSON.stringify({
            name,
        }));
    } catch (err) {
        console.error(JSON.stringify({
            name,
            err: JSON.stringify({
                 message: err.message,
            }) // todo specify format for global communication
        }));
    }
}