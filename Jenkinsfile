properties([disableConcurrentBuilds()])
// bump for testing
node("dind") {
    container('dind') {
        compat.test_build()
    }
}
