#!/usr/bin/env groovy

def deployAppinhouse(String glob, String sls) {
    def targetEnvName = 'appinhouse'
    def artifact = saltArtifact glob: glob, env: targetEnvName, job: env.JOB_NAME, build: env.BUILD_NUMBER
    saltDeploy(env: targetEnvName,
        artifact: artifact,
        sls: sls,
        nodegroup: 'appinhouse'
    )
}

pipeline {
    agent {
        label "os:linux"
    }
    options {
        skipDefaultCheckout()
        disableConcurrentBuilds()
        buildDiscarder(logRotator(
            daysToKeepStr: '15',
            artifactNumToKeepStr: '20'
        ))
        ansiColor('xterm')
    }
    parameters {
        booleanParam(name: 'DEPLOY_SERVER',
            defaultValue: false,
            description: 'When checked, will automatically deploy server (backend).')
        booleanParam(name: 'DEPLOY_WEB',
            defaultValue: false,
            description: 'When checked, will automatically deploy web (frontend).')
    }
    environment {
        GITHUB_URL = 'https://github.com/rog2/appinhouse'
        GO_VERSION = '1.10'
        GOPATH = "${env.WORKSPACE}"
        SERVER_ROOT = "${env.WORKSPACE}/src/appinhouse/server"
        WEB_ROOT = "${env.WORKSPACE}/src/appinhouse/web"
        SERVER_TARBALL = artifactName(name: 'appinhouse', extension: 'server.tar.gz')
        WEB_TARBALL = artifactName(name: 'appinhouse', extension: 'web.tar.gz')
    }
    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM',
                    userRemoteConfigs: [[url: env.GITHUB_URL]],
                    branches: [[name: env.BRANCH_NAME ?: 'master']],
                    browser: [$class: 'GithubWeb', repoUrl: env.GITHUB_URL],
                    extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'src/appinhouse']]
                ])
            }
        }
        stage('Remove Last Build') {
            steps {
                sh 'rm -vf *.tar.gz'
            }
        }
        stage('Build Server') {
            steps {
                dir (env.SERVER_ROOT) {
                    withGo(env.GO_VERSION) {
                        sh """
                            go get -v github.com/beego/bee
                            go get -v
                            go build -o appinhouse
                            ${env.GOPATH}/bin/bee pack -o ${env.WORKSPACE} -exr pack.sh -exr server -exr test.conf -exr bin
                        """
                    }
                }
                sh "mv server.tar.gz ${env.SERVER_TARBALL}"
            }
        }
        stage('Build Web') {
            steps {
                dir (env.WEB_ROOT) {
                    sh "tar czvf ${env.WORKSPACE}/${env.WEB_TARBALL} static"
                }
            }
        }
        stage('Archive') {
            steps {
                archiveArtifacts artifacts: '*.tar.gz', onlyIfSuccessful: true
            }
        }
        stage('Deploy Server') {
            when {
                expression {
                    return params.DEPLOY_SERVER
                }
            }
            steps {
                script {
                    deployAppinhouse(env.SERVER_TARBALL, 'appinhouse.server')
                }
            }
        }
        stage('Deploy Web') {
            when {
                expression {
                    return params.DEPLOY_WEB
                }
            }
            steps {
                script {
                    deployAppinhouse(env.WEB_TARBALL, 'appinhouse.web')
                }
            }
        }
    }
}
