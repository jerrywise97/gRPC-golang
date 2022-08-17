
pipeline {
    // install golang 1.15 on Jenkins node
    agent any
    tools {
        go 'go1.17'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
// 		sh 'docker version'
                sh 'go get ./...'
//                 sh 'docker build -t jerrywise97/gprc-coffee-shop:v1 .'
			    script {
                    sh 'docker build -t jerrywise97/gprc-coffee-shop:v1 .'
			    }
            }
        }
    }
}
