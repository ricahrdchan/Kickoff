# Kickoff
First AWS docker project

Deploy to EC2

	1. docker-machine create --driver amazonec2 --amazonec2-open-port 8080 --amazonec2-region <region> <name>
		- For me: docker-machine create --driver amazonec2 --amazonec2-open-port 8080 --amazonec2-region us-east-2 aws-test1
		- This will start an ec2 instance with basic config (free aws ec2)  

	2. Ensure having ~/.aws/credentials

	3. docker-machine env <name>

	4. eval $(docker-machine env <name>)

	5. docker-compose up --build -d
		- Create docker container according to docker-compose
