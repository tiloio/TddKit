export const WRONG_PHASE = (currentPhase: string, allowedPhases: string) => {
    throw new Error(`Test-Framework: Wrong phase in environment variable "TEST_FRAMEWORK_PHASE=${currentPhase}". 
    Only following values are allowed: ${allowedPhases}.`);
}
