export type Dependency = {
    id: string,
    dependencies: Dependency[]
}

export type Resource = {
    id: string,
    create: () => Promise<void>,
    init: () => Promise<void>,
    resources: Resource[]
}

export type TestSuiteType = {
    id: string;
    dependencies: Dependency[];
    resources: Resource[];
}

export type TestSuiteCreateOptions = {
    dependencies?: Dependency[],
    resources?: Resource[]
}

export type DiscoveredResource = {
    id: string,
    resources: DiscoveredResource[]
}

export enum DiscoveryType {
    TestSuite = 'TESTSUITE',
    Test = 'TEST'
}