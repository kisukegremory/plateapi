# Car Plate Service

A ideia do projeto é praticar golang com algo que nunca havia feito previamente e que possa ser utilizado por qualquer empresa que precise desse serviço futuramente e adicionar em seu banco de dados com pouca configuração

Para tal o nosso fluxo base será:

1. Client requisita os dados de uma placa
1. Consulta a placa no banco de dados
    1. Caso Exista na base e tenha dados -> retorna 200 e os dados do veículo
    
    1. Caso Não exista na base -> Retorna 202 e que será buscado
        1. Envia uma mensagem para o rabbit MQ
        1. O consumidor recebe a mensagem e consulta no serviço API Placa
        1. Caso tenha sucesso, Armaneza no banco de dados na tabela Plates e Vehicles
        1. Caso Não tenha sucesso, Armazena no banco de dados na tabela Plates
    1. Caso Exista na base e não tenha dados: 404


Vamos tentar estruturar como será feito o projeto:

1. Criar uma API utilizando gin, com o endpoint de v1/vehicles/{plate} que valida a placa e retorna um json padrão, 202
1. Adicionar um endpoint de autenticação que retorna um jwt válido por x minutos v1/auth
    ```
    Como o serviço só deve ser utilizado uma empresa, será consumido uma key única que gerará os jwt para serem utilizados como autenticação
    ```
1. Criar um docker-compose com o serviço do RabbitMQ
1. Alterar o endpoint de plate para enviar uma requisição para a fila vehicles.search
1. Criar um Consumer do rabbitMQ, que recebe a mensagem e busca na API Placa o modelo correspondente
    1. Caso tenha sucesso, envia uma mensagem para a fila vehicles.store, com o objeto VehiclePlates e VehicleAttributes(Para serem consumidas na mesma transaction, e enviar problemas de existir objeto Plate e não Vehicle)
    1. Caso não tenha sucesso, envia uma mensagem para a fila vehicles.store, com o objeto Plate 
1. Criar um serviço que receba as mensagens das filas vehicles.store
1. Adicionar no docker-compose o serviço do Mysql
1. Criar o ORM com Gorm e armazenar as informações nas tabelas VehiclePlates, VehicleAttributes
1. Alterar o serviço da API para consultar a placa na base e:
    1. Caso exista VehiclePlates & VehicleAttributes, retornar o objeto VehicleAttributes
    1. Caso exista o VehiclePlates apenas, retornar 404
    1. Caso não exista VehiclePlates, enviar a mensagem para a fila
1. Criar um Docker com o serviço funcional
1. Adicionar esse docker no docker-compose
1. Gravar um vídeo do serviço funcionando