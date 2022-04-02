exports.complexCalc = (x, y) => {
    let result = 100000;

    for (i = 0; i < 200000; i++) {
        result += Math.pow(x, y) - y * y - Math.pow(y,x) + Math.random(); 
    }

    return result;
}

