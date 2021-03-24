#!/usr/bin/env groovy
pipeline {
    agent any
    tools {
        go "golang"
    }
    environment {
        GOPATH = "${env.WORKSPACE}/../go"
        PATH = "${PATH}:${GOPATH}/bin"
    }
    options {
        buildDiscarder(logRotator(daysToKeepStr: '90', artifactNumToKeepStr: '3'))
    }
    stages {
        stage('Preflight') {
            steps {
                echo sh(returnStdout: true, script: 'env')
                sh 'git log --reverse -1'
            }
        }
        stage('Init') {
            steps {
                sh 'rm -f go.sum'
                sh 'go get -u golang.org/x/tools/cmd/cover@latest'
                sh 'go get -u golang.org/x/tools/cmd/stringer@latest'
                sh 'go get -u github.com/t-yuki/gocover-cobertura@latest'
                sh 'go get -u github.com/jstemmer/go-junit-report@latest'
                sh 'stringer -type=DeviceType,BrowserName,OSName,Platform -output=vars/const_string.go vars/const.go'
                sh 'gofmt -s -w .'

            }
        }
        stage('Test') {
            steps {
                // Test with benchmark and coverage report
                sh 'go test -v ./... -bench=. -coverpkg=./... -coverprofile=coverage.txt.tmp -covermode=atomic 2>&1 | go-junit-report > test_results.xml'
                // remove auto-generated file of "go stringer" from coverage report
                sh 'cat coverage.txt.tmp | grep -v "vars/const_string.go" > coverage.txt'
                //
                sh 'gocover-cobertura < coverage.txt > coverage.xml'
            }
        }
    }
    post {
        always {
            junit 'test_results.xml'
            cobertura coberturaReportFile: 'coverage.xml'
            emailext subject: '$DEFAULT_SUBJECT',
                    body: '''${SCRIPT, template="report.template"}''',
                    from: 'jenkins@redeviation.com',
                    to: '$DEFAULT_RECIPIENTS'
        }
//        cleanup {
//            cleanWs()
//        }
    }
}