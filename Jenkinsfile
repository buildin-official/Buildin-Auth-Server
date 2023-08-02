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
		stage("Build and Push") {
			steps {
				withCredentials([[$class: 'UsernamePasswordMultiBinding',
								credentialsId: 'docker-hub',
								usernameVariable: 'DOCKER_USER_ID',
								passwordVariable: 'DOCKER_USER_PASSWORD'
				]]){
					sh 'docker login -u $DOCKER_USER_ID -p $DOCKER_USER_PASSWORD'
					sh 'docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --tag $DOCKER_USER_ID/buildin-auth::$BUILD_NUMBER --push .'
					sh 'docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --tag $DOCKER_USER_ID/buildin-auth::latest --push .'
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
