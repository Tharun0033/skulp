pipeline {
    agent any

    stages {
        stage('Clone Repository') {
            steps {
                // Checkout code from your GitHub repository
                git branch: 'main', url: 'https://github.com/Tharun0033/skulp.git'
            }
        }

        stage('Set Up Maven') {
            steps {
                // Ensure Maven is installed and set up
                script {
                    if (!isUnix()) {
                        error('This script is designed for a Unix-like environment.')
                    }
                    sh 'mvn -version'
                }
            }
        }

        stage('Build') {
            steps {
                // Run Maven clean and package
                sh 'mvn clean package'
            }
        }

        stage('Test') {
            steps {
                // Run Maven tests
                sh 'mvn test'
            }
        }

        stage('Archive Artifacts') {
            steps {
                // Archive the generated JAR or WAR files
                archiveArtifacts artifacts: '**/target/*.jar', fingerprint: true
            }
        }
    }

    post {
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed. Please check the logs for details.'
        }
    }
}
