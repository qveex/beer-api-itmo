FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 8080
EXPOSE 8081

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /protos
COPY ../../protos .
WORKDIR /src
COPY ["CsharpBeer/CsharpBeer.OrderService/CsharpBeer.OrderService.csproj", "CsharpBeer.OrderService/"]
RUN dotnet restore "CsharpBeer.OrderService/CsharpBeer.OrderService.csproj"
COPY CsharpBeer/CsharpBeer.OrderService/ CsharpBeer.OrderService/
WORKDIR "/src/CsharpBeer.OrderService"
RUN dotnet build "CsharpBeer.OrderService.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "CsharpBeer.OrderService.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "CsharpBeer.OrderService.dll"]
