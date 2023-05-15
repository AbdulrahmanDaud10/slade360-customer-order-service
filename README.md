# Slade360-Backend-Screening-Challange

<!-- ABOUT THE PROJECT -->
## About The Project

This project was designed to test the skills/knowledge/experience that a SIL Backend Developer should possess for the purposes of recruitment. It can also be used as a yardstick for training and improving SIL’s existing backend developers.

The technical interview was built around a coding assignment that is designed to screen for the following basic skills:
* Experience in developing REST and GraphQL APIs in Python/Go/C#/Java
* Experience with a configuration management tool e.g. Chef, Puppet, Ansible etc
* Experience working with containers and container orchestration tools:smile:
* Experience writing automated tests at all levels - unit, integration and acceptance testing
* Experience with CI/CD (any CI/CD platform)

They were particularly interested in:
- Testing + coverage + CI/CD
- HTTP and APIs e.g. REST
- OAuth2
- Web security e.g. XSS
- Logic / flow of thought
- Setting up a database
- Version control

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [x] Create a simple Go service
- [x] Design a simple customers and orders database (Keeping it simple)
- [x] Add REST API to input/upload customers and orders:
    - [x] Customers should have details e.g. names and code.
    - [x] Orders have simple details e.g. amount and time
- [x] Impliment authentication and authorization via JWT
- [ ] When an order is added, send the customer an SMS alterting them (You can use the Africa's talking SMS gateway and sandbox)
- [ ] Write unit test (with coverage checking) and set up CI + automated CD. You can deploy to any PAAS/FAAS/IAAS of your choice.
- [x] Write a README for the project and host it on your GitHub.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
###  Setup
You need to have Golang installed at the very least `Go 1.20.3`. 


The easiest way to get started with the rest api in your local environment is to clone it using git:

```
git clone https://github.com/AbdulrahmanDaud10/slade360-customer-order-service.git
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

**Environment Variables**

You need to configure environment variables for the rest api to work as expected.

Copy the content on .env_example to .env file 
 `cat .env_example > .env`

The .env file should look like this:

```
# database configurations
DB_HOST = "127.0.0.1"
DB_PORT = "5432"
DB_USER = "postgres"
DB_NAME = "customer_order"
DB_PASSWORD = ""

# port for app to listen and serve
LISTEN_ADDRESS = "127.0.0.1"
LISTEN_PORT = "8080"

```
You can change the values to your own preference.
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/slade360-customer-order-service`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/slade360-customer-order-service`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

 Abdulrahman Daud Miraj- [@AbdulDaudMiraj](https://twitter.com/AbdulDaudMiraj) - daudabdulrahman22@gmail.com

Project Link: https://github.com/AbdulrahmanDaud10/slade360-customer-order-service

<p align="right">(<a href="#readme-top">back to top</a>)</p>
