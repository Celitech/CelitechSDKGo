pipeline {
    agent {
        docker {
            image 'golang:1.18'
        }
    }
    environment {
        MANIFEST_PATH = '.manifest.json'
    }
    stages {
        stage('Run Only on Tag') {
            when {
                expression {
                    // Matches tags like rls_0.1.0, rls_0.1.1, etc.
                    return env.BRANCH_NAME ==~ /^rls_\d+\.\d+\.\d+$/  
                }
            }
            steps {
                echo "Running on tag: $env.BRANCH_NAME"
            }
            stages {
                stage('Pre-check for Environment Variables') {
                    steps {
                        script {
                            if (!env.MANIFEST_PATH) {
                                error("Error: MANIFEST_PATH is not defined")
                            }
                            echo "MANIFEST_PATH is defined."
                        }
                    }
                }
                stage('Publish Go Module') {
                    steps {
                        sh '''
                            apt-get update && apt-get install -y jq
                            
                            # Extract the Go module name from the manifest file
                            MODULE_NAME=$(jq -r '.config.languageOptions.go.goModuleName' $MANIFEST_PATH)
                            echo "Go module name: $MODULE_NAME"
                            
                            if [ -z "$MODULE_NAME" ]; then
                                echo "Error: MODULE_NAME is empty"
                                exit 1
                            fi
                            
                            # Check if the go.mod file exists
                            if [ ! -f "go.mod" ]; then
                                echo "Error: go.mod file not found"
                                exit 1
                            fi
                            
                            # Download Go modules
                            go mod download
                            
                            # List the Go module
                            GOPROXY=proxy.golang.org go list -m $MODULE_NAME
                        '''
                    }
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}
