# Calculator Enviorment Database (CED)

CED is my Highschool exam work to [WSP](https://wsp.com), that are a consulting company.
The is development with [WSP](https://wsp.com) and their needs to the new climate declration the sweidsh goverment have added for the building industry.

## Getting Started

Follow steps below to get started!
* Clone the repo LINK
* Edit Server.gRPC/appsettings.json to fit your need.
* Go back to the root folder and run `docker-compose up -d`

### Prerequisites

[Docker](https://www.docker.com/)
[.NET-Core 3.0 or above](https://dotnet.microsoft.com/download)
[Microsoft SQL 2019](https://www.microsoft.com/sv-se/sql-server/sql-server-downloads)

### Installing

```
Clone repo
```

```
Open folder or sln in preferd editor/IDE of choice.
```

```
Edit DataLibrary config.json to your needs
```

```
Run docker-compose up -d
```

To access the application, locate to browser and enter the public port to the container.
To create projects and update you need to create your own gRPC client, at the moment. Will maybe change it at a later point.

## Deployment

### WIP 

## Built With

* [Blazor](https://dotnet.microsoft.com/apps/aspnet/web-apps/blazor) - The web is written in Blazor
* [.NET-Core](https://dotnet.microsoft.com/download) - Base framework
* [Docker](https://www.docker.com/) - Used to run development containers and publish docker images
* [gRPC](https://grpc.io/) - gRPC
* [Azure DevOps](dev.azure.com) - WIll be used as CI/CD

## Contributing

### WIP

## Authors

* **Leo Rönnebro** - *Initial work* - [Github](https://github.com/lilahamstern) [Twitch](https://twitch.tv/lilahamstern)

## License

This project is licensed under the GNU General Public License v3.0 License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Plan and create a bigger application.
* Database design
