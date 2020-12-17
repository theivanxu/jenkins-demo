@Library('github.com/releaseworks/jenkinslib') _
node("haimaxy-jnlp") {
    stage("Clone") {
        echo "1. Clone Stage"
        git url: "https://github.com/theivanxu/jenkins-demo.git"
        script {
            build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        }
    }
    stage("Test") {
        echo "2. Test Stage"
    }
    stage("Build") {
        echo "3. Build Docker Image Stage"
        sh "docker build -t 011383927026.dkr.ecr.us-west-2.amazonaws.com/jenkins-demo:${build_tag} ."
    }
    stage("Push") {
        echo "4. Push Docker Image Stage"
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'aws-key', usernameVariable: 'AWS_ACCESS_KEY_ID', passwordVariable: 'AWS_SECRET_ACCESS_KEY']]) {
            aws_password = AWS("--region=us-west-2 ecr get-login-password")
        }
        echo "aws password: ${aws_password}"
        sh "docker login --username AWS --password ${aws_password} 011383927026.dkr.ecr.us-west-2.amazonaws.com"
        sh "docker push 011383927026.dkr.ecr.us-west-2.amazonaws.com/jenkins-demo:${build_tag}"
    }
    stage("Deploy") {
        echo "5. Deploy Stage"
        def userInput = input(
            id: "userInput",
            message: "Choose a deploy environment",
            parameters: [
                [
                    $class: "ChoiceParameterDefinition",
                    choices: "DEV\nQA\nPROD",
                    name: "Env"
                ]
            ]
        )
        echo "This is a deploy step to ${userInput.Env}"
        sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        if (userInput.Env == "DEV") {
            echo "Deploy to dev environment"
        } else if (userInput.Env == "QA") {
            echo "Deploy to qa environment"
        } else if (userInput.Env == "PROD") {
            echo "Deploy to prod environment"
        } else {
            echo "No choose"
        }
        sh "kubectl apply -f k8s.yaml"
    }
}