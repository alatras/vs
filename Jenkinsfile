coverageSummaryText = 'unknown'

pipeline {
    agent { label 'master' }
    environment {
        LOG_LEVEL = "error"
    }
    stages {
        stage('Docker build') {
            steps {
                configFileProvider([configFile(fileId: '.env', targetLocation: '.env')]) {
                    sh 'docker build --target build -t validation-service .'
                    sh 'docker create --name validation-service validation-service'
                    sh 'docker cp validation-service:/build_artifacts .'
                }
            }
        }
    }
    post {
        always {
            junit './build_artifacts/coverage.xml'

            archiveArtifacts artifacts: 'build_artifacts/**', fingerprint: true

            configFileProvider([configFile(fileId: '.env', targetLocation: '.env')]) {
                sh 'docker rm -f -v validation-service'
            }

            cleanWs()
        }
        failure {
            slackSend channel: '#ci', color: 'danger', message: "*Validation Service FAILED* `${env.JOB_NAME}` on ${env.STAGE_NAME}\n- ${env.BUILD_DISPLAY_NAME}: ${env.RUN_DISPLAY_URL}\n- Commit: `${env.GIT_PREVIOUS_COMMIT}`", tokenCredentialId: 'slack-api'
        }
        success {
            slackSend channel: '#ci', color: 'good', message: "*Validation Service SUCCESS* `${env.JOB_NAME}`\n- ${env.BUILD_DISPLAY_NAME}: ${env.RUN_DISPLAY_URL}\n- Commit: `${env.GIT_PREVIOUS_COMMIT}`", tokenCredentialId: 'slack-api'
        }
    }
}
