const { aTestDependency } = require('./a.test');
const { TestDependencies: TestDependency } = require('test-framework');

const bTestDependency = TestDependency("b", aTestDependency);

console.log('b', JSON.stringify(bTestDependency));