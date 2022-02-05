export type AnyJson = boolean | number | string | null | undefined | JsonArray | JsonMap;
interface JsonMap { [key: string]: AnyJson; }
interface JsonArray extends Array<AnyJson> { }

export type Dependency = {
    id: string,
    dependencies: Dependency[]
}

export type ResourceOptions = {
    runId: number;
}

export type ResourceCreated = {
    data: AnyJson
    runId: number;
}

export type ResourceCreatedOutput = ResourceCreated | ResourceCreatedProzessWaitForLog;

export type Resource = {
    id: string,
    create: (options: ResourceOptions) => Promise<ResourceCreatedOutput>,
    init: (options: ResourceCreated) => Promise<void>,
    resources: Resource[]
}

export type TestSuiteType = {
    id: string;
    dependencies: Dependency[];
    resources: ResourceCreated[];
}

export type TestSuiteDiscoveryType = {
    id: string;
    dependencies: Dependency[];
    resources: Resource[];
}

export type TestSuiteCreateOptions = {
    dependencies?: Dependency[],
    resources?: Resource[]
}

export enum DiscoveryType {
    TestSuite = 'TESTSUITE',
    Test = 'TEST'
}

export enum ResourceCreationType {
    Local = 'Local',
    ProzessWaitForLog = 'ProzessWaitForLog'
}

export type ResourceCreatedProzessWaitForLog = {
    cmd: string,
    cwd: string,
    env: {
        [key: string]: string
    }
} & ResourceCreated;