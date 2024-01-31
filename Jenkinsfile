pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    // Install buildx component
                    sh 'docker buildx version || docker run --rm --privileged multiarch/qemu-user-static --reset -p yes'
                    sh 'docker buildx create --use'
                    sh 'docker buildx inspect --bootstrap'

                    // Build the Docker image with the specified Dockerfile
                    sh 'docker buildx build -t my-go-app .'
                }
            }
        }

        stage('Test') {
            steps {
                // Run unit tests
                sh 'docker run -rm my-go-app go test ./tests'
            }
        }

        stage('Deploy') {
            steps {
                // Deploy the image to the remote server
                // Assume the remote server has Docker installed and is running
                sh 'docker save my-go-app | ssh user@remote_server "docker load"'
                sh 'ssh user@remote_server ‘docker stop my-go-app || true'
                sh 'ssh user@remote_server ‘docker rm my-go-app || true'
                sh 'ssh user@remote_server ‘docker run -d --name my-go-app -p 8080:8080 my-go-app'
            }
        }
    }
}
