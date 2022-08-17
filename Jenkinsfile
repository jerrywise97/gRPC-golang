
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
        // PROJECT_ID = 'muqtar-project'
        // CLUSTER_NAME = ''
        // LOCATION = ''
        // CREDENTIALS_ID = 'kubernetes'
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
                sh 'go get ./...'
                sh 'docker build -t jerrywise97/gPRC_coffee_shop'.
            }
        }
    }
}
    
    
    
    
    
    
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push shadowshotx/product-go-micro'
        //         }
        //     }
        // }
    
//       stage("unit-test") {
//             steps {
//                 echo 'UNIT TEST EXECUTION STARTED'
//                 sh 'make unit-tests'
//             }
//         }
//         stage("functional-test") {
//             steps {
//                 echo 'FUNCTIONAL TEST EXECUTION STARTED'
//                 sh 'make functional-tests'
//             }
//         }
