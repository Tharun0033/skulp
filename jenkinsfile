pipeline {
    agent any

    environment {
        // Define environment variables if needed
        REPO_URL = 'https://github.com/Tharun0033/skulp.git'
        BRANCH = 'main' // Change this to your desired branch name
    }

    stages {
        stage('Clone Repository') {
            steps {
                echo 'Cloning the repository...'
                git branch: "$BRANCH", url: "$REPO_URL"
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies...'
                // Add the commands to install your project dependencies
                sh 'pip install -r requirements.txt' // Example for Python projects
            }
        }

        stage('Build') {
            steps {
                echo 'Building the project...'
                // Add the commands to build your project
                sh 'python setup.py build' // Replace with your build command
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
                // Add the commands to run your tests
                sh 'pytest' // Replace with your test command
            }
        }
    }

    post {
        success {
            echo 'Build and tests completed successfully!'
        }
        failure {
            echo 'Build or tests failed. Please check the logs.'
        }
    }
}
