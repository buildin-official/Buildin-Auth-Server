def component = [
		Preprocess: false,
		Hyper: false,
		Train: false,
		Test: false,
		Bento: false
]
pipeline {
	agent any
	stages {
		stage("Checkout") {
			steps {
				checkout scm
			}
		}
		stage("Build") {
      steps {
        sh "docker buildx build --platform linux/amd64 --tag buildin-auth:latest ."

      }
		}
	  stage("Tag and Push") {
			steps {
        withCredentials([[$class: 'UsernamePasswordMultiBinding',
								          credentialsId: 'docker-hub',
								          usernameVariable: 'DOCKER_USER_ID',
								          passwordVariable: 'DOCKER_USER_PASSWORD'
				]]){
					sh 'docker login -u $DOCKER_USER_ID -p $DOCKER_USER_PASSWORD'
					sh 'docker tag buildin-auth:latest $DOCKER_USER_ID/buildin-auth:$BUILD_NUMBER'
					sh 'docker push $DOCKER_USER_ID/buildin-auth:$BUILD_NUMBER'
					sh 'docker tag $DOCKER_USER_ID/buildin-auth:$BUILD_NUMBER $DOCKER_USER_ID/buildin-auth:latest'
					sh 'docker push $DOCKER_USER_ID/buildin-auth:latest'
				}	
			}
		}
		stage("SSH-Deploy") {
			steps {
				script {
					// Secret Text 타입의 자격 증명을 불러옵니다.
					withCredentials([
						string(credentialsId: 'buildin-server-host', variable: 'HOST'),
						string(credentialsId: 'buildin-server-port', variable: 'PORT'),
						sshUserPrivateKey(credentialsId: 'buildin-server', keyFileVariable: 'identity', passphraseVariable: '', usernameVariable: 'userName'),
					]) {
						sshagent (credentials: ['buildin-server']) {
                sh '''
								ssh -o StrictHostKeyChecking=no -p $PORT $userName@$HOST '
								cd ~/docker-compose/Buildin-Auth-Server
								git pull origin main
								docker pull implude/buildin-auth:latest
								doppler run -- docker compose -f docker-compose.prod.yml up -d
								'
								'''
						}
					}
				}
			}
		}
  }
}
